package gameliftgo

type GameSessionStatus int

const (
	StatusNotSet      GameSessionStatus = iota
	StatusActive                        = iota
	StatusActivating                    = iota
	StatusTerminated                    = iota
	StatusTerminating                   = iota
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
