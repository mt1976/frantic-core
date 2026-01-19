# frantic-core

frantic-core is a Go utility library and application toolkit.

It includes a large set of helper packages (banking, dates, logging, etc.) as well as a Storm-backed DAO layer (`dao/`) with caching and a code generator (`cmd/dao-gen`) for generating new DAO packages.

Docs:

- [CODE-GEN.md](CODE-GEN.md) (DAO generator)
- [logHandler/README.md](logHandler/README.md) (logging)
- [commonConfig/README.md](commonConfig/README.md) (configuration)
- [commonErrors/README.md](commonErrors/README.md) (shared errors)
- [contextHandler/README.md](contextHandler/README.md) (context/session)
- [idHelpers/README.md](idHelpers/README.md) (IDs)
- [importExportHelper/README.md](importExportHelper/README.md) (import/export)
- [jobs/README.md](jobs/README.md) (scheduler)
- [timing/README.md](timing/README.md) (timing)
- [chiMiddleware/README.md](chiMiddleware/README.md) (chi middleware)

## Features

- Banking utilities (IBAN, ISIN, LEI, UTI, etc.)
- Color manipulation
- Centralized configuration management
- Common error handling
- Context and session helpers
- Data import/export helpers
- Date and time utilities
- DAO / database layer (Storm-backed) with optional caching
- Typed (generic) helpers to reduce DAO boilerplate
- `go generate`-friendly DAO code generator (`cmd/dao-gen`)
- Docker helpers
- Email handling
- Financial calculations
- Logging and audit
- Math helpers
- Messaging and notification
- Path and file I/O helpers
- String manipulation
- TUI input helpers

## Directory Structure

```text
application/         # Application entry points
banking/             # Banking utilities (IBAN, ISIN, LEI, UTI, etc.)
cmd/                 # Tools (includes dao-gen)
colours/             # Color utilities
commonConfig/        # Configuration management
commonErrors/        # Error handling
contextHandler/      # Context and session helpers
dao/                 # Data access objects
dateHelpers/         # Date utilities
dockerHelpers/       # Docker support
emailHandler/        # Email utilities
financial/           # Financial calculations
frantic/             # Core frantic logic
htmlHelpers/         # HTML helpers
idHelpers/           # ID generation and validation
importExportHelper/  # Import/export helpers
ioHelpers/           # File and I/O helpers
jobs/                # Job scheduling and helpers
logHandler/          # Logging and audit
mathHelpers/         # Math utilities
messageHelpers/      # Messaging helpers
mockData/            # Mock/test data
netHandler/          # Networking helpers
notificationHandler/ # Notification utilities
paths/               # Path helpers
stringHelpers/       # String manipulation
timeHelpers/         # Time utilities
timing/              # Timing helpers
tuiInputHelper/      # TUI input helpers
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
    iban := "GB82WEST12345698765432"
    valid := banking.ValidateIBAN(iban)
    fmt.Printf("IBAN %s valid: %v\n", iban, valid)
}
```

### Example: Logging

```go
import "github.com/mt1976/frantic-core/logHandler"

func main() {
    logHandler.Info("Application started")
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
    fmt.Println("Today's date is:", today)
}
```

## DAO layer & code generation

If you’re using the DAO stack in `dao/`, these docs are the best entry points:

- DAO generator (`dao-gen`): see [CODE-GEN.md](CODE-GEN.md)
- Typed DB helpers (generics): see [dao/database/README.md](dao/database/README.md)
- Example generated/templated DAO: see [dao/test/templateStoreV2/README.md](dao/test/templateStoreV2/README.md)

Quick example (from this repo root):

```bash
mkdir -p dao/test/fred
go run ./cmd/dao-gen -out dao/test/fred -pkg fred -type Fred -table Fred -namespace main -force
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
