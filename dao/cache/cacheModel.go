package cache

import (
	"github.com/mt1976/frantic-core/dao/fields"
)

type cache struct {
	tablesActive map[string]bool
	key          map[string]fields.Field
	indices      map[string][]fields.Field
	cache        map[string]entrys // in-memory storage, indexde by table then by cache key
}

type entrys map[any]any // Map indexed by keyfield, storing one record per slot

var Cache = cache{}
