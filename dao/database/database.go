package database

import (
	"fmt"
	"reflect"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"
	"github.com/mt1976/frantic-core/timing"

	"github.com/mt1976/frantic-core/logHandler"
)

// Retrieve retrieves a single record from the database based on the specified fields.Field and value.
//
// Parameters:
//   - fields.Field: The fields.Field to be used for filtering the record.
//   - value: The value of the specified fields.Field to filter the record.
//   - to: A pointer to the struct where the retrieved record will be stored.
//
// Returns:
//   - any: The retrieved record.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
//
// DEPRECATED: Use Get instead.
func (db *DB) Retrieve(field entities.Field, value, to any) (any, error) {
	logHandler.WarningLogger.Printf("Retrieve is DEPRECATED, use Get instead")
	panic("Retrieve is DEPRECATED, use Get instead")
	//return db.get(field, value, to)
}

// Get retrieves a single record from the database based on the specified fields.Field and value.
//
// Parameters:
//   - fields.Field: The fields.Field to be used for filtering the record.
//   - value: The value of the specified fields.Field to filter the record.
//   - to: A pointer to the struct where the retrieved record will be stored.
//
// Returns:
//   - any: The retrieved record.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func (db *DB) Get(field entities.Field, value, to any) (any, error) {
	//logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db]", GetStructType(to), fields.Field, value, GetStructType(to), db.Name)
	return db.get(field, value, to)
}

// get is the internal implementation for retrieving a single record from the database.
//
// Parameters:
//   - fields.Field: The fields.Field to be used for filtering the record.
//   - value: The value of the specified fields.Field to filter the record.
//   - to: A pointer to the struct where the retrieved record will be stored.
//
// Returns:
//   - any: The retrieved record.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func (db *DB) get(field entities.Field, value, to any) (any, error) {
	logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db] {%+v}", GetStructType(to), field.String(), value, GetStructType(to), db.Name, db)

	// if db.isCaching(to) {

	// 	// TODO: What this should do is look through the cache for the appropriate entry, and check if it can find one with the matching fields.Field/value
	// 	// Not just using the cache key

	// 	for key, val := range inMemoryCache[GetStructType(to)] {
	// 		logHandler.CacheLogger.Printf("[GET]<%v>{CACHE SCAN} Checking cache entry [%v] of type [%v] against expected type [%v]", GetStructType(to), key, reflect.TypeOf(val), reflect.TypeOf(to))
	// 		// Val is a pointer to struct
	// 		valValue := reflect.ValueOf(val)
	// 		if valValue.Type() == reflect.TypeOf(to) {
	// 			fieldValue := valValue.Elem().FieldByName(field.String())
	// 			if fields.FieldValue.IsValid() {
	// 				logHandler.CacheLogger.Printf("[GET]<%v>{CACHE SCAN} Found matching type for cache entry [%v]. Checking fields.Field [%v] value [%v] against expected value [%v]", GetStructType(to), key, fields.Field.String(), fields.FieldValue.Interface(), value)
	// 				if fields.FieldValue.Interface() == value {
	// 					reflect.ValueOf(to).Elem().Set(valValue.Elem())
	// 					logHandler.CacheLogger.Printf("[GET]<%v>{HIT} (%v=%v) [%+v] [...%v.db] on %v - {CACHE HIT} DB Access", GetStructType(to), fields.Field, value, GetStructType(to), db.Name, "Retrieve")
	// 					return val, nil
	// 				}
	// 			} else {
	// 				logHandler.CacheLogger.Printf("[GET]<%v>{CACHE SCAN} fields.Field [%v] not found in cached entry [%v]", GetStructType(to), fields.Field.String(), key)
	// 			}
	// 		} else {
	// 			logHandler.CacheLogger.Printf("[GET]<%v>{CACHE SCAN} Type mismatch for cache entry [%v]: expected [%v], got [%v]", GetStructType(to), key, reflect.TypeOf(to), valValue.Type())
	// 		}

	// 	}

	// cacheKeyValue := fmt.Sprintf("%v", value)
	// if cachedValue, found := inMemoryCache[cacheKeyValue]; found {
	// 	reflect.ValueOf(to).Elem().Set(reflect.ValueOf(cachedValue).Elem())
	// 	logHandler.CacheLogger.Printf("[GET]<%v>{HIT} (%v=%v) [%+v] [...%v.db] on %v - {SKIP} DB Access", GetStructType(to), db.withCacheKey, cacheKeyValue, GetStructType(to), db.Name, "Retrieve")
	// 	return cachedValue, nil
	// }
	// logHandler.CacheLogger.Printf("[GET]<%v>{MISS} (%v=%v) [%+v] [...%v.db] on %v - Accessing DB", GetStructType(to), db.withCacheKey, cacheKeyValue, GetStructType(to), db.Name, "Retrieve")
	// }

	logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db] - From Database", GetStructType(to), field.String(), value, GetStructType(to), db.Name)

	// [GET] from database
	err := db.connection.One(field.String(), value, to)

	logHandler.DatabaseLogger.Printf("[GET]<%v> (%+v=%+v)[%+v] [%v.db] - Completed", GetStructType(to), field.String(), value, GetStructType(to), db.Name)
	if err != nil {
		// On error, do not attempt to use or populate the cache
		logHandler.CacheLogger.Printf("[GET]<%v>{ERR} (%v=%v) [%+v] [...%v.db] on %v - Error from DB: %v", GetStructType(to), field.String(), value, GetStructType(to), db.Name, "Retrieve", err)
		return nil, err
	}
	// Store in cache if caching is enabled and retrieval was successful
	//HydrateCache(db, err, to, fields.FieldName, value)
	//db.hydrateCache(err, to, "Retrieve", GetStructType(to))
	// Type assert the result to *TemplateStore

	return to, err
}

// GetAll retrieves all records of the specified type from the database.
//
// Parameters:
//   - to: A pointer to a slice where the retrieved records will be stored.
//
// Returns:
//   - []any: A slice of all retrieved records.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func (db *DB) GetAll(to any, options ...func(*index.Options)) ([]any, error) {
	logHandler.InfoLogger.Printf("[GET]<%v>{ALL} [%+v][%+v] [%v.db]", GetStructType(to), GetStructType(to), options, db.Name)

	// if db.isCaching(to) && db.isCacheInitialised(to) {
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning from cache", GetStructType(to), GetStructType(to), db.Name, "GetAll")
	// 	// return all cached entries of the appropriate type
	// 	sliceValue := reflect.ValueOf(to).Elem()
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - Preparing to read into slice of type %v", GetStructType(to), sliceValue.Type())
	// 	if sliceValue.Kind() != reflect.Slice {
	// 		logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - Expected slice when reading from cache, got %v", GetStructType(to), sliceValue.Kind())
	// 		return nil, commonErrors.ErrInvalidTypeWrapper("GetAll", string(GetStructType(to)), sliceValue.Kind().String())
	// 	}
	// 	elemType := sliceValue.Type().Elem()
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - Slice element type is %v", GetStructType(to), elemType)
	// 	// Clear the slice before populating
	// 	sliceValue.Set(reflect.MakeSlice(sliceValue.Type(), 0, 0))
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - Iterating through in-memory cache with %d entries", GetStructType(to), len(inMemoryCache))

	// 	// Get the entries for this table from the in-memory cache
	// 	entries, found := inMemoryCache[GetStructType(to)]
	// 	if !found {
	// 		logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - No cached entries found for type %v", GetStructType(to), GetStructType(to))
	// 		return nil, nil
	// 	}

	// 	// Iterate through the cached entries and add matching types to the result slice
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} - Found %d cached entries for type %v", GetStructType(to), len(entries), GetStructType(to))
	// 	for _, v := range entries {
	// 		cachedValue := reflect.ValueOf(v)
	// 		//	logHandler.EventLogger.Printf("[GET]<%v>{ALL} Checking cached entry [%v] of type [%v] against expected type [%v]", GetStructType(to), i, cachedValue.Type(), elemType)
	// 		if cachedValue.Type().Elem() == elemType {
	// 			//		logHandler.EventLogger.Printf("[GET]<%v>{ALL} Adding cached entry [%v] to result set [%v][%v]", GetStructType(to), i, elemType, cachedValue.Type())
	// 			sliceValue.Set(reflect.Append(sliceValue, cachedValue.Elem()))
	// 		}
	// 	}

	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning %d cached entries", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())

	// 	// Convert the typed slice (e.g. []TemplateStore) into []any
	// 	result := make([]any, sliceValue.Len())
	// 	for i := 0; i < sliceValue.Len(); i++ {
	// 		result[i] = sliceValue.Index(i).Interface()
	// 	}
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL}{HIT} [%+v] [...%v.db] on %v - Returning %d cached entries", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())
	// 	return result, nil
	// }

	logHandler.InfoLogger.Printf("DATABASE[GET]<%v>{ALL} [%+v] [%v.db] - From Database %+v", GetStructType(to), GetStructType(to), db.Name, options)
	// [GET] from database
	err := db.connection.All(to, options...)
	if err != nil {
		// On error, do not attempt to use or populate the cache
		logHandler.InfoLogger.Printf("DATABASE[GET]<%v>{ALL}{ERR} [%+v] [...%v.db] on %v - Error from DB: %v", GetStructType(to), GetStructType(to), db.Name, "GetAll", err)
		return nil, err
	}

	// Completed DB retrieval, wait 1 seconds
	//logHandler.TraceLogger.Printf("[GET]<%v>{ALL} [%+v] [%v.db] - Pausing before processing", GetStructType(to), GetStructType(to), db.Name)
	//time.Sleep(1 * time.Second)

	logHandler.InfoLogger.Printf("DATABASE[GET]<%v>{ALL} [%+v] [%v.db] - Completed", GetStructType(to), GetStructType(to), db.Name)

	// Use reflection to iterate through the slice without assuming its concrete type
	sliceValue := reflect.ValueOf(to).Elem()
	if sliceValue.Kind() != reflect.Slice {
		logHandler.InfoLogger.Printf("DATABASE[GET]<%v>{ALL} - Expected slice, got %v", GetStructType(to), sliceValue.Kind())
		return nil, commonErrors.ErrInvalidTypeWrapper("GetAll", string(GetStructType(to)), sliceValue.Kind().String())
	}

	logHandler.InfoLogger.Printf("DATABASE[GET]<%v>{ALL} [%+v] [...%v.db] on %v - Retrieved %d entries from DB", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())

	// Optionally hydrate cache if caching is enabled and initialised
	// if db.isCaching(to) {
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL}{POPULATE} [%+v] [...%v.db] on %v - Populating Cache", GetStructType(to), GetStructType(to), db.Name, "GetAll")
	// 	for i := 0; i < sliceValue.Len(); i++ {
	// 		item := sliceValue.Index(i)
	// 		// Get the address of the item so we can pass it to hydrateCache
	// 		itemPtr := item.Addr().Interface()
	// 		db.hydrateCache(err, itemPtr, "GetAll", GetStructType(to))
	// 	}
	// } else {
	// 	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL}{SKIP} [%+v] [...%v.db] on %v - Caching Disabled or Not Initialised", GetStructType(to), GetStructType(to), db.Name, "GetAll")
	// }

	// Convert the typed slice (e.g. []TemplateStore) into []any
	result := make([]any, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		result[i] = sliceValue.Index(i).Interface()
	}

	logHandler.InfoLogger.Printf("CACHE[GET]<%v>{ALL} [%+v] [...%v.db] on %v - Returning %d entries", GetStructType(to), GetStructType(to), db.Name, "GetAll", sliceValue.Len())
	return result, nil
}

// GetAllWhere retrieves all TemplateStore records that match the specified fields.Field and value.
//
// Parameters:
//   - fields.Field: The fields.Field to be used for filtering records.
//   - value: The value of the specified fields.Field to filter records.
//
// Returns:
//   - []TemplateStore: A slice of TemplateStore records that match the specified criteria.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func (db *DB) GetAllWhere(field entities.Field, value, to any) ([]any, error) {
	tableName := GetStructType(to)
	logHandler.DatabaseLogger.Printf("SELECT %v WHERE (%v=%v)", tableName, field.String(), value)

	clock := timing.Start(tableName, "GetAll", fmt.Sprintf("%v=%v", field.String(), value))

	//logHandler.DatabaseLogger.Printf("SELECT %v WHERE %v=%v", Domain, fields.Field, value)
	if err := IsValidFieldInStruct(field, to); err != nil {
		logHandler.ErrorLogger.Printf("Field validation error for fields.Field '%v': %v", field.String(), err)
		clock.Stop(0)
		return nil, err
	}

	if err := IsValidTypeForField(field, value, to); err != nil {
		logHandler.ErrorLogger.Printf("Type validation error for fields.Field '%v': %v", field.String(), err)
		clock.Stop(0)
		return nil, err
	}

	// If caching is enabled and initialised, use the existing GetAll + in-memory filter,
	// which operates on the in-memory cache and avoids hitting the database.
	// if db.isCaching(to) && db.isCacheInitialised(to) {
	// 	var resultList []any
	// 	// Get all records from cache
	// 	allRecords := db.getCachedEntries(tableName)
	// 	// Filter records based on the specified fields.Field and value
	// 	for _, record := range allRecords {
	// 		reflectValue := reflect.ValueOf(record)
	// 		if reflectValue.Kind() == reflect.Ptr {
	// 			reflectValue = reflectValue.Elem()
	// 		}
	// 		fieldValue := reflectValue.FieldByName(field.String())
	// 		if fields.FieldValue.IsValid() && fields.FieldValue.Interface() == value {
	// 			resultList = append(resultList, record)
	// 		}
	// 	}
	// 	logHandler.TraceLogger.Printf("[GET]<%v>{WHERE}{HIT} (%v=%v) [%+v] [...%v.db] on %v - Returning %d cached entries", tableName, fields.Field.String(), value, tableName, db.Name, "GetAllWhere", len(resultList))
	// 	// Log the type of the records in the result list
	// 	//for i, rec := range resultList {
	// 	//	logHandler.EventLogger.Printf("[GET]<%v>{WHERE}{HIT} Result[%d] Type: %v", tableName, i, reflect.TypeOf(rec))
	// 	//}
	// 	clock.Stop(len(resultList))
	// 	return resultList, nil
	// }

	// Otherwise, use Storm's indexed query to retrieve matching records directly.
	query := db.connection.Select(q.Eq(field.String(), value))
	err := query.Find(to)
	if err != nil {
		if err == storm.ErrNotFound {
			clock.Stop(0)
			return []any{}, nil
		}
		logHandler.ErrorLogger.Printf("Error querying %v where %v=%v: %v", tableName, field.String(), value, err)
		clock.Stop(0)
		return nil, err
	}

	sliceValue := reflect.ValueOf(to).Elem()
	if sliceValue.Kind() != reflect.Slice {
		logHandler.DatabaseLogger.Printf("[GET]<%v>{WHERE} - Expected slice pointer, got %v", tableName, sliceValue.Kind())
		clock.Stop(0)
		return nil, commonErrors.ErrInvalidTypeWrapper("GetAllWhere", tableName, sliceValue.Kind().String())
	}

	resultList := make([]any, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		resultList[i] = sliceValue.Index(i).Interface()
	}

	// hydrateerr := db.hydrateCacheBulk(resultList)
	// if hydrateerr != nil {
	// 	logHandler.ErrorLogger.Printf("[CCH]<%v>{AddBulk} Error hydrating cache: %v", tableName, hydrateerr)
	// }

	clock.Stop(len(resultList))
	return resultList, nil
}

// Delete removes the specified record from the database.
//
// Parameters:
//   - data: A pointer to the struct representing the record to be deleted.
//
// Returns:
//   - error: An error object if any issues occur during the deletion process; otherwise, nil.
func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("[DEL]<%v>{DELETE} Delete [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	err := db.connection.DeleteStruct(data)
	//removeFromCache(db, data, "Delete", GetStructType(data))
	return err
}

// Drop removes the entire bucket or collection associated with the specified struct from the database.
//
// Parameters:
//   - data: A pointer to the struct representing the type whose bucket or collection is to be dropped.
//
// Returns:
//   - error: An error object if any issues occur during the drop process; otherwise, nil.
func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("[DRP]<%v>{DROP} Drop [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	err := db.connection.Drop(data)
	//removeFromCache(db, data, "Drop", GetStructType(data))
	return err
}

// Update modifies an existing record in the database.
//
// Parameters:
//   - data: A pointer to the struct representing the record to be updated.
//
// Returns:
//   - error: An error object if any issues occur during the update process; otherwise, nil.
func (db *DB) Update(data any) error {
	logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - Start", GetStructType(data), GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - Error", GetStructType(data), GetStructType(data), db.Name)
		return commonErrors.ErrWrapper(err)
	}
	logHandler.DatabaseLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - End", GetStructType(data), GetStructType(data), db.Name)
	// if db.isCaching(data) {
	// 	go func() {
	// 		err = db.connection.Update(data)
	// 		if err != nil {
	// 			logHandler.ErrorLogger.Printf("[UPD]<%v>{UPDATE} Update [%+v] [%v.db] - Error: %v", GetStructType(data), GetStructType(data), db.Name, err)
	// 			return
	// 		}
	// 	}()
	// } else {
	err = db.connection.Update(data)
	// }
	// db.hydrateCache(err, data, "Update,", GetStructType(data))

	return err
}

// Create adds a new record to the database.
//
// Parameters:
//   - data: A pointer to the struct representing the record to be created.
//
// Returns:
//   - error: An error object if any issues occur during the creation process; otherwise, nil.
func (db *DB) Create(data any) error {
	logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - Start", GetStructType(data), GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - Error", GetStructType(data), GetStructType(data), db.Name)
		return commonErrors.ErrCreateWrapper(err)
	}
	logHandler.DatabaseLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - End", GetStructType(data), GetStructType(data), db.Name)

	err = db.connection.Save(data)
	//db.hydrateCache(err, data, "Create", GetStructType(data))
	if err != nil {
		logHandler.ErrorLogger.Printf("[NEW]<%v>{CREATE} Create [%+v] [%v.db] - Error: %v", GetStructType(data), GetStructType(data), db.Name, err)
	}
	return err
}

// Count returns the total number of records of the specified type in the database.
//
// Parameters:
//   - data: A pointer to the struct representing the type whose records are to be counted.
//
// Returns:
//   - int: The total number of records.
//   - error: An error object if any issues occur during the counting process; otherwise, nil.
func (db *DB) Count(data any) (int, error) {
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} Count [%+v] [%v.db]", GetStructType(data), GetStructType(data), db.Name)
	// if db.isCaching(data) {
	// 	logHandler.CacheLogger.Printf("[CNT]<%v>{SKIP} Count [%+v] [%v.db] - Caching Enabled", GetStructType(data), GetStructType(data), db.Name)
	// 	return len(inMemoryCache[GetStructType(data)]), nil
	// }
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CON]<%v>{CONNECTION POOL} Connection Pool [%v] [%v] [codec=%v]", GetStructType(data), key, value.databaseName, value.connection.Node.Codec().Name())
	}
	return db.connection.Count(data)
}

// CountWhere returns the number of records that match the specified fields.Field and value.
//
// Parameters:
//   - fields.FieldName: The fields.Field to be used for filtering records.
//   - value: The value of the specified fields.Field to filter records.
//   - to: A pointer to the struct representing the type whose records are to be counted.
//
// Returns:
//   - int: The number of records that match the specified criteria.
//   - error: An error object if any issues occur during the counting process; otherwise, nil.
func (db *DB) CountWhere(field entities.Field, value any, to any) (int, error) {
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db]", GetStructType(to), field.String(), value, GetStructType(to), db.Name)
	if err := IsValidFieldInStruct(field, to); err != nil {
		logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db] - Error", GetStructType(to), field.String(), value, GetStructType(to), db.Name)
		return 0, err
	}
	// if db.isCaching(to) {
	// 	logHandler.CacheLogger.Printf("[CNT]<%v>{SKIP} CountWhere (%+v=%+v)[%+v] [%v.db] - Caching Enabled", GetStructType(to), field.String(), value, GetStructType(to), db.Name)
	// 	// Range through inMemoryCache and count matching entries
	// 	count := 0
	// 	for _, v := range inMemoryCache[GetStructType(to)] {
	// 		val := reflect.ValueOf(v).Elem().FieldByName(string(fieldName))
	// 		if val.IsValid() && val.Interface() == value {
	// 			count++
	// 		}
	// 	}
	// 	return count, nil
	// }
	query := db.connection.Select(q.Eq(field.String(), value))
	count, err := query.Count(to)
	logHandler.DatabaseLogger.Printf("[CNT]<%v>{COUNT} CountWhere (%+v=%+v)[%+v] [%v.db] - Result: %d", GetStructType(to), field.String(), value, GetStructType(to), db.Name, count)
	return count, err
}
