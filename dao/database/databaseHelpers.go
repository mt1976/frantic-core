package database

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
)

// GetFunctionName returns the name of the function passed as an argument
// It uses reflection to obtain the function's program counter and retrieves its name.
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func GetStructType(data any) string {
	rtnType := reflect.TypeOf(data).String()
	// If the type is a pointer, get the underlying type
	if strings.Contains(rtnType, "*") {
		rtnType = reflect.TypeOf(data).Elem().String()
	}
	// If the type is a struct, get the name of the struct
	if strings.Contains(rtnType, ".") {
		rtnType = strings.Split(rtnType, ".")[1]
	}
	return rtnType
}

func IsValidFieldInStruct(fromField Field, data any) error {
	// Check if the field exists in the struct
	// if the data parameter passed in is an array or slice, get the element type, also need to cope if the slice is empty
	if reflect.TypeOf(data).Kind() == reflect.Slice || reflect.TypeOf(data).Kind() == reflect.Array {
		if reflect.ValueOf(data).Len() == 0 {
			return commonErrors.WrapInvalidFieldError(fromField.String())
		}
		data = reflect.ValueOf(data).Index(0).Interface()
	}
	_, isValidField := reflect.TypeOf(data).FieldByName(fromField.String())
	if !isValidField {
		logHandler.ErrorLogger.Println(commonErrors.WrapInvalidFieldError(fromField.String()))
		return commonErrors.WrapInvalidFieldError(fromField.String())
	}
	return nil
}

func IsValidTypeForField(field Field, data, forStruct any) error {
	dataType := reflect.TypeOf(data).String()
	structField, found := reflect.TypeOf(forStruct).FieldByName(field.String())
	if !found {
		return commonErrors.WrapInvalidFieldError(field.String())
	}
	structFieldType := structField.Type.String()
	if dataType != structFieldType {
		return commonErrors.WrapInvalidTypeError(field.String(), dataType, structFieldType)
	}
	return nil
}
