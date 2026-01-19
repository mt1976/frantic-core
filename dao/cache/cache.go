// Package cache provides in-memory caching functionalities for data access objects (DAOs),
// including cache management, record storage, retrieval, and synchronization with the underlying database.
package cache

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"
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
	table := entities.GetStructType(data)
	enabled, exists := Cache.tablesActive[table]
	if !exists {
		return false
	}
	return enabled
}

func Disable(data any) error {
	table := entities.GetStructType(data)
	Cache.tablesActive[table] = false
	Cache.cache[table] = make(entrys)
	Cache.indices[table] = []entities.Field{}
	Cache.key[table] = ""
	return nil
}

func IsDisabled(data any) bool {
	table := entities.GetStructType(data)
	enabled, exists := Cache.tablesActive[table]
	if !exists {
		return true
	}
	return !enabled
}

func Activate(data any) error {
	table := entities.GetStructType(data)
	logHandler.InfoLogger.Printf("Activating Cache for Table [%v]", table)
	Cache.tablesActive[table] = true
	Cache.cache[table] = make(entrys)
	Cache.indices[table] = []entities.Field{}
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
	table := entities.GetStructType(data)
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
	logHandler.InfoLogger.Printf("Setting Cache Expiry for Table [%v] to %v", entities.GetStructType(data), duration)
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("set expiry", "", string(entities.GetStructType(data)))
	}

	Cache.expiry[entities.GetStructType(data)] = duration
	logHandler.InfoLogger.Printf("Cache Expiry for Table [%v] set to %v", entities.GetStructType(data), duration)
	return nil
}

func GetExpiry(data any) (time.Duration, error) {
	if !IsEnabled(data) {
		return 0, ce.ErrCacheNotEnabledWrapper("get expiry", "", string(entities.GetStructType(data)))
	}

	return Cache.expiry[entities.GetStructType(data)], nil
}

func RegisterKey(data any, key entities.Field) error {
	logHandler.InfoLogger.Printf("Adding Cache Key [%v] for Table [%v]", key.String(), entities.GetStructType(data))
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add key", key.String(), string(entities.GetStructType(data)))
	}

	Cache.key[entities.GetStructType(data)] = key
	//	godump.Dump(Cache)
	logHandler.InfoLogger.Printf("Cache Key [%v] added for Table [%v]", key.String(), entities.GetStructType(data))
	return nil
}

func RegisterIndex(data any, key entities.Field) error {

	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("add index", key.String(), string(entities.GetStructType(data)))
	}

	// Find the index in the list of indices
	indesList := Cache.indices[entities.GetStructType(data)]
	for _, existingIndex := range indesList {
		if existingIndex.String() == key.String() {
			logHandler.WarningLogger.Printf("index %v already exists for %v", key.String(), entities.GetStructType(data))
			return nil
		}
	}
	Cache.indices[entities.GetStructType(data)] = append(Cache.indices[entities.GetStructType(data)], key)

	return nil
}

func RemoveIndex(data any, key entities.Field) error {
	if !IsEnabled(data) {
		return ce.ErrCacheNotEnabledWrapper("remove index", key.String(), string(entities.GetStructType(data)))
	}

	// Find the index in the list of indices
	indesList := Cache.indices[entities.GetStructType(data)]
	for i, existingIndex := range indesList {
		if existingIndex.String() == key.String() {
			// Remove the index from the slice
			Cache.indices[entities.GetStructType(data)] = append(indesList[:i], indesList[i+1:]...)
			return nil
		}
	}
	logHandler.WarningLogger.Printf("index %v does not exist for %v", key.String(), entities.GetStructType(data))

	return nil
}

func AddEntry(data any) error {
	if data == nil {
		logHandler.WarningLogger.Println("Cannot add <nil> data to cache")
		return ce.ErrCacheNilDataWrapper("add")
	}
	table := entities.GetStructType(data)
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
	table := entities.GetStructType(data)

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
	table := entities.GetStructType(data)
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

func Get[T any](data T, key any) (T, error) {
	// Find and return the record from the cache
	var zero T
	table := entities.GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return zero, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return zero, ce.ErrCacheNoKeyDefinedWrapper("get", table.String())
	}

	record, exists := inMemoryCacheEntry[key]
	if !exists {
		return zero, ce.ErrCacheRecordNotFoundWrapper(table.String(), key)
	}

	targetType := reflect.TypeOf((*T)(nil)).Elem()
	converted, ok := coerceCacheValue[T](record.dataRecord, targetType)
	if !ok {
		return zero, fmt.Errorf("cache contains unexpected type for table %v: got %T, want %v", table.String(), record.dataRecord, targetType)
	}

	return converted, nil
}

func GetAll[T any](data T) ([]T, error) {
	// Get all records from the cache
	table := entities.GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return nil, ce.ErrCacheNoKeyDefinedWrapper("getall", table.String())
	}

	// Range through the cache and build a strongly-typed return slice.
	targetType := reflect.TypeOf((*T)(nil)).Elem()
	rtn := make([]T, 0, len(inMemoryCacheEntry))
	for _, record := range inMemoryCacheEntry {
		converted, ok := coerceCacheValue[T](record.dataRecord, targetType)
		if !ok {
			return nil, fmt.Errorf("cache contains unexpected type for table %v: got %T, want %v", table.String(), record.dataRecord, targetType)
		}
		rtn = append(rtn, converted)
	}

	return rtn, nil
}

func coerceCacheValue[T any](value any, targetType reflect.Type) (T, bool) {
	var zero T
	if value == nil {
		return zero, false
	}
	if v, ok := value.(T); ok {
		return v, true
	}

	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return zero, false
	}

	// Direct assign/convert.
	if rv.Type().AssignableTo(targetType) {
		v, ok := rv.Interface().(T)
		return v, ok
	}
	if rv.Type().ConvertibleTo(targetType) {
		converted := rv.Convert(targetType).Interface()
		v, ok := converted.(T)
		return v, ok
	}

	// Pointer/value bridging.
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		ev := rv.Elem()
		if ev.IsValid() {
			if ev.Type().AssignableTo(targetType) {
				v, ok := ev.Interface().(T)
				return v, ok
			}
			if ev.Type().ConvertibleTo(targetType) {
				converted := ev.Convert(targetType).Interface()
				v, ok := converted.(T)
				return v, ok
			}
		}
	}
	if targetType.Kind() == reflect.Ptr && rv.Type().AssignableTo(targetType.Elem()) {
		pv := reflect.New(targetType.Elem())
		pv.Elem().Set(rv)
		converted := pv.Interface()
		v, ok := converted.(T)
		return v, ok
	}

	return zero, false
}

func Count(data any) (int64, error) {
	// Get count of records from the cache
	table := entities.GetStructType(data)
	_, exists := Cache.cache[table]
	if !exists {
		return 0, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	return Cache.count[table], nil
}

func FindByKey[T any](data T, key any) (T, error) {
	// Find and return the record from the cache
	return Get(data, key)
}

func FindByIndex[T any](data T, index entities.Field, value any) ([]T, error) {
	// Find and return the record(s) from the cache by index
	table := entities.GetStructType(data)
	inMemoryCacheEntry, exists := Cache.cache[table]
	if !exists {
		return nil, ce.ErrCacheDoesNotExistWrapper(table.String())
	}

	if !isKeyRegistered(table) {
		logHandler.WarningLogger.Printf("No Key registered for Table [%v]", table)
		return nil, ce.ErrCacheNoKeyDefinedWrapper("findbyindex", table.String())
	}

	targetType := reflect.TypeOf((*T)(nil)).Elem()
	rtn := make([]T, 0)
	for _, record := range inMemoryCacheEntry {
		rv := reflect.ValueOf(record.dataRecord)
		if !rv.IsValid() {
			continue
		}
		if rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				continue
			}
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Struct {
			continue
		}
		fv := rv.FieldByName(index.String())
		if !fv.IsValid() {
			continue
		}
		if fv.Interface() != value {
			continue
		}

		converted, ok := coerceCacheValue[T](record.dataRecord, targetType)
		if !ok {
			return nil, fmt.Errorf("cache contains unexpected type for table %v: got %T, want %v", table.String(), record.dataRecord, targetType)
		}
		rtn = append(rtn, converted)
	}

	return rtn, nil
}

func RegisterSynchroniser(data any, synchroniser func(any) error) {
	Cache.synchroniser = make(map[entities.Table]func(any) error)
	table := entities.GetStructType(data)
	Cache.synchroniser[table] = synchroniser
	// Get the name of the function passed in
	funcname := runtime.FuncForPC(reflect.ValueOf(synchroniser).Pointer()).Name()
	logHandler.WarningLogger.Printf("Registered Function %v as Synchroniser for Table [%v]", funcname, table)
	//
}

func RegisterHydrator(data any, hydrator func() ([]any, error)) {
	if Cache.hydrator == nil {
		Cache.hydrator = make(map[entities.Table]func() ([]any, error))
	}
	if data == nil {
		logHandler.WarningLogger.Println("Cannot register hydrator for <nil> data")
		return
	}
	table := entities.GetStructType(data)
	Cache.hydrator[table] = hydrator
	// Get the name of the function passed in
	funcname := runtime.FuncForPC(reflect.ValueOf(hydrator).Pointer()).Name()
	logHandler.WarningLogger.Printf("Registered Function %v as Hydrator for Table [%v]", funcname, table)
}

func HydrateForType(data any) error {
	if data == nil {
		return ce.ErrCacheNilDataWrapper("hydrate")
	}
	table := entities.GetStructType(data)

	return hydrateCacheByTable(table)
}

func Hydrate(table entities.Table) error {
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

func hydrateCacheByTable(table entities.Table) error {
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

func SynchroniseForType(data any) error {
	table := entities.GetStructType(data)
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

func Synchronise(table entities.Table) error {
	return SynchroniseForType(table)
}

func SynchroniseAll() error {
	for table, synchroniserFunc := range Cache.synchroniser {
		if synchroniserFunc == nil {
			continue
		}
		if err := SynchroniseForType(table); err != nil {
			return err
		}
	}
	return nil
}

func isKeyRegistered(table entities.Table) bool {
	keyField, exists := Cache.key[table]
	if !exists || keyField.String() == "" {
		return false
	}
	return true
}
