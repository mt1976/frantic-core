package cache

import (
	"time"

	"github.com/mt1976/frantic-core/dao/entities"
)

const defaultCacheExpiry = 100 * 365 * 24 * time.Hour // 100 years

// Initialise sets up the cache system.

func Initialise() {
	// Initialise global cache maps.
	// Note: do not use ':=' here, that would shadow the package-level Cache.
	Cache.created = time.Now()
	Cache.updated = time.Time{}
	Cache.cache = make(map[entities.Table]entrys)
	Cache.indices = make(map[entities.Table][]entities.Field)
	Cache.key = make(map[entities.Table]entities.Field)
	Cache.tablesActive = make(map[entities.Table]bool)
	Cache.count = make(map[entities.Table]int64)
	Cache.expiry = make(map[entities.Table]time.Duration)
	Cache.synchroniser = make(map[entities.Table]func(any) error)
	Cache.hydrator = make(map[entities.Table]func() ([]any, error))
}
