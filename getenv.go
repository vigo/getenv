package getenv

import (
	"errors"
	"fmt"
	"os"
)

// sentinel errors.
var (
	ErrEnvironmentVariableNotFound = errors.New("not found")
	ErrEnvironmentVariableIsEmpty  = errors.New("is empty")
)

var environmentVariableSetInstance = newEnvironmentVariableSet() //nolint:gochecknoglobals

// Value defines environment variable's value behaviours.
type Value interface {
	Set(s string) error
	Get() any
}

var (
	_ Value = (*boolValue)(nil)
	_ Value = (*intValue)(nil)
	_ Value = (*stringValue)(nil)
	_ Value = (*int64Value)(nil)
	_ Value = (*float64Value)(nil)
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

// Parse fetches environment variable, creates required Value, sets and stores.
func (e *EnvironmentVariableSet) Parse() error {
	for name, envVar := range e.variables {
		envValue := os.Getenv(name)
		if envValue != "" {
			if err := envVar.Value.Set(envValue); err != nil {
				return fmt.Errorf("error setting %s %w", name, err)
			}
		}

		if v, ok := envVar.Value.Get().(string); ok && v == "" {
			return ErrEnvironmentVariableIsEmpty
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
		return fmt.Errorf("can not set/assign environment variables %w", err)
	}

	return nil
}

// Reset resets/clears variables storage.
func Reset() {
	environmentVariableSetInstance.Reset()
}
