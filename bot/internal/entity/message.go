package entity

type Message struct {
	ChatID int64
	Text   string
}

// NewMessage returns a new message
func NewMessage(chatID int64, text string) Message {
	return Message{ChatID: chatID, Text: text}
}
