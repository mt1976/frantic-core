package database

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

type Field dao.Field

type DB struct {
	connection   *storm.DB
	Name         string
	databaseName string
	initialised  bool
}

func Connect() *DB {
	return connect("main")
}

func ConnectToNamedDB(name string) *DB {
	return connect(name)
}

func (db *DB) Reconnect() {
	logHandler.DatabaseLogger.Printf("[RECONNECT] Reconnecting [%v.db] data - %+v", db.Name, db)
	logHandler.DatabaseLogger.Printf("[RECONNECT] Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[RECONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	connect(db.Name)
	logHandler.DatabaseLogger.Printf("[RECONNECT] Reconnected [%v.db] data", db.Name)
}

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

func (db *DB) Disconnect() {
	timer := timing.Start(db.Name, actions.DISCONNECT.Code, db.databaseName)
	logHandler.DatabaseLogger.Printf("[DISCONNECT] Disconnecting [%v.db] connection", db.Name)
	err := db.connection.Close()
	if err != nil {
		logHandler.DatabaseLogger.Panicf("[DISCONNECT] Closing [%v.db] %v ", db.Name, err.Error())
		panic(commonErrors.WrapDisconnectError(err))
	}
	releaseFromConnectionPool(db)
	logHandler.DatabaseLogger.Printf("[DISCONNECT] Closed [%v.db] connection", db.Name)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("[DISCONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	timer.Stop(1)
}

func (db *DB) Retrieve(fieldName Field, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)
	return db.connection.One(string(fieldName), value, to)
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("GetAll [%+v][%+v] [%v.db]", dao.GetStructType(to), options, db.Name)
	return db.connection.All(to, options...)
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.DeleteStruct(data)
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.Drop(data)
}

func (db *DB) Update(data any) error {
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - Start", dao.GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - Error", dao.GetStructType(data), db.Name)
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db] - End", dao.GetStructType(data), db.Name)
	return db.connection.Update(data)
}

func (db *DB) Create(data any) error {
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - Start", dao.GetStructType(data), db.Name)
	err := validate(data, db)
	if err != nil {
		logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - Error", dao.GetStructType(data), db.Name)
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db] - End", dao.GetStructType(data), db.Name)
	return db.connection.Save(data)
}

func (db *DB) Count(data any) (int, error) {
	logHandler.DatabaseLogger.Printf("Count [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	return db.connection.Count(data)
}

func (db *DB) CountWhere(fieldName dao.Field, value any, to any) (int, error) {
	logHandler.DatabaseLogger.Printf("CountWhere (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)
	if err := dao.IsValidFieldInStruct(fieldName, to); err != nil {
		logHandler.DatabaseLogger.Printf("CountWhere (%+v=%+v)[%+v] [%v.db] - Error", fieldName, value, dao.GetStructType(to), db.Name)
		return 0, err
	}
	query := db.connection.Select(q.Eq(fieldName.String(), value))
	count, err := query.Count(to)
	return count, err
}
