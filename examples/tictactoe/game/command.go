package game

type CommandType string

const (
	CommandTypeAddToGame      = "addToGame"
	CommandTypeRemoveFromGame = "removeFromGame"
)

type Command struct {
	SenderID    string      `json:"sender"`
	CommandType CommandType `json:"type"`
}
