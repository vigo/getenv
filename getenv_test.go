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

func ExampleInt() {
	port := getenv.Int("PORT", 8000)
	if err := getenv.Parse(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(*port)
	// Outputs: 8000
}

func ExampleInt64() {
	long := getenv.Int64("LONG", 9223372036854775806)
	if err := getenv.Parse(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(*long)
	// Outputs: 9223372036854775806
}

func ExampleString() {
	hmacHeader := getenv.String("HMAC_HEADER", "X-Foo-Signature")
	if err := getenv.Parse(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(*hmacHeader)
	// Outputs: X-Foo-Signature
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
			getenv.Reset()
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
			getenv.Reset()
		})
	}
}

func TestInt64(t *testing.T) {
	os.Unsetenv("TEST_INT64_NON_EXISTING_1")

	os.Setenv("TEST_INT64_1", "-123456789012")
	os.Setenv("TEST_INT64_2", "abc")
	os.Setenv("TEST_INT64_3", "9223372036854775808123")

	defer func() {
		os.Unsetenv("TEST_INT64_1")
		os.Unsetenv("TEST_INT64_2")
		os.Unsetenv("TEST_INT64_3")
	}()

	tcs := []struct {
		testName      string
		envName       string
		defaultValue  int64
		exceptedValue int64
		expectedErr   error
	}{
		{
			testName:      "non existing env-var has default '1' should have '1'",
			envName:       "TEST_INT64_NON_EXISTING_1",
			defaultValue:  int64(1),
			exceptedValue: int64(1),
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has '-123456789012' default '1' should have '-123456789012'",
			envName:       "TEST_INT64_1",
			defaultValue:  int64(1),
			exceptedValue: int64(-123456789012),
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has 'abc' default '0' should have an error",
			envName:       "TEST_INT64_2",
			defaultValue:  0,
			exceptedValue: 0,
			expectedErr:   strconv.ErrSyntax,
		},
		{
			testName:      "existing env-var has '9223372036854775808123' default '0' should have an error",
			envName:       "TEST_INT64_3",
			defaultValue:  0,
			exceptedValue: 0,
			expectedErr:   strconv.ErrRange,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			val := getenv.Int64(tc.envName, tc.defaultValue)
			err := getenv.Parse()

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want %v, got: %v", tc.expectedErr, err)
			}
			if err == nil {
				if *val != tc.exceptedValue {
					t.Errorf("want %d, got: %d", tc.exceptedValue, *val)
				}
			}
			getenv.Reset()
		})
	}
}

func TestFloat64(t *testing.T) {
	os.Unsetenv("TEST_FLOAT64_NON_EXISTING_1")

	os.Setenv("TEST_FLOAT64_1", "-1")
	os.Setenv("TEST_FLOAT64_2", "abc")
	os.Setenv("TEST_FLOAT64_3", "1e500")

	defer func() {
		os.Unsetenv("TEST_INT64_1")
		os.Unsetenv("TEST_INT64_2")
	}()

	tcs := []struct {
		testName      string
		envName       string
		defaultValue  float64
		exceptedValue float64
		expectedErr   error
	}{
		{
			testName:      "non existing env-var has default '3.14' should have '3.14'",
			envName:       "TEST_FLOAT64_NON_EXISTING_1",
			defaultValue:  float64(3.14),
			exceptedValue: float64(3.14),
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has '-1' value should have '-1.0'",
			envName:       "TEST_FLOAT64_1",
			defaultValue:  0,
			exceptedValue: float64(-1.0),
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has invalid value should have an error",
			envName:       "TEST_FLOAT64_2",
			defaultValue:  0,
			exceptedValue: 0,
			expectedErr:   strconv.ErrSyntax,
		},
		{
			testName:      "existing env-var has invalid value (range) should have an error",
			envName:       "TEST_FLOAT64_3",
			defaultValue:  0,
			exceptedValue: 0,
			expectedErr:   strconv.ErrRange,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			val := getenv.Float64(tc.envName, tc.defaultValue)
			err := getenv.Parse()

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want %v, got: %v", tc.expectedErr, err)
			}
			if err == nil {
				if *val != tc.exceptedValue {
					t.Errorf("want %f, got: %f", tc.exceptedValue, *val)
				}
			}
			getenv.Reset()
		})
	}
}

func TestString(t *testing.T) {
	os.Unsetenv("TEST_STRING_NON_EXISTING_1")

	os.Setenv("TEST_STRING_1", "application/json")
	os.Setenv("TEST_STRING_2", "")

	defer func() {
		os.Unsetenv("TEST_STRING_1")
		os.Unsetenv("TEST_STRING_2")
	}()

	tcs := []struct {
		testName      string
		envName       string
		defaultValue  string
		exceptedValue string
		expectedErr   error
	}{
		{
			testName:      "non existing env-var has default 'X-Foo' should have 'X-Foo'",
			envName:       "TEST_STRING_NON_EXISTING_1",
			defaultValue:  "X-Foo",
			exceptedValue: "X-Foo",
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has 'application/json' default 'text/plain' should have 'application/json'",
			envName:       "TEST_STRING_1",
			defaultValue:  "text/plain",
			exceptedValue: "application/json",
			expectedErr:   nil,
		},
		{
			testName:      "existing env-var has empty string as default should have an error",
			envName:       "TEST_STRING_2",
			defaultValue:  "",
			exceptedValue: "",
			expectedErr:   getenv.ErrEnvironmentVariableIsEmpty,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			val := getenv.String(tc.envName, tc.defaultValue)
			err := getenv.Parse()
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("want %v, got: %v", tc.expectedErr, err)
			}
			if err == nil {
				if *val != tc.exceptedValue {
					t.Errorf("want %s, got: %s", tc.exceptedValue, *val)
				}
			}
			getenv.Reset()
		})
	}
}
