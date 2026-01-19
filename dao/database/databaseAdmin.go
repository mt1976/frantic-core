package database

import (
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Backup creates a backup of the database to the specified location
// It disconnects the database, performs the backup, and then reconnects.
// The backup process is timed and logged.
func (db *DB) Backup(loc string) {
	timer := timing.Start(db.Name, "Backup", db.databaseName)
	// Ensure all data is flushed to disk before backup
	//db.Flush()
	logHandler.DatabaseLogger.Printf("[ADM] Backup [...%v.db] data started... %v", db.Name, loc)
	db.Disconnect()
	logHandler.DatabaseLogger.Printf("[ADM] Backup [...%v.db] disconnected", db.Name)
	ioHelpers.Backup(db.Name, loc)
	logHandler.DatabaseLogger.Printf("[ADM] Backup [...%v.db] backup done ends", db.Name)
	db.Reconnect()
	logHandler.DatabaseLogger.Printf("[ADM] Backup [...%v.db] (re)connected", db.Name)
	timer.Stop(1)
	logHandler.DatabaseLogger.Printf("[ADM] Backup [...%v.db] data connection", db.Name)
}
