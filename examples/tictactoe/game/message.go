package game

type MessageType string

const (
	MessageTypePlayerAddedToGame     = "playerAddedToGame"
	MessageTypePlayerRemovedFromGame = "playerRemovedFromGame"
	MessageTypeGameStarted           = "gameStarted"
	MessageTypeGameEnded             = "gameEnded"
)

type JSON map[string]interface{}

type Message struct {
	MessageType MessageType `json:"type"`
	MessageData JSON        `json:"data"`
}

func newEmptyMessage(messageType MessageType) Message {
	return Message{
		MessageType: messageType,
		MessageData: JSON{},
	}
}

func newMessage(messageType MessageType, messageData JSON) Message {
	return Message{
		MessageType: messageType,
		MessageData: messageData,
	}
}
