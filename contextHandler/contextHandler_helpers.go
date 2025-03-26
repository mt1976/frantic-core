package contextHandler

import (
	"context"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
)

func Debug(c context.Context, name string) {
	name = strings.ToUpper("DEBUG-" + name)
	logHandler.TraceLogger.Printf("[DEBUG] Session Context: %+v", c)

	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionID", GetSession_ID(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionIdentifier", GetSession_Identifier())

	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "userCode", GetSession_UserCode(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "userKey", GetSession_UserKey(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "userRole", GetSession_UserRole(c))

	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionExpiry", GetSession_Expiry(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionTheme", GetSession_Theme(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionLocale", GetSession_Locale(c))
	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%v]", strings.ToUpper(name), "sessionTimezone", GetSession_Timezone(c))

	logHandler.TraceLogger.Printf("[%v] Session Context: [%v]=[%+v]", strings.ToUpper(name), "sessionToken", GetSession_Token(c))
	logHandler.TraceLogger.Printf("[DEBUG] Session Context: %+v", c)

}
