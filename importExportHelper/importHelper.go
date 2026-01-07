package importExportHelper

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

func ImportCSV[T any](importName string, entryTypeToInsert T, importProcessor func(*T) (string, error)) error {

	clock := timing.Start(importName, "Import", "")
	// Create a slice of entryTypeToInsert to hold the data from the CSV file
	insertEntriesList := []T{}

	importFile := openTargetFile(importName, importString, logHandler.ImportLogger, "csv", paths.Defaults().String())
	defer importFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)    // Allows use pipe as delimiter
		r.Comma = FIELDSEPARATOR  // Use tab-delimited format
		r.Comment = '#'           // Ignore comment lines
		r.TrimLeadingSpace = true // Trim leading space
		return r                  // Allows use pipe as delimiter
	})

	if err := gocsv.UnmarshalFile(importFile, &insertEntriesList); err != nil { // Load clients from file
		logHandler.ImportLogger.Printf("Importing %v: %v - No Content, nothing to import.", importName, err.Error())
		importFile.Close()
		clock.Stop(0)
		return nil
	}

	if _, err := importFile.Seek(0, 0); err != nil { // Go to the start of the file
		logHandler.ImportLogger.Panicf("Importing %v: %v - Unable to fet to start of file.", importName, err.Error())
		clock.Stop(0)
		panic(err)
	}

	totalImportEntries := len(insertEntriesList)

	count := 0
	for thisPos, insertEntry := range insertEntriesList {
		//logHandler.ImportLogger.Printf("Import %v (%v/%v)", importName, thisPos+1, totalImportEntries)
		// the load function is a helper function to create a new entry instance and save it to the database
		// the parameters should be customised to suit the specific requirements of the entryination table/DAO.
		entryIdentifier, err := importProcessor(&insertEntry)
		if err != nil {
			logHandler.ImportLogger.Panicf("Error importing %v [%v] Error=[%v]", importName, entryIdentifier, err.Error())
			continue
		}
		logHandler.ImportLogger.Printf("Imported %v (%v/%v) %v=[%v]", importName, thisPos+1, totalImportEntries, importName, entryIdentifier)

		count++
	}

	logHandler.ImportLogger.Printf("Imported (%v/%v) %v(s) from [%v]", count, totalImportEntries, importName, importFile.Name())
	logHandler.EventLogger.Printf("Imported (%v/%v) %v(s) from [%v]", count, totalImportEntries, importName, importFile.Name())
	importFile.Close()
	clock.Stop(count)
	return nil
}
