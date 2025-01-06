package getenv

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// sentinel errors.
var (
	ErrEnvironmentVariableNotFound = errors.New("not found")
	ErrEnvironmentVariableIsEmpty  = errors.New("is empty")
	ErrInvalid                     = errors.New("invalid")
)

var environmentVariableSetInstance = newEnvironmentVariableSet() //nolint:gochecknoglobals

// Value defines environment variable's value behaviours.
type Value interface {
	Set(s string) error
	Get() any
}

// compile time proofs.
var (
	_ Value = (*boolValue)(nil)
	_ Value = (*intValue)(nil)
	_ Value = (*stringValue)(nil)
	_ Value = (*int64Value)(nil)
	_ Value = (*float64Value)(nil)
	_ Value = (*durationValue)(nil)
	_ Value = (*tcpAddrValue)(nil)
)

// EnvironmentVariable represents environment variable.
type EnvironmentVariable struct {
	Value Value
	Name  string
}

// EnvironmentVariableSet mimics flag.FlagSet type.
type EnvironmentVariableSet struct {
	variables map[string]*EnvironmentVariable
}

// Var stores EnvironmentVariable type.
func (e *EnvironmentVariableSet) Var(value Value, name string) {
	envVar := &EnvironmentVariable{
		Name:  name,
		Value: value,
	}
	if e.variables == nil {
		e.variables = make(map[string]*EnvironmentVariable)
	}
	e.variables[name] = envVar
}

// Bool creates new bool.
func (e *EnvironmentVariableSet) Bool(name string, value bool) *bool {
	p := new(bool)
	e.BoolVar(p, name, value)

	return p
}

// Int creates new int.
func (e *EnvironmentVariableSet) Int(name string, value int) *int {
	p := new(int)
	e.IntVar(p, name, value)

	return p
}

// Int64 creates new int64.
func (e *EnvironmentVariableSet) Int64(name string, value int64) *int64 {
	p := new(int64)
	e.Int64Var(p, name, value)

	return p
}

// Float64 creates new float64.
func (e *EnvironmentVariableSet) Float64(name string, value float64) *float64 {
	p := new(float64)
	e.Float64Var(p, name, value)

	return p
}

// String creates new string.
func (e *EnvironmentVariableSet) String(name string, value string) *string {
	p := new(string)
	e.StringVar(p, name, value)

	return p
}

// Duration creates new duration.
func (e *EnvironmentVariableSet) Duration(name string, value time.Duration) *time.Duration {
	p := new(time.Duration)
	e.DurationVar(p, name, value)

	return p
}

// TCPAddr creates new tcp addr.
func (e *EnvironmentVariableSet) TCPAddr(name string, value string) *string {
	p := new(string)
	e.TCPAddrVar(p, name, value)

	return p
}

// BoolVar creates new bool variable.
func (e *EnvironmentVariableSet) BoolVar(p *bool, name string, value bool) {
	e.Var(newBoolValue(value, p), name)
}

// IntVar creates new int variable.
func (e *EnvironmentVariableSet) IntVar(p *int, name string, value int) {
	e.Var(newIntValue(value, p), name)
}

// Int64Var creates new int64 variable.
func (e *EnvironmentVariableSet) Int64Var(p *int64, name string, value int64) {
	e.Var(newInt64Value(value, p), name)
}

// Float64Var creates new float64 variable.
func (e *EnvironmentVariableSet) Float64Var(p *float64, name string, value float64) {
	e.Var(newFloat64Value(value, p), name)
}

// StringVar creates new string variable.
func (e *EnvironmentVariableSet) StringVar(p *string, name string, value string) {
	e.Var(newStringValue(value, p), name)
}

// DurationVar creates new duration variable.
func (e *EnvironmentVariableSet) DurationVar(p *time.Duration, name string, value time.Duration) {
	e.Var(newDurationValue(value, p), name)
}

// TCPAddrVar creates new string variable for tcp address value.
func (e *EnvironmentVariableSet) TCPAddrVar(p *string, name string, value string) {
	e.Var(newTCPAddrValue(value, p), name)
}

// Parse fetches environment variable, creates required Value, sets and stores.
func (e *EnvironmentVariableSet) Parse() error {
	for name, envVar := range e.variables {
		envValue := os.Getenv(name)

		// if environment variable is not empty.
		if envValue != "" {
			// set the environment variable's value.
			if err := envVar.Value.Set(envValue); err != nil {
				return fmt.Errorf("%q %w", name, err)
			}
		}

		// if the current value is empty?
		if envVar.Value.Get() == "" {
			return fmt.Errorf("%q %w", name, ErrEnvironmentVariableIsEmpty)
		}

		if v, ok := envVar.Value.(*tcpAddrValue); ok {
			if val, okay := v.Get().(string); okay {
				if _, err := ValidateTCPNetworkAddress(val); err != nil {
					return fmt.Errorf("%q [%w] %w", name, ErrInvalid, err)
				}
			}
		}
	}

	return nil
}

// Reset resets variables storage.
func (e *EnvironmentVariableSet) Reset() {
	if e.variables != nil {
		e.variables = make(map[string]*EnvironmentVariable)
	}
}

func newEnvironmentVariableSet() *EnvironmentVariableSet {
	return &EnvironmentVariableSet{}
}

// Parse handles environment variable set/assign operations.
func Parse() error {
	if err := environmentVariableSetInstance.Parse(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Reset resets/clears variables storage.
func Reset() {
	environmentVariableSetInstance.Reset()
}
