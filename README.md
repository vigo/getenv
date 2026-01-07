![Version](https://img.shields.io/badge/version-0.0.2-orange.svg)
![Go](https://img.shields.io/github/go-mod/go-version/vigo/getenv)
[![Test go code](https://github.com/vigo/getenv/actions/workflows/test.yml/badge.svg)](https://github.com/vigo/getenv/actions/workflows/test.yml)
[![golangci-lint](https://github.com/vigo/getenv/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/vigo/getenv/actions/workflows/golangci-lint.yml)
[![Documentation](https://godoc.org/github.com/vigo/getenv?status.svg)](https://pkg.go.dev/github.com/vigo/getenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/vigo/getenv)](https://goreportcard.com/report/github.com/vigo/getenv)
[![codecov](https://codecov.io/github/vigo/getenv/graph/badge.svg?token=OHI46PIBLN)](https://codecov.io/github/vigo/getenv)
![Powered by Rake](https://img.shields.io/badge/powered_by-rake-blue?logo=ruby)

# getenv

A minimalist Go library for type-safe environment variable parsing, inspired
by the standard `flag` package.

This library allows you to effortlessly use environment variables within your
code. It reads the values of environment variables using built-in types,
assigns fallback (*default*) values if no value is set, and performs type-based
validation to ensure correctness.

The currently supported built-in types are:

```go
getenv.Bool
getenv.Int
getenv.Int64
getenv.Float64
getenv.String
getenv.Duration
getenv.TCPAddr
getenv.StringSlice
getenv.LogLevel
```

---

## Installation

```bash
go get -u github.com/vigo/getenv
```

---

## Usage

For `getenv.Bool`:

```go
color := getenv.Bool("COLOR", false)  // COLOR doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*color) // false as bool
```

For `getenv.Int`:

```go
port := getenv.Int("PORT", 8000) // PORT doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*port) // 8000 as int
```

For `getenv.Int64`:

```go
long := getenv.Int64("LONG", 9223372036854775806) // LONG doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*long) // 9223372036854775806 as int64
```

For `getenv.Float64`:

```go
xFactor := getenv.Float64("X_FACTOR", 1.1) // X_FACTOR doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*xFactor) // 1.1 as float64
```

For `getenv.String`:

```go
hmacHeader := getenv.String("HMAC_HEADER", "X-Foo-Signature") // HMAC_HEADER doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*hmacHeader) // X-Foo-Signature as string
```

For `getenv.Duration`:

```go
timeout := getenv.Duration("SERVER_TIMEOUT", 5*time.Second) // SERVER_TIMEOUT doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*timeout) // 5s as time.Duration
```

For `getenv.TCPAddr`:

```go
listen := getenv.TCPAddr("LISTEN", ":4000") // LISTEN doesn't exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*listen) // :4000 as string
```

For `getenv.StringSlice`:

```go
// BROKERS doesn't exist in the environment
brokers := getenv.StringSlice("BROKERS", []string{":9092", ":9093"})
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*brokers) // [:9092 :9093] as []string

// if BROKERS=":9092,:9093,127.0.0.1:9094" exists in the environment
// result will be: [:9092 :9093 127.0.0.1:9094]
// - values are split by comma
// - whitespace is trimmed
// - empty values are filtered out
```

For `getenv.LogLevel`:

```go
// define your custom log levels
levels := map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"FATAL": 4,
}

// LOG_LEVEL doesn't exist in the environment, default is INFO (1)
logLevel := getenv.LogLevel("LOG_LEVEL", levels, 1)
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*logLevel) // 1 as int

// if LOG_LEVEL=debug exists in the environment (case-insensitive)
// result will be: 0
```

For all of them together:

```go
color := getenv.Bool("COLOR", false)
port := getenv.Int("PORT", 8000)
long := getenv.Int64("LONG", 9223372036854775806)
xFactor := getenv.Float64("X_FACTOR", 1.1)
hmacHeader := getenv.String("HMAC_HEADER", "X-Foo-Signature")
timeout := getenv.Duration("SERVER_TIMEOUT", 5*time.Second)
listen := getenv.TCPAddr("LISTEN", ":4000")
brokers := getenv.StringSlice("BROKERS", []string{":9092", ":9093"})

levels := map[string]int{"DEBUG": 0, "INFO": 1, "WARN": 2, "ERROR": 3}
logLevel := getenv.LogLevel("LOG_LEVEL", levels, 1)

if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

// now you have all the variables accessible via pointer...
```

Package also provides error types:

```go
getenv.ErrInvalid
getenv.ErrEnvironmentVariableIsEmpty
```

Use with:

- `errors.Is(err, getenv.ErrInvalid)`
- `errors.Is(err, getenv.ErrEnvironmentVariableIsEmpty)`

Feel free to contribute!

---

## Rake Tasks

```bash
rake -T

rake coverage  # show test coverage
rake test      # run test
```

---

## License

This project is licensed under MIT (MIT)

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com/vigo/getenv/blob/main/CODE_OF_CONDUCT.md
