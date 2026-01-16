package cache

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mt1976/frantic-core/logHandler"
)

func Spew() {

	//godump.Dump(Cache)
	//fmt.Printf("Cache Dump: %+v", Cache)

	logHandler.InfoBanner("Cache", "Report", "Starting Cache Report")
	if len(Cache.tablesActive) == 0 {
		logHandler.InfoLogger.Println("No tables are currently cached")
	}
	msg := ". Cached Tables: "
	for tableName := range Cache.tablesActive {
		msg += tableName + " "
	}

	if len(Cache.tablesActive) == 0 {
		logHandler.InfoBanner("Cache", "Report", "End Report")
		return
	}
	logHandler.InfoLogger.Println(msg)

	logHandler.InfoLogger.Println(". Cached Keys Summary")
	for tableName, keyField := range Cache.key {
		logHandler.InfoLogger.Printf(". 	Table [%v] has Key Field [%v]", tableName, keyField.String())
	}

	// Display A COUNT OF THE RECORDS IN THE CACHE
	logHandler.InfoLogger.Println(". Cached Records Summary")
	for tableName := range Cache.tablesActive {
		cachedEntry, exists := Cache.cache[tableName]
		if !exists {
			logHandler.WarningLogger.Printf(". 	Table [%v] has 0 cached records", tableName)
			continue
		}
		cacheExpiry := Cache.expiry[tableName]
		//
		logHandler.InfoLogger.Printf(". 	Table [%v] has [%d] cached records and expiry set to [%v]", tableName, len(cachedEntry), cacheExpiry)
		for key, record := range cachedEntry {
			logHandler.InfoLogger.Printf(".       %v>%v: %v - expires: %v(%v)", tableName, Cache.key[tableName].String(), key, record.cacheTimestamp.Format(time.RFC3339Nano), humanize.Time(record.cacheTimestamp))
		}
	}
	logHandler.InfoBanner("Cache", "Report", "End Report")
}
