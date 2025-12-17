package contextHandler

import (
	"context"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
)

// sessionKey      = new(cfg.GetSecuritySessionKey_Session())
// userKeyKey      = new(cfg.GetSecuritySessionKey_UserKey())
// userCodeKey     = new(cfg.GetSecuritySessionKey_UserCode())
// tokenKey        = new(cfg.GetSecuritySessionKey_Token())
// expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())

func GetSession_UserCode(ctx context.Context) string {
	value := ctx.Value(userCodeKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("User code (%v) requested but not found in context, returning empty string", userCodeKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_UserKey(ctx context.Context) string {
	value := ctx.Value(userKeyKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("User key (%v) requested but not found in context, returning empty string", userKeyKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_ID(ctx context.Context) string {
	value := ctx.Value(sessionIDKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session ID (%v) requested but not found in context, returning empty string", sessionIDKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_Token(ctx context.Context) any {
	value := ctx.Value(tokenKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session token (%v) requested but not found in context, returning nil", tokenKey.name)
		return nil
	}
	return value
}

func GetSession_Expiry(ctx context.Context) time.Time {
	value := ctx.Value(expiryPeriodKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session expiry (%v) requested but not found in context, returning zero time", expiryPeriodKey.name)
		return time.Time{}
	}
	return value.(time.Time)
}

func GetSession_Identifier() string {
	return sessionIDKey.name
}

func GetSession_Locale(ctx context.Context) string {
	value := ctx.Value(localeKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session locale (%v) requested but not found in context, returning default", localeKey.name)
		return cfg.GetApplication_Locale()
	}
	return value.(string)
}

func GetSession_Theme(ctx context.Context) string {
	value := ctx.Value(themeKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session theme (%v) requested but not found in context, returning empty string", themeKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_Timezone(ctx context.Context) string {
	value := ctx.Value(timezoneKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("Session timezone (%v) requested but not found in context, returning empty string", timezoneKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_UserRole(ctx context.Context) string {
	value := ctx.Value(userRoleKey.name)
	if value == nil {
		logHandler.WarningLogger.Printf("User role (%v) requested but not found in context, returning empty string", userRoleKey.name)
		return ""
	}
	return value.(string)
}

// Setters

func SetSession_ID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func SetSession_UserKey(ctx context.Context, userKey string) context.Context {
	return context.WithValue(ctx, userKeyKey, userKey)
}

func SetSession_UserCode(ctx context.Context, userCode string) context.Context {
	return context.WithValue(ctx, userCodeKey, userCode)
}

func SetSession_Token(ctx context.Context, token any) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func SetSession_Expiry(ctx context.Context, expiry time.Time) context.Context {
	return context.WithValue(ctx, expiryPeriodKey, expiry)
}

func SetSession_Locale(ctx context.Context, locale string) context.Context {
	return context.WithValue(ctx, localeKey, locale)
}

func SetSession_Theme(ctx context.Context, theme string) context.Context {
	return context.WithValue(ctx, themeKey, theme)
}

func SetSession_Timezone(ctx context.Context, timezone string) context.Context {
	return context.WithValue(ctx, timezoneKey, timezone)
}

func SetSession_UserRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, userRoleKey, role)
}
