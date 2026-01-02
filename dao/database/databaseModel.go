package database

import (
	"github.com/asdine/storm/v3"
)

// Field represents a database field used for queries
type field string
type Field field

// DB represents a database connection and its configuration
type DB struct {
	connection       *storm.DB
	Name             string
	databaseName     string
	initialised      bool
	withCaching      bool
	withCacheKey     Field
	verbose          bool
	timeout          int
	poolSize         int
	withEncryption   bool
	indices          []Field
	cacheInitialised bool
}

func (f Field) String() string {
	return string(f)
}
