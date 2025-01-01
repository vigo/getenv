package getenv

import (
	"errors"
	"fmt"
	"strconv"
)

type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val

	return (*float64Value)(p)
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return errors.Join(ErrInvalid, fmt.Errorf("parse float error: %w", err))
	}
	*f = float64Value(v)

	return nil
}

func (f *float64Value) Get() any { return float64(*f) }

// Float64 sets environment variable and returns the pointer of value.
func Float64(name string, value float64) *float64 {
	return environmentVariableSetInstance.Float64(name, value)
}
