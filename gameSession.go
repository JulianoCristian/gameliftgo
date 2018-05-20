package gameliftgo

type GameSessionStatus string

const (
	GameSessionStatusNotSet      GameSessionStatus = "NOT_SET"
	GameSessionStatusActive                        = "ACTIVE"
	GameSessionStatusActivating                    = "ACTIVATING"
	GameSessionStatusTerminated                    = "TERMINATED"
	GameSessionStatusTerminating                   = "TERMINATING"
)

type GameSession struct {
	GameSessionID             string
	Name                      string
	FleetID                   string
	MaximumPlayerSessionCount int
	Status                    GameSessionStatus
	GameProperties            map[string]string
	IPAddress                 string
	Port                      int
	GameSessionData           string
	MatchmakerData            string
	DNSName                   string
}
