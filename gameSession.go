package gameliftgo

type GameSessionStatus int

const (
	NotSet      GameSessionStatus = iota
	Active      GameSessionStatus = iota
	Activating  GameSessionStatus = iota
	Terminated  GameSessionStatus = iota
	Terminating GameSessionStatus = iota
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
