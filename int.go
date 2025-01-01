package getenv

import (
	"errors"
	"fmt"
	"strconv"
)

type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val

	return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return errors.Join(ErrInvalid, fmt.Errorf("parse int error: %w", err))
	}
	*i = intValue(v)

	return nil
}

func (i *intValue) Get() any { return int(*i) }

// Int sets environment variable and returns the pointer of value.
func Int(name string, value int) *int {
	return environmentVariableSetInstance.Int(name, value)
}
