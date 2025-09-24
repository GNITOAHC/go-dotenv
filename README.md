# go-dotenv

A lightweight Golang library for loading environment variables from `.env` files, with support for variable expansion and required variable checks.

## Features

- Load environment variables from one or more `.env` files.
- Supports variable expansion (e.g., `FOO=${BAR}`).
- Check for required environment variables and panic if missing.
- Simple API.

## Installation

```sh
go get github.com/gnitoahc/go-dotenv
```

## Usage

### Basic Usage

Create a `.env` file in your project root:

```
DB_HOST=localhost
DB_USER=root
DB_PASS=secret
```

Load the variables in your Go code:

```go
import "github.com/gnitoahc/go-dotenv"

func main() {
    // Load .env file (default)
    err := dotenv.Load()
    if err != nil {
        panic(err)
    }
}
```

### Load Multiple Files

```go
err := dotenv.Load(".env", ".env.local")
```

### Require Variables

Ensure certain variables are set (panics if missing):

```go
err := dotenv.Must([]string{"DB_HOST", "DB_USER"})
```

### Example

```go
package main

import "github.com/gnitoahc/go-dotenv"

func main() {
    err := dotenv.Must([]string{"DB_HOST"})
    if err != nil {
        print(err.Error())
    }
}
```

## API

### [`dotenv.Load`](dotenv.go)

```go
func Load(paths ...string) error
```

Loads environment variables from the specified files. Defaults to `.env` if no path is given.

### [`dotenv.Must`](must.go)

```go
func Must(vars []string, paths ...string) error
```

Loads environment variables and panics if any of the specified variables are not set.

## License

[MIT](LICENSE)
