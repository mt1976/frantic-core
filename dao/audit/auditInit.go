package audit

import (
	"github.com/mt1976/frantic-core/commonConfig"
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
	SYNC         Action
)

func init() {
	CREATE = Action{code: "CREATE", description: "Create Data", silent: false, short: "CREATE"}
	DELETE = Action{code: "DELETE", description: "Delete Data", silent: false, short: "DELETE"}
	UPDATE = Action{code: "UPDATE", description: "Update Data", silent: false, short: "UPDATE"}
	ERASE = DELETE
	CLONE = Action{code: "CLONE", description: "Clone Data", silent: false, short: "CLONE"}
	NOTIFICATION = Action{code: "NOTIFY", description: "Notify Sent", silent: false, short: "NOTIFY"}
	SERVICE = Action{code: "RUN", description: "Service", silent: false, short: "RUN"}
	SILENT = Action{code: "SIL", description: "Silent Action", silent: true, short: "Silent"}
	GRANT = Action{code: "GRANT", description: "Grant Permission", silent: false, short: "GRANT"}
	REVOKE = Action{code: "REVOKE", description: "Revoke Permission", silent: false, short: "REVOKE"}
	PROCESS = Action{code: "PROCESS", description: "Process Run", silent: false, short: "PROCESS"}
	IMPORT = Action{code: "IMPORT", description: "Import Data", silent: false, short: "IMPORT"}
	EXPORT = Action{code: "EXPORT", description: "Export Data", silent: false, short: "EXPORT"}
	REPAIR = Action{code: "REPAIR", description: "Repair Data", silent: false, short: "REPAIR"}
	audit = Action{code: "AUDIT", description: "Audit", silent: true, short: "AUDIT"}
	BACKUP = Action{code: "BACKUP", description: "Backup Data", silent: true, short: "BACKUP"}
	LOGIN = Action{code: "LOGIN", description: "User Login", silent: false, short: "LOGIN"}
	LOGOUT = Action{code: "LOGOUT", description: "User Logout", silent: false, short: "LOGOUT"}
	SYNC = Action{code: "SYNC", description: "Data Synchronisation", silent: false, short: "SYNC"}
}
