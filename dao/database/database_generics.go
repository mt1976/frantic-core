package database

import (
	"fmt"
	"reflect"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"
	"github.com/mt1976/frantic-core/logHandler"
)

// GetTyped retrieves a single record from the database and returns it as a concrete type.
//
// NOTE: Go does not currently allow methods with type parameters, so these helpers are
// package-level functions.
//
// T is expected to be a struct type (not a pointer). Storm expects a pointer to a struct
// destination; if you call this with a pointer type, it will return an invalid-type error.
func GetTyped[T any](db *DB, field entities.Field, value any) (T, error) {
	var zero T
	var record T

	if reflect.TypeOf(record) != nil && reflect.TypeOf(record).Kind() == reflect.Ptr {
		return zero, commonErrors.ErrInvalidTypeWrapper("GetTyped", fmt.Sprintf("%T", record), "non-pointer struct")
	}

	logHandler.DatabaseLogger.Printf("[GET_TYPED]<%v> (%+v=%+v) [%v.db]", GetStructType(record), field.String(), value, db.Name)
	if err := db.connection.One(field.String(), value, &record); err != nil {
		return zero, err
	}
	return record, nil
}

// GetAllTyped retrieves all records for type T and returns a typed slice.
//
// NOTE: T is expected to be a struct type (not a pointer).
func GetAllTyped[T any](db *DB, options ...func(*index.Options)) ([]T, error) {
	var record T
	if reflect.TypeOf(record) != nil && reflect.TypeOf(record).Kind() == reflect.Ptr {
		return nil, commonErrors.ErrInvalidTypeWrapper("GetAllTyped", fmt.Sprintf("%T", record), "non-pointer struct")
	}

	logHandler.DatabaseLogger.Printf("[GETALL_TYPED]<%v> [%v.db]", GetStructType(record), db.Name)
	result := []T{}
	if err := db.connection.All(&result, options...); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllWhereTyped retrieves all matching records for type T filtered by field/value.
//
// NOTE: T is expected to be a struct type (not a pointer).
func GetAllWhereTyped[T any](db *DB, field entities.Field, value any) ([]T, error) {
	var record T
	if reflect.TypeOf(record) != nil && reflect.TypeOf(record).Kind() == reflect.Ptr {
		return nil, commonErrors.ErrInvalidTypeWrapper("GetAllWhereTyped", fmt.Sprintf("%T", record), "non-pointer struct")
	}

	if err := IsValidFieldInStruct(field, record); err != nil {
		return nil, err
	}
	if err := IsValidTypeForField(field, value, record); err != nil {
		return nil, err
	}

	logHandler.DatabaseLogger.Printf("[GETALLWHERE_TYPED]<%v> WHERE (%+v=%+v) [%v.db]", GetStructType(record), field.String(), value, db.Name)
	result := []T{}
	query := db.connection.Select(q.Eq(field.String(), value))
	if err := query.Find(&result); err != nil {
		if err == storm.ErrNotFound {
			return []T{}, nil
		}
		return nil, err
	}
	return result, nil
}
