package cache

import (
	"fmt"

	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
)

// Enable enables caching for a table, but does not initialise it.
func Enable(data any) error {
	Cache.tablesActive[GetStructType(data)] = false
	Cache.cache[GetStructType(data)] = make(entrys)
	Cache.indices[GetStructType(data)] = []database.Field{}
	Cache.key[GetStructType(data)] = ""
	return nil
}

func IsEnabled(data any) bool {
	table := GetStructType(data)
	enabled, exists := Cache.tablesActive[table]
	if !exists {
		return false
	}
	return enabled
}

func Disable(data any) error {
	table := GetStructType(data)
	Cache.tablesActive[table] = false
	Cache.cache[table] = make(entrys)
	Cache.indices[table] = []database.Field{}
	Cache.key[table] = ""
	return nil
}

func IsDisabled(data any) bool {
	table := GetStructType(data)
	enabled, exists := Cache.tablesActive[table]
	if !exists {
		return true
	}
	return !enabled
}

func Initialise(data any) error {
	table := GetStructType(data)
	Cache.tablesActive[table] = true
	return nil
}

func IsInitialised(data any) bool {
	table := GetStructType(data)
	enabled, exists := Cache.tablesActive[table]
	if !exists {
		return false
	}
	return enabled
}

func DeInitialise(data any) error {
	return Disable(data)
}

func IsDeInitialised(data any) bool {
	return IsDisabled(data)
}

func AddKey(data any, key database.Field) error {

	if !IsEnabled(data) {
		return fmt.Errorf("cannot add key %v to %v - caching not enabled", key.String(), GetStructType(data))
	}

	Cache.key[GetStructType(data)] = key
	return nil
}

func AddIndex(data any, key database.Field) error {

	if !IsEnabled(data) {
		return fmt.Errorf("cannot add index %v to %v - caching not enabled", key.String(), GetStructType(data))
	}

	// Find the index in the list of indices
	indesList := Cache.indices[GetStructType(data)]
	for _, existingIndex := range indesList {
		if existingIndex.String() == key.String() {
			logHandler.WarningLogger.Printf("index %v already exists for %v", key.String(), GetStructType(data))
			return nil
		}
	}
	Cache.indices[GetStructType(data)] = append(Cache.indices[GetStructType(data)], key)

	return nil
}

func RemoveIndex(data any, key database.Field) error {
	// TODO: Write Code

	return nil
}

func Add(data any) error {
	// TODO: Write Code

	return nil
}

func Load(data []any) error {
	// TODO: Write Code

	return nil
}

func Remove(data any) error {
	// TODO: Write Code

	return nil
}

func Update(data any) error {
	// TODO: Write Code

	return nil
}

func Get(data any, key database.Field) (any, error) {
	// TODO: Write Code
	var rtn any
	return rtn, nil
}

func GetAll(data any) ([]any, error) {
	var rtn []any
	return rtn, nil
}

func Count(data any) (int, error) {
	//TODO: Write Code
	return 0, nil
}
