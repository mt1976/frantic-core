# frantic-core

frantic-core is a Go utility library and application toolkit.

It includes a large set of helper packages for banking, dates, logging, configuration, financial calculations, and more.

## Package Documentation

| Package | Description |
|---|---|
| [application](application/README.md) | OS/platform detection and system identity |
| [banking](banking/README.md) | IBAN, ISIN, LEI, UTI validation and GLEIF API lookup |
| [chiMiddleware](chiMiddleware/README.md) | Chi router middleware (Brotli, minify, method conversion, context injection) |
| [colours](colours/README.md) | ANSI colour escape sequences for terminal styling |
| [commonConfig](commonConfig/README.md) | TOML-based application configuration |
| [commonErrors](commonErrors/README.md) | Shared sentinel errors and wrapping helpers |
| [contextHandler](contextHandler/README.md) | Request-scoped context and session management |
| [dateHelpers](dateHelpers/README.md) | Date formatting, parsing, arithmetic, and business-day logic |
| [dockerHelpers](dockerHelpers/README.md) | Docker environment detection and payload deployment |
| [emailHandler](emailHandler/README.md) | Email composition and SMTP sending |
| [financial](financial/README.md) | Tenor calculations, settlement dates, and amount formatting |
| [frantic](frantic/README.md) | Project identity utilities and validation |
| [htmlHelpers](htmlHelpers/README.md) | HTML value conversion and URL-safe encoding |
| [idHelpers](idHelpers/README.md) | ID generation (KSUID), hashing, and composite IDs |
| [ioHelpers](ioHelpers/README.md) | File I/O, copy, backup, and directory management |
| [logHandler](logHandler/README.md) | Multi-channel logging with file rotation |
| [mathHelpers](mathHelpers/README.md) | Random numbers, min/max, coin toss |
| [messageHelpers](messageHelpers/README.md) | Inter-service message types (user, session, authority, translation) |
| [mockData](mockData/README.md) | Mock reference datasets for testing (countries, currencies, etc.) |
| [netHandler](netHandler/README.md) | Network host reachability checking |
| [notificationHandler](notificationHandler/README.md) | Push notifications via Pushover |
| [paths](paths/README.md) | Filesystem path helpers and directory structure |
| [stringHelpers](stringHelpers/README.md) | String formatting, quoting, encoding, and manipulation |
| [timeHelpers](timeHelpers/README.md) | Timezone inference from locale |
| [timing](timing/README.md) | Stopwatch timing and snooze utilities |
| [tuiInputHelper](tuiInputHelper/README.md) | Terminal UI input prompts |

## Features

- Banking utilities (IBAN, ISIN, LEI, UTI validation and formatting)
- GLEIF API integration for LEI lookups
- Color terminal output
- Centralized TOML-based configuration management
- Common error handling with sentinel errors and wrappers
- Context and session helpers with worker pool support
- Date and time utilities with business-day awareness
- Docker environment detection
- Email sending via SMTP
- Financial calculations (tenors, settlement dates, rate ladders, amount formatting)
- Project identity validation
- HTML value helpers and URL-safe encoding
- ID generation (KSUID), hashing (SHA3-256), and composite IDs
- File I/O, backup, and directory management
- Multi-channel logging with file rotation (lumberjack)
- Math helpers (random, min/max)
- Inter-service messaging types
- Mock reference data (countries, currencies, genders, titles, rate ladders)
- Network host reachability checks
- Push notifications (Pushover)
- Filesystem path management
- String manipulation and formatting
- Timezone inference
- Stopwatch timing and sleep utilities
- Terminal UI input helpers
- Chi router middleware (Brotli, HTML minification, HTTP method conversion)

## Directory Structure

```text
application/         # OS/platform detection and system identity
banking/             # Financial identifier validation (IBAN, ISIN, LEI, UTI)
chiMiddleware/       # Chi router middleware
colours/             # ANSI colour escape sequences
commonConfig/        # TOML-based configuration management
commonErrors/        # Shared sentinel errors and wrappers
contextHandler/      # Context and session management
dateHelpers/         # Date utilities and business-day logic
dockerHelpers/       # Docker environment support
emailHandler/        # Email sending via SMTP
financial/           # Financial calculations and tenor logic
frantic/             # Project identity utilities
htmlHelpers/         # HTML value helpers and URL encoding
idHelpers/           # ID generation and hashing
ioHelpers/           # File I/O and directory management
logHandler/          # Multi-channel logging with rotation
mathHelpers/         # Numeric helpers
messageHelpers/      # Inter-service message types
mockData/            # Mock reference datasets
netHandler/          # Network reachability checking
notificationHandler/ # Push notifications (Pushover)
paths/               # Filesystem path helpers
stringHelpers/       # String manipulation
timeHelpers/         # Timezone inference
timing/              # Stopwatch and snooze utilities
tuiInputHelper/      # Terminal UI input helpers
```

## Installation

To use frantic-core in your Go project:

```bash
go get github.com/mt1976/frantic-core
```

## Usage

Import the package or specific modules as needed:

```go
import (
    "github.com/mt1976/frantic-core/banking"
    "github.com/mt1976/frantic-core/dateHelpers"
    "github.com/mt1976/frantic-core/logHandler"
    // ...other imports as needed
)
```

### Example: Validate an IBAN

```go
package main

import (
    "fmt"
    "github.com/mt1976/frantic-core/banking"
)

func main() {
    iban, err := banking.NewIBAN("GB82WEST12345698765432")
    if err != nil {
        fmt.Println("Invalid IBAN:", err)
        return
    }
    fmt.Println("IBAN:", iban.String())
}
```

### Example: Logging

```go
import "github.com/mt1976/frantic-core/logHandler"

func main() {
    logHandler.Info.Println("Application started")
}
```

### Example: Date Helper

```go
import (
    "fmt"
    "github.com/mt1976/frantic-core/dateHelpers"
)

func main() {
    today := dateHelpers.Today()
    fmt.Println("Today:", dateHelpers.FormatHuman(today))
    fmt.Println("Is working day:", dateHelpers.IsWorkingDay(today))
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
