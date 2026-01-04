package database

import (
	"fmt"
	"reflect"

	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/timing"

	"github.com/mt1976/frantic-core/logHandler"
)

func (db *DB) Retrieve(fieldName Field, value, to any) (any, error) {
	logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db] {%+v}", GetStructType(to), fieldName, value, GetStructType(to), db.Name, db)

	if db.withCaching {
		cacheKeyValue := value.(string)
		if cachedValue, found := inMemoryCache[cacheKeyValue]; found {
			reflect.ValueOf(to).Elem().Set(reflect.ValueOf(cachedValue).Elem())
			logHandler.CacheLogger.Printf("[GET]<%v>{HIT} (%v=%v) [%+v] [...%v.db] on %v - {SKIP} DB Access", GetStructType(to), db.withCacheKey, cacheKeyValue, GetStructType(to), db.Name, "Retrieve")
			return cachedValue, nil
		}
		logHandler.CacheLogger.Printf("[GET]<%v>{MISS} (%v=%v) [%+v] [...%v.db] on %v - Accessing DB", GetStructType(to), db.withCacheKey, cacheKeyValue, GetStructType(to), db.Name, "Retrieve")
	}

	logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db] - From Database", GetStructType(to), fieldName, value, GetStructType(to), db.Name)

	// [GET] from database
	err := db.connection.One(string(fieldName), value, to)

	// Store in cache if caching is enabled and retrieval was successful
	//HydrateCache(db, err, to, fieldName, value)
	hydrateCache(db, err, to, "Retrieve", GetStructType(to))
	// Type assert the result to *TemplateStore

	return to, err
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) ([]any, error) {
	logHandler.DatabaseLogger.Printf("[GET]<%v>{ALL} [%+v][%+v] [%v.db] caching: %t initialised: %t", GetStructType(to), GetStructType(to), options, db.Name, db.withCaching, db.cacheInitialised)

	if db.withCaching && db.cacheInitialised {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning from cache", GetStructType(to), GetStructType(to), db.Name, "GetAll")
		// return all cached entries of the appropriate type
		sliceValue := reflect.ValueOf(to).Elem()
		if sliceValue.Kind() != reflect.Slice {
			logHandler.CacheLogger.Printf("[GET]<%v>{ALL} - Expected slice when reading from cache, got %v", GetStructType(to), sliceValue.Kind())
			return nil, fmt.Errorf("GetAll expected slice pointer, got %v", sliceValue.Kind())
		}
		elemType := sliceValue.Type().Elem()

		// Clear the slice before populating
		sliceValue.Set(reflect.MakeSlice(sliceValue.Type(), 0, 0))

		for i, v := range inMemoryCache {
			cachedValue := reflect.ValueOf(v)
			logHandler.CacheLogger.Printf("[GET]<%v>{ALL} Checking cached entry [%v] of type [%v] against expected type [%v]", GetStructType(to), i, cachedValue.Type(), elemType)
			if cachedValue.Type().Elem() == elemType {
				logHandler.CacheLogger.Printf("[GET]<%v>{ALL} Adding cached entry [%v] to result set [%v][%v]", GetStructType(to), i, elemType, cachedValue.Type())
				sliceValue.Set(reflect.Append(sliceValue, cachedValue.Elem()))
			}
		}

		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning %d cached entries", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())

		// Convert the typed slice (e.g. []TemplateStore) into []any
		result := make([]any, sliceValue.Len())
		for i := 0; i < sliceValue.Len(); i++ {
			result[i] = sliceValue.Index(i).Interface()
		}
		return result, nil
	}

	// [GET] from database
	err := db.connection.All(to, options...)
	if err != nil {
		// On error, do not attempt to use or populate the cache
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{ERR} [%+v] [...%v.db] on %v - Error from DB: %v", GetStructType(to), GetStructType(to), db.Name, "GetAll", err)
		return nil, err
	}

	// Use reflection to iterate through the slice without assuming its concrete type
	sliceValue := reflect.ValueOf(to).Elem()
	if sliceValue.Kind() != reflect.Slice {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL} - Expected slice, got %v", GetStructType(to), sliceValue.Kind())
		return nil, fmt.Errorf("GetAll expected slice pointer, got %v", sliceValue.Kind())
	}

	// Optionally hydrate cache if caching is enabled and initialised
	if db.withCaching && db.cacheInitialised {
		for i := 0; i < sliceValue.Len(); i++ {
			item := sliceValue.Index(i)
			// Get the address of the item so we can pass it to hydrateCache
			itemPtr := item.Addr().Interface()
			hydrateCache(db, err, itemPtr, "GetAll", GetStructType(to))
		}
	} else {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{SKIP} [%+v] [...%v.db] on %v - Caching Disabled or Not Initialised", GetStructType(to), GetStructType(to), db.Name, "GetAll")
	}

	// Convert the typed slice (e.g. []TemplateStore) into []any
	result := make([]any, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		result[i] = sliceValue.Index(i).Interface()
	}

	return result, nil
}

// GetAllWhere retrieves all TemplateStore records that match the specified field and value.
//
// Parameters:
//   - field: The field to be used for filtering records.
//   - value: The value of the specified field to filter records.
//
// Returns:
//   - []TemplateStore: A slice of TemplateStore records that match the specified criteria.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func (db *DB) GetAllWhere(field Field, value any, to any) ([]any, error) {
	Domain := GetStructType(to)
	logHandler.EventLogger.Printf("SELECT %v WHERE (%v=%v)", Domain, field.String(), value)

	clock := timing.Start(Domain, actions.GETALL.GetCode(), fmt.Sprintf("%v=%v", field, value))

	//logHandler.DatabaseLogger.Printf("SELECT %v WHERE %v=%v", Domain, field, value)
	logHandler.EventLogger.Println("Check IsValidFieldInStruct")
	if err := IsValidFieldInStruct(field, to); err != nil {
		return nil, err
	}

	logHandler.EventLogger.Println("Check IsValidTypeForField")
	if err := IsValidTypeForField(field, value, to); err != nil {
		return nil, err
	}

	//err := activeDB.Retrieve(field, value, &recordList)
	var resultList []any
	recordList, err := db.GetAll(to)
	if err != nil {
		return nil, err
	}
	count := 0

	for _, record := range recordList {
		if reflect.ValueOf(record).FieldByName(field.String()).Interface() == value {
			count++
			resultList = append(resultList, record)
		}
	}

	clock.Stop(len(resultList))

	return resultList, nil
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("[DEL]<%v>{DELETE} Delete [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	err := db.connection.DeleteStruct(data)
	removeFromCache(db, data, "Delete", GetStructType(data))
	return err
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("[DRP]<%v>{DROP} Drop [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	err := db.connection.Drop(data)
	removeFromCache(db, data, "Drop", GetStructType(data))
	return err
}

func (db *DB) Update(data any) error {
	logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - Start", GetStructType(data), GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - Error", GetStructType(data), GetStructType(data), db.Name)
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - End", GetStructType(data), GetStructType(data), db.Name)
	err = db.connection.Update(data)
	hydrateCache(db, err, data, "Update,", GetStructType(data))

	return err
}

func (db *DB) Create(data any) error {
	logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - Start", GetStructType(data), GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - Error", GetStructType(data), GetStructType(data), db.Name)
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - End", GetStructType(data), GetStructType(data), db.Name)
	err = db.connection.Save(data)

	hydrateCache(db, err, data, "Create", GetStructType(data))

	return err
}

func (db *DB) Count(data any) (int, error) {
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} Count [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	if db.withCaching {
		logHandler.CacheLogger.Printf("[CNT]<%v>{SKIP} Count [%+v] [%v.db] - Caching Enabled", GetStructType(data), GetStructType(data), db.Name)
		return len(inMemoryCache), nil
	}
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CON]<%v>{CONNECTION POOL} Connection Pool [%v] [%v] [codec=%v]", GetStructType(data), key, value.databaseName, value.connection.Node.Codec().Name())
	}
	return db.connection.Count(data)
}

func (db *DB) CountWhere(fieldName Field, value any, to any) (int, error) {
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db]", GetStructType(to), fieldName, value, GetStructType(to), db.Name)
	if err := IsValidFieldInStruct(fieldName, to); err != nil {
		logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db] - Error", GetStructType(to), fieldName, value, GetStructType(to), db.Name)
		return 0, err
	}
	if db.withCaching {
		logHandler.CacheLogger.Printf("[CNT]<%v>{SKIP} CountWhere (%+v=%+v)[%+v] [%v.db] - Caching Enabled", GetStructType(to), fieldName, value, GetStructType(to), db.Name)
		// Range through inMemoryCache and count matching entries
		count := 0
		for _, v := range inMemoryCache {
			val := reflect.ValueOf(v).Elem().FieldByName(string(fieldName))
			if val.IsValid() && val.Interface() == value {
				count++
			}
		}
		return count, nil
	}
	query := db.connection.Select(q.Eq(fieldName.String(), value))
	count, err := query.Count(to)
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db] - Result: %d", GetStructType(to), fieldName, value, GetStructType(to), db.Name, count)
	return count, err
}
