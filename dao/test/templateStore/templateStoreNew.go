package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the New function to implement the creation of a new domain entity
//TODO: Create any new functions required to support the domain entity

import (
	"context"
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// New creates a new template instance
func New() TemplateStore {
	return TemplateStore{}
}

// Create creates a new template instance in the database
// It takes name as a parameter and returns the created template instance or an error if any occurs
// It also checks if the DAO is ready for operations
// It records the creation action in the audit data and saves the instance to the database
// Parameters: (all but ctx are used to create a new template instance and should be replaced as needed)
//   - ctx context.Context: The context for managing request-scoped values, cancellation signals, and deadlines.
//   - userName string: The user name for the new template instance.
//   - uid string: The UID for the new template instance.
//   - realName string: The real name for the new template instance.
//   - email string: The email for the new template instance.
//   - gid string: The GID for the new template instance.
//
// Returns:
//   - TemplateStore: The created template instance.
//   - error: An error object if any issues occur during the creation process; otherwise, nil.
func Create(ctx context.Context, userName, uid, realName, email, gid string) (TemplateStore, error) {

	dao.CheckDAOReadyState(Domain, audit.CREATE, initialised) // Check the DAO has been initialised, Mandatory.

	//logHandler.InfoLogger.Printf("New %v (%v=%v)", domain, Fields.ID, field1)
	clock := timing.Start(Domain, "Create", fmt.Sprintf("%v", userName))

	sessionID := idHelpers.GetUUID()

	// Create a new struct
	record := TemplateStore{}
	record.Key = idHelpers.Encode(sessionID)
	record.Raw = sessionID
	record.UserName = userName
	record.UID = uid
	record.RealName = realName
	record.Email = email
	record.GID = gid
	record.Active.IsTrue()
	record.LastLogin = time.Time{}
	record.LastHost = ""
	record.UserCode = record.buildUserCode()

	// Record the create action in the audit data
	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", Domain, userName)))
	if auditErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOUpdateAuditWrapper(Domain, record.ID, auditErr))
	}

	// Save the status instance to the database
	writeErr := activeDB.Create(&record)
	if writeErr != nil {
		// Log and panic if there is an error creating the status instance
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOCreateWrapper(Domain, record.ID, writeErr))
		//	panic(writeErr)
	}

	//logHandler.AuditLogger.Printf("[%v] [%v] ID=[%v] Notes[%v]", audit.CREATE, domain, record.ID, fmt.Sprintf("New %v: %v", domain, field1))
	clock.Stop(1)
	return record, nil
}
