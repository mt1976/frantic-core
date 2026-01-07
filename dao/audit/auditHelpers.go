package audit

import (
	"context"
	"sync"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/contextHandler"
	"github.com/mt1976/frantic-core/logHandler"
)

var hostNameOnce sync.Once
var cachedHostName string

func getHostName() string {
	hostNameOnce.Do(func() {
		cachedHostName = application.HostName()
	})
	return cachedHostName
}

// getDBVersion retrieves the current database version
func getDBVersion() int {
	// Implement the logic to get the DB version without importing the dao package
	return cfg.GetDatabase_Version()
}

func (a *Action) popMessage() string {
	message := a.description
	a.description = ""
	return message
}

func (a *Audit) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Updates)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := range a.Updates {
			upd := a.Updates[i]
			logHandler.TraceLogger.Printf(AUDITMSG, upperName, upd.UpdateAction, upd.UpdatedAtDisplay, upd.UpdatedBy, upd.UpdatedOn, upd.UpdateNotes)
		}
	}
	return nil
}

func getAuditUserCode(ctx context.Context) (string, error) {

	defaultUser := cfg.GetServiceUser_UserCode()

	if ctx == context.Background() {
		usr := "svc_" + getHostName()
		return usr, nil
	}

	// Implement the logic to get the user without importing the dao package
	if ctx == context.TODO() || ctx == nil {
		usr := "sys_" + getHostName()
		return usr, nil
	}

	// Get the current user from the context
	sessionUser := contextHandler.GetSession_UserCode(ctx)
	//ctx.Value(cfg.GetSecuritySessionKey_UserCode())
	if sessionUser != "" {
		return sessionUser, nil
	}
	return defaultUser, commonErrors.ErrContextCannotGetUserCode
}
