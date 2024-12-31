package getenv

import (
	"fmt"
	"strconv"
)

type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val

	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return fmt.Errorf("bool parse error: %w", err)
	}
	*b = boolValue(v)

	return nil
}

func (b *boolValue) Get() any { return bool(*b) }

// Bool sets environment variable and returns the pointer of value.
func Bool(name string, value bool) *bool {
	return environmentVariableSetInstance.Bool(name, value)
}
