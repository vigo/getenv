package getenv

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

// sentinel errors.
var (
	ErrEnvironmentVariableIsNotSet = errors.New("environment variable is not set")
	ErrInvalidValue                = errors.New("invalid value")
)

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

// Bool retrieves the value of the environment variable named by the key
// and converts it to a boolean. It returns false if the variable is not set
// or cannot be converted to a boolean.
func Bool(name string) bool {
	value := strings.ToLower(os.Getenv(name))

	return slices.Contains([]string{"1", "t", "true"}, value)
}

// BoolWithDefault retrieves the value of the environment variable named by the key
// and converts it to a boolean. If the variable is not present or cannot be converted,
// it returns the provided default value.
func BoolWithDefault(name string, defaultValue bool) bool {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	value = strings.ToLower(value)

	return slices.Contains([]string{"1", "t", "true"}, value)
}

// BoolWithError retrieves the value of the environment variable named by the key
// and converts it to a boolean. If the variable is not present, it returns an error.
// Valid true values are "1", "t", "T", "true", "TRUE", "True"
// Valid false values are "0", "f", "F", "false", "FALSE", "False"
// .
func BoolWithError(name string) (bool, error) {
	value := os.Getenv(name)
	if value == "" {
		return false, fmt.Errorf("%s %w", name, ErrEnvironmentVariableIsNotSet)
	}

	switch strings.ToLower(value) {
	case "1", "t", "true":
		return true, nil
	case "0", "f", "false":
		return false, nil
	default:
		return false, fmt.Errorf("%s %w", name, ErrInvalidValue)
	}
}
