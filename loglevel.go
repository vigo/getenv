package getenv

import (
	"fmt"
	"strings"
)

type logLevelValue struct {
	val    *int
	levels map[string]int
}

func newLogLevelValue(levels map[string]int, def int, p *int) *logLevelValue {
	*p = def

	return &logLevelValue{val: p, levels: levels}
}

func (l *logLevelValue) Set(s string) error {
	key := strings.ToUpper(strings.TrimSpace(s))
	if v, ok := l.levels[key]; ok {
		*l.val = v

		return nil
	}

	return fmt.Errorf("[%w] unknown log level %q", ErrInvalid, s)
}

func (l *logLevelValue) Get() any { return *l.val }

// LogLevel sets environment variable and returns the pointer of value.
func LogLevel(name string, levels map[string]int, defaultValue int) *int {
	return environmentVariableSetInstance.LogLevel(name, levels, defaultValue)
}
