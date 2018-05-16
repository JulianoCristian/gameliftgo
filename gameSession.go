package gameliftgo

type GameSessionStatus int

const (
	GameSessionStatusNotSet      GameSessionStatus = iota
	GameSessionStatusActive                        = iota
	GameSessionStatusActivating                    = iota
	GameSessionStatusTerminated                    = iota
	GameSessionStatusTerminating                   = iota
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
