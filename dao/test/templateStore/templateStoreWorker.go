package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

import (
	"context"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Worker is a job that is scheduled to run at a predefined interval
func Worker(j jobs.Job, db *database.DB) {
	clock := timing.Start(jobs.CodedName(j), actions.INITIALISE.GetCode(), j.Description())
	oldDB := activeDB
	dbSwitched := false
	// Overide the default database connection if one is passed

	if db != nil {
		if activeDB.Name != db.Name {
			logHandler.EventLogger.Printf("Switching to %v.db", db.Name)
			activeDB = db
			dbSwitched = true
		}
	}

	templateJobProcessor(j)

	if dbSwitched {
		logHandler.EventLogger.Printf("Switching back to %v.db from %v.db", oldDB.Name, activeDB.Name)
		activeDB = oldDB
	}
	clock.Stop(1)
}

// templateJobProcessor processes jobs related to the TemplateStore domain entity.
// This function is triggered by the job scheduler to perform specific operations on TemplateStore records.
func templateJobProcessor(j jobs.Job) {
	clock := timing.Start(jobs.CodedName(j), actions.PROCESS.GetCode(), j.Description())
	count := 0

	//TODO: Add your job processing code here

	// Get all the sessions
	// For each session, check the expiry date
	// If the expiry date is less than now, then delete the session

	templateEntries, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Error: '%v'", jobs.CodedName(j), err.Error())
		return
	}

	notemplateEntries := len(templateEntries)
	if notemplateEntries == 0 {
		logHandler.ServiceLogger.Printf("[%v] No %vs to process", jobs.CodedName(j), Domain)
		clock.Stop(0)
		return
	}

	for templateEntryIndex, templateRecord := range templateEntries {
		logHandler.ServiceLogger.Printf("[%v] %v(%v/%v) %v", jobs.CodedName(j), Domain, templateEntryIndex+1, notemplateEntries, templateRecord.Raw)
		templateRecord.UpdateWithAction(context.TODO(), audit.GRANT, "Job Processing")
		templateRecord.UpdateWithAction(context.TODO(), audit.SERVICE, "Job Processing "+j.Name())
		count++
	}
	clock.Stop(count)
}
