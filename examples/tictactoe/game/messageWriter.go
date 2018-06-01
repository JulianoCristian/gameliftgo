package game

type MessageWriter interface {
	Write(message Message, playerIDs ...string)
	WriteAll(message Message)
}
