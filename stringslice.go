package getenv

import "strings"

type stringSliceValue []string

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	*p = val

	return (*stringSliceValue)(p)
}

func (s *stringSliceValue) Set(val string) error {
	parts := strings.Split(val, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	*s = result

	return nil
}

func (s *stringSliceValue) Get() any { return []string(*s) }

// StringSlice sets environment variable and returns the pointer of value.
func StringSlice(name string, value []string) *[]string {
	return environmentVariableSetInstance.StringSlice(name, value)
}
