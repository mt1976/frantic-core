package cache

import "github.com/mt1976/frantic-core/dao/database"

func init() {
	// Lets clear everything up
	Cache := cache{}
	Cache.cache = make(map[string]entrys)
	Cache.indices = make(map[string][]database.Field)
	Cache.key = make(map[string]database.Field)
	Cache.tablesActive = make(map[string]bool)
}
