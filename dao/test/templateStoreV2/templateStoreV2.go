package templateStoreV2

import (
	"context"
	"fmt"
	"reflect"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dao/entities"
	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

func Count() (int, error) {
	logHandler.DatabaseLogger.Printf("COUNT %v", tableName)
	return activeDB.Count(&TemplateStore{})
}

func CountWhere(field entities.Field, value any) (int, error) {
	logHandler.DatabaseLogger.Printf("COUNT %v WHERE (%v=%v)", tableName, field.String(), value)
	clock := timing.Start(tableName, "Count", fmt.Sprintf("%v=%v", field.String(), value))
	count, err := activeDB.CountWhere(field, value, &TemplateStore{})
	if err != nil {
		clock.Stop(0)
		return 0, err
	}
	clock.Stop(count)
	return count, nil
}

func GetBy(field entities.Field, value any) (TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v WHERE (%v=%v)", tableName, field.String(), value)
	clock := timing.Start(tableName, "Get", fmt.Sprintf("%v=%v", field, value))

	dao.CheckDAOReadyState(tableName, audit.GET, dbIsReady)

	if field == Fields.ID && reflect.TypeOf(value).Name() != "int" {
		msg := "invalid data type. Expected type of %v is int"
		clock.Stop(0)
		return TemplateStore{}, ce.ErrGetWrapper(tableName, field.String(), value, fmt.Errorf(msg, value))
	}

	record, err := database.GetTyped[TemplateStore](activeDB, field, value)
	if err != nil {
		clock.Stop(0)
		return TemplateStore{}, ce.ErrRecordNotFoundWrapper(tableName, field.String(), fmt.Sprintf("%v", value))
	}
	if err := record.postGet(); err != nil {
		clock.Stop(0)
		return TemplateStore{}, ce.ErrGetWrapper(tableName, field.String(), value, err)
	}

	clock.Stop(1)
	return record, nil
}

func GetAll() ([]TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v ALL", tableName)
	dao.CheckDAOReadyState(tableName, audit.GET, dbIsReady)

	clock := timing.Start(tableName, "GetAll", "ALL")
	records, err := database.GetAllTyped[TemplateStore](activeDB)
	if err != nil {
		clock.Stop(0)
		return nil, ce.ErrNotFoundWrapper(tableName, err)
	}
	result, err := postGetList(records)
	if err != nil {
		clock.Stop(0)
		return nil, err
	}
	clock.Stop(len(result))
	return result, nil
}

func GetAllWhere(field entities.Field, value any) ([]TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v WHERE (%v=%v)", tableName, field.String(), value)
	dao.CheckDAOReadyState(tableName, audit.GET, dbIsReady)

	clock := timing.Start(tableName, "GetAllWhere", fmt.Sprintf("%v=%v", field, value))
	records, err := database.GetAllWhereTyped[TemplateStore](activeDB, field, value)
	if err != nil {
		clock.Stop(0)
		return nil, err
	}
	result, err := postGetList(records)
	if err != nil {
		clock.Stop(0)
		return nil, err
	}
	clock.Stop(len(result))
	return result, nil
}

func Delete(ctx context.Context, id int, note string) error {
	return DeleteBy(ctx, Fields.ID, id, note)
}

func DeleteBy(ctx context.Context, field entities.Field, value any, note string) error {
	logHandler.DatabaseLogger.Printf("DELETE %v WHERE %v=%v", tableName, field, value)
	dao.CheckDAOReadyState(tableName, audit.DELETE, dbIsReady)

	clock := timing.Start(tableName, "Delete", fmt.Sprintf("%v=%v", field.String(), value))

	record, err := GetBy(field, value)
	if err != nil {
		clock.Stop(0)
		return ce.ErrDAODeleteWrapper(tableName, field.String(), value, err)
	}

	if err := record.Audit.Action(ctx, audit.DELETE.WithMessage(note)); err != nil {
		clock.Stop(0)
		return ce.ErrDAOUpdateAuditWrapper(tableName, value, err)
	}

	if err := record.preDeleteProcessing(); err != nil {
		clock.Stop(0)
		return ce.ErrDAODeleteWrapper(tableName, field.String(), value, err)
	}

	if err := activeDB.Delete(&record); err != nil {
		clock.Stop(0)
		return ce.ErrDAODeleteWrapper(tableName, field.String(), value, err)
	}

	clock.Stop(1)
	return nil
}

func (record *TemplateStore) Validate() error {
	return record.validationProcessing()
}

func (record *TemplateStore) Update(ctx context.Context, note string) error {
	return record.insertOrUpdate(ctx, note, "Update", audit.UPDATE, "Update")
}

func (record *TemplateStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) error {
	return record.insertOrUpdate(ctx, note, "Update", auditAction, "Update")
}

func (record *TemplateStore) Create(ctx context.Context, note string) error {
	return record.insertOrUpdate(ctx, note, "Create", audit.CREATE, "Create")
}

func (record *TemplateStore) Clone(ctx context.Context) (TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("CLONE %v ID=%v", tableName, record.Key)
	return templateClone(ctx, *record)
}

func GetDefaultLookup() (lookup.Lookup, error) {
	return GetLookup(Fields.Key, Fields.Raw)
}

func GetLookup(field, value entities.Field) (lookup.Lookup, error) {
	dao.CheckDAOReadyState(tableName, audit.PROCESS, dbIsReady)

	clock := timing.Start(tableName, "Lookup", "BUILD")

	recordList, err := GetAll()
	if err != nil {
		lkpErr := ce.ErrDAOLookupWrapper(tableName, field.String(), value, err)
		logHandler.ErrorLogger.Print(lkpErr.Error())
		clock.Stop(0)
		return lookup.Lookup{}, lkpErr
	}

	var rtnLookup lookup.Lookup
	rtnLookup.Data = make([]lookup.LookupData, 0)

	for _, a := range recordList {
		key := reflect.ValueOf(a).FieldByName(field.String()).Interface().(string)
		val := reflect.ValueOf(a).FieldByName(value.String()).Interface().(string)
		rtnLookup.Data = append(rtnLookup.Data, lookup.LookupData{Key: key, Value: val})
	}

	clock.Stop(len(rtnLookup.Data))
	return rtnLookup, nil
}

func Drop() error {
	logHandler.DatabaseLogger.Printf("DROP %v", tableName)
	return activeDB.Drop(TemplateStore{})
}

func ClearDown(ctx context.Context) error {
	logHandler.DatabaseLogger.Printf("CLEARFILE %v", tableName)

	dao.CheckDAOReadyState(tableName, audit.PROCESS, dbIsReady)

	clock := timing.Start(tableName, "Clear", "INITIALISE")

	recordList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Print(ce.ErrDAOInitialisationWrapper(tableName, err).Error())
		clock.Stop(0)
		return ce.ErrDAOInitialisationWrapper(tableName, err)
	}

	count := 0
	logHandler.DatabaseLogger.Printf("Clearing %v records", len(recordList))

	for i, record := range recordList {
		logHandler.DatabaseLogger.Printf("(%v/%v) DELETE %v WHERE %v=%v", i+1, len(recordList), tableName, Fields.ID, record.ID)

		delErr := Delete(ctx, record.ID, fmt.Sprintf("Clearing %v %v @ initialisation ", tableName, record.ID))
		if delErr != nil {
			logHandler.ErrorLogger.Print(ce.ErrDAOInitialisationWrapper(tableName, delErr).Error())
			continue
		}
		count++
	}

	clock.Stop(count)
	logHandler.DatabaseLogger.Printf("Cleared down %v", tableName)
	return nil
}
