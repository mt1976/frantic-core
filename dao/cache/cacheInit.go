package cache

import (
	"time"

	"github.com/mt1976/frantic-core/dao/fields"
)

const defaultCacheExpiry = 100 * 365 * 24 * time.Hour // 100 years

// Initialise sets up the cache system.

func Initialise() {
	// Initialise global cache maps.
	// Note: do not use ':=' here, that would shadow the package-level Cache.
	Cache.at = time.Now()
	Cache.cache = make(map[string]entrys)
	Cache.indices = make(map[string][]fields.Field)
	Cache.key = make(map[string]fields.Field)
	Cache.tablesActive = make(map[string]bool)
	Cache.count = make(map[string]int64)
	Cache.expiry = make(map[string]time.Duration)
}
