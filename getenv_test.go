package getenv_test

import (
	"errors"
	"os"
	"testing"

	"github.com/vigo/getenv"
)

func TestString(t *testing.T) {
	setEnvVars := map[string]string{
		"TEST_STRING_ONE": "One",
		"TEST_STRING_TWO": "2",
	}

	for k, v := range setEnvVars {
		os.Setenv(k, v)
	}
	os.Unsetenv("TEST_STRING_NON_EXISTENT")

	defer func() {
		for k := range setEnvVars {
			os.Unsetenv(k)
		}
	}()

	testcases := []struct {
		name     string
		envName  string
		expected string
	}{
		{
			name:     "get TEST_STRING_ONE",
			envName:  "TEST_STRING_ONE",
			expected: "One",
		},
		{
			name:     "get TEST_STRING_TWO",
			envName:  "TEST_STRING_TWO",
			expected: "2",
		},
		{
			name:     "get TEST_STRING_NON_EXISTENT",
			envName:  "TEST_STRING_NON_EXISTENT",
			expected: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.String(tc.envName); val != tc.expected {
				t.Errorf("want: %s, got: %s", tc.expected, val)
			}
		})
	}
}

func TestStringWithDefault(t *testing.T) {
	os.Setenv("TEST_STRING_FOO", "foo")
	os.Unsetenv("TEST_STRING_NON_EXISTENT")

	defer func() { os.Unsetenv("TEST_STRING_FOO") }()

	testcases := []struct {
		name         string
		envName      string
		defaultValue string
		expected     string
	}{
		{
			name:         "get TEST_STRING_NON_EXISTENT with default value",
			envName:      "TEST_STRING_NON_EXISTENT",
			defaultValue: "None",
			expected:     "None",
		},
		{
			name:         "get TEST_STRING_FOO with default value",
			envName:      "TEST_STRING_FOO",
			defaultValue: "Baz",
			expected:     "foo",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.StringWithDefault(tc.envName, tc.defaultValue); val != tc.expected {
				t.Errorf("want: %s, got: %s", tc.expected, val)
			}
		})
	}
}

func TestStringWithError(t *testing.T) {
	os.Setenv("TEST_STRING_FOO", "foo")
	os.Unsetenv("TEST_STRING_NON_EXISTENT")

	defer func() { os.Unsetenv("TEST_STRING_FOO") }()

	testcases := []struct {
		name          string
		envName       string
		expectedValue string
		expectedError error
	}{
		{
			name:          "get TEST_STRING_NON_EXISTENT with error",
			envName:       "TEST_STRING_NON_EXISTENT",
			expectedValue: "",
			expectedError: getenv.ErrEnvironmentVariableIsNotSet,
		},
		{
			name:          "get TEST_STRING_FOO with no error",
			envName:       "TEST_STRING_FOO",
			expectedValue: "foo",
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.StringWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %s, got: %s", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestBool(t *testing.T) {
	setEnvVars := map[string]string{
		"TEST_BOOL_TRUE":    "true",
		"TEST_BOOL_T":       "t",
		"TEST_BOOL_1":       "1",
		"TEST_BOOL_FALSE":   "false",
		"TEST_BOOL_F":       "f",
		"TEST_BOOL_0":       "0",
		"TEST_BOOL_INVALID": "invalid",
	}

	for k, v := range setEnvVars {
		os.Setenv(k, v)
	}
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")

	defer func() {
		for k := range setEnvVars {
			os.Unsetenv(k)
		}
	}()

	testcases := []struct {
		name     string
		envName  string
		expected bool
	}{
		{
			name:     "true value",
			envName:  "TEST_BOOL_TRUE",
			expected: true,
		},
		{
			name:     "t value",
			envName:  "TEST_BOOL_T",
			expected: true,
		},
		{
			name:     "1 value",
			envName:  "TEST_BOOL_1",
			expected: true,
		},
		{
			name:     "false value",
			envName:  "TEST_BOOL_FALSE",
			expected: false,
		},
		{
			name:     "f value",
			envName:  "TEST_BOOL_F",
			expected: false,
		},
		{
			name:     "0 value",
			envName:  "TEST_BOOL_0",
			expected: false,
		},
		{
			name:     "invalid value",
			envName:  "TEST_BOOL_INVALID",
			expected: false,
		},
		{
			name:     "non-existent value",
			envName:  "TEST_BOOL_NON_EXISTENT",
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.Bool(tc.envName); val != tc.expected {
				t.Errorf("want: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestBoolWithDefault(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	os.Setenv("TEST_BOOL_FALSE", "false")
	os.Setenv("TEST_BOOL_INVALID", "invalid")
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_BOOL_TRUE")
		os.Unsetenv("TEST_BOOL_FALSE")
		os.Unsetenv("TEST_BOOL_INVALID")
	}()

	testcases := []struct {
		name         string
		envName      string
		defaultValue bool
		expected     bool
	}{
		{
			name:         "true value with false default",
			envName:      "TEST_BOOL_TRUE",
			defaultValue: false,
			expected:     true,
		},
		{
			name:         "false value with true default",
			envName:      "TEST_BOOL_FALSE",
			defaultValue: true,
			expected:     false,
		},
		{
			name:         "invalid value with true default",
			envName:      "TEST_BOOL_INVALID",
			defaultValue: true,
			expected:     false,
		},
		{
			name:         "non-existent value with true default",
			envName:      "TEST_BOOL_NON_EXISTENT",
			defaultValue: true,
			expected:     true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.BoolWithDefault(tc.envName, tc.defaultValue); val != tc.expected {
				t.Errorf("want: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestBoolWithError(t *testing.T) {
	os.Setenv("TEST_BOOL_TRUE", "true")
	os.Setenv("TEST_BOOL_FALSE", "false")
	os.Setenv("TEST_BOOL_INVALID", "invalid")
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_BOOL_TRUE")
		os.Unsetenv("TEST_BOOL_FALSE")
		os.Unsetenv("TEST_BOOL_INVALID")
	}()

	testcases := []struct {
		name          string
		envName       string
		expectedValue bool
		expectedError error
	}{
		{
			name:          "true value",
			envName:       "TEST_BOOL_TRUE",
			expectedValue: true,
			expectedError: nil,
		},
		{
			name:          "false value",
			envName:       "TEST_BOOL_FALSE",
			expectedValue: false,
			expectedError: nil,
		},
		{
			name:          "invalid value",
			envName:       "TEST_BOOL_INVALID",
			expectedValue: false,
			expectedError: getenv.ErrInvalidValue,
		},
		{
			name:          "non-existent value",
			envName:       "TEST_BOOL_NON_EXISTENT",
			expectedValue: false,
			expectedError: getenv.ErrEnvironmentVariableIsNotSet,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.BoolWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %v, got: %v", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestInt(t *testing.T) {
	setEnvVars := map[string]string{
		"TEST_INT_POSITIVE": "123",
		"TEST_INT_NEGATIVE": "-456",
		"TEST_INT_ZERO":     "0",
		"TEST_INT_INVALID":  "invalid",
	}

	for k, v := range setEnvVars {
		os.Setenv(k, v)
	}
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		for k := range setEnvVars {
			os.Unsetenv(k)
		}
	}()

	testcases := []struct {
		name     string
		envName  string
		expected int
	}{
		{
			name:     "get positive integer",
			envName:  "TEST_INT_POSITIVE",
			expected: 123,
		},
		{
			name:     "get negative integer",
			envName:  "TEST_INT_NEGATIVE",
			expected: -456,
		},
		{
			name:     "get zero value",
			envName:  "TEST_INT_ZERO",
			expected: 0,
		},
		{
			name:     "get invalid integer",
			envName:  "TEST_INT_INVALID",
			expected: 0,
		},
		{
			name:     "get non-existent integer",
			envName:  "TEST_INT_NON_EXISTENT",
			expected: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.Int(tc.envName); val != tc.expected {
				t.Errorf("want: %d, got: %d", tc.expected, val)
			}
		})
	}
}

func TestIntWithDefault(t *testing.T) {
	os.Setenv("TEST_INT_POSITIVE", "123")
	os.Setenv("TEST_INT_INVALID", "invalid")
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_INT_POSITIVE")
		os.Unsetenv("TEST_INT_INVALID")
	}()

	testcases := []struct {
		name         string
		envName      string
		defaultValue int
		expected     int
	}{
		{
			name:         "get valid integer with default",
			envName:      "TEST_INT_POSITIVE",
			defaultValue: 999,
			expected:     123,
		},
		{
			name:         "get invalid integer with default",
			envName:      "TEST_INT_INVALID",
			defaultValue: 999,
			expected:     0,
		},
		{
			name:         "get non-existent integer with default",
			envName:      "TEST_INT_NON_EXISTENT",
			defaultValue: 999,
			expected:     999,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.IntWithDefault(tc.envName, tc.defaultValue); val != tc.expected {
				t.Errorf("want: %d, got: %d", tc.expected, val)
			}
		})
	}
}

func TestIntWithError(t *testing.T) {
	os.Setenv("TEST_INT_POSITIVE", "123")
	os.Setenv("TEST_INT_INVALID", "invalid")
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_INT_POSITIVE")
		os.Unsetenv("TEST_INT_INVALID")
	}()

	testcases := []struct {
		name          string
		envName       string
		expectedValue int
		expectedError error
	}{
		{
			name:          "get valid integer",
			envName:       "TEST_INT_POSITIVE",
			expectedValue: 123,
			expectedError: nil,
		},
		{
			name:          "get invalid integer",
			envName:       "TEST_INT_INVALID",
			expectedValue: 0,
			expectedError: getenv.ErrInvalidValue,
		},
		{
			name:          "get non-existent integer",
			envName:       "TEST_INT_NON_EXISTENT",
			expectedValue: 0,
			expectedError: getenv.ErrEnvironmentVariableIsNotSet,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.IntWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %d, got: %d", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
