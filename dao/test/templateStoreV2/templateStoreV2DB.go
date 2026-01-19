package templateStoreV2

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/cache"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var activeDB *database.DB
var dbIsReady bool
var cfg *commonConfig.Settings

func Initialise(ctx context.Context, cached bool) {
	logHandler.DatabaseLogger.Printf("Opening connection to %v", tableName)
	logHandler.TraceLogger.Printf("Initialising %v DAO Caching: %t", tableName, cached)

	clock := timing.Start(tableName, "Initialise", "Initialise")
	cfg = commonConfig.Get()
	_ = cfg

	activeDB = database.Connect(TemplateStore{}, database.WithVerbose(false), database.WithCaching(cached), database.WithCacheKey(Fields.Key), database.WithNameSpace("cheeseOnToast"))
	dbIsReady = true

	clock.Stop(1)
	logHandler.DatabaseLogger.Printf("Opened connection to %v", tableName)
}

func IsInitialised() bool {
	return dbIsReady
}

func Close() {
	logHandler.DatabaseLogger.Printf("Closing connection to %v", tableName)

	flusherr2 := cache.SynchroniseForType(TemplateStore{})
	if flusherr2 != nil {
		logHandler.ErrorLogger.Printf("Error flushing cache: %v", flusherr2)
	} else {
		logHandler.InfoLogger.Printf("Cache flushed successfully")
	}

	if activeDB != nil {
		activeDB.Disconnect()
	}
	dbIsReady = false
	logHandler.DatabaseLogger.Printf("Closed connection to %v", tableName)
}

func GetDatabaseConnections() func() ([]*database.DB, error) {
	return func() ([]*database.DB, error) {
		return []*database.DB{activeDB}, nil
	}
}
