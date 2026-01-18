package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY

import (
	"context"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/cache"
	"github.com/mt1976/frantic-core/logHandler"
)

// PreLoad preloads the cache for the TemplateStore DAO.
//
// This function preloads the cache for the TemplateStore Data Access Object (DAO).
// It retrieves all records from the database and stores them in the cache for faster access.
// The function logs the start and completion of the preload process.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, cancellation signals, and deadlines.
//
// Returns:
//   - error: An error object if any issues occur during the preload process; otherwise, nil.
func PreLoad(ctx context.Context) error {
	logHandler.CacheLogger.Printf("PreLoad [%+v]", Domain)
	//err := activeDB.PreLoadCache(&[]TemplateStore{})
	logHandler.CacheLogger.Printf("PreLoad [%+v] complete", Domain)
	return nil
}

func CacheSpew() {
	logHandler.CacheLogger.Printf("CacheSpew [%+v]", Domain)
	//activeDB.CacheSpew()
	logHandler.CacheLogger.Printf("CacheSpew [%+v] complete", Domain)
}

func FlushCache() error {
	logHandler.CacheLogger.Printf("FlushCache [%+v]", Domain)
	err := cache.Synchronise(TemplateStore{})
	logHandler.CacheLogger.Printf("FlushCache [%+v] complete", Domain)
	return err
}

func HydrateCache() error {
	logHandler.CacheLogger.Printf("HydrateCache [%+v]", Domain)
	err := cache.Hydrate(TemplateStore{})
	logHandler.CacheLogger.Printf("HydrateCache [%+v] complete", Domain)
	return err
}

// HydratorFuncTemplateStore returns a function that retrieves all TemplateStore records for cache hydration.
//
// Parameters:
//   - ctx: The context for managing request-scoped values, cancellation signals, and deadlines.
//
// Returns:
//   - func() ([]any, error): A function that retrieves all TemplateStore records and returns them as a slice of any type along with an error if any issues occur.
func CacheHydrator(ctx context.Context) func() ([]any, error) {
	return func() ([]any, error) {
		records, err := GetAll()
		if err != nil {
			return nil, err
		}
		result := make([]any, len(records))
		for i, v := range records {
			result[i] = v
		}
		return result, nil
	}
}

// Sync returns a function that flushes the cache for TemplateStore records.
//
// Returns:
//   - error: An error object if any issues occur during the cache flush process; otherwise, nil.
func CacheSynchroniser(ctx context.Context) func(any) error {
	logHandler.InfoLogger.Printf("Defining Sync function for %v", Domain)
	return func(data any) error {
		// Lets update the record in the db from the cache
		// The any parameter is expected to be of type *TemplateStore
		// Type assert the any parameter to *TemplateStore

		record := data.(TemplateStore)
		logHandler.InfoLogger.Printf("Sync cache for %v Key: %v", Domain, record.Key)
		return record.UpdateWithAction(ctx, audit.SYNC, "Cache Sync Update")
	}
}
