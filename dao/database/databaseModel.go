package database

import (
	"github.com/asdine/storm/v3"
)

// DB represents a database connection and its configuration
type DB struct {
	connection   *storm.DB
	Name         string
	databaseName string
	initialised  bool
	//	withCaching      bool
	//	withCacheKey     Field
	verbose        bool
	timeout        int
	poolSize       int
	withEncryption bool
	//indices        []Field
	//	cacheInitialised bool
	// cachedTables  map[string]bool
	// cacheKeyField map[string]Field
}
