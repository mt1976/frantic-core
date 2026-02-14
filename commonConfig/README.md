# commonConfig

`commonConfig` defines the application configuration model and accessors.

Configuration is loaded from a TOML file (`common.toml`) located under the application’s config directory (derived from `paths`).

## Quick start

```go
import "github.com/mt1976/frantic-core/commonConfig"

func loadConfig() *commonConfig.Settings {
    return commonConfig.Get()
}
```

## Key points

- `Get()` loads and unmarshals TOML into a `Settings` struct.
- If the config file cannot be read or parsed, `Get()` will panic.
- `Settings` includes sections for Application, Database, Server, Logging, Security, Translation, and more.

## Configuration Sections

| Section | Key Accessors |
|---|---|
| **Application** | `GetApplication_Name()`, `GetApplication_Prefix()`, `GetApplication_HomePath()`, `GetApplication_Description()`, `GetApplication_Version()`, `GetApplication_Environment()`, `GetApplication_ReleaseDate()`, `GetApplication_Copyright()`, `GetApplication_Author()`, `GetApplication_License()`, `GetApplication_Locale()`, `GetApplication_Theme()`, `GetApplication_Timezone()`, `IsApplicationMode(MODE) bool` |
| **Database** | `GetDatabase_Version()`, `GetDatabase_Type()`, `GetDatabase_Name()`, `GetDatabase_Path()`, `GetDatabase_Host()`, `GetDatabase_Port()`, `GetDatabase_User()`, `GetDatabase_Password()`, `GetDatabase_PoolSize()`, `GetDatabase_Timeout()` |
| **Server** | `GetServer_Port()`, `GetServer_Protocol()`, `GetServer_Host()`, `GetServer_Environment()`, `GetServer_Compression()` |
| **Dates** | `GetDateFormat_DateTime()`, `GetDateFormat_Date()`, `GetDateFormat_Time()`, `GetDateFormat_Human()`, `GetDateFormat_Internal()`, `GetDateFormat_DMY2()`, `GetDateFormat_YMD()`, `GetDateFormat_Backup()`, `GetDateFormat_BackupDirectory()` |
| **Logging** | `Is{General,Timing,Service,Audit,Translation,Trace,Warning,Event,Security,Database,Api,Import,Export,Communications,Cache,Lock,Logging}Disabled()`, `GetLogging_MaxSize()`, `GetLogging_MaxBackups()`, `GetLogging_MaxAge()`, `IsLogCompressionEnabled()` |
| **Security** | `GetSecuritySession_ExpiryPeriod()`, `GetSecuritySessionKey_Session()`, `GetServiceUser_Name()`, `GetServiceUser_UID()`, `GetServiceUser_UserCode()` |
| **Communications** | `GetCommunicationsPushover_UserKey()`, `GetCommunicationsPushover_APIToken()`, `GetCommunicationsEmail_Host()`, `GetCommunicationsEmail_Port()`, `GetCommunicationsEmail_Sender()`, `GetCommunicationsEmail_AdminEmail()` |
| **Translation** | `GetTranslationServer_Host()`, `GetTranslationServer_Port()`, `GetTranslation_Locale()`, `GetTranslation_PermittedOrigins()`, `GetTranslation_PermittedLocales()`, `IsPermittedTranslationLocale(string) bool`, `IsPermittedTranslationOrigin(string) bool` |
| **Assets** | `GetAssets_LogoPath()`, `GetAssets_FaviconPath()` |
| **Backups** | `GetBackup_RetainForDays()` |
| **Status** | `GetStatusList()`, `GetStatus_Unknown()`, `GetStatus_Online()`, `GetStatus_Offline()`, `GetStatus_Error()`, `GetStatus_Warning()` |
| **Display** | `GetDefault_Delimiter()`, `Delimiter()` |
| **General** | `GetHistory_MaxHistoryEntries()`, `GetWorkerPoolSize()` |
| **Hosts** | `GetValidHosts()` |

## Application Modes

- `MODE_DEVELOPMENT`
- `MODE_PRODUCTION`
- `MODE_TEST`

## Debugging

- `(*Settings).Spew()` prints a summary of the loaded config.
