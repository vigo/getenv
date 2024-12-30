package getenv_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/vigo/getenv"
)

var errTestOnly = errors.New("error")

func TestString(t *testing.T) {
	os.Setenv("TEST_STRING_1", "test1")
	os.Unsetenv("TEST_STRING_NON_EXISTENT")

	defer func() { os.Unsetenv("TEST_STRING_1") }()

	testcases := []struct {
		testName    string
		envName     string
		defaultVal  string
		verifyFunc  getenv.StringVerifyFunc
		expectedVal string
		expectedErr error
	}{
		{
			testName:    "non existing var with default",
			envName:     "TEST_STRING_NON_EXISTENT",
			defaultVal:  "default",
			verifyFunc:  nil,
			expectedVal: "default",
			expectedErr: nil,
		},
		{
			testName:    "existing var with default",
			envName:     "TEST_STRING_1",
			defaultVal:  "test-default",
			verifyFunc:  nil,
			expectedVal: "test1",
			expectedErr: nil,
		},
		{
			testName:   "existing var with default and verifier, should return an error",
			envName:    "TEST_STRING_1",
			defaultVal: "test-default",
			verifyFunc: func(val string) error {
				if val == "test1" {
					return fmt.Errorf("%w", errTestOnly)
				}
				return nil
			},
			expectedVal: "",
			expectedErr: errTestOnly,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := getenv.String(tc.envName, tc.defaultVal, tc.verifyFunc)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want: %v, got: %v", tc.expectedErr, err)
			}
			if val != tc.expectedVal {
				t.Errorf("want: %s, got: %s", tc.expectedVal, val)
			}
		})
	}
}

func ExampleString_non_existing_var_without_error() {
	value, _ := getenv.String("NON_EXISTENT", "default-value", nil)
	fmt.Println(value)
	// Output:
	// default-value
}

func TestBool(t *testing.T) {
	os.Setenv("TEST_BOOL_1", "TRUE")
	os.Unsetenv("TEST_BOOL_NON_EXISTENT")
	defer func() { os.Unsetenv("TEST_BOOL_1") }()

	testcases := []struct {
		testName    string
		envName     string
		defaultVal  any
		verifyFunc  getenv.BoolVerifyFunc
		expectedVal bool
		expectedErr error
	}{
		{
			testName:    "test existing env-var with TRUE value",
			envName:     "TEST_BOOL_1",
			defaultVal:  "False",
			verifyFunc:  nil,
			expectedVal: true,
			expectedErr: nil,
		},
		{
			testName:    "test non existing env-var with default True value",
			envName:     "TEST_BOOL_NON_EXISTENT",
			defaultVal:  "True",
			verifyFunc:  nil,
			expectedVal: true,
			expectedErr: nil,
		},
		{
			testName:   "test non existing env-var with default errorious value",
			envName:    "TEST_BOOL_NON_EXISTENT",
			defaultVal: "invalid",
			verifyFunc: func(val bool) error {
				if !val {
					return fmt.Errorf("%w", errTestOnly)
				}
				return nil
			},
			expectedVal: false,
			expectedErr: errTestOnly,
		},
		{
			testName:    "test non existing env-var with default bool value",
			envName:     "TEST_BOOL_NON_EXISTENT",
			defaultVal:  true,
			verifyFunc:  nil,
			expectedVal: true,
			expectedErr: nil,
		},
		{
			testName:    "test non existing env-var with default int value",
			envName:     "TEST_BOOL_NON_EXISTENT",
			defaultVal:  999,
			verifyFunc:  nil,
			expectedVal: false,
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := getenv.Bool(tc.envName, tc.defaultVal, tc.verifyFunc)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want: %v, got: %v", tc.expectedErr, err)
			}
			if val != tc.expectedVal {
				t.Errorf("want: %t, got: %t", tc.expectedVal, val)
			}
		})
	}
}
