package templateStoreV2

import (
	"context"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/cache"
	"github.com/mt1976/frantic-core/logHandler"
)

func PreLoad(ctx context.Context) error {
	logHandler.CacheLogger.Printf("PreLoad [%+v]", tableName)
	_ = ctx
	logHandler.CacheLogger.Printf("PreLoad [%+v] complete", tableName)
	return nil
}

func CacheSpew() {
	logHandler.CacheLogger.Printf("CacheSpew [%+v]", tableName)
	cache.SpewFor(TableName)
	logHandler.CacheLogger.Printf("CacheSpew [%+v] complete", tableName)
}

func FlushCache() error {
	logHandler.CacheLogger.Printf("FlushCache [%+v]", tableName)
	err := cache.SynchroniseForType(TemplateStore{})
	logHandler.CacheLogger.Printf("FlushCache [%+v] complete", tableName)
	return err
}

func HydrateCache() error {
	logHandler.CacheLogger.Printf("HydrateCache [%+v]", tableName)
	err := cache.HydrateForType(TemplateStore{})
	logHandler.CacheLogger.Printf("HydrateCache [%+v] complete", tableName)
	return err
}

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

func CacheSynchroniser(ctx context.Context) func(any) error {
	logHandler.InfoLogger.Printf("Defining Sync function for %v", tableName)
	return func(data any) error {
		record := data.(TemplateStore)
		logHandler.InfoLogger.Printf("Sync cache for %v Key: %v", tableName, record.Key)
		return record.UpdateWithAction(ctx, audit.SYNC, "Cache Sync Update")
	}
}
