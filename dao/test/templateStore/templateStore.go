package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
/// DO NOT CHANGE THIS FILE OTHER THAN TO RENAME "template" TO THE NAME OF THE DOMAIN ENTITY

import (
	"context"
	"fmt"
	"reflect"

	"github.com/goforj/godump"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dao/fields"
	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Count returns the total number of TemplateStore records in the database.
//
// Returns:
//   - int: The total count of TemplateStore records.
//   - error: An error object if any issues occur during the counting process; otherwise, nil.
func Count() (int, error) {
	logHandler.DatabaseLogger.Printf("COUNT %v", Domain)
	return activeDB.Count(&TemplateStore{})
}

// CountWhere counts the number of TemplateStore records that match the specified field and value.
//
// Parameters:
//   - field: The field to be used for filtering records.
//   - value: The value of the specified field to filter records.
//
// Returns:
//   - int: The count of TemplateStore records that match the specified criteria.
//   - error: An error object if any issues occur during the counting process; otherwise, nil.
func CountWhere(field fields.Field, value any) (int, error) {
	logHandler.DatabaseLogger.Printf("COUNT %v WHERE (%v=%v)", Domain, field.String(), value)
	clock := timing.Start(Domain, "Count", fmt.Sprintf("%v=%v", field.String(), value))
	count, err := activeDB.CountWhere(field, value, &TemplateStore{})
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
		clock.Stop(0)
		return 0, err
	}
	logHandler.DatabaseLogger.Printf("COUNT RESULT %v WHERE (%v=%v) = %d", Domain, field.String(), value, count)
	clock.Stop(count)
	return count, nil
}

// GetById retrieves a TemplateStore record from the database based on the specified ID.
//
// Parameters:
//   - id: The ID of the record to retrieve.
//
// Returns:
//   - TemplateStore: The TemplateStore record that matches the specified ID.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
//
// DEPRECATED: Please use GetBy with Fields.ID instead.
func GetById(id any) (TemplateStore, error) {
	logHandler.WarningLogger.Printf("GetById is deprecated, please use GetBy with Fields.ID instead")
	panic("GetById is deprecated, please use GetBy with Fields.ID instead")
	//	return GetBy(Fields.ID, id)
}

// GetByKey retrieves a TemplateStore record from the database based on the specified key.
//
// Parameters:
//   - key: The key of the record to retrieve.
//
// Returns:
//   - TemplateStore: The TemplateStore record that matches the specified key.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
//
// DEPRECATED: Please use GetBy with Fields.Key instead.
func GetByKey(key any) (TemplateStore, error) {
	logHandler.WarningLogger.Printf("GetByKey is deprecated, please use GetBy with Fields.Key instead")
	panic("GetByKey is deprecated, please use GetBy with Fields.Key instead")
	// return GetBy(Fields.Key, key)
}

// GetBy retrieves a TemplateStore record from the database based on the specified field and value.
//
// Parameters:
//   - field: The field to be used for filtering records.
//   - value: The value of the specified field to filter records.
//
// Returns:
//   - TemplateStore: The TemplateStore record that matches the specified criteria.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func GetBy(field fields.Field, value any) (TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v WHERE (%v=%v)", Domain, field.String(), value)

	clock := timing.Start(Domain, "Get", fmt.Sprintf("%v=%v", field, value))

	dao.CheckDAOReadyState(Domain, audit.GET, initialised)

	if field == Fields.ID && reflect.TypeOf(value).Name() != "int" {
		msg := "invalid data type. Expected type of %v is int"
		logHandler.ErrorLogger.Printf(msg, value)
		return TemplateStore{}, ce.ErrGetWrapper(Domain, field.String(), value, fmt.Errorf(msg, value))
	}

	// if err := database.IsValidFieldInStruct(field, TemplateStore{}); err != nil {
	// 	return TemplateStore{}, err
	// }

	// if err := database.IsValidTypeForField(field, value, TemplateStore{}); err != nil {
	// 	return TemplateStore{}, err
	// }

	record := TemplateStore{}

	// activeDB.Retrieve returns any, so we need to type-assert to *TemplateStore
	result, err := activeDB.Get(field, value, &record)
	if err != nil {
		clock.Stop(0)
		return TemplateStore{}, ce.ErrRecordNotFoundWrapper(Domain, field.String(), fmt.Sprintf("%v", value))
	}

	// Type assert the result to *TemplateStore
	x, err := assertTemplateStore(result, field, value)
	if err != nil {
		return TemplateStore{}, err
	}

	if err := x.postGet(); err != nil {
		clock.Stop(0)
		return TemplateStore{}, ce.ErrGetWrapper(Domain, field.String(), value, err)
	}

	clock.Stop(1)
	return *x, nil
}

// GetAll retrieves all TemplateStore records from the database.
//
// Returns:
//   - []TemplateStore: A slice of all TemplateStore records.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func GetAll() ([]TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v ALL", Domain)
	dao.CheckDAOReadyState(Domain, audit.GET, initialised) // Check the DAO has been initialised, Mandatory.

	recordList := []TemplateStore{}
	resultList := []TemplateStore{}

	clock := timing.Start(Domain, "GetAll", "ALL")
	logHandler.DatabaseLogger.Printf("SELECT %v ALL", Domain)

	recordListAny, errG := activeDB.GetAll(&recordList)

	logHandler.DatabaseLogger.Printf("Got %v records from database (%v)", len(recordListAny), len(recordList))

	if errG != nil {
		clock.Stop(0)
		return []TemplateStore{}, ce.ErrNotFoundWrapper(Domain, errG)
	}

	for _, rec := range recordListAny {
		ts, ok := rec.(TemplateStore)
		if !ok {
			clock.Stop(0)
			return []TemplateStore{}, ce.ErrInvalidTypeWrapper("GetAll", Domain, fmt.Sprintf("%T", rec))
		}
		resultList = append(resultList, ts)
	}

	var errPost error
	if resultList, errPost = postGetList(&resultList); errPost != nil {
		clock.Stop(0)
		return nil, errPost
	}

	clock.Stop(len(resultList))

	logHandler.DatabaseLogger.Printf("RETRIEVED %v records from %v", len(resultList), database.GetStructType(recordList))

	return resultList, nil
}

// GetAllWhere retrieves all TemplateStore records that match the specified field and value.
//
// Parameters:
//   - field: The field to be used for filtering records.
//   - value: The value of the specified field to filter records.
//
// Returns:
//   - []TemplateStore: A slice of TemplateStore records that match the specified criteria.
//   - error: An error object if any issues occur during the retrieval process; otherwise, nil.
func GetAllWhere(field fields.Field, value any) ([]TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("SELECT %v WHERE (%v=%v)", Domain, field.String(), value)
	dao.CheckDAOReadyState(Domain, audit.GET, initialised) // Check the DAO has been initialised, Mandatory.

	//recordList := []TemplateStore{}
	resultList := []TemplateStore{}

	clock := timing.Start(Domain, "GetAllWhere", fmt.Sprintf("%v=%v", field, value))

	//logHandler.DatabaseLogger.Printf("SELECT %v WHERE %v=%v", Domain, field, value)

	// logHandler.InfoLogger.Printf("Check IsValidFieldInStruct %v", field.String())
	// if err := database.IsValidFieldInStruct(field, TemplateStore{}); err != nil {
	// 	return nil, err
	// }

	// logHandler.InfoLogger.Printf("Check IsValidTypeForField %v", field.String())
	// if err := database.IsValidTypeForField(field, value, TemplateStore{}); err != nil {
	// 	return nil, err
	// }

	//err := activeDB.Retrieve(field, value, &recordList)
	//logHandler.DatabaseLogger.Println("Call GetAllWhere")
	recordListAny, err := activeDB.GetAllWhere(field, value, &[]TemplateStore{})
	if err != nil {
		logHandler.ErrorLogger.Print(err.Error())
		return nil, err
	}
	//logHandler.DatabaseLogger.Println("Process returned records")
	for _, rec := range recordListAny {
		//logHandler.InfoLogger.Printf("Processing record of type %v", database.GetStructType(rec))
		if database.GetStructType(rec) != database.GetStructType(TemplateStore{}) {
			logHandler.ErrorLogger.Printf("Invalid record type returned from GetAllWhere wanted %v, got %v", database.GetStructType(TemplateStore{}), database.GetStructType(rec))
			panic(fmt.Sprintf("invalid record type returned from GetAllWhere wanted %v, got %v", database.GetStructType(TemplateStore{}), database.GetStructType(rec)))
		}
		if reflect.TypeOf(rec).Kind() == reflect.Ptr {
			//	logHandler.InfoLogger.Printf("Dereferencing pointer to get TemplateStore value")
			rec = reflect.ValueOf(rec).Elem().Interface()
		}
		resultList = append(resultList, rec.(TemplateStore))
	}
	// count := 0

	// for _, record := range recordList {
	// 	if reflect.ValueOf(record).FieldByName(field.String()).Interface() == value {
	// 		count++
	// 		resultList = append(resultList, record)
	// 	}
	// }

	var errPost error
	if resultList, errPost = postGetList(&resultList); errPost != nil {
		clock.Stop(0)
		logHandler.ErrorLogger.Print(errPost.Error())
		return nil, errPost
	}

	clock.Stop(len(resultList))

	return resultList, nil
}

// Delete removes a TemplateStore record from the database based on the specified ID.
//
// Parameters:
//   - ctx: The context for the operation.
//   - id: The ID of the record to delete.
//   - note: A note describing the delete action.
//
// Returns:
//   - error: An error object if any issues occur during the deletion process; otherwise, nil.
func Delete(ctx context.Context, id int, note string) error {
	return DeleteBy(ctx, Fields.ID, id, note)
}

// DeleteByKey deletes a TemplateStore record from the database based on the specified key.
//
// Parameters:
//   - ctx: The context for the operation.
//   - key: The key of the record to delete.
//   - note: A note describing the delete action.
//
// Returns:
//   - error: An error object if any issues occur during the deletion process; otherwise, nil.
func DeleteByKey(ctx context.Context, key string, note string) error {
	logHandler.WarningLogger.Printf("DeleteByKey is deprecated, please use DeleteBy with Fields.Key instead")
	panic("DeleteByKey is deprecated, please use DeleteBy with Fields.Key instead")
	//	return DeleteBy(ctx, Fields.Key, key, note)
}

// DeleteBy deletes a TemplateStore record from the database based on the specified field and value.
//
// Parameters:
//   - ctx: The context for the operation.
//   - field: The field to be used for identifying the record to delete.
//   - value: The value of the specified field to identify the record.
//   - note: A note describing the delete action.
//
// Returns:
//   - error: An error object if any issues occur during the deletion process; otherwise, nil.
func DeleteBy(ctx context.Context, field fields.Field, value any, note string) error {
	logHandler.DatabaseLogger.Printf("DELETE %v WHERE %v=%v", Domain, field, value)

	dao.CheckDAOReadyState(Domain, audit.DELETE, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(Domain, "Delete", fmt.Sprintf("%v=%v", field.String(), value))

	// if err := database.IsValidFieldInStruct(field, TemplateStore{}); err != nil {
	// 	logHandler.ErrorLogger.Print(commonErrors.WrapDAODeleteError(Domain, field.String(), value, err).Error())
	// 	clock.Stop(0)
	// 	return commonErrors.WrapDAODeleteError(Domain, field.String(), value, err)
	// }

	// if err := database.IsValidTypeForField(field, value, TemplateStore{}); err != nil {
	// 	logHandler.ErrorLogger.Print(commonErrors.WrapDAODeleteError(Domain, field.String(), value, err).Error())
	// 	clock.Stop(0)
	// 	return err
	// }

	record, err := GetBy(field, value)

	if err != nil {
		getErr := ce.ErrDAODeleteWrapper(Domain, field.String(), value, err)
		logHandler.ErrorLogger.Panic(getErr.Error(), err)
		clock.Stop(0)
		return getErr
	}

	auditErr := record.Audit.Action(ctx, audit.DELETE.WithMessage(note))
	if auditErr != nil {
		audErr := ce.ErrDAOUpdateAuditWrapper(Domain, value, auditErr)
		logHandler.ErrorLogger.Print(audErr.Error())
		clock.Stop(0)
		return audErr
	}

	preDeleteErr := record.preDeleteProcessing()
	if preDeleteErr != nil {
		logHandler.ErrorLogger.Print(ce.ErrDAODeleteWrapper(Domain, field.String(), value, preDeleteErr).Error())
		clock.Stop(0)
		return preDeleteErr
	}
	//logHandler.WarningLogger.Printf("Deleting %v record where %v=%v %v", Domain, field.String(), value, audit.DELETE.ShortName())
	//record.ExportRecordAsJSON(audit.DELETE.ShortName())

	if err := activeDB.Delete(&record); err != nil {
		delErr := ce.ErrDAODeleteWrapper(Domain, field.String(), value, err)
		logHandler.ErrorLogger.Panic(delErr.Error())
		clock.Stop(0)
		return delErr
	}

	clock.Stop(1)

	return nil
}

// Spew outputs the contents of the TemplateStore record to the Info log.
func (record *TemplateStore) Spew() {
	logHandler.TraceLogger.Printf("[%v] Record=[%+v]", Domain, godump.DumpStr(record))
}

// Validate checks if the TemplateStore record is valid.
//
// Returns:
//   - error: An error object if the record is invalid; otherwise, nil.
func (record *TemplateStore) Validate() error {
	return record.validationProcessing()
}

// Update updates the TemplateStore record in the database.
//
// Parameters:
//   - ctx: The context for the operation.
//   - note: A note describing the update action.
//
// Returns:
//   - error: An error object if any issues occur during the update process; otherwise, nil.
func (record *TemplateStore) Update(ctx context.Context, note string) error {
	return record.insertOrUpdate(ctx, note, "Update", audit.UPDATE, "Update")
}

// UpdateWithAction updates the TemplateStore record in the database with a specified audit action.
//
// Parameters:
//   - ctx: The context for the operation.
//   - auditAction: The audit action to be recorded during the update.
//   - note: A note describing the update action.
//
// Returns:
//   - error: An error object if any issues occur during the update process; otherwise, nil.
func (record *TemplateStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) error {
	return record.insertOrUpdate(ctx, note, "Update", auditAction, "Update")
}

// Create inserts a new TemplateStore record into the database.
//
// Parameters:
//   - ctx: The context for the operation.
//   - note: A note describing the creation action.
//
// Returns:
//   - error: An error object if any issues occur during the creation process; otherwise, nil.
func (record *TemplateStore) Create(ctx context.Context, note string) error {
	return record.insertOrUpdate(ctx, note, "Create", audit.CREATE, "Create")
}

// Clone creates a duplicate of the current TemplateStore record in the database.
//
// Parameters:
//   - ctx: The context for the operation.
//
// Returns:
//   - TemplateStore: A new TemplateStore instance that is a clone of the current record.
//   - error: An error object if any issues occur during the cloning process; otherwise, nil.
func (record *TemplateStore) Clone(ctx context.Context) (TemplateStore, error) {
	logHandler.DatabaseLogger.Printf("CLONE %v ID=%v", Domain, record.Key)
	return templateClone(ctx, *record)
}

// GetDefaultLookup builds a default Lookup structure using Key as the key field and Raw as the value field.
//
// Returns:
//   - lookup.Lookup: A Lookup structure containing key-value pairs based on the Key and Raw fields.
//   - error: An error object if any issues occur during the lookup process; otherwise, nil.
func GetDefaultLookup() (lookup.Lookup, error) {
	return GetLookup(Fields.Key, Fields.Raw)
}

// GetLookup builds a Lookup structure for the specified field and value.
//
// Parameters:
//   - field: The field to be used as the key in the lookup.
//   - value: The field to be used as the value in the lookup.
//
// Returns:
//   - lookup.Lookup: A Lookup structure containing key-value pairs based on the specified fields.
//   - error: An error object if any issues occur during the lookup process; otherwise, nil.
func GetLookup(field, value fields.Field) (lookup.Lookup, error) {

	dao.CheckDAOReadyState(Domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(Domain, "Lookup", "BUILD")

	// Get all status
	recordList, err := GetAll()
	if err != nil {
		lkpErr := ce.ErrDAOLookupWrapper(Domain, field.String(), value, err)
		logHandler.ErrorLogger.Print(lkpErr.Error())
		clock.Stop(0)
		return lookup.Lookup{}, lkpErr
	}

	// Create a new Lookup
	var rtnLookup lookup.Lookup
	rtnLookup.Data = make([]lookup.LookupData, 0)

	// range through Behaviour list, if status code is found and deletedby is empty then return error
	for _, a := range recordList {
		key := reflect.ValueOf(a).FieldByName(field.String()).Interface().(string)
		val := reflect.ValueOf(a).FieldByName(value.String()).Interface().(string)
		rtnLookup.Data = append(rtnLookup.Data, lookup.LookupData{Key: key, Value: val})
	}

	clock.Stop(len(rtnLookup.Data))

	return rtnLookup, nil
}

// Drop removes the DAO's database entirely.
//
// This function is typically used during development or testing phases to completely remove
// the database associated with the Data Access Object (DAO).
// It logs the drop action and invokes the database's drop method for the specific domain entity.
//
// Returns:
//   - error: An error object if any issues occur during the drop process; otherwise, nil.
func Drop() error {
	logHandler.DatabaseLogger.Printf("DROP %v", Domain)
	return activeDB.Drop(TemplateStore{})
}

// ClearDown removes all records from the DAO's database.
//
// This function is typically used during the initialisation phase to ensure that the database
// is clean and free of any existing records related to the DAO's domain entity.
// It retrieves all records and deletes them one by one, logging the process.
// Parameters:
//   - ctx: The context for the operation.
//
// Returns:
//   - error: An error object if any issues occur during the clear down process; otherwise, nil.
func ClearDown(ctx context.Context) error {
	logHandler.DatabaseLogger.Printf("CLEARFILE %v", Domain)

	dao.CheckDAOReadyState(Domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(Domain, "Clear", "INITIALISE")

	// Delete all active session recordList
	recordList, err := GetAll()
	if err != nil {
		logHandler.ErrorLogger.Print(ce.ErrDAOInitialisationWrapper(Domain, err).Error())
		clock.Stop(0)
		return ce.ErrDAOInitialisationWrapper(Domain, err)
	}

	//noRecords := len(recordList)
	count := 0
	logHandler.DatabaseLogger.Printf("Clearing %v records", len(recordList))

	for i, record := range recordList {
		logHandler.DatabaseLogger.Printf("(%v/%v) DELETE %v WHERE %v=%v", i+1, len(recordList), Domain, Fields.ID, record.ID)

		delErr := Delete(ctx, record.ID, fmt.Sprintf("Clearing %v %v @ initialisation ", Domain, record.ID))
		if delErr != nil {
			logHandler.ErrorLogger.Print(ce.ErrDAOInitialisationWrapper(Domain, delErr).Error())
			continue
		}
		count++
	}

	clock.Stop(count)
	logHandler.DatabaseLogger.Printf("Cleared down %v", Domain)
	return nil
}
