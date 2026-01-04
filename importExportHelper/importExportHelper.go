package importExportHelper

import (
	"fmt"
	"log"
	"os"

	"github.com/mt1976/frantic-core/paths"
)

// TemplateImportData is a struct to hold the data from the CSV file
// it is used to import the data into the database
// The struct tags are used to map the fields to the CSV columns
// this struct should be customised to suit the specific requirements of the entryination table/DAO.

var FIELDSEPARATOR = '|'
var importString = "Import"
var exportString = "Export"

func openTargetFile(in, action string, useLog *log.Logger, fileExtension string, path string) *os.File {
	//defaultPath := paths.Defaults()
	templateDataFileName := in + "." + fileExtension
	fileName := fmt.Sprintf("%s%s/%s", paths.Application().String(), path, templateDataFileName)

	dataFileHandle, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		useLog.Fatalf("error opening file: %v", err)
		panic(err)
	}
	//useLog.Printf("%ving %vs from File=[%v]", action, in, dataFileHandle.Name())
	return dataFileHandle
}
