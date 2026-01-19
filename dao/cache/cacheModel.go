package cache

import (
	"time"

	"github.com/mt1976/frantic-core/dao/entities"
)

type cache struct {
	created      time.Time
	updated      time.Time
	tablesActive map[entities.Table]bool
	key          map[entities.Table]entities.Field
	indices      map[entities.Table][]entities.Field
	cache        map[entities.Table]entrys // in-memory storage, indexde by table then by cache key
	count        map[entities.Table]int64
	expiry       map[entities.Table]time.Duration
	synchroniser map[entities.Table]func(any) error
	hydrator     map[entities.Table]func() ([]any, error)
}

type entrys map[any]dataCache // Map indexed by keyfield, storing one record per slot
// dataCache is the structure stored in each cache entry
type dataCache struct {
	dataRecord     any
	cacheTimestamp time.Time
}

var Cache = cache{}
