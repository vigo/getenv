package getenv

import (
	"fmt"
	"os"
)

// StringVerifyFunc is a verify function type for String.
type StringVerifyFunc func(string) error

// String retrieves the value of the environment variable identified by the given key.
// If the environment variable is not set, it falls back to the provided default value.
// Optionally, a custom validation function (vf) of type StringVerifyFunc can be supplied
// to perform additional checks on the resulting value.
// If the validation fails, an error is returned.
func String(name, defaultValue string, vf StringVerifyFunc) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	if vf != nil {
		if err := vf(val); err != nil {
			return "", fmt.Errorf("verify error: %w", err)
		}
	}

	return val, nil
}

// // BoolVerifyFunc is a verify function type for Bool.
// type BoolVerifyFunc func(bool) error
//
// // Bool retrieves the value of the environment variable identified by the given
// // name as a boolean. If the environment variable is not set, it falls back to
// // the provided defaultValue. Optionally, a custom validation function (vf) of
// // type BoolVerifyFunc can be supplied to perform additional checks on the
// // resulting boolean value. If the validation fails, an error is returned.
// func Bool(name string, defaultValue any, vf BoolVerifyFunc) (bool, error) {
// 	val := os.Getenv(name)
// 	if val == "" {
// 		defVal := parseBool(defaultValue)
// 		if vf != nil {
// 			if err := vf(defVal); err != nil {
// 				return false, fmt.Errorf("verify error: %w", err)
// 			}
// 		}
//
// 		return defVal, nil
// 	}
//
// 	val = strings.ToLower(val)
//
// 	return isTrue(val), nil
// }
//
// func parseBool(val any) bool {
// 	switch v := val.(type) {
// 	case bool:
// 		return v
// 	case string:
// 		return isTrue(strings.ToLower(v))
// 	default:
// 		return false
// 	}
// }
//
// func isTrue(value string) bool {
// 	trueValues := []string{"1", "t", "true", "T", "True", "TRUE"}
//
// 	return contains(trueValues, value)
// }
//
// func contains(slice []string, value string) bool {
// 	for _, v := range slice {
// 		if v == value {
// 			return true
// 		}
// 	}
//
// 	return false
// }

// // IntVerifyFunc is a verify function type for Bool.
// type IntVerifyFunc func(int) error
//
// // Int ...
// func Int(name string, defaultValue any, vf IntVerifyFunc) (int, error) {
// 	val := os.Getenv(name)
// 	if val == "" {
// 		defVal := parseInt(defaultValue)
// 		if vf != nil {
// 			if err := vf(defVal); err != nil {
// 				return 0, fmt.Errorf("verify error: %w", err)
// 			}
// 		}
//
// 		return defVal, nil
// 	}
//
// 	i, err := strconv.Atoi(val)
// 	if err != nil {
// 		return 0, fmt.Errorf("convert error: %w", err)
// 	}
//
// 	return i, nil
// }
//
// func parseInt(val any) int {
// 	switch v := val.(type) {
// 	case int:
// 		return v
// 	case string:
// 		i, err := strconv.Atoi(v)
// 		if err != nil {
// 			return 0
// 		}
//
// 		return i
// 	default:
// 		return 0
// 	}
// }
