package usecase

import (
	"github.com/Relkos12/bot-voice2txt/internal/entity"
	"github.com/Relkos12/bot-voice2txt/internal/port"
)

type MessageService struct {
	sender port.MessageSender // dependence on port.MessageSender
}

// NewMessageService is constructor for MessageService
func NewMessageService(sender port.MessageSender) *MessageService {
	return &MessageService{sender: sender}
}

// SendMessage sending a message to user
func (s *MessageService) SendMessage(message entity.Message) error {
	if err := s.sender.SendMessage(message.ChatID, message.Text); err != nil {
		return err
	}
	return nil
}
