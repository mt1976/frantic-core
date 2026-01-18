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
	Cache.created = time.Now()
	Cache.updated = time.Time{}
	Cache.cache = make(map[entity]entrys)
	Cache.indices = make(map[entity][]fields.Field)
	Cache.key = make(map[entity]fields.Field)
	Cache.tablesActive = make(map[entity]bool)
	Cache.count = make(map[entity]int64)
	Cache.expiry = make(map[entity]time.Duration)
	Cache.synchroniser = make(map[entity]func(any) error)
	Cache.hydrator = make(map[entity]func() ([]any, error))
}
