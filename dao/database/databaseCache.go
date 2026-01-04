package database

import (
	"reflect"

	"github.com/asdine/storm/v3/index"
	"github.com/mt1976/frantic-core/logHandler"
)

// hydrateCache adds the retrieved data to the in-memory cache if caching is enabled
// and the retrieval was successful
func hydrateCache(db *DB, err error, to any, action string, structType string) {
	if db.withCaching && err == nil {
		cacheKeyValue := db.getCacheKeyValue(to)
		inMemoryCache[cacheKeyValue] = to
		logHandler.CacheLogger.Printf("[HYD]<%v>{Add} (%v=%v) [%+v] [...%v.db] on %v initialised: %t", structType, db.withCacheKey, cacheKeyValue, structType, db.Name, action, db.cacheInitialised)
	}
}

// removeFromCache removes the specified data from the in-memory cache if caching is enabled
// This is typically called after a delete operation
func removeFromCache(db *DB, data any, action string, structType string) {
	if db.withCaching {
		cacheKeyValue := db.getCacheKeyValue(data)
		delete(inMemoryCache, cacheKeyValue)
		logHandler.CacheLogger.Printf("[CAC]<%v>{REMOVE} Remove (%v=%v) from Cache [%+v] [...%v.db] on %v %v", structType, db.withCacheKey, cacheKeyValue, structType, db.Name, action, GetStructType(data))
	}
}

func (db *DB) getCacheKeyValue(to any) string {
	// For simplicity, using a static cache key in this example.
	// In a real implementation, you would generate a unique key based on the query parameters.

	keyField := db.withCacheKey
	logHandler.TraceLogger.Printf("<%v>{getCacheKey} [%v.db] - Key Field: '%v' to:%+v", GetStructType(to), db.Name, keyField, reflect.TypeOf(to))

	reflectValue := reflect.ValueOf(to).Elem().FieldByName(keyField.String())
	logHandler.TraceLogger.Printf("<%v>{getCacheKey} [%v.db] - Reflect Value: '%v'", GetStructType(to), db.Name, reflectValue)

	return reflectValue.String()
}

func (db *DB) PreLoadCache(to any, options ...func(*index.Options)) error {
	logHandler.CacheLogger.Printf("[HYD]<%v>{LOAD} Hydrate Cache  [%+v] [...%v.db] on %v", GetStructType(to), GetStructType(to), db.Name, "PreLoadCache")

	// [GET] from database
	// err := db.connection.All(to, options...)
	// // Store in cache if caching is enabled and retrieval was successful
	if !db.withCaching {
		logHandler.WarningLogger.Printf("[HYD]<%v>{LOAD} CACHING NOT ENABLED [...%v.db] on %v", GetStructType(to), db.Name, "PreLoadCache")
		return nil
	}
	db.cacheInitialised = false

	res, err := db.GetAll(to, options...)

	logHandler.CacheLogger.Printf("[HYD]<%v>{LOAD}{GET} COMPLETE [%+v] [...%v.db] on %v", GetStructType(to), GetStructType(to), db.Name, "PreLoadCache")
	if err != nil {
		return err
	}

	// res is []any containing concrete values (e.g. []TemplateStore)
	for _, item := range res {
		val := reflect.ValueOf(item)
		var itemPtr any
		if val.Kind() == reflect.Ptr {
			itemPtr = item
		} else {
			ptr := reflect.New(val.Type())
			ptr.Elem().Set(val)
			itemPtr = ptr.Interface()
		}
		hydrateCache(db, err, itemPtr, "PreLoadCache", GetStructType(to))
	}

	db.cacheInitialised = true

	logHandler.CacheLogger.Printf("[HYD]<%v>{LOAD} COMPLETE [%+v] [...%v.db] on %v - Cached %d entries - Cache: %t Initialised: %t %v", GetStructType(to), GetStructType(to), db.Name, "PreLoadCache", len(res), db.withCaching, db.cacheInitialised, db.withCacheKey)

	return nil
}
