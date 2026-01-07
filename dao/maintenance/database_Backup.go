package maintenance

import (
	"time"

	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

type DatabaseBackupJob struct {
	databaseAccessors []func() ([]*database.DB, error)
}

func (job *DatabaseBackupJob) Run() error {
	jobs.PreRun(job)
	performDatabaseBackup(job)
	jobs.PostRun(job)
	return nil
}

func (job *DatabaseBackupJob) Service() func() {
	return func() {
		_ = job.Run()
	}
}

func (job *DatabaseBackupJob) Schedule() string {
	return "55 23 * * *"
}

func (job *DatabaseBackupJob) Name() string {
	return "Maintenance - Backup Database"
}

func performDatabaseBackup(job *DatabaseBackupJob) {
	logHandler.ServiceLogger.Printf("[%v] [%v] Started", domain, job.Name())

	// Get a coded name for the job
	name := jobs.CodedName(job)

	j := timing.Start(name, "Backup", job.Description())

	dateTime := time.Now().Format(dateHelpers.Format.BackupFolder)
	logHandler.ServiceLogger.Printf("[%v] Backup Date: [%v]", domain, dateTime)

	destPath := paths.Backups().String() + paths.Seperator() + dateTime
	fullBackupPath := paths.Application().String() + destPath

	//create a folder
	err := ioHelpers.MkDir(fullBackupPath)
	if err != nil {
		logHandler.ServiceLogger.Panicf("[%v] [%v] Error: [%v]", domain, name, err.Error())
	}
	count := 0
	for _, thisFunc := range job.databaseAccessors {
		dbList, err := thisFunc()
		if err != nil {
			logHandler.ServiceLogger.Panicf("[%v] [%v] Error: [%v]", domain, name, err.Error())
			panic(err)
		}
		for _, db := range dbList {
			count++
			logHandler.ServiceLogger.Printf("[%v] [%v] Backup [%v]", domain, name, db.Name)
			db.Disconnect()
			logHandler.ServiceLogger.Printf("[%v] [%v] Disconnected [%v]", domain, name, db.Name)
			db.Backup(fullBackupPath)
			logHandler.ServiceLogger.Printf("[%v] [%v] Done [%v]", domain, name, db.Name)
			db.Reconnect()
			logHandler.ServiceLogger.Printf("[%v] [%v] Reconnected [%v]", domain, name, db.Name)
		}
	}
	j.Stop(count)
	logHandler.ServiceLogger.Printf("[%v] [%v] Completed", domain, job.Name())
}

func (job *DatabaseBackupJob) AddDatabaseAccessFunctions(fn func() ([]*database.DB, error)) {
	logHandler.ServiceLogger.Printf("[%v] [%v] Adding Function", domain, job.Name())
	job.databaseAccessors = append(job.databaseAccessors, fn)
	logHandler.ServiceLogger.Printf("[%v] [%v] Function Added - No Funcs=(%v)", domain, job.Name(), len(job.databaseAccessors))
}

func (job *DatabaseBackupJob) Description() string {
	return "Database Backup, runs at 11:55 daily"
}
