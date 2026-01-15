package dao

import (
	"os"
	"strings"

	"github.com/asdine/storm/v3"
	"github.com/mt1976/frantic-core/commonConfig"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var name = "DAO"
var DBVersion = 1
var DB *storm.DB
var DBName string = "default"

func Initialise(cfg *commonConfig.Settings) error {
	clock := timing.Start(name, "Initialise", "")
	logHandler.InfoLogger.Printf("[%v] Initialising...", strings.ToUpper(name))

	DBVersion = cfg.GetDatabase_Version()
	DBName = cfg.GetDatabase_Name()

	logHandler.InfoLogger.Printf("[%v] Initialised", strings.ToUpper(name))
	clock.Stop(1)
	return nil
}

func GetDBNameFromPath(t string) string {
	dbName := t
	// split dbName on "/"
	dbNameArr := strings.Split(dbName, string(os.PathSeparator))
	noparts := len(dbNameArr)
	dbName = dbNameArr[noparts-1]
	logHandler.InfoLogger.Printf("dbName: %v\n", dbName)
	return dbName
}

// func IsValidFieldInStruct(fromField database.Field, data any) error {
// 	_, isValidField := reflect.TypeOf(data).FieldByName(fromField.String())
// 	if !isValidField {
// 		logHandler.ErrorLogger.Println(commonErrors.WrapInvalidFieldError(fromField.String()))
// 		return commonErrors.WrapInvalidFieldError(fromField.String())
// 	}
// 	return nil
// }

// func IsValidTypeForField(field database.Field, data, forStruct any) error {
// 	dataType := reflect.TypeOf(data).String()
// 	structField, found := reflect.TypeOf(forStruct).FieldByName(field.String())
// 	if !found {
// 		return commonErrors.WrapInvalidFieldError(field.String())
// 	}
// 	structFieldType := structField.Type.String()
// 	if dataType != structFieldType {
// 		return commonErrors.WrapInvalidTypeError(field.String(), dataType, structFieldType)
// 	}
// 	return nil
// }

func CheckDAOReadyState(table string, action audit.Action, isDaoReady bool) {
	if !isDaoReady {
		err := ce.ErrDAONotInitialisedWrapper(table, action.Description())
		logHandler.ErrorLogger.Panic(err)
	}
}

// func GetStructType(data any) string {
// 	rtnType := reflect.TypeOf(data).String()
// 	// If the type is a pointer, get the underlying type
// 	if strings.Contains(rtnType, "*") {
// 		rtnType = reflect.TypeOf(data).Elem().String()
// 	}
// 	// If the type is a struct, get the name of the struct
// 	if strings.Contains(rtnType, ".") {
// 		rtnType = strings.Split(rtnType, ".")[1]
// 	}
// 	return rtnType
// }
