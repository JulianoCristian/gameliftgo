package game

type SendMode int

const (
	SendModeEveryone   = iota
	SendModeRecipients = iota
)

type Message struct {
	SendMode     SendMode
	RecipientIDs []string
}
