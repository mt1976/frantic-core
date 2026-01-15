package cache

import (
	"github.com/mt1976/frantic-core/dao/fields"
)

func New() {
	// Initialise global cache maps.
	// Note: do not use ':=' here, that would shadow the package-level Cache.
	Cache.cache = make(map[string]entrys)
	Cache.indices = make(map[string][]fields.Field)
	Cache.key = make(map[string]fields.Field)
	Cache.tablesActive = make(map[string]bool)
}
