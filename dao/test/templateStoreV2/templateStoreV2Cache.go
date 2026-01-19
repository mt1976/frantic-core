package templateStoreV2

import (
	"context"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/cache"
	"github.com/mt1976/frantic-core/logHandler"
)

// PreLoad performs any pre-load work required before the cache is hydrated.
func PreLoad(ctx context.Context) error {
	logHandler.CacheLogger.Printf("PreLoad [%+v]", tableName)
	_ = ctx
	logHandler.CacheLogger.Printf("PreLoad [%+v] complete", tableName)
	return nil
}

// CacheSpew writes the current cache state for this table to the logs.
func CacheSpew() {
	logHandler.CacheLogger.Printf("CacheSpew [%+v]", tableName)
	cache.SpewFor(TableName)
	logHandler.CacheLogger.Printf("CacheSpew [%+v] complete", tableName)
}

// FlushCache synchronises cached records back to the underlying database.
func FlushCache() error {
	logHandler.CacheLogger.Printf("FlushCache [%+v]", tableName)
	err := cache.SynchroniseForType(TemplateStore{})
	logHandler.CacheLogger.Printf("FlushCache [%+v] complete", tableName)
	return err
}

// HydrateCache loads all records into cache.
func HydrateCache() error {
	logHandler.CacheLogger.Printf("HydrateCache [%+v]", tableName)
	err := cache.HydrateForType(TemplateStore{})
	logHandler.CacheLogger.Printf("HydrateCache [%+v] complete", tableName)
	return err
}

// CacheHydrator returns the cache hydrator function for this table.
func CacheHydrator(ctx context.Context) func() ([]any, error) {
	_ = ctx
	return func() ([]any, error) {
		records, err := GetAll()
		if err != nil {
			return nil, err
		}
		result := make([]any, len(records))
		for i := range records {
			result[i] = records[i]
		}
		return result, nil
	}
}

// CacheSynchroniser returns the cache synchroniser function for this table.
func CacheSynchroniser(ctx context.Context) func(any) error {
	logHandler.InfoLogger.Printf("Defining Sync function for %v", tableName)
	return func(data any) error {
		switch rec := data.(type) {
		case TemplateStore:
			logHandler.CacheLogger.Printf("Sync cache for %v Key: %v", tableName, rec.Key)
			return rec.UpdateWithAction(ctx, audit.SYNC, "Cache Sync Update")
		case *TemplateStore:
			if rec == nil {
				return nil
			}
			logHandler.CacheLogger.Printf("Sync cache for %v Key: %v", tableName, rec.Key)
			return rec.UpdateWithAction(ctx, audit.SYNC, "Cache Sync Update")
		default:
			logHandler.WarningLogger.Printf("Sync cache for %v received unexpected type %T", tableName, data)
			return nil
		}
	}
}
