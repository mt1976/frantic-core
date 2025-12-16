package dao

import (
	"os"
	"reflect"
	"strconv"
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

type Int struct {
	Value string
}

type Float struct {
	Value string
}

type Bool struct {
	Value string
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

func (i *Int) Set(in int) Int {
	return Int{Value: strconv.Itoa(in)}
}

func (i *Int) Int() int {
	val, err := strconv.Atoi(i.Value)
	if err != nil {
		logHandler.ErrorLogger.Panic(commonErrors.WrapInvalidTypeError("Int", i.Value, "int"))
	}
	return val
}

func (i *Int) Get() int {
	return i.Int()
}

func (i *Int) String() string {
	return i.Value
}

func (f *Float) Set(in float64) Float {
	return Float{Value: strconv.FormatFloat(in, 'f', -1, 64)}
}

func (f *Float) Float() float64 {
	val, err := strconv.ParseFloat(f.Value, 64)
	if err != nil {
		logHandler.ErrorLogger.Panic(commonErrors.WrapInvalidTypeError("Float", f.Value, "float64"))
	}
	return val
}

func (f *Float) Get() float64 {
	return f.Float()
}

func (f *Float) String() string {
	return f.Value
}

func (b *Bool) Set(in bool) Bool {
	if in {
		return Bool{Value: "true"}
	} else {
		return Bool{Value: "false"}
	}
}

func (b *Bool) Bool() bool {
	return b.Value == "true"
}

func (b *Bool) Get() bool {
	return b.Bool()
}

func (b *Bool) String() string {
	return b.Value
}

func (b *Bool) IsTrue() bool {
	return b.Bool()
}

func (b *Bool) IsFalse() bool {
	return !b.Bool()
}
