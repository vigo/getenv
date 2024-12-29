package getenv_test

import (
	"errors"
	"os"
	"testing"
	"time"

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
		"TEST_BOOL_TRUE_1":         "1",
		"TEST_BOOL_TRUE_T":         "t",
		"TEST_BOOL_TRUE_TCAP":      "T",
		"TEST_BOOL_TRUE_TRUE":      "true",
		"TEST_BOOL_TRUE_TCAP2":     "True",
		"TEST_BOOL_TRUE_TRUECAP":   "TRUE",
		"TEST_BOOL_FALSE_0":        "0",
		"TEST_BOOL_FALSE_F":        "f",
		"TEST_BOOL_FALSE_FCAP":     "F",
		"TEST_BOOL_FALSE_FALSE":    "false",
		"TEST_BOOL_FALSE_FALSECAP": "FALSE",
		"TEST_BOOL_FALSE_FCAP2":    "False",
		"TEST_BOOL_INVALID":        "invalid",
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
			name:     "true value 1",
			envName:  "TEST_BOOL_TRUE_1",
			expected: true,
		},
		{
			name:     "true value t",
			envName:  "TEST_BOOL_TRUE_T",
			expected: true,
		},
		{
			name:     "true value T",
			envName:  "TEST_BOOL_TRUE_TCAP",
			expected: true,
		},
		{
			name:     "true value true",
			envName:  "TEST_BOOL_TRUE_TRUE",
			expected: true,
		},
		{
			name:     "true value True",
			envName:  "TEST_BOOL_TRUE_TCAP2",
			expected: true,
		},
		{
			name:     "true value TRUE",
			envName:  "TEST_BOOL_TRUE_TRUECAP",
			expected: true,
		},
		{
			name:     "false value 0",
			envName:  "TEST_BOOL_FALSE_0",
			expected: false,
		},
		{
			name:     "false value f",
			envName:  "TEST_BOOL_FALSE_F",
			expected: false,
		},
		{
			name:     "false value F",
			envName:  "TEST_BOOL_FALSE_FCAP",
			expected: false,
		},
		{
			name:     "false value false",
			envName:  "TEST_BOOL_FALSE_FALSE",
			expected: false,
		},
		{
			name:     "false value FALSE",
			envName:  "TEST_BOOL_FALSE_FALSECAP",
			expected: false,
		},
		{
			name:     "false value False",
			envName:  "TEST_BOOL_FALSE_FCAP2",
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
	os.Setenv("TEST_BOOL_TRUE_1", "1")
	os.Setenv("TEST_BOOL_TRUE_T", "t")
	os.Setenv("TEST_BOOL_TRUE_T_UPPER", "T")
	os.Setenv("TEST_BOOL_TRUE_TRUE", "true")
	os.Setenv("TEST_BOOL_TRUE_TRUE_UPPER", "TRUE")
	os.Setenv("TEST_BOOL_TRUE_TRUE_CAPITALIZED", "True")
	os.Setenv("TEST_BOOL_FALSE_0", "0")
	os.Setenv("TEST_BOOL_FALSE_F", "f")
	os.Setenv("TEST_BOOL_FALSE_F_UPPER", "F")
	os.Setenv("TEST_BOOL_FALSE_FALSE", "false")
	os.Setenv("TEST_BOOL_FALSE_FALSE_UPPER", "FALSE")
	os.Setenv("TEST_BOOL_FALSE_FALSE_CAPITALIZED", "False")
	os.Setenv("TEST_BOOL_INVALID", "invalid")
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_BOOL_TRUE_1")
		os.Unsetenv("TEST_BOOL_TRUE_T")
		os.Unsetenv("TEST_BOOL_TRUE_T_UPPER")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE_UPPER")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE_CAPITALIZED")
		os.Unsetenv("TEST_BOOL_FALSE_0")
		os.Unsetenv("TEST_BOOL_FALSE_F")
		os.Unsetenv("TEST_BOOL_FALSE_F_UPPER")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE_UPPER")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE_CAPITALIZED")
		os.Unsetenv("TEST_BOOL_INVALID")
	}()

	testcases := []struct {
		name         string
		envName      string
		defaultValue any
		expected     bool
	}{
		{"true value '1'", "TEST_BOOL_TRUE_1", false, true},
		{"true value 't'", "TEST_BOOL_TRUE_T", false, true},
		{"true value 'T'", "TEST_BOOL_TRUE_T_UPPER", false, true},
		{"true value 'true'", "TEST_BOOL_TRUE_TRUE", false, true},
		{"true value 'TRUE'", "TEST_BOOL_TRUE_TRUE_UPPER", false, true},
		{"true value 'True'", "TEST_BOOL_TRUE_TRUE_CAPITALIZED", false, true},

		{"false value '0'", "TEST_BOOL_FALSE_0", true, false},
		{"false value 'f'", "TEST_BOOL_FALSE_F", true, false},
		{"false value 'F'", "TEST_BOOL_FALSE_F_UPPER", true, false},
		{"false value 'false'", "TEST_BOOL_FALSE_FALSE", true, false},
		{"false value 'FALSE'", "TEST_BOOL_FALSE_FALSE_UPPER", true, false},
		{"false value 'False'", "TEST_BOOL_FALSE_FALSE_CAPITALIZED", true, false},

		{"invalid value", "TEST_BOOL_INVALID", true, false},

		{"non-existent value with true default", "TEST_BOOL_NON_EXISTENT", true, true},
		{"non-existent value with false default", "TEST_BOOL_NON_EXISTENT", false, false},
		{"non-existent value with string true default", "TEST_BOOL_NON_EXISTENT", "true", true},
		{"non-existent value with string false default", "TEST_BOOL_NON_EXISTENT", "false", false},
		{"non-existent value with invalid string default", "TEST_BOOL_NON_EXISTENT", "invalid", false},
		{"non-existent value with non-boolean default", "TEST_BOOL_NON_EXISTENT", 123, false},
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
	os.Setenv("TEST_BOOL_TRUE_1", "1")
	os.Setenv("TEST_BOOL_TRUE_T", "t")
	os.Setenv("TEST_BOOL_TRUE_T_UPPER", "T")
	os.Setenv("TEST_BOOL_TRUE_TRUE", "true")
	os.Setenv("TEST_BOOL_TRUE_TRUE_UPPER", "TRUE")
	os.Setenv("TEST_BOOL_TRUE_TRUE_CAPITALIZED", "True")
	os.Setenv("TEST_BOOL_FALSE_0", "0")
	os.Setenv("TEST_BOOL_FALSE_F", "f")
	os.Setenv("TEST_BOOL_FALSE_F_UPPER", "F")
	os.Setenv("TEST_BOOL_FALSE_FALSE", "false")
	os.Setenv("TEST_BOOL_FALSE_FALSE_UPPER", "FALSE")
	os.Setenv("TEST_BOOL_FALSE_FALSE_CAPITALIZED", "False")
	os.Setenv("TEST_BOOL_INVALID", "invalid")
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_BOOL_TRUE_1")
		os.Unsetenv("TEST_BOOL_TRUE_T")
		os.Unsetenv("TEST_BOOL_TRUE_T_UPPER")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE_UPPER")
		os.Unsetenv("TEST_BOOL_TRUE_TRUE_CAPITALIZED")
		os.Unsetenv("TEST_BOOL_FALSE_0")
		os.Unsetenv("TEST_BOOL_FALSE_F")
		os.Unsetenv("TEST_BOOL_FALSE_F_UPPER")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE_UPPER")
		os.Unsetenv("TEST_BOOL_FALSE_FALSE_CAPITALIZED")
		os.Unsetenv("TEST_BOOL_INVALID")
	}()

	testcases := []struct {
		name          string
		envName       string
		expectedValue bool
		expectedError error
	}{
		{"true value '1'", "TEST_BOOL_TRUE_1", true, nil},
		{"true value 't'", "TEST_BOOL_TRUE_T", true, nil},
		{"true value 'T'", "TEST_BOOL_TRUE_T_UPPER", true, nil},
		{"true value 'true'", "TEST_BOOL_TRUE_TRUE", true, nil},
		{"true value 'TRUE'", "TEST_BOOL_TRUE_TRUE_UPPER", true, nil},
		{"true value 'True'", "TEST_BOOL_TRUE_TRUE_CAPITALIZED", true, nil},

		{"false value '0'", "TEST_BOOL_FALSE_0", false, nil},
		{"false value 'f'", "TEST_BOOL_FALSE_F", false, nil},
		{"false value 'F'", "TEST_BOOL_FALSE_F_UPPER", false, nil},
		{"false value 'false'", "TEST_BOOL_FALSE_FALSE", false, nil},
		{"false value 'FALSE'", "TEST_BOOL_FALSE_FALSE_UPPER", false, nil},
		{"false value 'False'", "TEST_BOOL_FALSE_FALSE_CAPITALIZED", false, nil},

		{"invalid value", "TEST_BOOL_INVALID", false, getenv.ErrInvalidValue},

		{"non-existent value", "TEST_BOOL_NON_EXISTENT", false, getenv.ErrEnvironmentVariableIsNotSet},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.BoolWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %v, got: %v", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestInt(t *testing.T) {
	os.Setenv("TEST_INT_POSITIVE", "123")
	os.Setenv("TEST_INT_NEGATIVE", "-456")
	os.Setenv("TEST_INT_ZERO", "0")
	os.Setenv("TEST_INT_INVALID", "invalid")
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_INT_POSITIVE")
		os.Unsetenv("TEST_INT_NEGATIVE")
		os.Unsetenv("TEST_INT_ZERO")
		os.Unsetenv("TEST_INT_INVALID")
	}()

	testcases := []struct {
		name     string
		envName  string
		expected int
	}{
		{"positive integer value", "TEST_INT_POSITIVE", 123},
		{"negative integer value", "TEST_INT_NEGATIVE", -456},
		{"zero integer value", "TEST_INT_ZERO", 0},
		{"invalid integer value", "TEST_INT_INVALID", 0},
		{"non-existent variable", "TEST_INT_NON_EXISTENT", 0},
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
	os.Setenv("TEST_INT_NEGATIVE", "-456")
	os.Setenv("TEST_INT_ZERO", "0")
	os.Setenv("TEST_INT_INVALID", "invalid")
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_INT_POSITIVE")
		os.Unsetenv("TEST_INT_NEGATIVE")
		os.Unsetenv("TEST_INT_ZERO")
		os.Unsetenv("TEST_INT_INVALID")
	}()

	testcases := []struct {
		name         string
		envName      string
		defaultValue any
		expected     int
	}{
		{"positive integer value", "TEST_INT_POSITIVE", 10, 123},
		{"negative integer value", "TEST_INT_NEGATIVE", 10, -456},
		{"zero integer value", "TEST_INT_ZERO", 10, 0},

		{"invalid integer value with int default", "TEST_INT_INVALID", 10, 10},
		{"invalid integer value with string default", "TEST_INT_INVALID", "15", 15},

		{"non-existent variable with int default", "TEST_INT_NON_EXISTENT", 10, 10},
		{"non-existent variable with string default", "TEST_INT_NON_EXISTENT", "20", 20},
		{"non-existent variable with invalid string default", "TEST_INT_NON_EXISTENT", "invalid", 0},

		{"non-existent variable with non-integer default", "TEST_INT_NON_EXISTENT", 1.5, 0},
		{"non-existent variable with nil default", "TEST_INT_NON_EXISTENT", nil, 0},
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
	os.Setenv("TEST_INT_NEGATIVE", "-456")
	os.Setenv("TEST_INT_ZERO", "0")
	os.Setenv("TEST_INT_INVALID", "invalid")
	os.Unsetenv("TEST_INT_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_INT_POSITIVE")
		os.Unsetenv("TEST_INT_NEGATIVE")
		os.Unsetenv("TEST_INT_ZERO")
		os.Unsetenv("TEST_INT_INVALID")
	}()

	testcases := []struct {
		name          string
		envName       string
		expectedValue int
		expectedError error
	}{
		{"positive integer value", "TEST_INT_POSITIVE", 123, nil},
		{"negative integer value", "TEST_INT_NEGATIVE", -456, nil},
		{"zero integer value", "TEST_INT_ZERO", 0, nil},

		{"invalid integer value", "TEST_INT_INVALID", 0, getenv.ErrInvalidValue},

		{"non-existent variable", "TEST_INT_NON_EXISTENT", 0, getenv.ErrEnvironmentVariableIsNotSet},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.IntWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %d, got: %d", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	os.Setenv("TEST_DURATION_VALID", "5s")
	os.Setenv("TEST_DURATION_MINUTES", "2m")
	os.Setenv("TEST_DURATION_HOURS", "1h")
	os.Setenv("TEST_DURATION_ZERO", "0s")
	os.Setenv("TEST_DURATION_INVALID", "invalid")
	os.Unsetenv("TEST_DURATION_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_DURATION_VALID")
		os.Unsetenv("TEST_DURATION_MINUTES")
		os.Unsetenv("TEST_DURATION_HOURS")
		os.Unsetenv("TEST_DURATION_ZERO")
		os.Unsetenv("TEST_DURATION_INVALID")
	}()

	testcases := []struct {
		name     string
		envName  string
		expected time.Duration
	}{
		{"valid duration seconds", "TEST_DURATION_VALID", 5 * time.Second},
		{"valid duration minutes", "TEST_DURATION_MINUTES", 2 * time.Minute},
		{"valid duration hours", "TEST_DURATION_HOURS", 1 * time.Hour},

		{"zero duration", "TEST_DURATION_ZERO", 0},
		{"invalid duration value", "TEST_DURATION_INVALID", 0},
		{"non-existent variable", "TEST_DURATION_NON_EXISTENT", 0},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.Duration(tc.envName); val != tc.expected {
				t.Errorf("want: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestDurationWithDefault(t *testing.T) {
	os.Setenv("TEST_DURATION_VALID", "10s")
	os.Setenv("TEST_DURATION_INVALID", "invalid")
	os.Setenv("TEST_DURATION_EMPTY", "")
	os.Unsetenv("TEST_DURATION_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_DURATION_VALID")
		os.Unsetenv("TEST_DURATION_INVALID")
		os.Unsetenv("TEST_DURATION_EMPTY")
	}()

	testcases := []struct {
		name         string
		envName      string
		defaultValue any
		expected     time.Duration
	}{
		{"valid duration (seconds)", "TEST_DURATION_VALID", "5s", 10 * time.Second},

		{"invalid duration with string default", "TEST_DURATION_INVALID", "15s", 15 * time.Second},
		{"invalid duration with time.Duration default", "TEST_DURATION_INVALID", 20 * time.Second, 20 * time.Second},

		{"empty duration with string default", "TEST_DURATION_EMPTY", "30s", 30 * time.Second},
		{"empty duration with time.Duration default", "TEST_DURATION_EMPTY", 45 * time.Second, 45 * time.Second},

		{"non-existent variable with string default", "TEST_DURATION_NON_EXISTENT", "1m", 1 * time.Minute},
		{
			"non-existent variable with time.Duration default",
			"TEST_DURATION_NON_EXISTENT",
			90 * time.Second,
			90 * time.Second,
		},

		{"non-existent variable with invalid string default", "TEST_DURATION_NON_EXISTENT", "invalid", 0},
		{"non-existent variable with unsupported default type", "TEST_DURATION_NON_EXISTENT", 123, 0},
		{"non-existent variable with nil default", "TEST_DURATION_NON_EXISTENT", nil, 0},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if val := getenv.DurationWithDefault(tc.envName, tc.defaultValue); val != tc.expected {
				t.Errorf("want: %v, got: %v", tc.expected, val)
			}
		})
	}
}

func TestDurationWithError(t *testing.T) {
	os.Setenv("TEST_DURATION_VALID", "10s")
	os.Setenv("TEST_DURATION_INVALID", "invalid")
	os.Unsetenv("TEST_DURATION_NON_EXISTENT")

	defer func() {
		os.Unsetenv("TEST_DURATION_VALID")
		os.Unsetenv("TEST_DURATION_INVALID")
	}()

	testcases := []struct {
		name          string
		envName       string
		expectedValue time.Duration
		expectedError error
	}{
		{"valid duration", "TEST_DURATION_VALID", 10 * time.Second, nil},

		{"invalid duration", "TEST_DURATION_INVALID", 0, getenv.ErrInvalidValue},

		{"non-existent variable", "TEST_DURATION_NON_EXISTENT", 0, getenv.ErrEnvironmentVariableIsNotSet},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := getenv.DurationWithError(tc.envName)

			if val != tc.expectedValue {
				t.Errorf("want: %v, got: %v", tc.expectedValue, val)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("want error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
