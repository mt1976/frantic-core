package contextHandler

import (
	"context"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
)

func Debug(c context.Context, name string) {
	name = strings.ToUpper("DEBUG-" + name)
	logHandler.Trace.Printf("[DEBUG] Session Context %v", name)
	logHandler.Trace.Printf("[DEBUG] Session Context %v", name)

	logHandler.Trace.Printf("[%v] Session Context: %+v", name, c)

	logHandler.Trace.Printf("[%v] Session ID: [%v]=[%v]", strings.ToUpper(name), sessionIDKey.name, GetSession_ID(c))

	logHandler.Trace.Printf("[%v] Session U Code: [%v]=[%v]", strings.ToUpper(name), userCodeKey.name, GetSession_UserCode(c))
	logHandler.Trace.Printf("[%v] Session U Key: [%v]=[%v]", strings.ToUpper(name), userKeyKey.name, GetSession_UserKey(c))
	logHandler.Trace.Printf("[%v] Session U Role: [%v]=[%v]", strings.ToUpper(name), userRoleKey.name, GetSession_UserRole(c))

	logHandler.Trace.Printf("[%v] Session Expiry: [%v]=[%v]", strings.ToUpper(name), expiryPeriodKey.name, GetSession_Expiry(c))
	logHandler.Trace.Printf("[%v] Session Theme: [%v]=[%v]", strings.ToUpper(name), themeKey.name, GetSession_Theme(c))
	logHandler.Trace.Printf("[%v] Session Locale: [%v]=[%v]", strings.ToUpper(name), localeKey.name, GetSession_Locale(c))
	logHandler.Trace.Printf("[%v] Session Timezone: [%v]=[%v]", strings.ToUpper(name), timezoneKey.name, GetSession_Timezone(c))
	logHandler.Trace.Printf("[%v] Session TOKEN", name)
	logHandler.Trace.Printf("[%v] Session Token: [%v]=[%+v]", strings.ToUpper(name), tokenKey.name, GetSession_Token(c))
	logHandler.Trace.Printf("[%v] DEBUG ENDS", name)
	logHandler.Trace.Printf("[%v] DEBUG ENDS", name)
}
