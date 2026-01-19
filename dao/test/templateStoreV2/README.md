# TemplateStoreV2

`templateStoreV2` is an example DAO package intended as a modernised template (typed DB reads, cache hooks, worker, import/export, etc.).

## Public API

### Exported types/vars

- `type TemplateStore struct { ... }`
- `var TableName entities.Table`
- `var Fields fieldNames`

### Database lifecycle

- `func Initialise(ctx context.Context, cached bool)`
- `func IsInitialised() bool`
- `func Close()`
- `func GetDatabaseConnections() func() ([]*database.DB, error)`

### Queries

- `func Count() (int, error)`
- `func CountWhere(field entities.Field, value any) (int, error)`
- `func GetBy(field entities.Field, value any) (TemplateStore, error)`
- `func GetAll() ([]TemplateStore, error)`
- `func GetAllWhere(field entities.Field, value any) ([]TemplateStore, error)`

### Mutations

- `func Delete(ctx context.Context, id int, note string) error`
- `func DeleteBy(ctx context.Context, field entities.Field, value any, note string) error`
- `func Drop() error`
- `func ClearDown(ctx context.Context) error`

### Record methods

- `func (record *TemplateStore) Validate() error`
- `func (record *TemplateStore) Update(ctx context.Context, note string) error`
- `func (record *TemplateStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) error`
- `func (record *TemplateStore) Create(ctx context.Context, note string) error`
- `func (record *TemplateStore) Clone(ctx context.Context) (TemplateStore, error)`

### Lookups

- `func GetDefaultLookup() (lookup.Lookup, error)`
- `func GetLookup(field, value entities.Field) (lookup.Lookup, error)`

### Cache integration

- `func PreLoad(ctx context.Context) error`
- `func CacheSpew()`
- `func FlushCache() error`
- `func HydrateCache() error`
- `func CacheHydrator(ctx context.Context) func() ([]any, error)`
- `func CacheSynchroniser(ctx context.Context) func(any) error`

### Construction

- `func New() TemplateStore`
- `func Create(ctx context.Context, userName, uid, realName, email, gid string) (TemplateStore, error)`

### Import / Export

- `func (record *TemplateStore) ExportRecordAsJSON(name string)`
- `func ExportAllAsJSON(message string)`
- `func (record *TemplateStore) ExportRecordAsCSV(name string) error`
- `func ExportAllAsCSV(msg string) error`
- `func ImportAllFromCSV() error`

### Worker

- `func Worker(j jobs.Job, db *database.DB)`

### Example business logic

- `func Login(ctx context.Context, sq string) (TemplateStore, error)`
- `func Add(ctx context.Context, sq string) (TemplateStore, error)`
- `func (u *TemplateStore) SetName(name string) error`

### Debug

- `func (record *TemplateStore) Spew()`

### Deprecated wrappers

These are kept for compatibility and live in `templateStoreV2Deprecated.go`.

- `func GetById(id int) (TemplateStore, error)`
- `func GetByKey(key string) (TemplateStore, error)`
- `func DeleteByKey(key any) error`
