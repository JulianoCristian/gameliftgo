package util

import (
	"fmt"
	"math"
	"net"
)

const maxTries = math.MaxUint16

func GetPort(min int, max int) (int, error) {
	return getPort(min, max, 1)
}

func getPort(min int, max int, tries int) (int, error) {
	if tries > maxTries {
		return 0, fmt.Errorf("Unable to find open port after %d tries", maxTries)
	}
	port, err := getOpenPort()
	if err != nil {
		return 0, err
	}
	if port < min || port > max {
		return getPort(min, max, tries+1)
	}
	return port, nil
}

func getOpenPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
