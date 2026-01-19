package templateStoreV2

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/cache"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var activeDBConnection *database.DB
var databaseConnectionActive bool
var cfg *commonConfig.Settings

// Initialise opens the database connection for TemplateStoreV2 and optionally enables caching.
func Initialise(ctx context.Context, cached bool) {
	logHandler.DatabaseLogger.Printf("Opening connection to %v", tableName)
	logHandler.TraceLogger.Printf("Initialising %v DAO Caching: %t", tableName, cached)

	clock := timing.Start(tableName, "Initialise", "Initialise")
	cfg = commonConfig.Get()
	_ = cfg

	activeDBConnection = database.Connect(TemplateStore{}, database.WithVerbose(false), database.WithCaching(cached), database.WithCacheKey(Fields.Key), database.WithNameSpace("cheeseOnToast"))
	databaseConnectionActive = true

	clock.Stop(1)
	logHandler.DatabaseLogger.Printf("Opened connection to %v", tableName)
}

// IsInitialised reports whether the DAO has an active database connection.
func IsInitialised() bool {
	return databaseConnectionActive
}

// Close flushes the cache (if enabled) and closes the active database connection.
func Close() {
	logHandler.DatabaseLogger.Printf("Closing connection to %v", tableName)

	flusherr2 := cache.SynchroniseForType(TemplateStore{})
	if flusherr2 != nil {
		logHandler.ErrorLogger.Printf("Error flushing cache: %v", flusherr2)
	} else {
		logHandler.InfoLogger.Printf("Cache flushed successfully")
	}

	if activeDBConnection != nil {
		activeDBConnection.Disconnect()
	}
	databaseConnectionActive = false
	logHandler.DatabaseLogger.Printf("Closed connection to %v", tableName)
}

// GetDatabaseConnections returns a function that supplies the database connections used by this DAO.
func GetDatabaseConnections() func() ([]*database.DB, error) {
	return func() ([]*database.DB, error) {
		return []*database.DB{activeDBConnection}, nil
	}
}
