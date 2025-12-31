package database

import (
	"strings"

	"github.com/asdine/storm/v3"
	"github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var (
	connectionPool        map[string]*DB         = make(map[string]*DB)                                 // map of database connections, indexed by domain.
	connectionPoolMaxSize int                    = 10                                                   // maximum number of connections
	cfg                   *commonConfig.Settings = commonConfig.Get()                                   // configuration settings
	dataValidator         *validator.Validate    = validator.New(validator.WithRequiredStructEnabled()) // data validator
	inMemoryCache         map[string]any         = make(map[string]any)                                 // in-memory storage
)

func init() {
	//	Connect()
	//	cfg = commonConfig.Get()
	//dataValidator = validator.New(validator.WithRequiredStructEnabled())
	//connectionPool = make(map[string]*DB)
	connectionPoolMaxSize = cfg.GetDatabase_PoolSize()
	logHandler.DatabaseLogger.Printf("Database Connection Pool Size [%v]", connectionPoolMaxSize)
}

func Connect(options ...Option) *DB {
	logHandler.InfoLogger.Printf("Options : '%+v'", options)
	return connect(options...)
}

func ConnectToNamedDB(name string, options ...Option) *DB {
	logHandler.WarningLogger.Println("DEPRECATED: ConnectToNamedDB - Use Connect with WithNameSpace option instead")
	options = append(options, WithNameSpace(name))
	return connect(options...)
}

func (db *DB) Disconnect() {
	timer := timing.Start(db.Name, actions.DISCONNECT.Code, db.databaseName)
	logHandler.DatabaseLogger.Printf("[DISCONNECT] Disconnecting [%v.db] connection", db.Name)
	err := db.connection.Close()
	if err != nil {
		logHandler.DatabaseLogger.Panicf("[DISCONNECT] Closing [%v.db] %v ", db.Name, err.Error())
		panic(commonErrors.WrapDisconnectError(err))
	}
	releaseFromConnectionPool(db)
	logHandler.DatabaseLogger.Printf("[DISCONNECT] Closed [%v.db] connection", db.Name)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[DISCONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	timer.Stop(1)
}

func (db *DB) Reconnect() {
	logHandler.DatabaseLogger.Printf("[RECONNECT] Reconnecting [%v.db] data - %+v", db.Name, db)
	logHandler.DatabaseLogger.Printf("[RECONNECT] Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[RECONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	connect(WithNameSpace(db.Name))
	logHandler.DatabaseLogger.Printf("[RECONNECT] Reconnected [%v.db] data", db.Name)
}

func connect(options ...Option) *DB {
	// Create default configuration
	config := &connectionConfig{
		withCaching:      false,
		verbose:          false,
		timeout:          30,
		poolSize:         connectionPoolMaxSize,
		nameSpace:        "main",
		withEncryption:   false,
		indices:          []Field{},
		withCacheKey:     "ID",
		cacheInitialised: false,
	}

	// Clear the inmory cache
	inMemoryCache = make(map[string]any)

	// Apply all provided options
	for _, option := range options {
		option(config)
	}

	// Log the applied configuration
	logHandler.DatabaseLogger.Printf("[CONNECT] Configuration for [%v.db]: caching=%t, verbose=%t, timeout=%d, poolSize=%d, nameSpace=%s, encryption=%t, indices=%v",
		config.nameSpace, config.withCaching, config.verbose, config.timeout, config.poolSize, config.nameSpace, config.withEncryption, config.indices)

	// Ensure the name is lowercase
	config.nameSpace = strings.ToLower(config.nameSpace)
	logHandler.DatabaseLogger.Printf("[CONNECT] Opening Connection to [%v.db] data (%v)", config.nameSpace, len(connectionPool))
	// list the connection pool
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	// check if connection already exists
	if connectionPool[config.nameSpace] != nil && connectionPool[config.nameSpace].Name == config.nameSpace {
		logHandler.DatabaseLogger.Printf("[CONNECT] Connection already open [%v], using connection pool [%v] [codec=%v]", connectionPool[config.nameSpace].Name, connectionPool[config.nameSpace].databaseName, connectionPool[config.nameSpace].connection.Node.Codec().Name())
		return connectionPool[config.nameSpace]
	}

	logHandler.DatabaseLogger.Printf("[CONNECT] (re)Opening [%v.db] data connection", config.nameSpace)
	// Open a new connection

	db := DB{}
	db.Name = config.nameSpace
	db.databaseName = ioHelpers.GetDBFileName(db.Name)
	db.initialised = false
	db.withCaching = config.withCaching
	db.withCacheKey = config.withCacheKey
	db.verbose = config.verbose
	db.timeout = config.timeout
	db.poolSize = config.poolSize
	db.withEncryption = config.withEncryption
	db.indices = config.indices
	db.cacheInitialised = false
	logHandler.DatabaseLogger.Printf("[CONNECT]  Opening [%v.db] data connection *%+v*", db.Name, db)
	connect := timing.Start(db.Name, actions.CONNECT.GetCode(), db.databaseName)
	var err error
	db.connection, err = storm.Open(db.databaseName, storm.BoltOptions(0666, nil))
	if err != nil {
		connect.Stop(0)
		logHandler.DatabaseLogger.Fatalf("[CONNECT] Opening [%v.db] connection Error=[%v]", strings.ToLower(db.databaseName), err.Error())
		panic(commonErrors.WrapConnectError(err))
	}
	db.initialised = true
	logHandler.DatabaseLogger.Printf("[CONNECT]  Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CONNECT]  Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	// Add to connection pool
	addConnectionToPool(db, db.Name)
	logHandler.DatabaseLogger.Printf("[CONNECT]  Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CONNECT]  Connection Pool [%v] [%v] [codec=%v] %v", key, value.databaseName, value.connection.Node.Codec().Name(), value.initialised)
	}
	logHandler.DatabaseLogger.Printf("[CONNECT] Opened [%v.db] data connection [codec=%v] %v", db.databaseName, db.connection.Node.Codec().Name(), db.initialised)
	connect.Stop(1)
	return &db
}

func addConnectionToPool(db DB, key string) {
	logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL] Adding [%v] to connection pool (%v)", key, db.databaseName)
	if len(connectionPool) >= connectionPoolMaxSize {
		logHandler.DatabaseLogger.Panicf("[CONNECTIONPOOL] Connection pool full [%v]", connectionPoolMaxSize)
		return
	}
	connectionPool[key] = &db
	logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL] Connection pool [size=%v]", len(connectionPool))
}

func releaseFromConnectionPool(db *DB) {
	logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL] Removing [%v] from connection pool (%v)", db.Name, db.databaseName)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL]  Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	connectionPool[db.Name] = nil
	delete(connectionPool, db.Name)
	logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL] Connection pool [size=%v]", len(connectionPool))
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[CONNECTIONPOOL]  Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
}

func validate(data any, db *DB) error {
	timer := timing.Start(db.Name, actions.VALIDATE.GetCode(), "")
	logHandler.DatabaseLogger.Printf("validate [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.DatabaseLogger.Printf("error walidating %v %v [%v.db]", err.Error(), dao.GetStructType(data), db.Name)
		timer.Stop(0)
		return commonErrors.WrapValidationError(err)
	}
	timer.Stop(1)
	return nil
}
