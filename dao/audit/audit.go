package audit

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

// Audit package handles auditing of actions performed on data entities.
var name = "Audit"
var cfg *commonConfig.Settings
var upperName = strings.ToUpper(name)

func (a *Action) WithMessage(in string) Action {
	a.description = in
	return *a
}

func (a *Audit) Action(ctx context.Context, action Action) error {

	message := action.popMessage()
	timingMessage := fmt.Sprintf("Action: %v(%v) Message: %v", action.Code(), action.ShortName(), message)
	clock := timing.Start(name, "Audit", timingMessage)

	auditTime := time.Now()
	auditDisplay := dateHelpers.FormatAudit(auditTime)
	// auditUser := support.GetActiveUserCode()
	auditUser, err := getAuditUserCode(ctx)
	if err != nil {
		logHandler.WarningLogger.Printf("Action: %v(%v) Message: %v Error: %v", action.code, action.short, message, err)
	}
	auditHost := getHostName()

	if auditUser == "" {
		//	logHandler.WarningLogger.Printf("[%v] Error: %v", strings.ToUpper(name), "No Active User")
		logHandler.WarningLogger.Printf("Action: %v(%v) Message: %v Error: %v", action.code, action.short, message, "No Active User")
		os.Exit(1)
	}
	//updateAction := action

	if action.Is(CREATE) {
		a.CreatedAt = auditTime
		a.CreatedBy = auditUser
		a.CreatedOn = auditHost
		a.CreatedAtDisplay = auditDisplay
	}

	if action.Is(DELETE) {
		a.DeletedAt = auditTime
		a.DeletedBy = auditUser
		a.DeletedOn = auditHost
		a.DeletedAtDisplay = auditDisplay
	}

	if a.AuditSequence == 0 {
		a.AuditSequence = 1
	} else {
		a.AuditSequence++
	}

	update := AuditUpdateInfo{}

	update.UpdatedAt = auditTime
	update.UpdatedBy = auditUser
	update.UpdatedOn = auditHost
	update.UpdatedAtDisplay = auditDisplay
	update.UpdateAction = action.code
	update.UpdateNotes = message
	// a.DBVersion = dao.Version
	dbVersion := getDBVersion()
	a.DBVersion = dbVersion
	if !(action.Is(SERVICE) || action.Is(SILENT) || action.IsSilent()) {
		a.Updates = append(a.Updates, update)
	}

	logHandler.AuditLogger.Printf(AUDITMSG, upperName, action.code, auditDisplay, auditUser, auditHost, message)
	clock.Stop(1)
	return nil
}

func (a *Action) Is(inAction Action) bool {
	return a.code == inAction.code
}

func (a *Action) Silent() Action {
	a.silent = true
	return *a
}

func (a *Action) NotSilent() Action {
	a.silent = false
	return *a
}

func (a *Action) unSilience() Action {
	a.silent = false
	return *a
}

func (a *Action) IsSilent() bool {
	return a.silent
}

func (a *Action) Description() string {
	return a.description
}

func (a *Action) ShortNameRaw() string {
	return a.short
}

func (a *Action) ShortName() string {
	return strings.ToUpper(a.ShortNameRaw())
}

func (a *Action) Text() string {
	return strings.ToUpper(a.code)
}

func (a *Action) SetMessage(in string) {
	a.description = in
}

func (a *Action) GetMessage() string {
	return a.description
}

func (a *Action) SetText(in string) {
	a.code = in
}

func (a *Action) Code() string {
	return a.code
}
