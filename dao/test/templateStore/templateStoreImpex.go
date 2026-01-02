package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the importProcessor function to process the domain entity

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/importExportHelper"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// ExportRecordAsJSON exports the TemplateStore record as a JSON file.
// Parameters:
//   - name: The name to be used for the exported file.
func (record *TemplateStore) ExportRecordAsJSON(name string) {

	ID := reflect.ValueOf(*record).FieldByName(Fields.ID.String())

	clock := timing.Start(Domain, actions.EXPORT.GetCode(), fmt.Sprintf("%v", ID))

	//ioHelpers.Dump(Domain, paths.Dumps(), name, fmt.Sprintf("%v", ID), record)
	err := importExportHelper.ExportJSON(name, []TemplateStore{*record}, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v record %v: %v", Domain, ID, err.Error())
	}

	clock.Stop(1)
}

// ExportAllAsJSON exports all TemplateStore records as JSON files.
// Parameters:
//   - message: A message to be included in the export process.
func ExportAllAsJSON(message string) {

	dao.CheckDAOReadyState(Domain, audit.EXPORT, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(Domain, actions.EXPORT.GetCode(), "ALL")
	recordList, _ := GetAll()
	if len(recordList) == 0 {
		logHandler.WarningLogger.Printf("[%v] %v data not found", Domain, Domain)
		clock.Stop(0)
		return
	}
	//SEP := "!"
	// for _, record := range recordList {
	// 	//msg := fmt.Sprintf("%v%v%v", audit.EXPORT.Description(), SEP, message)
	// 	//if message == "" {
	// 	//	msg = fmt.Sprintf("%v%v", audit.EXPORT.Description(), SEP)
	// 	//}
	err := importExportHelper.ExportJSON(message, recordList, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting all %v's: %v", Domain, err.Error())
	}
	// }
	clock.Stop(len(recordList))
}

// ExportRecordAsCSV exports the TemplateStore record as a CSV file.
// Parameters:
//   - name: The name to be used for the exported file.
func (record *TemplateStore) ExportRecordAsCSV(name string) error {

	ID := reflect.ValueOf(*record).FieldByName(Fields.ID.String())

	clock := timing.Start(Domain, actions.EXPORT.GetCode(), fmt.Sprintf("%v", ID))
	err := importExportHelper.ExportCSV(name, []TemplateStore{*record}, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Printf("Error exporting %v record %v: %v", Domain, ID, err.Error())
		clock.Stop(0)
		return err
	}

	clock.Stop(1)
	return nil
}

// ExportAllAsCSV exports all TemplateStore records as a CSV file.
// Parameters:
//   - none
func ExportAllAsCSV(msg string) error {

	exportListData, err := GetAll()
	if err != nil {
		logHandler.ExportLogger.Panicf("error Getting all %v's: %v", Domain, err.Error())
	}

	return importExportHelper.ExportCSV(msg, exportListData, Fields.ID)
}

func ImportAllFromCSV() error {
	return importExportHelper.ImportCSV(Domain, &TemplateStore{}, templateImportProcessor)
}

// templateImportProcessor is a helper function to create a new entry instance and save it to the database
// It should be customised to suit the specific requirements of the entryination table/DAO.
func templateImportProcessor(inOriginal **TemplateStore) (string, error) {
	//TODO: Build the import processing functionality for the template_Store data here
	//
	importedData := **inOriginal

	//	logHandler.ImportLogger.Printf("Importing %v [%v] [%v]", domain, original.Raw, original.Field1)

	//logger.InfoLogger.Printf("ACT: NEW New %v %v %v", tableName, name, entryination)
	// u := Behaviour_Store{}
	// u.Key = idHelpers.Encode(strings.ToUpper(original.Raw))
	// u.Raw = original.Raw
	// u.Text = original.Text
	// // u.Action = original.Action
	// u.Domain = original.Domain

	// importAction := actions.New(original.Action.Name)
	// bh, err := Declare(importAction, domains.Special(original.Domain), original.Text)
	// if err != nil {
	// 	logHandler.ImportLogger.Panicf("Error importing template: %v", err.Error())
	// }

	// Return the created entry and nil error
	//logHandler.ImportLogger.Printf("Imported %v [%+v]", domain, importedData)

	stringField1 := strconv.Itoa(importedData.ID)

	_, err := Create(context.TODO(), importedData.UserName, importedData.UID, importedData.RealName, importedData.Email, importedData.GID)
	if err != nil {
		logHandler.ImportLogger.Panicf("Error importing %v: %v", Domain, err.Error())
		return stringField1, err
	}

	return stringField1, nil
}
