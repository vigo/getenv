package getenv

import (
	"errors"
	"fmt"
	"net"
)

type tcpAddrValue string

func newTCPAddrValue(val string, p *string) *tcpAddrValue {
	*p = val

	return (*tcpAddrValue)(p)
}

func (s *tcpAddrValue) Set(val string) error {
	if err := ValidateTCPNetworkAddress(val); err != nil {
		return fmt.Errorf("invalid tcp address %w", err)
	}
	*s = tcpAddrValue(val)

	return nil
}

func (s *tcpAddrValue) Get() any { return string(*s) }

// TCPAddr sets environment variable and returns the pointer of value.
func TCPAddr(name string, value string) *string {
	return environmentVariableSetInstance.TCPAddr(name, value)
}

// ValidateTCPNetworkAddress validates given tcp address as string and
// returns an error if the provided arg is not a valid tcp address.
func ValidateTCPNetworkAddress(addr string) error {
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		return errors.Join(ErrInvalid, fmt.Errorf("invalid tcp addr: %w", err))
	}

	return nil
}
