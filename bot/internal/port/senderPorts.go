package port

//FileSender is port for sending file
type FileSender interface {
	SendFile(chatID int64, path string) error
}

//MessageSender is port for sending message
type MessageSender interface {
	SendMessage(chatID int64, text string) error
}
