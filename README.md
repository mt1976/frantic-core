# frantic-core

frantic-core is a comprehensive Go utility library providing a wide range of helpers and utilities for application development. It includes modules for banking, colors, configuration, error handling, date and time helpers, logging, math, messaging, networking, notifications, string manipulation, and more.

## Features

- Banking utilities (IBAN, ISIN, LEI, UTI, etc.)
- Color manipulation
- Centralized configuration management
- Common error handling
- Context and session helpers
- Data import/export helpers
- Date and time utilities
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

```
application/         # Application entry points
banking/             # Banking utilities (IBAN, ISIN, LEI, UTI, etc.)
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

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
