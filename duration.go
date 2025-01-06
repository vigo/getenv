package getenv

import (
	"fmt"
	"time"
)

type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val

	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return fmt.Errorf("[%w] %w", ErrInvalid, err)
	}
	*d = durationValue(v)

	return nil
}

func (d *durationValue) Get() any { return time.Duration(*d) }

// Duration sets environment variable and returns the pointer of value.
func Duration(name string, value time.Duration) *time.Duration {
	return environmentVariableSetInstance.Duration(name, value)
}
