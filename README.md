# Logging Generator

Generates logging middleware for golang interface

This golang generator can be used to generate a logging middleware with the zerolog logging library for an provided interface.

> For detected bugs please contact: marco-engstler@gmx.de

- [Logging Generator](#logging-generator)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Flags](#flags)
  - [Examples](#examples)
    - [Generate manually](#generate-manually)
    - [Generate by go generate](#generate-by-go-generate)
    - [Example output for *CompositeParamsInterface* in file interfaces/interfaces_test.go](#example-output-for-compositeparamsinterface-in-file-interfacesinterfaces_testgo)

## Installation

```bash
go get github.com/hanofzelbri/middleware-generator
```

## Usage

```bash
middleware-generator [flags]
```

### Flags

```bash
  -p, --emptyFunctionParamNamePrefix string         If there is no function parameter name provided this prefix will be used (default "param")
  -r, --emptyFunctionReturnParamNamePrefix string   If there is no function parameter return name provided this prefix will be used (default "ret")
  -h, --help                                        help for middleware-generator
  -i, --interface string                            Interface definition to generate logging middleware for.
  -f, --middlewareFunctionName string               Function name for middleware (default "WithMiddleware")
  -o, --output string                               Output file. If empty StdOut is used
  -w, --wrapper string                              Wrapper definition for implementation of middleware interface.
```

## Examples

### Generate manually

```bash
middleware-generator -i "io.Reader" -w "pkg.structname"
middleware-generator -i "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface" -o "logging-middleware.go"
```

### Generate by go generate

```go
//go:generate middleware-generator -i "github.com/hanofzelbri/middleware-generator/interfaces.CompositeParamsInterface" -o "logging-middleware.go"
```

### Example output for *CompositeParamsInterface* in file [interfaces/interfaces_test.go](interfaces/interfaces_test.go)

```go
// Code generated by middleware-generator; DO NOT EDIT

package interfaces

import (
    "time"

    "github.com/google/uuid"
    "github.com/rs/zerolog/log"
    "go/ast"
)

// CompositeParamsInterface is a dummy interface to test program
type compositeParamsInterface struct {
    wrapper CompositeParamsInterface
}

// WithMiddleware adds logging for interface CompositeParamsInterface
func WithMiddleware(wrapper CompositeParamsInterface) CompositeParamsInterface {
    return &compositeParamsInterface{
        wrapper: wrapper,
    }
}

// Array param types
func (l *compositeParamsInterface) Array(a [3]uuid.UUID) (r [10]bool) {
    defer func(begin time.Time) {
        log.Info().
            Interface("a", a).
            Dur("took", time.Since(begin)).
            Interface("r", r).
            Msg("Method Array called")
    }(time.Now())

    return l.wrapper.Array(a)
}

// Channel param types
func (l *compositeParamsInterface) Channel(param1 chan string) (ret1 chan int) {
    defer func(begin time.Time) {
        log.Info().
            Interface("param1", param1).
            Dur("took", time.Since(begin)).
            Interface("ret1", ret1).
            Msg("Method Channel called")
    }(time.Now())

    return l.wrapper.Channel(param1)
}

// Composite param types
func (l *compositeParamsInterface) Composite(m map[string]chan int, d [2]chan func(string) map[bool]*ast.MapType) (ret1 []chan func(string) error) {
    defer func(begin time.Time) {
        log.Info().
            Interface("m", m).
            Interface("d", d).
            Dur("took", time.Since(begin)).
            Interface("ret1", ret1).
            Msg("Method Composite called")
    }(time.Now())

    return l.wrapper.Composite(m, d)
}

// Map param types
func (l *compositeParamsInterface) Map(param1 map[string]uuid.UUID) (ret1 map[bool]int) {
    defer func(begin time.Time) {
        log.Info().
            Interface("param1", param1).
            Dur("took", time.Since(begin)).
            Interface("ret1", ret1).
            Msg("Method Map called")
    }(time.Now())

    return l.wrapper.Map(param1)
}

// Slice param types
func (l *compositeParamsInterface) Slice(param1 []uuid.UUID, param2 []int) (ret1 []bool) {
    defer func(begin time.Time) {
        log.Info().
            Interface("param1", param1).
            Interface("param2", param2).
            Dur("took", time.Since(begin)).
            Interface("ret1", ret1).
            Msg("Method Slice called")
    }(time.Now())

    return l.wrapper.Slice(param1, param2)
}

```
