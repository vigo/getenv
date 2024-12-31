package getenv

import (
	"fmt"
	"strconv"
)

type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val

	return (*int64Value)(p)
}

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return fmt.Errorf("parse int error: %w", err)
	}
	*i = int64Value(v)

	return nil
}

func (i *int64Value) Get() any { return int64(*i) }

// Int64 sets environment variable and returns the pointer of value.
func Int64(name string, value int64) *int64 {
	return environmentVariableSetInstance.Int64(name, value)
}
