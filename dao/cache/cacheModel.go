package cache

import (
	"time"

	"github.com/mt1976/frantic-core/dao/fields"
)

type cache struct {
	at           time.Time
	tablesActive map[string]bool
	key          map[string]fields.Field
	indices      map[string][]fields.Field
	cache        map[string]entrys // in-memory storage, indexde by table then by cache key
	count        map[string]int64
	expiry       map[string]time.Duration
	synchroniser map[string]func(any) error
}

type entrys map[any]dataCache // Map indexed by keyfield, storing one record per slot
// dataCache is the structure stored in each cache entry
type dataCache struct {
	dataRecord     any
	cacheTimestamp time.Time
}

var Cache = cache{}
