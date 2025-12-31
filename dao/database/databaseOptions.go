package database

import (
	"github.com/mt1976/frantic-core/dao"
)

// connectionConfig holds the configuration options for database connections
type connectionConfig struct {
	withCaching      bool
	withCacheKey     dao.Field
	verbose          bool
	timeout          int
	poolSize         int
	nameSpace        string
	withEncryption   bool
	indices          []Field
	cacheInitialised bool
}

// Option is a function that configures the database connection
type Option func(*connectionConfig)

// WithCaching enables or disables caching for the database connection
func WithCaching(enabled bool) Option {
	return func(c *connectionConfig) {
		c.withCaching = enabled
	}
}

func WithCacheKey(field dao.Field) Option {
	return func(c *connectionConfig) {
		c.withCacheKey = field
	}
}

// WithVerbose enables or disables verbose logging for the database connection
func WithVerbose(enabled bool) Option {
	return func(c *connectionConfig) {
		c.verbose = enabled
	}
}

// WithTimeout sets the connection timeout in seconds
func WithTimeout(seconds int) Option {
	return func(c *connectionConfig) {
		c.timeout = seconds
	}
}

// WithPoolSize sets the maximum connection pool size
func WithPoolSize(size int) Option {
	return func(c *connectionConfig) {
		c.poolSize = size
	}
}

func WithNameSpace(name string) Option {
	return func(c *connectionConfig) {
		c.nameSpace = name
	}
}

func WithEncryption(enabled bool) Option {
	return func(c *connectionConfig) {
		c.withEncryption = enabled
	}
}

func WithIndices(indices []Field) Option {
	return func(c *connectionConfig) {
		c.indices = indices
	}
}

func WithIndex(field Field) Option {
	return func(c *connectionConfig) {
		c.indices = append(c.indices, field)
	}
}
