package templateStoreV2

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

// New returns an empty TemplateStore record.
func New() TemplateStore {
	return TemplateStore{}
}

// Create constructs and inserts a new TemplateStore record.
func Create(ctx context.Context, userName, uid, realName, email, gid string) (TemplateStore, error) {
	dao.CheckDAOReadyState(tableName, audit.CREATE, databaseConnectionActive)

	clock := timing.Start(tableName, "Create", fmt.Sprintf("%v", userName))

	sessionID := idHelpers.GetUUID()

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

	auditErr := record.Audit.Action(ctx, audit.CREATE.WithMessage(fmt.Sprintf("New %v created %v", tableName, userName)))
	if auditErr != nil {
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOUpdateAuditWrapper(tableName, record.ID, auditErr))
	}

	writeErr := activeDBConnection.Create(&record)
	if writeErr != nil {
		logHandler.ErrorLogger.Panic(commonErrors.ErrDAOCreateWrapper(tableName, record.ID, writeErr))
	}

	clock.Stop(1)
	return record, nil
}
