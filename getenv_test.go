package getenv_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/vigo/getenv"
)

func ExampleBool() {
	color := getenv.Bool("COLOR", false)
	if err := getenv.Parse(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(*color)
	// Outputs: false
}

func TestBool(t *testing.T) {
	os.Unsetenv("TEST_BOOL_NON_EXISTING_1")
	os.Unsetenv("TEST_BOOL_NON_EXISTING_2")

	os.Setenv("TEST_BOOL_1", "T")
	os.Setenv("TEST_BOOL_2", "true")
	os.Setenv("TEST_BOOL_3", "t")
	os.Setenv("TEST_BOOL_4", "1")
	os.Setenv("TEST_BOOL_5", "TRUE")
	os.Setenv("TEST_BOOL_6", "invalid")

	defer func() {
		os.Unsetenv("TEST_BOOL_1")
		os.Unsetenv("TEST_BOOL_2")
		os.Unsetenv("TEST_BOOL_3")
		os.Unsetenv("TEST_BOOL_4")
		os.Unsetenv("TEST_BOOL_5")
		os.Unsetenv("TEST_BOOL_6")
	}()

	tcs := []struct {
		testName      string
		envName       string
		defaultValue  bool
		exceptedValue bool
		expectedErr   error
	}{
		{
			testName:      "non existing env-var has default true should have true",
			envName:       "TEST_BOOL_NON_EXISTING_1",
			defaultValue:  true,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "non existing env-var has default false should have false",
			envName:       "TEST_BOOL_NON_EXISTING_2",
			defaultValue:  false,
			exceptedValue: false,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var 'T' has default false should have true",
			envName:       "TEST_BOOL_1",
			defaultValue:  false,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var 'true' has default false should have true",
			envName:       "TEST_BOOL_2",
			defaultValue:  false,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var 't' has default false should have true",
			envName:       "TEST_BOOL_3",
			defaultValue:  false,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var '1' has default false should have true",
			envName:       "TEST_BOOL_4",
			defaultValue:  false,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var 'TRUE' has default false should have true",
			envName:       "TEST_BOOL_5",
			defaultValue:  false,
			exceptedValue: true,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var 'invalid' has default false should have and error",
			envName:       "TEST_BOOL_6",
			defaultValue:  false,
			exceptedValue: false,
			expectedErr:   strconv.ErrSyntax,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			val := getenv.Bool(tc.envName, tc.defaultValue)
			err := getenv.Parse()
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want %v, got: %v", tc.expectedErr, err)
			}
			if err == nil {
				if *val != tc.exceptedValue {
					t.Errorf("want %t, got: %t", tc.exceptedValue, *val)
				}
			}
		})
	}
}

func TestInt(t *testing.T) {
	os.Unsetenv("TEST_INT_NON_EXISTING_1")

	os.Setenv("TEST_INT_1", "8000")
	os.Setenv("TEST_INT_2", "invalid")

	defer func() {
		os.Unsetenv("TEST_INT_1")
		os.Unsetenv("TEST_INT_2")
	}()

	tcs := []struct {
		testName      string
		envName       string
		defaultValue  int
		exceptedValue int
		expectedErr   error
	}{
		{
			testName:      "non existing env-var has default 999 should have 999",
			envName:       "TEST_INT_NON_EXISTING_1",
			defaultValue:  999,
			exceptedValue: 999,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has 8000 default 4000 should have 8000",
			envName:       "TEST_INT_1",
			defaultValue:  4000,
			exceptedValue: 8000,
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has 'invalid' default 4000 should have an error",
			envName:       "TEST_INT_2",
			defaultValue:  4000,
			exceptedValue: 0,
			expectedErr:   strconv.ErrSyntax,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			val := getenv.Int(tc.envName, tc.defaultValue)
			err := getenv.Parse()
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want %v, got: %v", tc.expectedErr, err)
			}
			if err == nil {
				if *val != tc.exceptedValue {
					t.Errorf("want %d, got: %d", tc.exceptedValue, *val)
				}
			}
		})
	}
}
