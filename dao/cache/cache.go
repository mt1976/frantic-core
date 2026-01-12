package cache

import (
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
)

// Enable enables caching for a table, but does not initialise it.
func Enable(data any) error {
	Cache.tables[GetStructType(data)] = false
	return nil
}

func Disable(data any) error {
	table := GetStructType(data)
	delete(Cache.tables, table)
	delete(Cache.indices, table)
	delete(Cache.keys, table)
	delete(Cache.cache, table)
	return nil
}

func IsCaching(data any) (bool, error) {
	// TODO: Write Code
	return true, nil
}

func AddKey(data any, key database.Field) error {

	ok, err := IsCaching(data)

	if err != nil {
		return err
	}

	if ok {
		Cache.keys[GetStructType(data)] = key
	} else {
		logHandler.WarningLogger.Printf("cannot add key %v to %v", key.String(), GetStructType(data))
	}
	return nil
}

func AddIndex(data any, key database.Field) error {

	ok, err := IsCaching(data)

	if err != nil {
		return err
	}

	if ok {
		//TODO : Prevent duplicates
		Cache.indices[GetStructType(data)] = append(Cache.indices[GetStructType(data)], key)
	} else {
		logHandler.WarningLogger.Printf("cannot add key %v to %v", key.String(), GetStructType(data))
	}

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
