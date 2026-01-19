package templateStoreV2

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/importExportHelper"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// ExportRecordAsJSON exports the TemplateStore record as a JSON file.
func (record *TemplateStore) ExportRecordAsJSON(name string) {
	ID := reflect.ValueOf(*record).FieldByName(Fields.ID.String())
	clock := timing.Start(tableName, "Export", fmt.Sprintf("%v", ID))

	err := importExportHelper.ExportJSON(name, []TemplateStore{*record}, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v record %v: %v", tableName, ID, err.Error())
	}

	clock.Stop(1)
}

// ExportAllAsJSON exports all TemplateStore records as JSON files.
func ExportAllAsJSON(message string) {
	dao.CheckDAOReadyState(tableName, audit.EXPORT, databaseConnectionActive)

	clock := timing.Start(tableName, "Export", "ALL")
	recordList, _ := GetAll()
	if len(recordList) == 0 {
		logHandler.WarningLogger.Printf("[%v] %v data not found", tableName, tableName)
		clock.Stop(0)
		return
	}

	err := importExportHelper.ExportJSON(message, recordList, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting all %v's: %v", tableName, err.Error())
	}
	clock.Stop(len(recordList))
}

// ExportRecordAsCSV exports the TemplateStore record as a CSV file.
func (record *TemplateStore) ExportRecordAsCSV(name string) error {
	ID := reflect.ValueOf(*record).FieldByName(Fields.ID.String())
	clock := timing.Start(tableName, "Export", fmt.Sprintf("%v", ID))

	err := importExportHelper.ExportCSV(name, []TemplateStore{*record}, Fields.ID)
	if err != nil {
		logHandler.ExportLogger.Printf("Error exporting %v record %v: %v", tableName, ID, err.Error())
		clock.Stop(0)
		return err
	}

	clock.Stop(1)
	return nil
}

// ExportAllAsCSV exports all TemplateStore records as a CSV file.
func ExportAllAsCSV(msg string) error {
	exportListData, err := GetAll()
	if err != nil {
		logHandler.ExportLogger.Panicf("error Getting all %v's: %v", tableName, err.Error())
	}
	return importExportHelper.ExportCSV(msg, exportListData, Fields.ID)
}

// ImportAllFromCSV imports records for this table from a CSV file.
func ImportAllFromCSV() error {
	return importExportHelper.ImportCSV(tableName, &TemplateStore{}, templateImportProcessor)
}

// templateImportProcessor is called for each CSV row during import.
func templateImportProcessor(inOriginal **TemplateStore) (string, error) {
	importedData := **inOriginal
	stringField1 := strconv.Itoa(importedData.ID)

	_, err := Create(context.TODO(), importedData.UserName, importedData.UID, importedData.RealName, importedData.Email, importedData.GID)
	if err != nil {
		logHandler.ImportLogger.Panicf("Error importing %v: %v", tableName, err.Error())
		return stringField1, err
	}

	return stringField1, nil
}
