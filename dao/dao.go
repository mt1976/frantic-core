package dao

import (
	"os"
	"reflect"
	"strings"

	"github.com/asdine/storm/v3"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var name = "DAO"
var DBVersion = 1
var DB *storm.DB
var DBName string = "default"

// StormBool is a boolean type that can be marshalled to and from a string, this has been created as Storm does not support boolean types properly
type StormBool struct {
	State string
}

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

func IsValidFieldInStruct(fromField string, data any) error {
	_, isValidField := reflect.TypeOf(data).FieldByName(fromField)
	if !isValidField {
		logHandler.ErrorLogger.Panic(commonErrors.WrapInvalidFieldError(fromField))
		return commonErrors.WrapInvalidFieldError(fromField)
	}
	return nil
}

func IsValidTypeForField(field string, data, forStruct any) error {
	dataType := reflect.TypeOf(data).String()
	structField, found := reflect.TypeOf(forStruct).FieldByName(field)
	if !found {
		return commonErrors.WrapInvalidFieldError(field)
	}
	structFieldType := structField.Type.String()
	if dataType != structFieldType {
		return commonErrors.WrapInvalidTypeError(field, dataType, structFieldType)
	}
	return nil
}

func CheckDAOReadyState(table string, action audit.Action, isDaoReady bool) {
	if !isDaoReady {
		err := commonErrors.WrapDAONotInitialisedError(table, action.Description())
		logHandler.ErrorLogger.Panic(err)
	}
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

func (sb *StormBool) Set(b bool) {
	if b {
		sb.State = "true"
	} else {
		sb.State = "false"
	}
}

func (sb *StormBool) Bool() bool {
	return sb.State == "true"
}

func (sb *StormBool) String() string {
	return sb.State
}

func (sb *StormBool) IsTrue() bool {
	return sb.Bool()
}

func (sb *StormBool) IsFalse() bool {
	return !sb.Bool()
}
