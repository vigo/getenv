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

// func ExampleString_non_existing_var_without_error() {
// 	value, _ := getenv.String("NON_EXISTENT", "default-value", nil)
// 	fmt.Println(value)
// 	// Output:
// 	// default-value
// }

// func TestBool(t *testing.T) {
// 	os.Setenv("TEST_BOOL_1", "TRUE")
// 	os.Setenv("TEST_BOOL_2", "True")
// 	os.Setenv("TEST_BOOL_3", "true")
// 	os.Setenv("TEST_BOOL_4", "T")
// 	os.Setenv("TEST_BOOL_5", "t")
//
// 	os.Unsetenv("TEST_BOOL_NON_EXISTENT")
// 	defer func() {
// 		os.Unsetenv("TEST_BOOL_1")
// 		os.Unsetenv("TEST_BOOL_2")
// 		os.Unsetenv("TEST_BOOL_3")
// 		os.Unsetenv("TEST_BOOL_4")
// 		os.Unsetenv("TEST_BOOL_5")
// 	}()
//
// 	testcases := []struct {
// 		testName    string
// 		envName     string
// 		defaultVal  any
// 		verifyFunc  getenv.BoolVerifyFunc
// 		expectedVal bool
// 		expectedErr error
// 	}{
// 		{
// 			testName:    "test TEST_BOOL_1 with TRUE value",
// 			envName:     "TEST_BOOL_1",
// 			defaultVal:  "False",
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_2 with True value",
// 			envName:     "TEST_BOOL_2",
// 			defaultVal:  "",
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_3 with true value",
// 			envName:     "TEST_BOOL_3",
// 			defaultVal:  nil,
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_4 with T value",
// 			envName:     "TEST_BOOL_4",
// 			defaultVal:  nil,
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_t with t value",
// 			envName:     "TEST_BOOL_5",
// 			defaultVal:  nil,
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_NON_EXISTENT with default F value",
// 			envName:     "TEST_BOOL_NON_EXISTENT",
// 			defaultVal:  "F",
// 			verifyFunc:  nil,
// 			expectedVal: false,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_BOOL_NON_EXISTENT with default True value",
// 			envName:     "TEST_BOOL_NON_EXISTENT",
// 			defaultVal:  "True",
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:   "test TEST_BOOL_NON_EXISTENT with default errorious value",
// 			envName:    "TEST_BOOL_NON_EXISTENT",
// 			defaultVal: "invalid",
// 			verifyFunc: func(val bool) error {
// 				if !val {
// 					return fmt.Errorf("%w", errTestOnly)
// 				}
// 				return nil
// 			},
// 			expectedVal: false,
// 			expectedErr: errTestOnly,
// 		},
// 		{
// 			testName:    "test non existing env-var with default bool value",
// 			envName:     "TEST_BOOL_NON_EXISTENT",
// 			defaultVal:  true,
// 			verifyFunc:  nil,
// 			expectedVal: true,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test non existing env-var with default int value",
// 			envName:     "TEST_BOOL_NON_EXISTENT",
// 			defaultVal:  999,
// 			verifyFunc:  nil,
// 			expectedVal: false,
// 			expectedErr: nil,
// 		},
// 	}
//
// 	for _, tc := range testcases {
// 		t.Run(tc.testName, func(t *testing.T) {
// 			val, err := getenv.Bool(tc.envName, tc.defaultVal, tc.verifyFunc)
// 			if !errors.Is(err, tc.expectedErr) {
// 				t.Errorf("want: %v, got: %v", tc.expectedErr, err)
// 			}
// 			if val != tc.expectedVal {
// 				t.Errorf("want: %t, got: %t", tc.expectedVal, val)
// 			}
// 		})
// 	}
// }

// func TestInt(t *testing.T) {
// 	os.Setenv("TEST_INT_1", "1")
// 	os.Setenv("TEST_INT_2", "invalid")
// 	os.Unsetenv("TEST_INT_NON_EXISTENT")
//
// 	defer func() {
// 		os.Unsetenv("TEST_INT_1")
// 	}()
//
// 	testcases := []struct {
// 		testName    string
// 		envName     string
// 		defaultVal  any
// 		verifyFunc  getenv.IntVerifyFunc
// 		expectedVal int
// 		expectedErr error
// 	}{
// 		{
// 			testName:    "test TEST_INT_1 with '1' value",
// 			envName:     "TEST_INT_1",
// 			defaultVal:  nil,
// 			verifyFunc:  nil,
// 			expectedVal: 1,
// 			expectedErr: nil,
// 		},
// 		{
// 			testName:    "test TEST_INT_2 with 'invalid' value",
// 			envName:     "TEST_INT_2",
// 			defaultVal:  nil,
// 			verifyFunc:  nil,
// 			expectedVal: 0,
// 			expectedErr: nil,
// 		},
// 	}
//
// 	for _, tc := range testcases {
// 		t.Run(tc.testName, func(t *testing.T) {
// 			val, err := getenv.Int(tc.envName, tc.defaultVal, tc.verifyFunc)
// 			if !errors.Is(err, tc.expectedErr) {
// 				t.Errorf("want: %v, got: %v", tc.expectedErr, err)
// 			}
// 			if val != tc.expectedVal {
// 				t.Errorf("want: %d, got: %d", tc.expectedVal, val)
// 			}
// 		})
// 	}
// }
