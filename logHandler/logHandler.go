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
	WarningLogger        *log.Logger
	InfoLogger           *log.Logger
	ErrorLogger          *log.Logger
	PanicLogger          *log.Logger
	TimingLogger         *log.Logger
	EventLogger          *log.Logger
	ServiceLogger        *log.Logger
	TraceLogger          *log.Logger
	AuditLogger          *log.Logger
	TranslationLogger    *log.Logger
	SecurityLogger       *log.Logger
	DatabaseLogger       *log.Logger
	ApiLogger            *log.Logger
	ImportLogger         *log.Logger
	ExportLogger         *log.Logger
	CommunicationsLogger *log.Logger
	LockLogger           *log.Logger
)

var Reset string
var Red string
var Green string
var Yellow string
var Blue string
var Magenta string
var Cyan string
var Gray string
var White string

func init() {
	settings := commonConfig.Get()
	//applicationPath := "data/logs/"
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

	serviceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "service"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsServiceLoggingDisabled() || settings.IsLoggingDisabled() {
		serviceWriter = io.MultiWriter(io.Discard)
	}

	auditWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "audit"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsAuditLoggingDisabled() || settings.IsLoggingDisabled() {
		auditWriter = io.MultiWriter(io.Discard)
	}

	errorWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "error"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsLoggingDisabled() {
		errorWriter = io.MultiWriter(io.Discard)
	}

	panicWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "panic"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
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

	lockWriter := io.MultiWriter(&lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "lock"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsLockLoggingDisabled() || settings.IsLoggingDisabled() {
		lockWriter = io.MultiWriter(io.Discard)
	}

	communicationsWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: assembleLogFileName(applicationPath, "communications"), MaxSize: maxSize, MaxBackups: maxBackups, MaxAge: maxAge, Compress: compress})
	if settings.IsCommunicationsLoggingDisabled() || settings.IsLoggingDisabled() {
		communicationsWriter = io.MultiWriter(io.Discard)
	}

	msgStructure := log.Lmsgprefix | log.Ldate | log.Lmicroseconds | log.Lshortfile

	InfoLogger = log.New(generalWriter, formatNameWithColor(White, "Info"), msgStructure)
	WarningLogger = log.New(warningWriter, formatNameWithColor(Yellow, "Warning"), msgStructure)
	ErrorLogger = log.New(errorWriter, formatNameWithColor(Red, "Error"), msgStructure)
	PanicLogger = log.New(panicWriter, formatNameWithColor(Red, "Panic"), msgStructure)
	TimingLogger = log.New(timingWriter, formatNameWithColor(Gray, "Timing"), msgStructure)
	EventLogger = log.New(eventWriter, formatNameWithColor(Magenta, "Event"), msgStructure)
	ServiceLogger = log.New(serviceWriter, formatNameWithColor(Green, "Service"), msgStructure)
	TraceLogger = log.New(traceWriter, formatNameWithColor(White, "Trace"), msgStructure)
	AuditLogger = log.New(auditWriter, formatNameWithColor(Yellow, "Audit"), msgStructure)
	TranslationLogger = log.New(translationWriter, formatNameWithColor(Cyan, "Translation"), msgStructure)
	SecurityLogger = log.New(securityWriter, formatNameWithColor(Magenta, "Security"), msgStructure)
	DatabaseLogger = log.New(databaseWriter, formatNameWithColor(Gray, "Database"), msgStructure)
	ApiLogger = log.New(apiWriter, formatNameWithColor(Green, "API"), msgStructure)
	ImportLogger = log.New(importWriter, formatNameWithColor(Blue, "Import"), msgStructure)
	ExportLogger = log.New(exportWriter, formatNameWithColor(Cyan, "Export"), msgStructure)
	CommunicationsLogger = log.New(communicationsWriter, formatNameWithColor(White, "Communications"), msgStructure)
	LockLogger = log.New(lockWriter, formatNameWithColor(Blue, "Lock"), msgStructure)
}

func TestIt() {
	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	PanicLogger.Println("Something went wrong")
	ErrorLogger.Println("Something went wrong")
	TimingLogger.Println("Timing")
	EventLogger.Println("Events")
	ServiceLogger.Println("Service")
	TraceLogger.Println("Trace")
	AuditLogger.Println("Audit")
	TranslationLogger.Println("Translation")
	SecurityLogger.Println("Security")
	DatabaseLogger.Println("Database")
	ApiLogger.Println("API")
	ImportLogger.Println("Import")
	ExportLogger.Println("Export")
	CommunicationsLogger.Println("Communications")
	LockLogger.Println("Lock")
}

var hdr = "*------------------------------------------------------------------------*"

func banner(logCategory, logActivity, logMessage string, logger *log.Logger) {
	logger.Println(hdr)
	logger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(logCategory), logActivity, logMessage)
	logger.Println(hdr)
}

func InfoBanner(logCategory, logActivity, logMessage string) {
	banner(logCategory, logActivity, logMessage, InfoLogger)
}

// ServiceBanner - log a banner message to the service log
// Deprecated: No longer to be used
func ServiceBanner(logCategory, logActivity, logMessage string) {
	// banner(logCategory, logActivity, logMessage, ServiceLogger)
}

func Break() {
	InfoLogger.Println(formatNameWithColor(Cyan, hdr))
}
