package importExportHelper

import (
	"encoding/csv"
	"fmt"
	"io"
	"os/user"
	"reflect"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

var SEP = "-"

func ExportCSV[T any](exportName string, exportList []T, idField database.Field) error {
	clock := timing.Start(exportName, actions.EXPORT.GetCode(), "")

	logHandler.ExportLogger.Printf("Exporting %v record(s) as CSV '%v'", len(exportList), exportName)

	exportName = buildName(exportName, exportList, idField)

	logHandler.ExportLogger.Printf("Exporting %v.csv", exportName)
	exportFile := openTargetFile(exportName, exportString, logHandler.ExportLogger, "csv", paths.Defaults().String())
	defer exportFile.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = FIELDSEPARATOR // Use tab-delimited format
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	_, err := gocsv.MarshalString(exportList) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v: %v", exportName, err.Error())
	}

	err = gocsv.MarshalFile(exportList, exportFile) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v: %v", exportName, err.Error())
	}

	//Example: # Generated 4 Zones at 11:39:05 on 2025-02-25
	noItems := len(exportList)
	plurality := "s"
	if noItems == 1 {
		plurality = ""
	}
	u, _ := user.Current()
	var by string
	if u != nil {
		by = u.Uid + SEP + u.Username
	} else {
		by = "sys_" + application.SystemIdentity()
	}
	on := application.SystemIdentity()
	os := application.OS()
	msg := fmt.Sprintf("# Generated (%v) %v%v at %v %v by %v on %v(%v)", len(exportList), exportName, plurality, time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"), by, on, os)
	exportFile.WriteString(msg)

	exportFile.Close()

	logHandler.ExportLogger.Printf("Exported (%v/%v) %v(s) to [%v]", len(exportList), len(exportList), exportName, exportFile.Name())
	//logHandler.EventLogger.Printf("Exported (%v/%v) %v(s) to [%v]", len(exportList), len(exportList), exportName, exportFile.Name())
	clock.Stop(len(exportList))
	return nil
}

func ExportJSON[T any](exportName string, exportList []T, idField database.Field) error {
	clock := timing.Start(exportName, actions.EXPORT.GetCode(), "")

	//if exportName == "" {
	//	exportName = buildName(exportName, exportList, idField)
	//}
	logHandler.TraceLogger.Printf("Exporting %v record(s) as JSON '%v'", len(exportList), exportName)

	for _, record := range exportList {
		//ID := reflect.ValueOf(record).FieldByName(idField.String())
		outputName := buildNameForRecord(exportName, record, idField)
		logHandler.TraceLogger.Printf("Exporting %v.json", outputName)

		exportJSON(outputName, paths.Dumps(), record)
	}
	clock.Stop(1)
	return nil
}

func buildName[T any](baseName string, exportList []T, idField database.Field) string {
	if baseName == "" {
		baseName = "Export"
	}
	if len(exportList) == 0 {
		return idHelpers.GetUUID() + SEP + baseName
	}
	firstRecord := exportList[0]
	domainName := reflect.TypeOf(firstRecord).Name()
	domainName = idHelpers.GetUUID() + SEP + domainName
	if len(exportList) > 1 {
		return domainName
	}
	xx := reflect.ValueOf(firstRecord).FieldByName(idField.String()).Interface()
	if xx == reflect.Invalid {
		return domainName
	}
	fmt.Printf("xx=%v", xx)
	domainName = domainName + SEP + fmt.Sprintf("%v", xx) + SEP + baseName
	return domainName
}

func buildNameForRecord[T any](baseName string, record T, idField database.Field) string {
	logHandler.TraceLogger.Printf("buildNameForRecord IN: baseName=[%v]", baseName)
	if baseName == "" {
		baseName = "Export"
	}
	var domainName string
	typ := reflect.TypeOf(record)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	domainName = typ.Name()
	if domainName == "" {
		domainName = "Record"
	}
	domainName = idHelpers.GetUUID() + SEP + domainName

	val := reflect.ValueOf(record)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if !val.IsValid() || val.Kind() != reflect.Struct {
		return domainName
	}
	field := val.FieldByName(idField.String())
	if !field.IsValid() {
		return domainName
	}
	xx := field.Interface()
	domainName = domainName + SEP + fmt.Sprintf("%v", xx) + SEP + baseName
	logHandler.TraceLogger.Printf("buildNameForRecord OUT: domainName=[%v]", domainName)
	return domainName
}

func exportJSON[T any](exportName string, where paths.FileSystemPath, record T) {

	logHandler.TraceLogger.Printf("Exporting %v.json", exportName)
	logHandler.ExportLogger.Printf("Exporting %v %v.json", database.GetStructType(record), exportName)
	exportFile := openTargetFile(exportName, exportString, logHandler.ExportLogger, "json", where.String())
	defer exportFile.Close()

	exportFile.WriteString(godump.DumpJSONStr(record))

	exportFile.Close()

}
