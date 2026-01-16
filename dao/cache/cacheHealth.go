package cache

import (
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Remove any cached entries that have expired
func PurgeExpiredEntries() {
	watch := timing.Start("Cache", "Purge_Expired_Entries", "")
	now := time.Now()
	logHandler.ServiceLogger.Printf("Cache Purge Started at %v", now.Format(time.RFC3339Nano))
	noPurged := 0
	for tableName, entries := range Cache.cache {
		for key, record := range entries {
			if now.After(record.cacheTimestamp) {
				logHandler.InfoLogger.Printf("Cache Entry for Table [%v] with Key [%v] expired at [%v], removing it", tableName, key, record.cacheTimestamp.Format(time.RFC3339Nano))
				delete(entries, key)
				Cache.count[tableName]--
				noPurged++
			}
		}
	}
	watch.Stop(noPurged)
	logHandler.ServiceLogger.Printf("Cache Purge Completed at %v", now.Format(time.RFC3339Nano))
}
