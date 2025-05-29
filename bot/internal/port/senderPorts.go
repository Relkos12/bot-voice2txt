package port

import "os"

//FileSender is port for sending file
type FileSender interface {
	SendFile(chatID int64, file *os.File) error
}

//MessageSender is port for sending message
type MessageSender interface {
	SendMessage(chatID int64, text string) error
}
