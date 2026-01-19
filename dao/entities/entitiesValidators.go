package entities

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

func GetStructType(data any) Table {
	logHandler.TraceLogger.Printf("Resolving Struct Type for data: %v", data)
	rtnType := reflect.TypeOf(data).String()
	base := rtnType

	//base := rtnType
	// If the type is a pointer, get the underlying type
	if strings.Contains(rtnType, "*") {
		rtnType = reflect.TypeOf(data).Elem().String()
	}
	// If the type is a struct, get the name of the struct
	if strings.Contains(rtnType, ".") {
		rtnType = strings.Split(rtnType, ".")[1]
	}
	logHandler.TraceLogger.Printf("{TYPE} Resolved Struct Type: %v (base: %v)", rtnType, base)

	return Table(rtnType)
}

func IsValidFieldInStruct(fromField Field, data any) error {
	// Normalise the type: unwrap pointers, and if it's a slice/array, use the element type.
	if data == nil {
		logHandler.ErrorLogger.Printf("Cannot validate field '%v' on <nil> data", fromField.String())
		return commonErrors.ErrInvalidFieldWrapper(fromField.String())
	}

	t := reflect.TypeOf(data)
	for t.Kind() == reflect.Ptr {
		// Unwrap pointer types
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		logHandler.TraceLogger.Printf("Validating Field '%v' against slice/array element type '%v'", fromField.String(), t.Elem())
		t = t.Elem()
		for t.Kind() == reflect.Ptr {
			// Handle slices of pointers to structs
			t = t.Elem()
		}
	}

	logHandler.TraceLogger.Printf("Validating Field '%v' in Struct type '%v'", fromField.String(), t.Name())

	if t.Kind() != reflect.Struct {
		logHandler.ErrorLogger.Printf("Type '%v' is not a struct; cannot validate field '%v'", t, fromField.String())
		return commonErrors.ErrInvalidFieldWrapper(fromField.String())
	}

	if _, isValidField := t.FieldByName(fromField.String()); !isValidField {
		logHandler.ErrorLogger.Printf("Field '%v' not found in struct '%v'", fromField.String(), t.Name())
		logHandler.ErrorLogger.Println(commonErrors.ErrInvalidFieldWrapper(fromField.String()))
		return commonErrors.ErrInvalidFieldWrapper(fromField.String())
	}
	logHandler.TraceLogger.Printf("Field '%v' is valid in struct '%v'", fromField.String(), t.Name())
	return nil
}

func IsValidTypeForField(field Field, data, forStruct any) error {
	if forStruct == nil {
		logHandler.ErrorLogger.Printf("Cannot validate type for field '%v' on <nil> struct", field.String())
		return commonErrors.ErrInvalidFieldWrapper(field.String())
	}

	// Normalise the type of forStruct: unwrap pointers, and if it's a slice/array, use the element type.
	st := reflect.TypeOf(forStruct)
	for st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	if st.Kind() == reflect.Slice || st.Kind() == reflect.Array {
		st = st.Elem()
		for st.Kind() == reflect.Ptr {
			st = st.Elem()
		}
	}

	if st.Kind() != reflect.Struct {
		logHandler.ErrorLogger.Printf("Type '%v' is not a struct; cannot validate type for field '%v'", st, field.String())
		return commonErrors.ErrInvalidFieldWrapper(field.String())
	}

	structField, found := st.FieldByName(field.String())
	if !found {
		logHandler.ErrorLogger.Printf("Field '%v' not found in struct '%v' when validating type", field.String(), st.Name())
		return commonErrors.ErrInvalidFieldWrapper(field.String())
	}

	dataType := "<nil>"
	if data != nil {
		dataType = reflect.TypeOf(data).String()
	}
	structFieldType := structField.Type.String()
	if dataType != structFieldType {
		logHandler.ErrorLogger.Printf("Type mismatch for field '%v': expected '%v', got '%v'", field.String(), structFieldType, dataType)
		return commonErrors.ErrInvalidTypeWrapper(field.String(), dataType, structFieldType)
	}
	logHandler.TraceLogger.Printf("Type for field '%v' is valid: expected '%v', got '%v'", field.String(), structFieldType, dataType)
	return nil
}
