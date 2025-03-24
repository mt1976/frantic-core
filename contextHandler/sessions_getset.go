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
	value := ctx.Value(userCodeKey)
	if value == nil {
		logHandler.WarningLogger.Printf("User code requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSession_UserKey(ctx context.Context) string {
	value := ctx.Value(userKeyKey)
	if value == nil {
		logHandler.WarningLogger.Printf("User key requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSession_ID(ctx context.Context) string {
	value := ctx.Value(sessionIDKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session ID requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSession_Token(ctx context.Context) any {
	value := ctx.Value(tokenKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session token requested but not found in context, returning nil")
		return nil
	}
	return value
}

func GetSession_Expiry(ctx context.Context) time.Time {
	value := ctx.Value(expiryPeriodKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session expiry requested but not found in context, returning zero time")
		return time.Time{}
	}
	return value.(time.Time)
}

func GetSession_Identifier() string {
	return sessionIDKey.name
}

func GetSession_Locale(ctx context.Context) string {
	value := ctx.Value(localeKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session locale requested but not found in context, returning default")
		return cfg.GetApplication_Locale()
	}
	return value.(string)
}

func GetSession_Theme(ctx context.Context) string {
	value := ctx.Value(themeKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session theme requested but not found in context, returning empty string")
		return ""
	}
	return value.(string)
}

func GetSession_Timezone(ctx context.Context) string {
	value := ctx.Value(timezoneKey)
	if value == nil {
		logHandler.WarningLogger.Printf("Session timezone requested but not found in context, returning empty string")
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
