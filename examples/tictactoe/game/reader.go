package game

type CommandReader interface {
	Read(command Command)
}
