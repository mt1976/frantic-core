package database

import (
	"reflect"

	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"

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

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("[GET]<%v>{ALL} [%+v][%+v] [%v.db] caching: %t initialised: %t", GetStructType(to), GetStructType(to), options, db.Name, db.withCaching, db.cacheInitialised)

	if db.withCaching && db.cacheInitialised {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning from cache", GetStructType(to), GetStructType(to), db.Name, "GetAll")
		// return all cached entries of the appropriate type
		sliceValue := reflect.ValueOf(to).Elem()
		elemType := sliceValue.Type().Elem()
		//elemPtrType := reflect.PointerTo(elemType)

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

		//	godump.DumpJSON(sliceValue)

		// Set the output parameter
		to = sliceValue.Interface()

		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning %d cached entries", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())
		return nil
	}

	// [GET] from database
	err := db.connection.All(to, options...)
	// Store in cache if caching is enabled and retrieval was successful
	if !db.withCaching || err != nil || !db.cacheInitialised {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL}{SKIP} [%+v] [...%v.db] on %v - Caching Disabled or Not Initialised", GetStructType(to), GetStructType(to), db.Name, "GetAll")
		return err
	}

	// Use reflection to iterate through the slice without casting
	sliceValue := reflect.ValueOf(to).Elem()

	// Check if it's actually a slice
	if sliceValue.Kind() != reflect.Slice {
		logHandler.CacheLogger.Printf("[GET]<%v>{ALL} - Expected slice, got %v", GetStructType(to), sliceValue.Kind())
		return err
	}

	// Iterate through each element in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		// Get the address of the item so we can pass it to hydrateCache
		itemPtr := item.Addr().Interface()
		hydrateCache(db, err, itemPtr, "GetAll", GetStructType(to))
	}

	return err
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
