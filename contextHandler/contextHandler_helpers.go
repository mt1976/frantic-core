package contextHandler

import (
	"context"
	"fmt"
	"os/user"
	"runtime"
	"strings"

	"github.com/alitto/pond/v2"
	"github.com/goforj/godump"
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

func AddUserContext(ctx context.Context, userID, userName string) context.Context {
	existingUserID := GetSession_UserKey(ctx)
	existingUserName := GetSession_UserCode(ctx)

	logHandler.Trace.Printf("Existing User Context: %v=%v - %v=%v", "userID", existingUserID, "userName", existingUserName)
	if existingUserID != "" && existingUserName != "" {
		logHandler.Event.Printf("User Context already exists: %v=%v - %v=%v", "userID", existingUserID, "userName", existingUserName)
		return ctx
	}

	if userID == "" || userName == "" {
		currentUser := GetUserDetails()
		userID = currentUser.Uid
		userName = currentUser.Username
	}

	userCode := userID + "_" + userName
	userKey := userID

	ctx = SetSession_UserKey(ctx, userKey)
	ctx = SetSession_UserCode(ctx, userCode)

	logHandler.Trace.Printf("User Context Added: %v(%v)=%v(%v)", userCode, GetSession_UserCode(ctx), userKey, GetSession_UserKey(ctx))

	return ctx
}

func GetUserDetails() *user.User {
	currentUser, err := user.Current()
	if err != nil {
		logHandler.Error.Fatalln(err.Error())
	}
	// if running on windows, the UID is in the format of "S-1-5-21-3849575818-2607088806-4266749144", we need to convert it to a more readable format
	if runningOnWindows() {
		// Example UID on windows "S-1-5-21-3849575818-2607088806-4266749144"
		// Tokenize the UID
		tkID := strings.Split(currentUser.Uid, "-")
		// Concatenate the token in 0, 1, 2, 3
		currentUser.Uid = fmt.Sprintf("%s%s%s%s", tkID[0], tkID[1], tkID[2], tkID[3])
	}
	return currentUser
}

func runningOnWindows() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "windows")
}

func NewWorkerPool(ctx context.Context, poolName string, poolSize int) context.Context {
	// Define a reasonable number of workers for the worker pool. This can be adjusted based on the expected load and performance requirements. For a typical application, starting with 10 workers is a good balance between concurrency and resource usage.
	noWorkers := poolSize

	logHandler.Event.Println("Creating Worker Pool...")
	ctx = addWorkerPoolName(ctx, poolName)
	workerPool := pond.NewPool(noWorkers, pond.WithContext(ctx), pond.WithQueueSize(noWorkers*3))
	// Now add the pool to the context so it can be used by the jobs
	ctx = context.WithValue(ctx, WorkerPoolKey, workerPool)
	defer workerPool.StopAndWait()
	logHandler.Event.Printf("Worker Pool Created with %d workers", noWorkers)

	return ctx
}

var workerPoolNameKey = new("WorkerPoolName")

func addWorkerPoolName(ctx context.Context, poolName string) context.Context {
	logHandler.Trace.Printf("Setting Worker Pool Name in Context: %v=%v", workerPoolNameKey.name, poolName)
	return context.WithValue(ctx, workerPoolNameKey, poolName)
}

func GetWorkerPoolName(ctx context.Context) string {
	logHandler.Trace.Printf("Retrieving Worker Pool Name from Context: %v", godump.DumpStr(ctx))
	value := ctx.Value(workerPoolNameKey)
	if value == nil {
		logHandler.Warning.Printf("Worker pool name (%v) requested but not found in context, returning empty string", workerPoolNameKey.name)
		return ""
	}
	return value.(string)
}
