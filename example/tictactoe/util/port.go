package util

import (
	"fmt"
	"math"

	"github.com/phayes/freeport"
)

const maxTries = math.MaxUint16

func GetPort(min int, max int) (int, error) {
	return getPort(min, max, 1)
}

func getPort(min int, max int, tries int) (int, error) {
	if tries > maxTries {
		return 0, fmt.Errorf("Unable to find open port after %d tries", maxTries)
	}
	port, err := freeport.GetFreePort()
	if err != nil {
		return 0, err
	}
	if port < min || port > max {
		return getPort(min, max, tries+1)
	}
	return port, nil
}
