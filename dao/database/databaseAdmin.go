package database

import (
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

func (db *DB) Backup(loc string) {
	timer := timing.Start(db.Name, actions.BACKUP.GetCode(), db.databaseName)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] data started... %v", db.Name, loc)
	db.Disconnect()
	logHandler.DatabaseLogger.Printf("Backup [%v.db] disconnected", db.Name)
	ioHelpers.Backup(db.Name, loc)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] backup done ends", db.Name)
	db.Reconnect()
	logHandler.DatabaseLogger.Printf("Backup [%v.db] (re)connected", db.Name)
	timer.Stop(1)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] data connection", db.Name)
}
