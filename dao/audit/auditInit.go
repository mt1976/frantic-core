package audit

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
)

// init initializes the audit package
func init() {
	cfg = commonConfig.Get()
}

const (
	AUDITMSG = "[%v] Action: %s At: %v By: %v On: %v Notes: %v"
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
