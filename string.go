package getenv

type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val

	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)

	return nil
}

func (s *stringValue) Get() any { return string(*s) }

// String sets environment variable and returns the pointer of value.
func String(name string, value string) *string {
	return environmentVariableSetInstance.String(name, value)
}
