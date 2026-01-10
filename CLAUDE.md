# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This repository provides interface wrappers for Go standard library packages to enable easy mocking in tests. Each package mirrors the structure of its corresponding standard library package while exposing mockable interfaces.

## Common Commands

### Testing
- Run all tests: `make test` or `go test './...'`
- Run tests for a specific package: `go test ./path/to/package`
- Run a single test: `go test -run TestName ./path/to/package`

### Formatting
- Format all code: `make fmt` or `go fmt ./...`

## Architecture Patterns

### Facade Pattern

Every package follows a consistent facade pattern with these components:

1. **Interface Definition**: A primary interface (e.g., `IO`, `JSON`, `Sync`) that exposes all package functions and constructors
2. **Facade Implementation**: A concrete type (e.g., `ioFacade`, `jsonFacade`, `syncFacade`) that implements the interface by delegating to the real standard library
3. **Constructor**: A `New*` function that returns the facade implementation
4. **Type Aliases**: Standard library interfaces are re-exported as type aliases (e.g., `type Reader = io.Reader`)

Example structure from `io` package:
```go
type IO interface {
    Copy(Writer, Reader) (int64, error)
    // ... other functions
}

type ioFacade struct {}

func NewIO() ioFacade {
    return ioFacade{}
}

func (_ ioFacade) Copy(w Writer, r Reader) (int64, error) {
    return io.Copy(w, r)
}
```

### Wrapping Pattern

For types that wrap standard library structs, the pattern includes:

1. **Interface**: Defines the methods available on the type
2. **Facade Type**: A struct containing a pointer to the real standard library type (named `nub`, `realClient`, `realServer`, etc.)
3. **Constructor with Options**: A `New*` function that accepts functional options
4. **Wrap Function**: A `Wrap*` function to wrap existing standard library instances
5. **Delegation Methods**: Each method delegates to the underlying standard library type

Example from `sync/mutex.go`:
```go
type Mutex interface {
    Lock()
    Unlock()
    TryLock() bool
}

type mutexFacade struct {
    realMutex *sync.Mutex
}

type MutexOption func(mut *sync.Mutex)

func (_ syncFacade) NewMutex(options ...MutexOption) mutexFacade {
    var mut sync.Mutex
    for _, opt := range options {
        opt(&mut)
    }
    return mutexFacade{realMutex: &mut}
}

func (m mutexFacade) Lock() {
    m.realMutex.Lock()
}
```

### Functional Options Pattern

Types with configuration use functional options:
- Options are defined as functions that modify the underlying standard library type
- Option functions are named `With*` (e.g., `WithTimeout`, `WithTransport`)
- Options are applied during construction before wrapping

Example from `net/http/client/client.go`:
```go
type ClientOption func(cl *http.Client)

func WithTimeout(t time.Duration) ClientOption {
    return func(cl *http.Client) {
        cl.Timeout = t
    }
}
```

### Package-Level Functions

Some packages also expose direct package-level functions that bypass the facade for convenience (e.g., `encoding/json` has `Marshal`, `Unmarshal` functions that directly call `encoding/json.Marshal`, etc.).

## Package Structure

### Key Packages

- **encoding/json**: JSON encoding/decoding with `Encoder` and `Decoder` interfaces
- **io**: Core I/O primitives, reader/writer interfaces, and utilities
- **io/fs**: Filesystem interfaces (`FileInfo`, `DirEntry`, `FileMode`)
- **net**: Network dialing, listening, and connection interfaces
- **net/http/client**: HTTP client split into separate package from server
- **net/http/server**: HTTP server functionality
- **os**: File operations, process management
- **os/exec**: Command execution with `Cmd` interface
- **os/signal**: Signal handling
- **path** and **path/filepath**: Path manipulation
- **sync**: Synchronization primitives (Mutex, WaitGroup, Once, Pool, Map, Cond, RWMutex)

### http Package Split

The `net/http` package is uniquely split into `client` and `server` subdirectories to separate client-side and server-side HTTP concerns. See `net/http/README.md`.

## Testing Conventions

Tests follow standard Go testing practices:
- Test files are named `*_test.go`
- Each package has comprehensive unit tests
- Tests verify both interface compliance and actual functionality
- Concurrent behavior is tested where relevant (e.g., `sync` package)

## Important Implementation Details

### Interface Returns

When facade methods return types from the standard library that also have interface wrappers in this project, they must wrap those return values. For example:
- `os.File` methods return wrapped `File` interfaces, not `*os.File`
- HTTP `Client.Do()` returns wrapped `Response` interface
- Functions that return multiple standard library types wrap each appropriately

### Accessing Underlying Types

Facade types provide access to the underlying standard library type when needed:
- Method names vary: `Nub()`, `RealRequest()`, `GetUnderlyingResolver()`, etc.
- Used when passing to code that requires the actual standard library type

### Receiver Conventions

- Facade methods use blank identifier receivers `(_ facadeType)` when they don't need struct state
- Wrapper methods use named receivers when accessing the underlying type
