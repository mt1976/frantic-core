# timing

`timing` provides simple timing utilities and instrumentation helpers.

It’s used across the codebase to log durations (typically via `logHandler.TimingLogger`).

## Stopwatch

- `Start(domain, activity, notes string) Stopwatch`
- `(*Stopwatch).Stop(count int)`

Typical usage:

```go
import "github.com/mt1976/frantic-core/timing"

func doWork() {
    sw := timing.Start("users", "load", "")
    // ... work ...
    sw.Stop(1)
}
```

## Snooze helpers

- `SnoozeFor(noSeconds int)`
- `Snooze()` (random 0..10 seconds)
