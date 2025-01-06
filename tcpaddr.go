package getenv

import (
	"fmt"
	"net"
)

type tcpAddrValue string

func newTCPAddrValue(val string, p *string) *tcpAddrValue {
	*p = val

	return (*tcpAddrValue)(p)
}

func (s *tcpAddrValue) Set(val string) error {
	if val == "" {
		return fmt.Errorf("[%w]", ErrEnvironmentVariableIsEmpty)
	}

	tcpAddr, err := ValidateTCPNetworkAddress(val)
	if err != nil {
		return fmt.Errorf("[%w] %w", ErrInvalid, err)
	}

	*s = tcpAddrValue(tcpAddr.String())

	return nil
}

func (s *tcpAddrValue) Get() any { return string(*s) }

// TCPAddr sets environment variable and returns the pointer of value.
func TCPAddr(name string, value string) *string {
	return environmentVariableSetInstance.TCPAddr(name, value)
}

// ValidateTCPNetworkAddress validates given tcp address as string and
// returns an error if the provided arg is not a valid tcp address.
func ValidateTCPNetworkAddress(addr string) (*net.TCPAddr, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return tcpAddr, nil
}
