# Running Code-Gen (`dao-gen`)

This repo includes a small code generator at `cmd/dao-gen` that produces a new DAO package using the TemplateStoreV2 template set.

See also: [README.md](README.md)

## Quick start

From the repo root:

```bash
mkdir -p dao/test/fred
go run ./cmd/dao-gen -out dao/test/fred -pkg fred -type Fred -table Fred -namespace main -force
```

What those flags mean (high level):

- `-out`: where the generated package files are written
- `-pkg`: Go package name (usually lowercase)
- `-type`: exported Go struct/type name (usually PascalCase)
- `-table`: table name used by the DAO (defaults to `-type` if omitted)
- `-namespace`: value passed into `database.WithNameSpace(...)`
- `-force`: allow overwriting existing generated files (with some safety rules below)

## Usage

Run:

```bash
go run ./cmd/dao-gen [flags]
```

You can also see built-in help:

```bash
go run ./cmd/dao-gen -h
```

## Options

All flags supported by `dao-gen`:

- `-out` (string, default `.`)
  - Output directory for generated files.
- `-pkg` (string, required)
  - Go package name to generate (example: `fred`).
- `-type` (string, required)
  - Exported type/struct name (example: `Fred`).
- `-table` (string, optional)
  - Logical table name used to create `TableName`.
  - Defaults to `-type`.
- `-namespace` (string, optional)
  - Value passed into `database.WithNameSpace(...)` when connecting.
  - Defaults to `cheeseOnToast` if omitted.
- `-force` (bool, default `false`)
  - Allows overwriting existing generated files.
  - Safety rule: for `*Model.go`, `*Helpers.go`, and `*New.go`, the generator will **rename the existing file** to `.OLD` first (and warn), rather than overwriting in place.
- `-with-worker` (bool, default `true`)
  - Generate the worker file (`*Worker.go`).
- `-with-impex` (bool, default `true`)
  - Generate the import/export file (`*Impex.go`).
- `-with-debug` (bool, default `true`)
  - Generate the debug file (`*Debug.go`).

## Safety behaviour (important)

When regenerating into an existing package:

- If the destination `*Model.go`, `*Helpers.go`, or `*New.go` already exists:

  - The generator renames it from `.go` to `.OLD` (or `.OLD1`, `.OLD2`, etc.).
  - It prints a warning telling you to manually migrate any custom code.
- For all other generated files:
  - Without `-force`, generation will refuse to overwrite existing files.
  - With `-force`, it will overwrite those files.

This is intended to protect the files where people commonly add custom domain logic.

## Using with `go generate`

The recommended workflow is to commit a small `generate.go` file inside your target package:

```go
package fred

//go:generate go run ../../../cmd/dao-gen -out . -pkg fred -type Fred -table Fred -namespace main -force
```

Then you can regenerate with:

```bash
go generate ./...
```

## Using `dao-gen` from another project (Option 2)

If you want to use the generator from a different repo, you can run it directly from this module without copying any files.

Important: the generated code imports `database` and `cache` from frantic-core:

- `github.com/mt1976/frantic-core/dao/database`
- `github.com/mt1976/frantic-core/dao/cache`

So the project you generate into must also depend on frantic-core (it will compile against those packages).

### One-off usage

From your other project:

```bash
go get github.com/mt1976/frantic-core@latest
go run github.com/mt1976/frantic-core/cmd/dao-gen@latest -out .app/dao/fred -pkg fred -type Fred -table Fred -namespace main -force
```

In the example above replace fred/Fred with the name of the DAO, for example userStore/UserStore and the namespace used by the db instance, normally this will be "main"

For repeatable builds, pin a specific version instead of `@latest`.

### With `go generate`

In your other project’s package, you can use a directive like:

```go
package fred

//go:generate go run github.com/mt1976/frantic-core/cmd/dao-gen@latest -out . -pkg fred -type Fred -table Fred -namespace main -force
```

If you prefer to pin the generator version, replace `@latest` with a tag/commit (and keep your project’s `go.mod` aligned).

## Output files

`dao-gen` writes a set of `.go` files similar to TemplateStoreV2, including:

- `*Model.go` (type + Fields + TableName)
- `*DB.go` (initialise/close)
- `*Cache.go` (cache hooks)
- `*.go` (DAO CRUD/query functions)
- `*New.go` (constructor/create)
- `*Internals.go`, `*Helpers.go`
- Optional: `*Worker.go`, `*Impex.go`, `*Debug.go`
- `README.md`
