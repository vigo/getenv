package getenv

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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
// or cannot be converted to a boolean. true values are:
// "1", "t", "true", "T", "True", "TRUE".
func Bool(name string) bool {
	value := strings.ToLower(os.Getenv(name))
	trueValues := []string{"1", "t", "true", "T", "True", "TRUE"}

	return slices.Contains(trueValues, value)
}

// BoolWithDefault retrieves the value of the environment variable named by the key
// and converts it to a boolean. If the variable is not present or cannot be converted,
// it returns the provided default value, which can be either a bool or a string
// representing a boolean "1", "t", "true", "T", "True", "TRUE".
func BoolWithDefault(name string, defaultValue any) bool {
	value := os.Getenv(name)
	if value == "" {
		return parseDefaultBool(defaultValue)
	}

	value = strings.ToLower(value)

	return isTrue(value)
}

func parseDefaultBool(defaultValue any) bool {
	switch v := defaultValue.(type) {
	case bool:
		return v
	case string:
		return isTrue(strings.ToLower(v))
	default:
		return false
	}
}

// isTrue checks if a string represents a true value.
func isTrue(value string) bool {
	trueValues := []string{"1", "t", "true", "T", "True", "TRUE"}

	return contains(trueValues, value)
}

// contains checks if a slice contains a value.
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
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

// Int retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns an 0, zero value of int.
func Int(name string) int {
	value := os.Getenv(name)
	if value == "" {
		return 0
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return i
}

// IntWithDefault retrieves the value of the environment variable named by the key.
// If the variable is not present or cannot be converted, it returns the provided default value.
// The defaultValue can be an int or a string containing a valid integer.
func IntWithDefault(name string, defaultValue any) int {
	value := os.Getenv(name)
	if value == "" {
		return parseDefaultInt(defaultValue)
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return parseDefaultInt(defaultValue)
	}

	return i
}

func parseDefaultInt(defaultValue any) int {
	switch v := defaultValue.(type) {
	case int:
		return v
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}

		return i
	default:
		return 0
	}
}

// IntWithError retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns 0 and no error.
// If the variable's value cannot be converted to an integer, it returns an error.
func IntWithError(name string) (int, error) {
	value := os.Getenv(name)
	if value == "" {
		return 0, fmt.Errorf("%s %w", name, ErrEnvironmentVariableIsNotSet)
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s %w", name, ErrInvalidValue)
	}

	return i, nil
}

// Duration retrieves the value of the environment variable named by the key as a time.Duration.
// If the variable is not present or cannot be parsed, it returns 0.
func Duration(name string) time.Duration {
	value := os.Getenv(name)
	if value == "" {
		return 0
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		return 0
	}

	return d
}

// DurationWithDefault retrieves the value of the environment variable named by the key as a time.Duration.
// The defaultValue can be either a string (e.g., "5s") or a time.Duration. If it's a string,
// it will be parsed using time.ParseDuration. If the variable is not present or cannot be parsed,
// it returns the default value.
func DurationWithDefault(name string, defaultValue any) time.Duration {
	value := os.Getenv(name)
	if value == "" {
		return parseDefaultDuration(defaultValue)
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		return parseDefaultDuration(defaultValue)
	}

	return d
}

func parseDefaultDuration(defaultValue any) time.Duration {
	switch v := defaultValue.(type) {
	case time.Duration:
		return v
	case string:
		d, err := time.ParseDuration(v)
		if err != nil {
			return 0
		}

		return d
	default:
		return 0
	}
}

// DurationWithError retrieves the value of the environment variable named by the key as a time.Duration.
// If the variable is not present, it returns 0 and an ErrEnvironmentVariableIsNotSet error.
// If the variable's value cannot be parsed, it returns 0 and an ErrInvalidValue error.
func DurationWithError(name string) (time.Duration, error) {
	value := os.Getenv(name)
	if value == "" {
		return 0, fmt.Errorf("%s %w", name, ErrEnvironmentVariableIsNotSet)
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s %w", name, ErrInvalidValue)
	}

	return d, nil
}
