package dao

import (
	"fmt"
	"os"
	"strconv"
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

type Field string

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
	i.Value = strconv.Itoa(in)
	return *i
}

func (i *Int) Int() int {
	val, err := strconv.Atoi(i.Value)
	if i.Value == "" {
		return 0
	}
	if err != nil {
		logHandler.ErrorLogger.Panic(ce.ErrInvalidTypeWrapper("Int", i.Value, "int"))
	}
	return val
}

func (i *Int) Get() int {
	return i.Int()
}

func (i *Int) String() string {
	return i.Value
}

func (i *Int) Equals(other Int) bool {
	return i.Value == other.Value
}

func (i *Int) LessThan(other Int) bool {
	return i.Int() < other.Int()
}

func (i *Int) LessThanOrEqual(other Int) bool {
	return i.Int() <= other.Int()
}

func (i *Int) GreaterThan(other Int) bool {
	return i.Int() > other.Int()
}

func (i *Int) GreaterThanOrEqual(other Int) bool {
	return i.Int() >= other.Int()
}

func (i *Int) Add(other Int) Int {
	sum := i.Int() + other.Int()
	return i.Set(sum)
}

func (i *Int) Subtract(other Int) Int {
	diff := i.Int() - other.Int()
	return i.Set(diff)
}

func (i *Int) MultiplyBy(other Int) Int {
	prod := i.Int() * other.Int()
	return i.Set(prod)
}

func (i *Int) DivideBy(other Int) Int {
	if other.Int() == 0 {
		logHandler.ErrorLogger.Panic(fmt.Errorf("division by zero in Int.Divide"))
		return *i
	}
	quot := i.Int() / other.Int()
	return i.Set(quot)
}

func (i *Int) IncrumentBy(other Int) Int {
	sum := i.Int() + other.Int()
	return i.Set(sum)
}

func (i *Int) Increment() Int {
	sum := i.Int() + 1
	return i.Set(sum)
}

func (i *Int) DecrementBy(other Int) Int {
	diff := i.Int() - other.Int()
	return i.Set(diff)
}

func (i *Int) Decrement() Int {
	diff := i.Int() - 1
	return i.Set(diff)
}

func (f *Float) Set(in float64) Float {
	f.Value = strconv.FormatFloat(in, 'f', -1, 64)
	return *f
}

func (f *Float) Float() float64 {
	if f.Value == "" {
		return 0.0
	}
	val, err := strconv.ParseFloat(f.Value, 64)
	if err != nil {
		logHandler.ErrorLogger.Panic(ce.ErrInvalidTypeWrapper("Float", f.Value, "float64"))
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
	if b.Value == "" {
		return false
	}
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

func (f Field) String() string {
	return string(f)
}
