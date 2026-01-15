package cache

import (
	"reflect"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/fields"
	"github.com/mt1976/frantic-core/logHandler"
)

// Enable enables caching for a table, but does not initialise it.
func Enable(data any) error {
	Cache.tablesActive[GetStructType(data)] = false
	Cache.cache[GetStructType(data)] = make(entrys)
	Cache.indices[GetStructType(data)] = []fields.Field{}
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
	Cache.indices[table] = []fields.Field{}
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

func AddKey(data any, key fields.Field) error {

	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add key", key.String(), GetStructType(data))
	}

	Cache.key[GetStructType(data)] = key
	return nil
}

func AddIndex(data any, key fields.Field) error {

	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add index", key.String(), GetStructType(data))
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

func RemoveIndex(data any, key fields.Field) error {
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("remove index", key.String(), GetStructType(data))
	}

	// Find the index in the list of indices
	indesList := Cache.indices[GetStructType(data)]
	for i, existingIndex := range indesList {
		if existingIndex.String() == key.String() {
			// Remove the index from the slice
			Cache.indices[GetStructType(data)] = append(indesList[:i], indesList[i+1:]...)
			return nil
		}
	}
	logHandler.WarningLogger.Printf("index %v does not exist for %v", key.String(), GetStructType(data))

	return nil
}

func Add(data any) error {
	table := GetStructType(data)
	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return ce.ErrCacheNoKeyDefinedWrapper("add", table)
	}

	// Lets get the key value and build the cache entry
	// Get the key value, by using reflection to get the field value
	key := reflect.ValueOf(data).FieldByName(keyField.String()).Interface()

	Cache.cache[table][key] = data

	return nil
}

func Load(data []any) error {
	// Range through the data and add each record to the cache
	for _, record := range data {
		err := Add(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func Remove(data any) error {
	// Fine and remove the record from the cache
	table := GetStructType(data)
	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table)
	}

	// Get the key value, by using reflection to get the field value
	key := reflect.ValueOf(data).FieldByName(keyField.String()).Interface()

	delete(Cache.cache[table], key)

	return nil
}

func RemoveByKey(data any, key any) error {
	// Find and remove the record from the cache
	table := GetStructType(data)
	_, exists := Cache.key[table]
	if !exists {
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table)
	}

	delete(Cache.cache[table], key)

	return nil
}

func Update(data any) error {
	return Add(data)
}

func Get(data any, key any) (any, error) {
	// Find and return the record from the cache
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table)
	}

	record, exists := inMemoryCacheEntry[key]
	if !exists {
		return nil, ce.ErrCacheRecordNotFoundWrapper(table, key)
	}

	return record, nil
}

func GetAll(data any) ([]any, error) {
	// Get all records from the cache
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table)
	}

	// Range through the cache and build the return slice
	var rtn []any
	for _, record := range inMemoryCacheEntry {
		rtn = append(rtn, record)
	}
	return rtn, nil
}

func Count(data any) (int, error) {
	// Get count of records from the cache
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return 0, ce.ErrCacheDoesNotExistWrapper(table)
	}

	return len(inMemoryCacheEntry), nil
}
