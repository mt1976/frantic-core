package contextHandler

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/idHelpers"
)

type contextKey struct {
	name                string
	pointlessIdentifier string
}

var cfg *commonConfig.Settings

var (
	// / SECURITY SESSION KEYS
	sessionIDKey    contextKey
	userKeyKey      contextKey
	userCodeKey     contextKey
	tokenKey        contextKey
	expiryPeriodKey contextKey
	localeKey       contextKey
	themeKey        contextKey
	timezoneKey     contextKey
	userRoleKey     contextKey
)

// NewFartFarmer is a constructor for the fartFarmer struct
func new(in string) contextKey {
	var out contextKey
	out.name = in
	out.pointlessIdentifier = idHelpers.Encode(in)
	return out
}

func init() {
	cfg = commonConfig.Get()
	sessionIDKey = new(cfg.GetSecuritySessionKey_Session())
	userKeyKey = new(cfg.GetSecuritySessionKey_UserKey())
	userCodeKey = new(cfg.GetSecuritySessionKey_UserCode())
	tokenKey = new(cfg.GetSecuritySessionKey_Token())
	expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())
	localeKey = new(cfg.GetSecuritySessionKey_Locale())
	themeKey = new(cfg.GetSecuritySessionKey_Theme())
	timezoneKey = new(cfg.GetSecuritySessionKey_Timezone())
	userRoleKey = new(cfg.GetSecuritySessionKey_UserRole())
}
