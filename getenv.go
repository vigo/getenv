package getenv

import (
	"errors"
	"fmt"
	"os"
)

// sentinel errors.
var (
	ErrEnvironmentVariableNotFound = errors.New("not found")
)

var environmentVariableSetInstance = newEnvironmentVariableSet() //nolint:gochecknoglobals

// Value defines environment variable's value behaviours.
type Value interface {
	Set(s string) error
}

var (
	_ Value = (*boolValue)(nil)
	_ Value = (*intValue)(nil)
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

// BoolVar creates new bool variable.
func (e *EnvironmentVariableSet) BoolVar(p *bool, name string, value bool) {
	e.Var(newBoolValue(value, p), name)
}

// Int creates new int.
func (e *EnvironmentVariableSet) Int(name string, value int) *int {
	p := new(int)
	e.IntVar(p, name, value)

	return p
}

// IntVar creates new int variable.
func (e *EnvironmentVariableSet) IntVar(p *int, name string, value int) {
	e.Var(newIntValue(value, p), name)
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
	}

	return nil
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

// // Set sets the variable's vallue.
// func (e *EnvironmentVariableSet) Set(name, value string) error {
// 	envVar, ok := e.values[name]
// 	if !ok {
// 		return fmt.Errorf("%s %w", name, ErrEnvironmentVariableNotFound)
// 	}
//
// 	if err := envVar.Value.Set(value); err != nil {
// 		return fmt.Errorf("can not set the value %w", err)
// 	}
//
// 	if e.values == nil {
// 		e.values = make(map[string]*EnvironmentVariable)
// 	}
// 	e.values[name] = envVar
//
// 	return nil
// }
