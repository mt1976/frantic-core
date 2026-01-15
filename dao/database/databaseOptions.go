package database

import (
	"github.com/mt1976/frantic-core/dao/fields"
	"github.com/mt1976/frantic-core/logHandler"
)

// connectionConfig holds the configuration options for database connections
type connectionConfig struct {
	withCaching      bool
	withCacheKey     fields.Field
	Verbose          bool
	timeout          int
	poolSize         int
	nameSpace        string
	withEncryption   bool
	indices          []fields.Field
	cacheInitialised bool
}

// Option is a function that configures the database connection
type Option func(*connectionConfig)

// WithCaching enables or disables caching for the database connection
func WithCaching(enabled bool) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithCaching set to %v", enabled)
	return func(c *connectionConfig) {
		c.withCaching = enabled
	}
}

// WithCacheKey sets the cache key fields.Fields.Field for the database connection
func WithCacheKey(field fields.Field) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithCacheKey set to %v", field.String())
	return func(c *connectionConfig) {
		c.withCacheKey = field
	}
}

// WithVerbose enables or disables verbose logging for the database connection
func WithVerbose(enabled bool) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithVerbose set to %v", enabled)
	return func(c *connectionConfig) {
		c.Verbose = enabled
	}
}

// WithTimeout sets the connection timeout in seconds
func WithTimeout(seconds int) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithTimeout set to %d", seconds)
	return func(c *connectionConfig) {
		c.timeout = seconds
	}
}

// WithPoolSize sets the maximum connection pool size
func WithPoolSize(size int) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithPoolSize set to %d", size)
	return func(c *connectionConfig) {
		c.poolSize = size
	}
}

// WithNameSpace sets the namespace (database name) for the connection
func WithNameSpace(name string) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithNameSpace set to %s", name)
	return func(c *connectionConfig) {
		c.nameSpace = name
	}
}

// WithEncryption enables or disables encryption for the database connection
// By default, encryption is disabled.
// Reserved for Future use.
func WithEncryption(enabled bool) Option {
	logHandler.DatabaseLogger.Printf("[CON]{OPTION} WithEncryption set to %v", enabled)
	logHandler.WarningLogger.Printf("WithEncryption is not yet implemented")
	return func(c *connectionConfig) {
		c.withEncryption = enabled
	}
}
