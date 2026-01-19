package templateStoreV2

import (
	"context"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Worker is a job that is scheduled to run at a predefined interval.
func Worker(j jobs.Job, db *database.DB) {
	clock := timing.Start(jobs.CodedName(j), "Initialise", j.Description())
	oldDB := activeDBConnection
	dbSwitched := false

	if db != nil {
		if activeDBConnection.Name != db.Name {
			logHandler.EventLogger.Printf("Switching to %v.db", db.Name)
			activeDBConnection = db
			dbSwitched = true
		}
	}

	jobProcessor(j)

	if dbSwitched {
		logHandler.EventLogger.Printf("Switching back to %v.db from %v.db", oldDB.Name, activeDBConnection.Name)
		activeDBConnection = oldDB
	}
	clock.Stop(1)
}

// jobProcessor processes jobs related to the TemplateStore tableName entity.
func jobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), "Process", j.Description())
	count := 0

	templateEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error: '%v'", jobs.CodedName(j), err.Error())
		return
	}

	notemplateEntries := len(templateEntries)
	if notemplateEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), tableName)
		clock.Stop(0)
		return
	}

	for templateEntryIndex, templateRecord := range templateEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), tableName, templateEntryIndex+1, notemplateEntries, templateRecord.Raw)
		_ = templateRecord.UpdateWithAction(context.Background(), audit.SERVICE, "Job Processing "+j.Name())
		count++
	}
	clock.Stop(count)
}
