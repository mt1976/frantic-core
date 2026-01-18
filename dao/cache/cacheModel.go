package cache

import (
	"time"

	"github.com/mt1976/frantic-core/dao/fields"
)

type cache struct {
	created      time.Time
	updated      time.Time
	tablesActive map[entity]bool
	key          map[entity]fields.Field
	indices      map[entity][]fields.Field
	cache        map[entity]entrys // in-memory storage, indexde by table then by cache key
	count        map[entity]int64
	expiry       map[entity]time.Duration
	synchroniser map[entity]func(any) error
	hydrator     map[entity]func() ([]any, error)
}

type entity string
type entrys map[any]dataCache // Map indexed by keyfield, storing one record per slot
// dataCache is the structure stored in each cache entry
type dataCache struct {
	dataRecord     any
	cacheTimestamp time.Time
}

var Cache = cache{}
