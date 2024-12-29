package getenv

import (
	"errors"
	"fmt"
	"os"
)

// sentinel errors.
var (
	ErrEnvironmentVariableIsNotSet = errors.New("environment variable is not set")
)

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

// StringWithDefault retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the provided default value.
func StringWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}

// StringWithError retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns an error.
func StringWithError(name string) (string, error) {
	value := os.Getenv(name)
	if value == "" {
		return "", fmt.Errorf("%s %w", name, ErrEnvironmentVariableIsNotSet)
	}

	return value, nil
}
