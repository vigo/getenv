package getenv_test

import (
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
				t.Errorf("want: %s, got: %s", val, tc.expected)
			}
		})
	}
}
