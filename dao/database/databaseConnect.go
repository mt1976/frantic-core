package database

import (
	"strings"

	"github.com/asdine/storm/v3"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"

	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// connect establishes a database connection with the provided options
// It applies default settings and overrides them with any specified options.
// It also manages the connection pool to reuse existing connections.
func connect(table any, options ...Option) *DB {
	// Create default configuration
	config := &connectionConfig{
		withCaching:      false,
		Verbose:          false,
		timeout:          30,
		poolSize:         connectionPoolMaxSize,
		nameSpace:        "main",
		withEncryption:   false,
		indices:          []entities.Field{},
		withCacheKey:     "ID",
		cacheInitialised: false,
	}

	// Apply all provided options
	for _, option := range options {
		option(config)
	}

	// Log the applied configuration
	logHandler.DatabaseLogger.Printf("[CON]{CONNECT} Configuration for %v.db: caching: %t, cacheKey: %v, verbose: %t, timeout: %d, poolSize: %d, nameSpace: %s, encryption: %t, indices: %v",
		config.nameSpace, config.withCaching, config.withCacheKey, config.Verbose, config.timeout, config.poolSize, config.nameSpace, config.withEncryption, config.indices)

	if config.withCaching && config.withCacheKey == "" {
		logHandler.DatabaseLogger.Panicf("[CON]{CONNECT} Caching enabled but no cache key provided for [%v.db]", config.nameSpace)
		panic(commonErrors.ErrDBConnect)
	}

	// Ensure the name is lowercase
	config.nameSpace = strings.ToLower(config.nameSpace)
	logHandler.DatabaseLogger.Printf("[CON]{CONNECT} Opening Connection to [%v.db] data (%v)", config.nameSpace, len(connectionPool))
	// list the connection pool
	if config.Verbose {
		for key, value := range connectionPool {
			logHandler.DatabaseLogger.Printf("[CON]{CONNECT} Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
		}
	}
	// check if connection already exists
	if connectionPool[config.nameSpace] != nil && connectionPool[config.nameSpace].Name == config.nameSpace {
		logHandler.DatabaseLogger.Printf("[CON]{CONNECT} Connection already open [%v], using connection pool [%v] [codec=%v]", connectionPool[config.nameSpace].Name, connectionPool[config.nameSpace].databaseName, connectionPool[config.nameSpace].connection.Node.Codec().Name())
		rtn := connectionPool[config.nameSpace]
		// Update configuration in case options have changed
		//rtn.withCaching = config.withCaching
		//rtn.withCacheKey = config.withCacheKey
		// rtn.verbose = config.Verbose
		// rtn.timeout = config.timeout
		// rtn.poolSize = config.poolSize
		// rtn.withEncryption = config.withEncryption
		// rtn.indices = config.indices
		//rtn.cacheInitialised = config.cacheInitialised

		return rtn
	}

	logHandler.DatabaseLogger.Printf("[CON]{CONNECT} (re)Opening [%v.db] data connection", config.nameSpace)
	// Open a new connection

	db := DB{}
	db.Name = config.nameSpace
	db.databaseName = ioHelpers.GetDBFileName(db.Name)
	db.initialised = false
	// db.withCaching = config.withCaching
	// db.withCacheKey = config.withCacheKey
	db.verbose = config.Verbose
	db.timeout = config.timeout
	db.poolSize = config.poolSize
	db.withEncryption = config.withEncryption
	//db.indices = config.indices
	// db.cacheInitialised = false
	// logHandler.DatabaseLogger.Printf("[CON]{CONNECT}  Opening [%v.db] data connection *%+v*", db.Name, db)
	connect := timing.Start(db.Name, "Connect", db.databaseName)
	var err error
	db.connection, err = storm.Open(db.databaseName, storm.BoltOptions(0666, nil))
	if err != nil {
		connect.Stop(0)
		logHandler.DatabaseLogger.Fatalf("[CON]{CONNECT} Opening [%v.db] connection Error=[%v]", strings.ToLower(db.databaseName), err.Error())
		panic(commonErrors.ErrConnectWrapper(err))
	}
	if db.verbose {
		logHandler.DatabaseLogger.Printf("[CON]{CONNECT}  Connection Pool [%+v]", connectionPool)
		for key, value := range connectionPool {
			logHandler.DatabaseLogger.Printf("[CON]{CONNECT}  Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
		}
	}
	// Add to connection pool
	addConnectionToPool(db, db.Name)
	if db.verbose {
		logHandler.DatabaseLogger.Printf("[CON]{CONNECT}  Connection Pool [%+v]", connectionPool)
		for key, value := range connectionPool {
			logHandler.DatabaseLogger.Printf("[CON]{CONNECT}  Connection Pool [%v] [%v] [codec=%v] %v", key, value.databaseName, value.connection.Node.Codec().Name(), value.initialised)
		}
	}
	logHandler.DatabaseLogger.Printf("[CON]{CONNECT} Opened [%v.db] data connection [codec=%v] %v", db.databaseName, db.connection.Node.Codec().Name(), db.initialised)

	// // Enable caching for the specified table if caching is enabled
	// if config.withCaching && table != nil {
	// 	enableCachingForTable(&db, table)
	// 	if err != nil {
	// 		logHandler.DatabaseLogger.Panicf("[CON]{CONNECT} Error enabling caching for table %v: %v", GetStructType(table), err.Error())
	// 		panic(commonErrors.ErrConnectWrapper(err))
	// 	}
	// }

	connect.Stop(1)
	return &db
}

// validate checks the data against validation rules before database operations
// It uses a timing mechanism to log the duration of the validation process.
// If validation fails, it logs the error and returns a wrapped validation error.
func validate(data any, db *DB) error {
	timer := timing.Start(db.Name, "Validate", "")
	logHandler.DatabaseLogger.Printf("[CON]{VALIDATE} Validate [%+v] [%v.db]", GetStructType(data), db.Name)
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.DatabaseLogger.Printf("[CON]{VALIDATE} error validating %v %v [%v.db]", err.Error(), GetStructType(data), db.Name)
		timer.Stop(0)
		return commonErrors.ErrValidationWrapper(err)
	}
	timer.Stop(1)
	return nil
}

// Connect establishes a database connection with the provided options
// It is the primary function to initiate a connection using various configuration options.
func Connect(table any, options ...Option) *DB {
	logHandler.DatabaseLogger.Printf("[CON] %d Options ", len(options))
	return connect(table, options...)
}

// ConnectToNamedDB establishes a database connection to a named database
// DEPRECATED: ConnectToNamedDB - Use Connect with WithNameSpace option instead
func ConnectToNamedDB(name string, options ...Option) *DB {
	logHandler.WarningLogger.Println("[CON] DEPRECATED: ConnectToNamedDB - Use Connect with WithNameSpace option instead")
	panic("Deprecated: ConnectToNamedDB - Use Connect with WithNameSpace option instead")
	//return connect(options...)
}

// Disconnect closes the database connection and removes it from the connection pool
// It uses a timing mechanism to log the duration of the disconnection process.
// If disconnection fails, it logs the error and panics with a wrapped disconnect error.
func (db *DB) Disconnect() {
	timer := timing.Start(db.Name, "Disconnect", db.databaseName)
	logHandler.DatabaseLogger.Printf("[CON]{DISCONNECT} Disconnecting [%v.db] connection", db.Name)
	err := db.connection.Close()
	if err != nil {
		logHandler.DatabaseLogger.Panicf("[CON]{DISCONNECT} Closing [%v.db] %v ", db.Name, err.Error())
		panic(commonErrors.ErrDisconnectWrapper(err))
	}
	releaseFromConnectionPool(db)
	logHandler.DatabaseLogger.Printf("[CON]{DISCONNECT} Closed [%v.db] connection", db.Name)
	if db.verbose {
		for key, value := range connectionPool {
			logHandler.DatabaseLogger.Printf("[CON]{DISCONNECT} Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
		}
	}
	timer.Stop(1)
}

func (db *DB) Reconnect() {
	logHandler.DatabaseLogger.Printf("[CON]{RECONNECT} Reconnecting [%v.db] data - %+v", db.Name, db)
	logHandler.DatabaseLogger.Printf("[CON]{RECONNECT} Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CON]{RECONNECT} Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	connect(WithNameSpace(db.Name))
	logHandler.DatabaseLogger.Printf("[CON]{RECONNECT} Reconnected [%v.db] data", db.Name)
}
