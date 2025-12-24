package importExportHelper

import (
	"encoding/csv"
	"fmt"
	"io"
	"os/user"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

func ExportCSV[T any](exportName string, exportList []T) error {
	clock := timing.Start(exportName, actions.EXPORT.GetCode(), "")

	exportFile := openTargetFile(exportName, exportString, logHandler.ExportLogger)
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
		by = u.Uid + "_" + u.Username
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
