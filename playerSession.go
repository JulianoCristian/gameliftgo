package gameliftgo

import (
	"time"
)

type PlayerSessionStatus int

const (
	PlayerSessionStatusNotSet    PlayerSessionStatus = iota
	PlayerSessionStatusReserved  PlayerSessionStatus = iota
	PlayerSessionStatusActive    PlayerSessionStatus = iota
	PlayerSessionStatusCompleted PlayerSessionStatus = iota
	PlayerSessionStatusTimedout  PlayerSessionStatus = iota
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
