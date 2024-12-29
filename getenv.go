package getenv

import "os"

// Value is the interface for custom types that can parse and hold environment
// variable values.
type Value interface {
	String() string
	Set(value string) error
}

// String retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns an empty string.
func String(name string) string {
	return os.Getenv(name)
}
