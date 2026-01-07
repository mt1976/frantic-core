package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY

import (
	"context"

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
	err := activeDB.PreLoadCache(&[]TemplateStore{})
	logHandler.CacheLogger.Printf("PreLoad [%+v] complete", Domain)
	return err
}

func CacheSpew() {
	logHandler.CacheLogger.Printf("CacheSpew [%+v]", Domain)
	activeDB.CacheSpew()
	logHandler.CacheLogger.Printf("CacheSpew [%+v] complete", Domain)
}

func FlushCache() error {
	logHandler.CacheLogger.Printf("FlushCache [%+v]", Domain)
	err := activeDB.Flush()
	logHandler.CacheLogger.Printf("FlushCache [%+v] complete", Domain)
	return err
}
