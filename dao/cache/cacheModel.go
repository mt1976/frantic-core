package cache

import "github.com/mt1976/frantic-core/dao/database"

type cache struct {
	tables  map[string]bool
	keys    map[string]database.Field
	indices map[string][]database.Field
	cache   map[string]entrys // in-memory storage, indexde by table then by cache key
}

type entrys map[string]any // Map indexed by keyfield, storing one record per slot

var Cache = cache{}
