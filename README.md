# Void

Package `void` provides a simple and more readable way to work with no space consuming structs.

## About

This project is inspired by https://dave.cheney.net/2014/03/25/the-empty-struct and Go 1.18 `any` alias.

Working with empty structs in Go can increase the performance of your code by reducing memory consumption. Empty structs are zero-sized types in Go that do not occupy any memory space.
Similar to `interface{}`, the syntax for creating an empty struct `struct{}{}` is somewhat incomprehensible and not easy to understand, which makes code less readable.

To solve this problem introducing void package :)

## Installation

    go get -u github.com/c14h3/void

## Documentation

https://pkg.go.dev/github.com/c14h3/void

## Usage examples

```go
import "github.com/c14h3/void"

keyMap := make(map[string]void.Void)
keyMap["foo"] = void.Value

limiterChan := make(chan void.Void, concurrentRequestLimit)
limiterChan <- void.Value
```