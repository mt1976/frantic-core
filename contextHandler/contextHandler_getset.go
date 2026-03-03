package contextHandler

import (
	"context"
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
)

// sessionKey      = new(cfg.GetSecuritySessionKey_Session())
// userKeyKey      = new(cfg.GetSecuritySessionKey_UserKey())
// userCodeKey     = new(cfg.GetSecuritySessionKey_UserCode())
// tokenKey        = new(cfg.GetSecuritySessionKey_Token())
// expiryPeriodKey = new(cfg.GetSecuritySessionKey_ExpiryPeriod())

// func AddWorkerPoolToContext(ctx context.Context, pool pond.Pool) context.Context {
// 	logHandler.Trace.Printf("Setting Worker Pool in Context: %v=%v", WorkerPoolKey.name, pool)
// 	return context.WithValue(ctx, WorkerPoolKey, pool)
// }

func GetWorkerPool(ctx context.Context) (pond.Pool, error) {
	logHandler.Trace.Printf("Retrieving Worker Pool from Context: %v", godump.DumpStr(ctx))
	value := ctx.Value(WorkerPoolKey)
	if value == nil {
		logHandler.Error.Fatalf("Worker pool (%v) requested but not found in context, returning nil", WorkerPoolKey.name)
		return nil, fmt.Errorf("Worker pool requested but not found in context")
		// panic("Worker pool requested but not found in context")
	}
	// Coearce the value to the expected type (e.g., *pond.WorkerPool)
	return value.(pond.Pool), nil
}

func GetSession_UserCode(ctx context.Context) string {
	value := ctx.Value(userCodeKey)
	if value == nil {
		logHandler.Warning.Printf("User code (%v) requested but not found in context, returning empty string", userCodeKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_UserKey(ctx context.Context) string {
	value := ctx.Value(userKeyKey)
	if value == nil {
		logHandler.Warning.Printf("User key (%v) requested but not found in context, returning empty string", userKeyKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_ID(ctx context.Context) string {
	value := ctx.Value(sessionIDKey)
	if value == nil {
		logHandler.Warning.Printf("Session ID (%v) requested but not found in context, returning empty string", sessionIDKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_Token(ctx context.Context) any {
	value := ctx.Value(tokenKey)
	if value == nil {
		logHandler.Warning.Printf("Session token (%v) requested but not found in context, returning nil", tokenKey.name)
		return nil
	}
	return value
}

func GetSession_Expiry(ctx context.Context) time.Time {
	value := ctx.Value(expiryPeriodKey)
	if value == nil {
		logHandler.Warning.Printf("Session expiry (%v) requested but not found in context, returning zero time", expiryPeriodKey.name)
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
		logHandler.Warning.Printf("Session locale (%v) requested but not found in context, returning default", localeKey.name)
		return cfg.GetApplication_Locale()
	}
	return value.(string)
}

func GetSession_Theme(ctx context.Context) string {
	value := ctx.Value(themeKey)
	if value == nil {
		logHandler.Warning.Printf("Session theme (%v) requested but not found in context, returning empty string", themeKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_Timezone(ctx context.Context) string {
	value := ctx.Value(timezoneKey)
	if value == nil {
		logHandler.Warning.Printf("Session timezone (%v) requested but not found in context, returning empty string", timezoneKey.name)
		return ""
	}
	return value.(string)
}

func GetSession_UserRole(ctx context.Context) string {
	value := ctx.Value(userRoleKey)
	if value == nil {
		logHandler.Warning.Printf("User role (%v) requested but not found in context, returning empty string", userRoleKey.name)
		return ""
	}
	return value.(string)
}

// Setters

func SetSession_ID(ctx context.Context, sessionID string) context.Context {
	logHandler.Trace.Printf("Setting Session ID in Context: %v=%v", sessionIDKey.name, sessionID)
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func SetSession_UserKey(ctx context.Context, userKey string) context.Context {
	logHandler.Trace.Printf("Setting User Key in Context: %v=%v", userKeyKey.name, userKey)
	return context.WithValue(ctx, userKeyKey, userKey)
}

func SetSession_UserCode(ctx context.Context, userCode string) context.Context {
	logHandler.Trace.Printf("Setting User Code in Context: %v=%v", userCodeKey.name, userCode)
	return context.WithValue(ctx, userCodeKey, userCode)
}

func SetSession_Token(ctx context.Context, token any) context.Context {
	logHandler.Trace.Printf("Setting Session Token in Context: %v=%v", tokenKey.name, token)
	return context.WithValue(ctx, tokenKey, token)
}

func SetSession_Expiry(ctx context.Context, expiry time.Time) context.Context {
	logHandler.Trace.Printf("Setting Session Expiry in Context: %v=%v", expiryPeriodKey.name, expiry)
	return context.WithValue(ctx, expiryPeriodKey, expiry)
}

func SetSession_Locale(ctx context.Context, locale string) context.Context {
	logHandler.Trace.Printf("Setting Session Locale in Context: %v=%v", localeKey.name, locale)
	return context.WithValue(ctx, localeKey, locale)
}

func SetSession_Theme(ctx context.Context, theme string) context.Context {
	logHandler.Trace.Printf("Setting Session Theme in Context: %v=%v", themeKey.name, theme)
	return context.WithValue(ctx, themeKey, theme)
}

func SetSession_Timezone(ctx context.Context, timezone string) context.Context {
	logHandler.Trace.Printf("Setting Session Timezone in Context: %v=%v", timezoneKey.name, timezone)
	return context.WithValue(ctx, timezoneKey, timezone)
}

func SetSession_UserRole(ctx context.Context, role string) context.Context {
	logHandler.Trace.Printf("Setting User Role in Context: %v=%v", userRoleKey.name, role)
	return context.WithValue(ctx, userRoleKey, role)
}

func AddWorkerPoolToContext(ctx context.Context, pool pond.Pool) context.Context {
	logHandler.Trace.Printf("Setting Worker Pool in Context: %v=%v", WorkerPoolKey.name, pool)
	return context.WithValue(ctx, WorkerPoolKey, pool)
}
