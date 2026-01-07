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
	"strings"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// insertOrUpdate inserts or updates the TemplateStore record in the database based on the specified operation.
//
// Parameters:
//
//	ctx context.Context: The context for managing request-scoped values, cancellation signals, and deadlines.
//	note string: A note to be included in the audit log.
//	activity string: The activity code for timing purposes.
//	auditAction audit.Action: The audit action to be recorded.
//	operation string: The operation type, either "CREATE" or "UPDATE".
//
// Returns:
//
//	error: An error object if any issues occur during the insert or update process; otherwise, nil.
func (record *TemplateStore) insertOrUpdate(ctx context.Context, note, activity string, auditAction audit.Action, operation string) error {

	isCreateOperation := false
	if strings.EqualFold(operation, actions.CREATE.GetCode()) {
		isCreateOperation = true
		if !strings.EqualFold(auditAction.Code(), actions.CREATE.GetCode()) {
			return commonErrors.WrapDAOUpdateError(Domain, fmt.Errorf("invalid audit action '%v' for create event '%v'", auditAction.Code(), operation))
		}
	}

	dao.CheckDAOReadyState(Domain, auditAction, initialised) // Check the DAO has been initialised, Mandatory.

	clock := timing.Start(Domain, activity, fmt.Sprintf("%v", record.ID))
	if isCreateOperation {
		if err := record.checkForDuplicate(); err != nil {
			clock.Stop(0)
			return commonErrors.WrapDAOCreateError(Domain, record.ID, err)
		}
	}

	if calculationError := record.defaultProcessing(); calculationError != nil {
		rtnErr := commonErrors.WrapDAOCaclulationError(Domain, calculationError)
		logHandler.ErrorLogger.Print(rtnErr.Error())
		clock.Stop(0)
		return rtnErr
	}

	if validationError := record.validationProcessing(); validationError != nil {
		valErr := commonErrors.WrapDAOValidationError(Domain, validationError)
		logHandler.ErrorLogger.Print(valErr.Error())
		clock.Stop(0)
		return valErr
	}

	auditErr := record.Audit.Action(ctx, auditAction.WithMessage(note))
	if auditErr != nil {
		audErr := commonErrors.WrapDAOUpdateAuditError(Domain, record.ID, auditErr)
		logHandler.ErrorLogger.Print(audErr.Error())
		clock.Stop(0)
		return audErr
	}
	var actionError error
	if isCreateOperation {
		logHandler.TraceLogger.Printf("Creating %v record %v %v", Domain, record.Key, record.ID)
		actionError = activeDB.Create(record)
	} else {
		logHandler.TraceLogger.Printf("Updating %v record %v %v", Domain, record.Key, record.ID)
		actionError = activeDB.Update(record)
	}
	if actionError != nil {

		godump.Dump(record)

		updErr := commonErrors.WrapDAOUpdateError(Domain, actionError)
		logHandler.ErrorLogger.Panic(updErr.Error(), actionError)
		clock.Stop(0)
		return updErr
	}

	clock.Stop(1)

	return nil
}

// postGetList performs post-retrieval processing on a list of TemplateStore records.
//
// Parameters:
//
//	recordList *[]TemplateStore: A pointer to a slice of TemplateStore records to be processed.
//
// Returns:
//
//	[]TemplateStore: A slice of processed TemplateStore records.
//	error: An error object if any issues occur during the post-retrieval processing; otherwise, nil.
func postGetList(recordList *[]TemplateStore) ([]TemplateStore, error) {
	clock := timing.Start(Domain, actions.PROCESS.GetCode(), "POSTGET")
	returnList := []TemplateStore{}
	for _, record := range *recordList {
		if err := record.postGet(); err != nil {
			return nil, err
		}
		returnList = append(returnList, record)
	}
	clock.Stop(len(returnList))
	return returnList, nil
}

// postGet performs post-retrieval processing on a TemplateStore record.
//
// Returns:
//
//	error: An error object if any issues occur during the post-retrieval processing; otherwise, nil.
func (record *TemplateStore) postGet() error {

	upgradeError := record.upgradeProcessing()
	if upgradeError != nil {
		return upgradeError
	}

	defaultingError := record.defaultProcessing()
	if defaultingError != nil {
		return defaultingError
	}

	validationError := record.validationProcessing()
	if validationError != nil {
		return validationError
	}

	return record.postGetProcessing()
}

// checkForDuplicate checks for duplicate TemplateStore records in the database.
//
// Returns:
//
//	error: An error object if a duplicate record is found; otherwise, nil.
func (record *TemplateStore) checkForDuplicate() error {

	dao.CheckDAOReadyState(Domain, audit.PROCESS, initialised) // Check the DAO has been initialised, Mandatory.

	// Get all status
	responseRecord, err := GetBy(Fields.Key, record.Key)
	if err != nil {
		// This is ok, no record could be read
		return nil
	}

	if responseRecord.Audit.DeletedBy != "" {
		return nil
	}

	logHandler.WarningLogger.Printf("Duplicate %v, %v already in use", Domain, record.ID)
	return commonErrors.ErrorDuplicate
}
