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
- `Settings` includes sections for Application, Database, Server, Logging, Security, Translation, etc.

## Debugging

- `(*Settings).Spew()` prints a summary of the loaded config.
