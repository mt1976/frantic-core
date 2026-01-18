package cache

import (
	"reflect"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/fields"
	"github.com/mt1976/frantic-core/logHandler"
)

// IsEnabled checks if the cache for the given data type is enabled.
//
// Parameters:
//   - data: The data type to check the cache status for.
//
// Returns:
//   - bool: True if the cache is enabled for the given data type; otherwise, false.
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

func Activate(data any) error {
	table := GetStructType(data)
	logHandler.InfoLogger.Printf("Activating Cache for Table [%v]", table)
	Cache.tablesActive[table] = true
	Cache.cache[table] = make(entrys)
	Cache.indices[table] = []fields.Field{}
	Cache.key[table] = ""
	Cache.count[table] = 0
	Cache.expiry[table] = defaultCacheExpiry
	Cache.synchroniser[table] = nil
	Cache.hydrator[table] = nil
	//	godump.Dump(Cache)
	logHandler.InfoLogger.Printf("Cache for Table [%v] Activated", table)
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

func RegisterExpiry(data any, duration time.Duration) error {
	logHandler.InfoLogger.Printf("Setting Cache Expiry for Table [%v] to %v", GetStructType(data), duration)
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("set expiry", "", string(GetStructType(data)))
	}

	Cache.expiry[GetStructType(data)] = duration
	logHandler.InfoLogger.Printf("Cache Expiry for Table [%v] set to %v", GetStructType(data), duration)
	return nil
}

func GetExpiry(data any) (time.Duration, error) {
	if !IsEnabled(data) {
		return 0, ce.ErrCacheNotEnabledWrapper("get expiry", "", string(GetStructType(data)))
	}

	return Cache.expiry[GetStructType(data)], nil
}

func RegisterKey(data any, key fields.Field) error {
	logHandler.InfoLogger.Printf("Adding Cache Key [%v] for Table [%v]", key.String(), GetStructType(data))
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add key", key.String(), string(GetStructType(data)))
	}

	Cache.key[GetStructType(data)] = key
	//	godump.Dump(Cache)
	logHandler.InfoLogger.Printf("Cache Key [%v] added for Table [%v]", key.String(), GetStructType(data))
	return nil
}

func RegisterIndex(data any, key fields.Field) error {

	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add index", key.String(), string(GetStructType(data)))
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
		return ce.ErrCacheNotEnabledWrapper("remove index", key.String(), string(GetStructType(data)))
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

func AddEntry(data any) error {
	if data == nil {
		logHandler.WarningLogger.Println("Cannot add <nil> data to cache")
		return ce.ErrCacheNilDataWrapper("add")
	}
	table := GetStructType(data)
	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return ce.ErrCacheNoKeyDefinedWrapper("add", table.String())
	}
	//logHandler.InfoLogger.Printf("Adding Cache Entry for Table [%v]", table)
	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return ce.ErrCacheNoKeyDefinedWrapper("add", table.String())
	}

	// Lets get the key value and build the cache entry
	// Get the key value, by using reflection to get the field value
	key := reflect.ValueOf(data).FieldByName(keyField.String()).Interface()

	// Get Cache Expiry
	expiryDuration, err := GetExpiry(data)
	if err != nil {
		expiryDuration = defaultCacheExpiry // Default to 30 years
	}
	// Add the record to the cache
	// check if the table cache exists
	_, exists = Cache.cache[table]
	if !exists {
		Cache.cache[table] = make(entrys)
		Cache.tablesActive[table] = true
		record := dataCache{cacheTimestamp: time.Now().Add(expiryDuration), dataRecord: data}
		Cache.cache[table][key] = record
		Cache.count[table] = 1
		Cache.updated = time.Now()
		return nil
	}
	record := dataCache{cacheTimestamp: time.Now().Add(expiryDuration), dataRecord: data}
	Cache.cache[table][key] = record
	Cache.count[table]++
	Cache.updated = time.Now()
	logHandler.InfoLogger.Printf("Cache Entry for Table [%v] added with Key [%v], expiry [%v] %v", table, key, record.cacheTimestamp.Format(time.RFC3339Nano), humanize.Time(record.cacheTimestamp))
	return nil
}

func Load(data []any) error {
	// Range through the data and add each record to the cache
	for _, record := range data {
		err := AddEntry(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveEntry(data any) error {
	// Fine and remove the record from the cache
	table := GetStructType(data)

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table.String())
	}

	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table.String())
	}

	// Get the key value, by using reflection to get the field value
	key := reflect.ValueOf(data).FieldByName(keyField.String()).Interface()

	delete(Cache.cache[table], key)
	Cache.count[table]--
	Cache.updated = time.Now()

	return nil
}

func RemoveByKey(data any, key any) error {
	// Find and remove the record from the cache
	table := GetStructType(data)
	_, exists := Cache.key[table]
	if !exists {
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return ce.ErrCacheNoKeyDefinedWrapper("remove", table.String())
	}

	delete(Cache.cache[table], key)
	Cache.count[table]--
	Cache.updated = time.Now()
	return nil
}

func Update(data any) error {
	return AddEntry(data)
}

func Get(data any, key any) (any, error) {
	// Find and return the record from the cache
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return nil, ce.ErrCacheNoKeyDefinedWrapper("get", table.String())
	}

	record, exists := inMemoryCacheEntry[key]
	if !exists {
		return nil, ce.ErrCacheRecordNotFoundWrapper(table.String(), key)
	}

	return record.dataRecord, nil
}

func GetAll(data any) ([]any, error) {
	// Get all records from the cache
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return nil, ce.ErrCacheNoKeyDefinedWrapper("getall", table.String())
	}

	// Range through the cache and build the return slice
	var rtn []any
	for _, record := range inMemoryCacheEntry {
		rtn = append(rtn, record.dataRecord)
	}
	return rtn, nil
}

func Count(data any) (int64, error) {
	// Get count of records from the cache
	table := GetStructType(data)
	_, exists := Cache.cache[table]
	if !exists {
		return 0, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	return Cache.count[table], nil
}

func FindByKey(data any, key any) (any, error) {
	// Find and return the record from the cache
	return Get(data, key)
}

func FindByIndex(data any, index fields.Field, value any) ([]any, error) {
	// Find and return the record(s) from the cache by index
	table := GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return nil, ce.ErrCacheNoKeyDefinedWrapper("findbyindex", table.String())
	}

	var rtn []any
	for _, record := range inMemoryCacheEntry {
		v := reflect.ValueOf(record.dataRecord).FieldByName(index.String()).Interface()
		if v == value {
			rtn = append(rtn, record.dataRecord)
		}
	}

	return rtn, nil
}

func RegisterSynchroniser(data any, synchroniser func(any) error) {
	Cache.synchroniser = make(map[entity]func(any) error)
	table := GetStructType(data)
	Cache.synchroniser[table] = synchroniser
	// Get the name of the function passed in
	funcname := runtime.FuncForPC(reflect.ValueOf(synchroniser).Pointer()).Name()
	logHandler.WarningLogger.Printf("Registered Function %v as Synchroniser for Table [%v]", funcname, table)
	//
}

func RegisterHydrator(data any, hydrator func() ([]any, error)) {
	if Cache.hydrator == nil {
		Cache.hydrator = make(map[entity]func() ([]any, error))
	}
	if data == nil {
		logHandler.WarningLogger.Println("Cannot register hydrator for <nil> data")
		return
	}
	table := GetStructType(data)
	Cache.hydrator[table] = hydrator
	// Get the name of the function passed in
	funcname := runtime.FuncForPC(reflect.ValueOf(hydrator).Pointer()).Name()
	logHandler.WarningLogger.Printf("Registered Function %v as Hydrator for Table [%v]", funcname, table)
}

func Hydrate(data any) error {
	if data == nil {
		return ce.ErrCacheNilDataWrapper("hydrate")
	}
	table := GetStructType(data)

	return hydrateCacheByTable(table)
}

func HydrateAll() error {
	if Cache.hydrator == nil {
		return nil
	}

	for table, hydratorFunc := range Cache.hydrator {
		if hydratorFunc == nil {
			continue
		}
		if err := hydrateCacheByTable(table); err != nil {
			return err
		}
	}

	return nil
}

func hydrateCacheByTable(table entity) error {
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return ce.ErrCacheNoKeyDefinedWrapper("hydrate", table.String())
	}

	hydratorFunc, exists := Cache.hydrator[table]
	if !exists || hydratorFunc == nil {
		return ce.ErrCacheNoHydratorDefinedWrapper(table.String())
	}

	count := len(inMemoryCacheEntry)
	countIndex := 0
	records, err := hydratorFunc()
	if err != nil {
		return err
	}
	for _, record := range records {
		err := AddEntry(record)
		if err != nil {
			return err
		}
		countIndex++
	}

	logHandler.InfoLogger.Printf("Cache for Table [%v] hydrated (%d/%d)", table, countIndex, count)
	return nil
}

func Synchronise(data any) error {
	table := GetStructType(data)
	//	logHandler.InfoLogger.Printf("Flushing Cache for Table [%v]", table)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return ce.ErrCacheNoKeyDefinedWrapper("synchronise", table.String())
	}

	synchroniserFunc, exists := Cache.synchroniser[table]
	if !exists {
		return ce.ErrCacheNoSynchroniserDefinedWrapper(table.String())
	}
	count := len(inMemoryCacheEntry)
	countIndex := 0
	for _, record := range inMemoryCacheEntry {
		err := synchroniserFunc(record.dataRecord)
		if err != nil {
			return err
		}
		countIndex++
	}

	logHandler.InfoLogger.Printf("Cache for Table [%v] synchronised (%d/%d)", table, countIndex, count)
	return nil
}

func SynchroniseAll() error {
	for table, synchroniserFunc := range Cache.synchroniser {
		if synchroniserFunc == nil {
			continue
		}
		if err := Synchronise(table); err != nil {
			return err
		}
	}
	return nil
}

func isKeyRegistered(table entity) bool {
	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return false
	}
	return true
}
