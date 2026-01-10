# go-interfaces

Interface wrappers for Go standard library packages to enable easy mocking in tests.

## Overview

`go-interfaces` provides mockable interface wrappers around Go's standard library packages. This allows you to write testable code without fighting Go's concrete types, making it easy to create mocks and stubs for unit testing.

Each package mirrors the structure of its corresponding standard library package while exposing mockable interfaces, following consistent architectural patterns throughout.

## Why Use This Library?

Go's standard library uses concrete types, making them difficult to mock in tests. This library solves that problem by:

- **Enabling Easy Mocking**: All functionality is exposed through interfaces that can be easily mocked
- **Maintaining Compatibility**: Interfaces mirror standard library APIs exactly
- **Zero Learning Curve**: If you know the standard library, you already know these interfaces
- **Clean Architecture**: Supports dependency injection and clean separation of concerns
- **Comprehensive Coverage**: Includes commonly-mocked packages like `io`, `net/http`, `os`, `sync`, and more

## Installation

```bash
go get github.com/pdutton/go-interfaces
```

Requires Go 1.24.0 or later.

## Quick Start

### Before (hard to test):

```go
func ReadConfig() ([]byte, error) {
    return os.ReadFile("config.json")
}
```

### After (easy to test):

```go
import "github.com/pdutton/go-interfaces/os"

type ConfigReader struct {
    os os.OS
}

func NewConfigReader(osInterface os.OS) *ConfigReader {
    return &ConfigReader{os: osInterface}
}

func (cr *ConfigReader) ReadConfig() ([]byte, error) {
    return cr.os.ReadFile("config.json")
}
```

Now in your tests, you can inject a mock `os.OS` implementation instead of hitting the real filesystem.

## Available Packages

- **encoding/json** - JSON encoding/decoding with `Encoder` and `Decoder` interfaces
- **io** - Core I/O primitives, reader/writer interfaces, and utilities
- **io/fs** - Filesystem interfaces (`FileInfo`, `DirEntry`, `FileMode`)
- **net** - Network dialing, listening, and connection interfaces
- **net/http/client** - HTTP client functionality
- **net/http/server** - HTTP server functionality
- **os** - File operations, process management, environment variables
- **os/exec** - Command execution with `Cmd` interface
- **os/signal** - Signal handling
- **path** - Path manipulation (slash-separated paths)
- **path/filepath** - Path manipulation (OS-specific paths)
- **sync** - Synchronization primitives (Mutex, WaitGroup, Once, Pool, Map, Cond, RWMutex)

## Usage Example

Here's a complete example showing dependency injection and testing:

```go
package myapp

import (
    "github.com/pdutton/go-interfaces/net/http/client"
    "github.com/pdutton/go-interfaces/io"
)

type APIClient struct {
    http http.Client
    io   io.IO
}

func NewAPIClient(httpClient http.Client, ioInterface io.IO) *APIClient {
    return &APIClient{
        http: httpClient,
        io:   ioInterface,
    }
}

func (a *APIClient) FetchData(url string) ([]byte, error) {
    resp, err := a.http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body().Close()

    return a.io.ReadAll(resp.Body())
}
```

In your tests, create mocks for `http.Client` and `io.IO` to test without real HTTP calls.

## License

MIT License - see [LICENSE](LICENSE) for details.

## TODO

- Updates and fixes based on usage feedback
- Additional standard library package coverage as needed
