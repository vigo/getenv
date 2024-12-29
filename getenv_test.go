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
