package audit

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/contextHandler"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var name = "Audit"
var cfg *commonConfig.Settings

type Action struct {
	code        string
	short       string
	description string
	silent      bool
}

func init() {
	cfg = commonConfig.Get()
}

func getDBVersion() int {
	// Implement the logic to get the DB version without importing the dao package
	return cfg.GetDatabase_Version()
}

type Audit struct {
	CreatedAt        time.Time
	CreatedBy        string
	CreatedOn        string
	CreatedAtDisplay string
	Updates          []AuditUpdateInfo
	DeletedAt        time.Time
	DeletedBy        string
	DeletedOn        string
	DeletedAtDisplay string
	AuditSequence    int
	DBVersion        int
	//Empty     time.Time // Convience Field - Used to avoid erros with dates.
}

type AuditUpdateInfo struct {
	UpdatedAt        time.Time
	UpdateAction     string
	UpdatedBy        string
	UpdatedOn        string
	UpdatedAtDisplay string
	UpdateNotes      string
}

const (
	AUDITMSG = "[%v] Action=[%s] At=[%v] By=[%v] On=[%v] Notes=[%v]"
)

var (
	CREATE       Action
	DELETE       Action
	UPDATE       Action
	ERASE        Action
	CLONE        Action
	NOTIFICATION Action
	SERVICE      Action
	SILENT       Action
	GRANT        Action
	REVOKE       Action
	PROCESS      Action
	IMPORT       Action
	EXPORT       Action
	GET          Action
	REPAIR       Action
	audit        Action
	BACKUP       Action
	LOGIN        Action
	LOGOUT       Action
)

func init() {
	CREATE = Action{code: actions.CREATE.GetCode(), description: actions.CREATE.GetDescription("Data"), silent: false, short: actions.CREATE.GetShortName()}
	DELETE = Action{code: actions.DELETE.GetCode(), description: actions.DELETE.GetDescription("Data"), silent: false, short: actions.DELETE.GetShortName()}
	UPDATE = Action{code: actions.UPDATE.GetCode(), description: actions.UPDATE.GetDescription("Data"), silent: false, short: actions.UPDATE.GetShortName()}
	ERASE = DELETE
	CLONE = Action{code: actions.CLONE.GetCode(), description: actions.CLONE.GetDescription("Data"), silent: false, short: actions.CLONE.GetShortName()}
	NOTIFICATION = Action{code: actions.NOTIFY.GetCode(), description: actions.NOTIFY.GetDescription("Sent"), silent: false, short: actions.NOTIFY.GetShortName()}
	SERVICE = Action{code: actions.RUN.GetCode(), description: actions.RUN.GetDescription("Service"), silent: false, short: actions.RUN.GetShortName()}
	SILENT = Action{code: "SIL", description: "Silent Action", silent: true, short: "Silent"}
	GRANT = Action{code: actions.GRANT.GetCode(), description: actions.GRANT.GetDescription(""), silent: false, short: actions.GRANT.GetShortName()}
	REVOKE = Action{code: actions.REVOKE.GetCode(), description: actions.REVOKE.GetDescription(""), silent: false, short: actions.REVOKE.GetShortName()}
	PROCESS = Action{code: actions.PROCESS.GetCode(), description: actions.PROCESS.GetDescription("Run"), silent: false, short: actions.PROCESS.GetShortName()}
	IMPORT = Action{code: actions.IMPORT.GetCode(), description: actions.IMPORT.GetDescription("Data"), silent: false, short: actions.IMPORT.GetShortName()}
	EXPORT = Action{code: actions.EXPORT.GetCode(), description: actions.EXPORT.GetDescription("Data"), silent: false, short: actions.EXPORT.GetShortName()}
	REPAIR = Action{code: actions.REPAIR.GetCode(), description: actions.REPAIR.GetDescription("Data"), silent: false, short: actions.REPAIR.GetShortName()}
	audit = Action{code: actions.AUDIT.GetCode(), description: actions.AUDIT.GetDescription("Audit"), silent: true, short: actions.AUDIT.GetShortName()}
	BACKUP = Action{code: actions.BACKUP.GetCode(), description: actions.BACKUP.GetDescription("Data"), silent: true, short: actions.BACKUP.GetShortName()}
	LOGIN = Action{code: actions.LOGIN.GetCode(), description: actions.LOGIN.GetDescription("User"), silent: false, short: actions.LOGIN.GetShortName()}
	LOGOUT = Action{code: actions.LOGOUT.GetCode(), description: actions.LOGOUT.GetDescription("User"), silent: false, short: actions.LOGOUT.GetShortName()}
}

func (a *Action) WithMessage(in string) Action {
	a.description = in
	return *a
}

func (a *Action) popMessage() string {
	message := a.description
	a.description = ""
	return message
}

func (a *Audit) Action(ctx context.Context, action Action) error {

	message := action.popMessage()
	timingMessage := fmt.Sprintf("Action=[%v] Message=[%v]", action.Code(), message)
	clock := timing.Start("Audit", actions.AUDIT.GetCode(), timingMessage)

	auditTime := time.Now()
	auditDisplay := dateHelpers.FormatAudit(auditTime)
	// auditUser := support.GetActiveUserCode()
	auditUser, err := getAuditUserCode(ctx)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err)
	}
	auditHost := application.HostName()

	if auditUser == "" {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), "No Active User")
		logHandler.InfoLogger.Printf("[%v] Action=[%v] Message=[%v]", strings.ToUpper(name), action.code, message)
		os.Exit(0)
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
	a.DBVersion = getDBVersion()
	if !(action.Is(SERVICE) || action.Is(SILENT) || action.IsSilent()) {
		a.Updates = append(a.Updates, update)
	}

	a.DBVersion = getDBVersion()

	logHandler.AuditLogger.Printf(AUDITMSG, strings.ToUpper(name), action.code, auditDisplay, auditUser, auditHost, message)
	clock.Stop(1)
	return nil
}

func (a *Audit) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Updates)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := 0; i < noAudit; i++ {
			upd := a.Updates[i]
			logHandler.TraceLogger.Printf(AUDITMSG, strings.ToUpper(name), upd.UpdateAction, upd.UpdatedAtDisplay, upd.UpdatedBy, upd.UpdatedOn, upd.UpdateNotes)
		}
	}
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

func getAuditUserCode(ctx context.Context) (string, error) {
	// Implement the logic to get the user without importing the dao package
	defaultUser := cfg.GetServiceUser_UserCode()
	if ctx == context.TODO() || ctx == nil {
		return defaultUser, nil
	}

	// Get the current user from the context
	sessionUser := contextHandler.GetUserCode(ctx)
	//ctx.Value(cfg.GetSecuritySessionKey_UserCode())
	if sessionUser != "" {
		return sessionUser, nil
	}
	return defaultUser, commonErrors.ErrContextCannotGetUserCode
}
