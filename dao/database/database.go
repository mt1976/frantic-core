package database

import (
	"reflect"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/logHandler"
)

type Field dao.Field

type DB struct {
	connection       *storm.DB
	Name             string
	databaseName     string
	initialised      bool
	withCaching      bool
	withCacheKey     dao.Field
	verbose          bool
	timeout          int
	poolSize         int
	withEncryption   bool
	indices          []Field
	cacheInitialised bool
}

func (db *DB) Retrieve(fieldName Field, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)

	if db.withCaching {
		cacheKeyValue := value.(string)
		if cachedValue, found := inMemoryCache[cacheKeyValue]; found {
			reflect.ValueOf(to).Elem().Set(reflect.ValueOf(cachedValue).Elem())
			logHandler.CacheLogger.Printf("[CACHE] HIT (%v=%v) [%+v] [...%v.db] on %v - Skipping DB Access", db.withCacheKey, cacheKeyValue, dao.GetStructType(to), db.Name, "Retrieve")
			return nil
		}
		logHandler.CacheLogger.Printf("[CACHE] MISS (%v=%v) [%+v] [...%v.db] on %v - Accessing DB", db.withCacheKey, cacheKeyValue, dao.GetStructType(to), db.Name, "Retrieve")
	}

	logHandler.DatabaseLogger.Printf("Retrieve (%+v=%+v)[%+v] [%v.db] - From Database", fieldName, value, dao.GetStructType(to), db.Name)

	// Retrieve from database
	err := db.connection.One(string(fieldName), value, to)

	// Store in cache if caching is enabled and retrieval was successful
	//HydrateCache(db, err, to, fieldName, value)
	hydrateCache(db, err, to, "Retrieve")
	return err
}

func hydrateCache(db *DB, err error, to any, action string) {
	if db.withCaching && err == nil {
		cacheKeyValue := db.getCacheKeyValue(to)
		inMemoryCache[cacheKeyValue] = to
		logHandler.CacheLogger.Printf("[CACHE] ADD (%v=%v) [%+v] [...%v.db] on %v initialised: %t", db.withCacheKey, cacheKeyValue, dao.GetStructType(to), db.Name, action, db.cacheInitialised)
		logHandler.DatabaseLogger.Printf("[DB] ADD (%v=%v) [%+v] [...%v.db] on %v initialised: %t", db.withCacheKey, cacheKeyValue, dao.GetStructType(to), db.Name, action, db.cacheInitialised)
	}
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.EventLogger.Printf("GetAll [%+v][%+v] [%v.db] caching: %t initialised: %t", dao.GetStructType(to), options, db.Name, db.withCaching, db.cacheInitialised)

	if db.withCaching && db.cacheInitialised {
		logHandler.CacheLogger.Printf("[CACHE] HIT GetAll [%+v] [...%v.db] on %v - Returning from cache", dao.GetStructType(to), db.Name, "GetAll")
		// return all cached entries of the appropriate type
		sliceValue := reflect.ValueOf(to).Elem()
		elemType := sliceValue.Type().Elem()
		//elemPtrType := reflect.PointerTo(elemType)

		// Clear the slice before populating
		sliceValue.Set(reflect.MakeSlice(sliceValue.Type(), 0, 0))

		for i, v := range inMemoryCache {
			cachedValue := reflect.ValueOf(v)
			logHandler.CacheLogger.Printf("[CACHE] GetAll Checking cached entry [%v] of type [%v] against expected type [%v]", i, cachedValue.Type(), elemType)
			if cachedValue.Type().Elem() == elemType {
				logHandler.CacheLogger.Printf("[CACHE] GetAll Adding cached entry [%v] to result set [%v][%v]", i, elemType, cachedValue.Type())
				sliceValue.Set(reflect.Append(sliceValue, cachedValue.Elem()))
			}
		}

		//	godump.DumpJSON(sliceValue)

		// Set the output parameter
		to = sliceValue.Interface()

		logHandler.CacheLogger.Printf("[CACHE] HIT GetAll [%+v] [...%v.db] on %v - Returning %d cached entries", dao.GetStructType(to), db.Name, "GetAll", sliceValue.Len())
		return nil
	}

	// Retrieve from database
	err := db.connection.All(to, options...)
	// Store in cache if caching is enabled and retrieval was successful
	if !db.withCaching || err != nil || !db.cacheInitialised {
		logHandler.CacheLogger.Printf("[CACHE] SKIP GetAll [%+v] [...%v.db] on %v - Caching Disabled or Not Initialised", dao.GetStructType(to), db.Name, "GetAll")
		return err
	}

	// Use reflection to iterate through the slice without casting
	sliceValue := reflect.ValueOf(to).Elem()

	// Check if it's actually a slice
	if sliceValue.Kind() != reflect.Slice {
		logHandler.CacheLogger.Printf("GetAll - Expected slice, got %v", sliceValue.Kind())
		return err
	}

	// Iterate through each element in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		// Get the address of the item so we can pass it to hydrateCache
		itemPtr := item.Addr().Interface()
		hydrateCache(db, err, itemPtr, "GetAll")
	}

	return err
}

func (db *DB) PreLoadCache(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("[DB] PreLoadCache [%+v][%+v] [%v.db]", dao.GetStructType(to), options, db.Name)
	logHandler.CacheLogger.Printf("[CACHE] PRELOAD [%+v] [...%v.db] on %v", dao.GetStructType(to), db.Name, "PreLoadCache")

	// Retrieve from database
	// err := db.connection.All(to, options...)
	// // Store in cache if caching is enabled and retrieval was successful
	// if !db.withCaching || err != nil {
	// 	return err
	// }
	db.cacheInitialised = false

	err := db.GetAll(to, options...)

	logHandler.CacheLogger.Printf("[CACHE] FETCH COMPLETE [%+v] [...%v.db] on %v", dao.GetStructType(to), db.Name, "PreLoadCache")
	// Use reflection to iterate through the slice without casting
	sliceValue := reflect.ValueOf(to).Elem()

	// Check if it's actually a slice
	if sliceValue.Kind() != reflect.Slice {
		logHandler.CacheLogger.Printf("GetAll - Expected slice, got %v", sliceValue.Kind())
		return err
	}

	// Iterate through each element in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		// Get the address of the item so we can pass it to hydrateCache
		itemPtr := item.Addr().Interface()
		hydrateCache(db, err, itemPtr, "PreLoadCache")
	}

	db.cacheInitialised = true

	logHandler.CacheLogger.Printf("[CACHE] PRELOAD COMPLETE [%+v] [...%v.db] on %v - Cached %d entries", dao.GetStructType(to), db.Name, "PreLoadCache", sliceValue.Len())

	return err
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	err := db.connection.DeleteStruct(data)
	removeFromCache(db, data, "Delete")
	return err
}

func removeFromCache(db *DB, data any, action string) {
	if db.withCaching {
		cacheKeyValue := db.getCacheKeyValue(data)
		delete(inMemoryCache, cacheKeyValue)
		logHandler.CacheLogger.Printf("[CACHE] DELETE (%v=%v) [%+v] [...%v.db] on %v", db.withCacheKey, cacheKeyValue, dao.GetStructType(data), db.Name, action)
	}
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	err := db.connection.Drop(data)
	removeFromCache(db, data, "Drop")
	return err
}

func (db *DB) Update(data any) error {
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - Start", dao.GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - Error", dao.GetStructType(data), db.Name)
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - End", dao.GetStructType(data), db.Name)
	err = db.connection.Update(data)
	hydrateCache(db, err, data, "Update")

	return err
}

func (db *DB) Create(data any) error {
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - Start", dao.GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - Error", dao.GetStructType(data), db.Name)
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - End", dao.GetStructType(data), db.Name)
	err = db.connection.Save(data)

	hydrateCache(db, err, data, "Create")

	return err
}

func (db *DB) Count(data any) (int, error) {
	logHandler.DatabaseLogger.Printf("Count [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	if db.withCaching {
		logHandler.DatabaseLogger.Printf("[CACHE] SKIP Count [%+v] [%v.db] - Caching Enabled", dao.GetStructType(data), db.Name)
		return len(inMemoryCache), nil
	}
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	return db.connection.Count(data)
}

func (db *DB) CountWhere(fieldName dao.Field, value any, to any) (int, error) {
	logHandler.DatabaseLogger.Printf("CountWhere (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)
	if err := dao.IsValidFieldInStruct(fieldName, to); err != nil {
		logHandler.DatabaseLogger.Printf("CountWhere (%+v=%+v)[%+v] [%v.db] - Error", fieldName, value, dao.GetStructType(to), db.Name)
		return 0, err
	}
	if db.withCaching {
		logHandler.CacheLogger.Printf("[CACHE] SKIP CountWhere (%+v=%+v)[%+v] [%v.db] - Caching Enabled", fieldName, value, dao.GetStructType(to), db.Name)
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
	return count, err
}

// // toString converts any value to a string for cache key generation
// func toString(value any) string {
// 	if value == nil {
// 		return "nil"
// 	}
// 	return dao.GetStructType(value)
// }

func (db *DB) getCacheKeyValue(to any) string {
	// For simplicity, using a static cache key in this example.
	// In a real implementation, you would generate a unique key based on the query parameters.

	keyField := db.withCacheKey
	logHandler.TraceLogger.Printf("getCacheKey [%v.db] - Key Field: '%v' to:%+v", db.Name, keyField, reflect.TypeOf(to))

	reflectValue := reflect.ValueOf(to).Elem().FieldByName(keyField.String())
	logHandler.TraceLogger.Printf("getCacheKey [%v.db] - Reflect Value: '%v'", db.Name, reflectValue)

	return reflectValue.String()
}
