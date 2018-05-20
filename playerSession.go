package gameliftgo

import (
	"time"
)

type PlayerSessionStatus string

const (
	PlayerSessionStatusNotSet    PlayerSessionStatus = "NOT_SET"
	PlayerSessionStatusReserved                      = "RESERVED"
	PlayerSessionStatusActive                        = "ACTIVE"
	PlayerSessionStatusCompleted                     = "COMPLETED"
	PlayerSessionStatusTimedout                      = "TIMEDOUT"
)

type PlayerSession struct {
	PlayerSessionID string
	GameSessionID   string
	FleetID         string
	CreationTime    time.Time
	TerminationTime time.Time
	Status          PlayerSessionStatus
	IPAddress       string
	Port            int
	PlayerData      string
	DNSName         string
}
