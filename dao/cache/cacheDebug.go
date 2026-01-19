package cache

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mt1976/frantic-core/dao/entities"
	"github.com/mt1976/frantic-core/logHandler"
)

func Spew() {

	//godump.Dump(Cache)
	//fmt.Printf("Cache Dump: %+v", Cache)

	logHandler.InfoBanner("Cache", "Report", "Starting Cache Report")

	if len(Cache.tablesActive) == 0 {
		logHandler.InfoLogger.Println("No tables are currently cached")
	}

	logHandler.InfoLogger.Printf("Cache created at: %v", Cache.created.Format(time.RFC3339Nano))
	logHandler.InfoLogger.Printf("Cache updated at: %v", Cache.updated.Format(time.RFC3339Nano))
	logHandler.InfoLogger.Printf("Cache Age: %v", humanize.Time(Cache.created))
	logHandler.InfoLogger.Printf("Cache Last Updated: %v", humanize.Time(Cache.updated))
	logHandler.InfoLogger.Println("")
	msg := ". Cached Tables: "
	for tableName := range Cache.tablesActive {
		msg += string(tableName) + " "
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
		SpewFor(tableName)
	}
	created, updated, noTables, noCacheEntries := Stats()
	logHandler.InfoLogger.Println("")
	logHandler.InfoLogger.Printf("Cache Stats - Created: %v, Updated: %v, Tables: %v, Entries: %v", created.Format(time.RFC3339Nano), updated.Format(time.RFC3339Nano), noTables, noCacheEntries)

	logHandler.InfoBanner("Cache", "Report", "End Report")
}

// SpewFor outputs the same cache report detail as Spew(), but scoped to a single table.
//
// The table name should match the cache's internal table key (typically the struct type name
// as derived by the cache package).
func SpewFor(tableName entities.Table) {
	spewForEntity(tableName)
}

// SpewForType outputs the same cache report detail as Spew(), but scoped to the table
// associated with the provided record type.
//
// Example:
//
//	cache.SpewForType(templateStoreV2.TemplateStore{})
func SpewForType(data any) {
	spewForEntity(entities.GetStructType(data))
}

func spewForEntity(table entities.Table) {
	tableNameStr := table.String()
	if len(Cache.tablesActive) == 0 {
		logHandler.WarningLogger.Printf(". \tTable [%v] is not cached (no active tables)", tableNameStr)
		return
	}

	if _, ok := Cache.tablesActive[table]; !ok {
		logHandler.WarningLogger.Printf(". \tTable [%v] is not currently cached", tableNameStr)
		return
	}

	cachedEntry, exists := Cache.cache[table]
	if !exists {
		logHandler.WarningLogger.Printf(". \tTable [%v] has 0 cached records", tableNameStr)
		return
	}

	cacheExpiry := Cache.expiry[table]
	logHandler.InfoLogger.Printf(". \tTable [%v] has [%d] cached records and expiry set to [%v]", tableNameStr, len(cachedEntry), cacheExpiry)
	for key, record := range cachedEntry {
		keyField, ok := Cache.key[table]
		if ok {
			logHandler.InfoLogger.Printf(".       %v>%v: %v - expires: %v(%v)", tableNameStr, keyField.String(), key, record.cacheTimestamp.Format(time.RFC3339Nano), humanize.Time(record.cacheTimestamp))
			continue
		}
		logHandler.InfoLogger.Printf(".       %v>%v: %v - expires: %v(%v)", tableNameStr, "<unknown-key>", key, record.cacheTimestamp.Format(time.RFC3339Nano), humanize.Time(record.cacheTimestamp))
	}
}

func Stats() (created time.Time, updated time.Time, noTables int64, noCacheEntries int64) {
	// Loop through Cache.cache to count total entries
	var totalEntries int64 = 0
	for _, entries := range Cache.cache {
		totalEntries += int64(len(entries))
	}
	return Cache.created, Cache.updated, int64(len(Cache.tablesActive)), totalEntries
	//
}
