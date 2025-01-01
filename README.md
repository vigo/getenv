![Version](https://img.shields.io/badge/version-0.0.0-orange.svg)
![Go](https://img.shields.io/github/go-mod/go-version/vigo/getenv)
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
```

There are also plans to add support for custom types in the future.

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
listen := getenv.TCPAddr("LISTEN", ":4000") // LISTEN doesn’t exist in the environment
if err := getenv.Parse(); err != nil {
	log.Fatal(err)
}

fmt.Println(*listen) // :4000 as string
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
