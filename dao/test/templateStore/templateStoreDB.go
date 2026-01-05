package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Add any initialisation code to the Initialise function

import (
	"context"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var activeDB *database.DB
var initialised bool = false // default to false
var cfg *commonConfig.Settings

// Initialise sets up the database connection and prepares the DAO for operations.
//
// This function establishes a connection to the database specified for the DAO.
// It configures the database connection with necessary parameters and indices.
// The function also sets the initialised flag to true, indicating that the DAO is ready for use.
//
// Parameters:
//
//	ctx context.Context: The context for managing request-scoped values, cancellation signals, and deadlines.
//
// Returns:
//
//	None
func Initialise(ctx context.Context) {
	logHandler.DatabaseLogger.Printf("Opening connection to %v", Domain)

	timing := timing.Start(Domain, actions.INITIALISE.GetCode(), "Initialise")
	cfg = commonConfig.Get()
	// For a specific database connection, use WithNameSpace("value"), otherwise don't specify a namespace
	// Example:
	//			activeDB = database.Connect(database.WithVerbose(false), database.WithCaching(true), database.WithCacheKey(Fields.Key), database.WithIndex(database.Field(Fields.RealName)), database.WithIndex(database.Field(Fields.UserName)))
	// Example to connect to a named database from config:
	//
	//			activeDB = database.Connect(database.WithVerbose(false), database.WithCaching(true), database.WithCacheKey(Fields.Key), database.WithNameSpace(cfg.DatabaseNamespace), database.WithIndex(database.Field(Fields.RealName)), database.WithIndex(database.Field(Fields.UserName)))
	// Example to connect to a specific named database
	//
	//			activeDB = database.Connect(database.WithVerbose(false), database.WithCaching(true), database.WithCacheKey(Fields.Key), database.WithNameSpace("BNK"), database.WithIndex(database.Field(Fields.RealName)), database.WithIndex(database.Field(Fields.UserName)))

	activeDB = database.Connect(database.WithVerbose(false), database.WithCaching(true), database.WithCacheKey(Fields.Key), database.WithNameSpace("cheeseOnToast"), database.WithIndex(database.Field(Fields.RealName)), database.WithIndex(database.Field(Fields.UserName)))
	initialised = true

	//TODO: Add any initialisation code here

	timing.Stop(1)
	logHandler.DatabaseLogger.Printf("Opened connection to %v", Domain)
}

// IsInitialised returns the initialisation status of the DAO.
//
// This function checks whether the Data Access Object (DAO) has been successfully initialised.
// It returns a boolean value indicating the initialisation status.
//
// Returns:
//
//	bool: A boolean value indicating whether the DAO is initialised (true) or not (false).
func IsInitialised() bool {
	return initialised
}

// Close terminates the connection to the database used by the DAO.
//
// This function closes the active database connection associated with the Data Access Object (DAO).
// It ensures that any resources related to the database connection are properly released.
func Close() {
	logHandler.EventLogger.Printf("Closing connection to %v", Domain)
	if activeDB != nil {
		activeDB.Disconnect()
	}
	initialised = false

	logHandler.EventLogger.Printf("Closed connection to %v", Domain)
}

// GetDatabaseConnections returns a function that fetches the current database instances.
//
// This function is used to retrieve the active database instances being used by the application.
// It returns a function that, when called, returns a slice of pointers to `database.DB` and an error.
//
// Returns:
//
//	func() ([]*database.DB, error): A function that returns a slice of pointers to `database.DB` and an error.
func GetDatabaseConnections() func() ([]*database.DB, error) {
	return func() ([]*database.DB, error) {
		return []*database.DB{activeDB}, nil
	}
}
