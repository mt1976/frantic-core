# logHandler

`logHandler` centralises application logging and log file rotation.

It exposes a set of preconfigured `*log.Logger` instances (Info, Warning, Error, Timing, Audit, etc.) which write to stdout/stderr and to rotating log files (via `lumberjack`).

## How it works

- Initialises automatically via `init()` when the package is imported.
- Log file location is derived from the application paths and the configured application name.
- Most log categories can be enabled/disabled via `commonConfig`.
- Uses ANSI colour prefixes on non-Windows platforms.

## Common loggers

A few of the exported loggers:

- `logHandler.InfoLogger`
- `logHandler.WarningLogger`
- `logHandler.ErrorLogger`
- `logHandler.PanicLogger`
- `logHandler.TimingLogger`
- `logHandler.AuditLogger`
- `logHandler.DatabaseLogger`
- `logHandler.CacheLogger`

(There are more categories in `logHandler.go`.)

## Usage

```go
package main

import (
    "github.com/mt1976/frantic-core/logHandler"
)

func main() {
    logHandler.Info.Println("hello")
    logHandler.Warning.Println("something to look at")
    logHandler.Error.Println("something failed")

    logHandler.InfoBanner("startup", "init", "application starting")
    logHandler.Break()
}
```

## Notes

- Because initialisation happens in `init()`, importing this package will create/open the log writers immediately.
- Rotation settings (max size, backups, age, compression) are taken from `commonConfig`.
