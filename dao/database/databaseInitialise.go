package database

import (
	"github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
)

var (
	connectionPool        map[string]*DB         = make(map[string]*DB)                                 // map of database connections, indexed by domain.
	connectionPoolMaxSize int                    = 10                                                   // maximum number of connections
	cfg                   *commonConfig.Settings = commonConfig.Get()                                   // configuration settings
	dataValidator         *validator.Validate    = validator.New(validator.WithRequiredStructEnabled()) // data validator
	inMemoryCache         map[string]any         = make(map[string]any)                                 // in-memory storage
)

// init initializes the database package
func init() {

	connectionPoolMaxSize = cfg.GetDatabase_PoolSize()
	logHandler.DatabaseLogger.Printf("[CON] Database Connection Pool Size [%v]", connectionPoolMaxSize)
	logHandler.CacheLogger.Printf("[CACHE] In-Memory Cache Initialised %d items", len(inMemoryCache))

}
