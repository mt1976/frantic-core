# jobs

`jobs` defines a small background job model plus helpers for scheduling jobs via cron.

It wraps `robfig/cron` for scheduling and uses `timing` + `logHandler` for instrumentation.

## Job interface

A job must implement:

- `Run() error`
- `Service() func()`
- `Schedule() string` (cron expression)
- `Name() string`
- `Description() string`
- `AddDatabaseAccessFunctions(f func() ([]*database.DB, error))`

## Scheduler lifecycle

- `Initialise(cfg *commonConfig.Settings) error`
- `AddJobToScheduler(j Job)` / `AddJobsToScheduler(jobs []Job)`
- `StartScheduler()`

## Helpers

- `StartOfDay(time.Time) time.Time`
- `GetHumanReadableCronFreq(freq string) string`
- `PreRun(job Job)` / `PostRun(job Job)`

## Example

```go
import (
    "github.com/mt1976/frantic-core/commonConfig"
    "github.com/mt1976/frantic-core/jobs"
)

func setupScheduler() {
    cfg := commonConfig.Get()
    _ = jobs.Initialise(cfg)

    // jobs.AddJobToScheduler(myJob)
    jobs.StartScheduler()
}
```
