package cache

import "github.com/mt1976/frantic-core/logHandler"

func CacheSpew() {
	logHandler.WarningLogger.Println("Caching Status")
	if len(Cache.tablesActive) == 0 {
		logHandler.WarningLogger.Println("No tables are currently cached")
	}
	msg := ". Cached Tables: "
	for tableName := range Cache.tablesActive {
		msg += tableName + " "
	}
	logHandler.WarningLogger.Println(msg)
	// Display A COUNT OF THE RECORDS IN THE CACHE
	logHandler.WarningLogger.Println(". Cached Records Summary")
	for tableName := range Cache.tablesActive {
		inMemoryCacheEntry, exists := Cache.cache[tableName]
		if !exists {
			logHandler.WarningLogger.Printf(". 	Table [%v] has 0 cached records", tableName)
			continue
		}
		//
		logHandler.WarningLogger.Printf(". 	Table [%v] has %d cached records", tableName, len(inMemoryCacheEntry))
	}
}
