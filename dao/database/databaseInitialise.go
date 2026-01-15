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
)

// init initializes the database package
func init() {

	connectionPoolMaxSize = cfg.GetDatabase_PoolSize()
	logHandler.DatabaseLogger.Printf("[CON] Database Connection Pool Size [%v]", connectionPoolMaxSize)

}
