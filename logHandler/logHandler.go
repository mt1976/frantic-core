package logHandler

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/paths"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Warning        *log.Logger
	Info           *log.Logger
	Error          *log.Logger
	Panic          *log.Logger
	Timing         *log.Logger
	Event          *log.Logger
	Service        *log.Logger
	Trace          *log.Logger
	Audit          *log.Logger
	Translation    *log.Logger
	Security       *log.Logger
	Database       *log.Logger
	Api            *log.Logger
	Import         *log.Logger
	Export         *log.Logger
	Communications *log.Logger
	Lock           *log.Logger
	Unlock         *log.Logger
	SkipLock       *log.Logger
	Cache          *log.Logger
	Web            *log.Logger
)

var (
	Reset   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	Gray    string
	White   string
)

func init() {
	settings := commonConfig.Get()
	// applicationPath := "data/logs/"
	applicationPath := paths.Application().String()
	applicationPath += paths.Logs().String()
	applicationPath += string(os.PathSeparator)
	applicationPath += settings.GetApplication_Name() + "-"

	maxSize := settings.GetLogging_MaxSize()
	maxBackups := settings.GetLogging_MaxBackups()
	maxAge := settings.GetLogging_MaxAge()
	compress := settings.IsLogCompressionEnabled()

	setColoursNormal()
	if runtime.GOOS == "windows" {
		setColoursWindows()
	}

	generalWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "general"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsGeneralLoggingDisabled() || settings.IsLoggingDisabled() {
		generalWriter = io.MultiWriter(io.Discard)
	}

	timingWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "timing"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsTimingLoggingDisabled() || settings.IsLoggingDisabled() {
		timingWriter = io.MultiWriter(io.Discard)
	}

	serviceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "worker"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsServiceLoggingDisabled() || settings.IsLoggingDisabled() {
		serviceWriter = io.MultiWriter(io.Discard)
	}

	auditWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "audit"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsAuditLoggingDisabled() || settings.IsLoggingDisabled() {
		auditWriter = io.MultiWriter(io.Discard)
	}

	errorWriter := io.MultiWriter(os.Stderr, os.Stdout, os.Stdout, os.Stdout, os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "error"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsLoggingDisabled() {
		errorWriter = io.MultiWriter(io.Discard)
	}

	panicWriter := io.MultiWriter(os.Stderr, os.Stdout, os.Stdout, os.Stdout, os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "panic"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsLoggingDisabled() {
		panicWriter = io.MultiWriter(io.Discard)
	}

	translationWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "translation"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsTranslationLoggingDisabled() || settings.IsLoggingDisabled() {
		translationWriter = io.MultiWriter(io.Discard)
	}

	traceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "trace"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsTraceLoggingDisabled() || settings.IsLoggingDisabled() {
		traceWriter = io.MultiWriter(io.Discard)
	}

	warningWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "warning"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsWarningLoggingDisabled() || settings.IsLoggingDisabled() {
		warningWriter = io.MultiWriter(io.Discard)
	}

	eventWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "event"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsEventLoggingDisabled() || settings.IsLoggingDisabled() {
		eventWriter = io.MultiWriter(io.Discard)
	}

	securityWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "security"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsSecurityLoggingDisabled() || settings.IsLoggingDisabled() {
		securityWriter = io.MultiWriter(io.Discard)
	}

	databaseWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "database"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsDatabaseLoggingDisabled() || settings.IsLoggingDisabled() {
		databaseWriter = io.MultiWriter(io.Discard)
	}

	apiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "api"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsApiLoggingDisabled() || settings.IsLoggingDisabled() {
		apiWriter = io.MultiWriter(io.Discard)
	}

	importWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "import"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsImportLoggingDisabled() || settings.IsLoggingDisabled() {
		importWriter = io.MultiWriter(io.Discard)
	}

	exportWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "export"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsExportLoggingDisabled() || settings.IsLoggingDisabled() {
		exportWriter = io.MultiWriter(io.Discard)
	}

	lockWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "lock"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsLockLoggingDisabled() || settings.IsLoggingDisabled() {
		lockWriter = io.MultiWriter(io.Discard)
	}

	communicationsWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "communications"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsCommunicationsLoggingDisabled() || settings.IsLoggingDisabled() {
		communicationsWriter = io.MultiWriter(io.Discard)
	}

	cacheWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "cache"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsCacheLoggingDisabled() || settings.IsLoggingDisabled() {
		cacheWriter = io.MultiWriter(io.Discard)
	}

	webWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "web"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsWebLoggingDisabled() || settings.IsLoggingDisabled() {
		webWriter = io.MultiWriter(io.Discard)
	}

	msgStructure := log.Lmsgprefix | log.Ldate | log.LstdFlags | log.Lshortfile

	Info = log.New(generalWriter, formatNameWithColor(Green, "Info"), msgStructure)
	Warning = log.New(warningWriter, formatNameWithColor(Yellow, "Warn"), msgStructure)
	Error = log.New(errorWriter, formatNameWithColor(Red, "Errr"), msgStructure)
	Panic = log.New(panicWriter, formatNameWithColor(Red, "Pnic"), msgStructure)
	Timing = log.New(timingWriter, formatNameWithColor(Gray, "Time"), msgStructure)
	Event = log.New(eventWriter, formatNameWithColor(Magenta, "Evnt"), msgStructure)
	Service = log.New(serviceWriter, formatNameWithColor(Cyan, "Work"), msgStructure)
	Trace = log.New(traceWriter, formatNameWithColor(Gray, "Trac"), msgStructure)
	Audit = log.New(auditWriter, formatNameWithColor(White, "Audt"), msgStructure)
	Translation = log.New(translationWriter, formatNameWithColor(Cyan, "Trl8"), msgStructure)
	Security = log.New(securityWriter, formatNameWithColor(Magenta, "SECR"), msgStructure)
	Database = log.New(databaseWriter, formatNameWithColor(Blue, "DBas"), msgStructure)
	Api = log.New(apiWriter, formatNameWithColor(Cyan, "API_"), msgStructure)
	Import = log.New(importWriter, formatNameWithColor(Gray, "Impo"), msgStructure)
	Export = log.New(exportWriter, formatNameWithColor(White, "Expo"), msgStructure)
	Communications = log.New(communicationsWriter, formatNameWithColor(Cyan, "Comm"), msgStructure)
	Lock = log.New(lockWriter, formatNameWithColor(Blue, "Lock"), msgStructure)
	Unlock = log.New(lockWriter, formatNameWithColor(Blue, "Unlk"), msgStructure)
	SkipLock = log.New(lockWriter, formatNameWithColor(Blue, "Skip"), msgStructure)

	Cache = log.New(cacheWriter, formatNameWithColor(Cyan, "Cche"), msgStructure)
	Web = log.New(webWriter, formatNameWithColor(Cyan, "Web "), msgStructure)
}

func TestIt() {
	Info.Println("Starting the application...")
	Info.Println("Something noteworthy happened")
	Warning.Println("There is something you should know about")
	Panic.Println("Something went wrong")
	Error.Println("Something went wrong")
	Timing.Println("Timing")
	Event.Println("Events")
	Service.Println("Service")
	Trace.Println("Trace")
	Audit.Println("Audit")
	Translation.Println("Translation")
	Security.Println("Security")
	Database.Println("Database")
	Api.Println("API")
	Import.Println("Import")
	Export.Println("Export")
	Communications.Println("Communications")
	Lock.Println("Lock")
	Unlock.Println("Unlock")
	SkipLock.Println("Skip Lock")
	Cache.Println("Cache")
}

var hdr = "*------------------------------------------------------------------------*"

func banner(logCategory, logActivity, logMessage string, logger *log.Logger) {
	Info.Println(hdr)
	Info.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(logCategory), logActivity, logMessage)
	Info.Println(hdr)
}

func Banner(logCategory, logActivity, logMessage string) {
	banner(logCategory, logActivity, logMessage, Info)
}

// ServiceBanner - log a banner message to the service log
// Deprecated: No longer to be used
func ServiceBanner(logCategory, logActivity, logMessage string) {
	// banner(logCategory, logActivity, logMessage, ServiceLogger)
}

func Break() {
	Info.Println(formatNameWithColor(Cyan, hdr))
}
