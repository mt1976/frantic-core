package database

import (
	"reflect"
	"time"

	"github.com/asdine/storm/v3/index"
	"github.com/mt1976/frantic-core/logHandler"
)

// hydrateCache adds the retrieved data to the in-memory cache if caching is enabled
// and the retrieval was successful
func (db *DB) hydrateCache(err error, to any, action string, structType string) {
	logHandler.TraceLogger.Printf("[CCH]<%v>{Add} (%v=%v) [%+v] [...%v.db] on %v - Caching Enabled: %t", structType, db.withCacheKey, db.getCacheKeyValue(to), structType, db.Name, action, db.withCaching)
	if db.withCaching && err == nil {
		cacheKeyValue := db.getCacheKeyValue(to)
		table := GetStructType(to)
		// Ensure the table map exists
		if _, exists := inMemoryCache[table]; !exists {
			inMemoryCache[table] = make(cacheEntrys)
		}
		// Add to the cache
		cacheEntry := cacheEntrys(inMemoryCache[table])
		cacheEntry[cacheKeyValue] = to
		inMemoryCache[table] = cacheEntry

		logHandler.CacheLogger.Printf("[CCH]<%v>{Add} (%v=%v) [%+v] [...%v.db] on %v initialised: %t", structType, db.withCacheKey, cacheKeyValue, structType, db.Name, action, db.cacheInitialised)
	}
}

// removeFromCache removes the specified data from the in-memory cache if caching is enabled
// This is typically called after a delete operation
func removeFromCache(db *DB, data any, action string, structType string) {
	logHandler.TraceLogger.Printf("[CCH]<%v>{REMOVE} (%v=%v) from Cache [%+v] [...%v.db] on %v - Caching Enabled: %t", structType, db.withCacheKey, db.getCacheKeyValue(data), structType, db.Name, action, db.withCaching)
	if db.withCaching {
		// Remove the entry from the cache, first find the entries for the table, then delete by cache key
		table := GetStructType(data)
		cacheEntry := cacheEntrys(inMemoryCache[table])

		cacheKeyValue := db.getCacheKeyValue(data)
		delete(cacheEntry, cacheKeyValue)
		inMemoryCache[table] = cacheEntry
		logHandler.CacheLogger.Printf("[CCH]<%v>{REMOVE} Remove (%v=%v) from Cache [%+v] [...%v.db] on %v %v", structType, db.withCacheKey, cacheKeyValue, structType, db.Name, action, GetStructType(data))
	}
}

func (db *DB) getCacheKeyValue(to any) string {
	// For simplicity, using a static cache key in this example.
	// In a real implementation, you would generate a unique key based on the query parameters.

	keyField := db.withCacheKey
	logHandler.TraceLogger.Printf("[CCH]<%v>{getCacheKey} [%v.db] - Key Field: '%v' to:%+v", GetStructType(to), db.Name, keyField, reflect.TypeOf(to))

	reflectValue := reflect.ValueOf(to).Elem().FieldByName(keyField.String())
	logHandler.TraceLogger.Printf("[CCH]<%v>{getCacheKey} [%v.db] - Reflect Value: '%v'", GetStructType(to), db.Name, reflectValue)

	return reflectValue.String()
}

func (db *DB) PreLoadCache(to any, options ...func(*index.Options)) error {
	logHandler.CacheLogger.Printf("[CCH]<%v>{LOAD} Hydrate Cache  [%+v] [...%v.db] on %v - Caching: %t", GetStructType(to), GetStructType(to), db.Name, "PreLoadCache", db.withCaching)

	// [GET] from database
	// err := db.connection.All(to, options...)
	// // Store in cache if caching is enabled and retrieval was successful
	if !db.withCaching {
		logHandler.WarningLogger.Printf("[CCH]<%v>{LOAD} CACHING NOT ENABLED [...%v.db] on %v", GetStructType(to), db.Name, "PreLoadCache")
		return nil
	}
	db.cacheInitialised = false

	res, err := db.GetAll(to, options...)
	if err != nil {
		return err
	}

	// At this point GetAll has already hydrated the cache when caching is enabled.
	db.cacheInitialised = true

	logHandler.CacheLogger.Printf("[CCH]<%v>{LOAD} COMPLETE [%+v] [...%v.db] on %v - Cached %d entries - Cache: %t Initialised: %t %v", GetStructType(to), GetStructType(to), db.Name, "PreLoadCache", len(res), db.withCaching, db.cacheInitialised, db.withCacheKey)

	return nil
}

func (db *DB) isCachedEnabled(tableName string) bool {
	if db.cachedTables == nil {
		return false
	}
	_, exists := db.cachedTables[tableName]
	return exists
}

func enableCachingForTable(db *DB, table any) {
	tableName := GetStructType(table)
	logHandler.CacheLogger.Printf("Enabling caching for table [%v] in database [%v]", tableName, db.Name)
	if db.cachedTables == nil {
		db.cachedTables = make(map[string]bool)
	}
	db.cachedTables[tableName] = true
	logHandler.CacheLogger.Printf("Caching enabled for table [%v] in database [%v]", tableName, db.Name)
}

func (db *DB) CacheSpew() {
	logHandler.InfoLogger.Printf("Caching Status for Database [%v]:", db.Name)
	if len(db.cachedTables) == 0 {
		logHandler.InfoLogger.Printf("No tables are currently cached in database [%v]", db.Name)
		return
	}
	msg := ". Cached Tables in Database [" + db.Name + "]: "
	for tableName := range db.cachedTables {
		msg += tableName + " "
	}
	logHandler.InfoLogger.Println(msg)
	// Display A COUNT OF THE RECORDS IN THE CACHE
	logHandler.InfoLogger.Printf(". Cached Records Summary for Database [%v]", db.Name)
	for tableName := range db.cachedTables {
		inMemoryCacheEntry, exists := inMemoryCache[tableName]
		if !exists {
			logHandler.InfoLogger.Printf(". 	Table [%v] has 0 cached records in database [%v]", tableName, db.Name)
			continue
		}
		//
		logHandler.InfoLogger.Printf(". 	Table [%v] has %d cached records in database [%v]", tableName, len(inMemoryCacheEntry), db.Name)
	}
}

func (db *DB) Flush() error {
	if !db.withCaching && !db.cacheInitialised {
		logHandler.WarningLogger.Printf("Caching not enabled for [%v.db], skipping flush", db.Name)
		return nil
	}

	recordCount := 0
	tableCount := 0
	start := time.Now()
	// Scan the inMemoryCache and update only entries for this DB
	for key, cacheEntry := range inMemoryCache {
		if db.verbose {
			logHandler.CacheLogger.Printf("{FLUSH} <%v> Scanning cache entry [%v] of type [%v]", db.Name, key, reflect.TypeOf(cacheEntry))
		}
		// Here we would need to check if the cache entry belongs to this DB
		// For simplicity, assuming all entries belong to this DB
		// Write the records back to the database
		for _, v := range cacheEntry {
			err := db.connection.Save(v)
			if err != nil {
				logHandler.ErrorLogger.Printf("{FLUSH} <%v> Error writing back cached record of type [%v] to database: %v", db.Name, reflect.TypeOf(v), err)
			}
			recordCount++
		}
		// Clear the cache entry
		delete(inMemoryCache, key)
		tableCount++
	}
	dur := time.Since(start)
	db.cacheInitialised = false
	logHandler.EventLogger.Printf("{FLUSH}<%v> Flushed cache for [%v.db] - Total records flushed: %d across %d tables in %v", db.Name, db.Name, recordCount, tableCount, dur.String())
	return nil
}
