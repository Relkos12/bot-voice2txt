package telegram

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotAdpter struct {
	Bot *tgbot.BotAPI
}

// NewBotAdapter is constructor for BotAdapter
func NewBotAdapter(token string) (*BotAdpter, error) {
	bot, err := tgbot.NewBotAPI(token)

	if err != nil {
		return nil, err
	}
	return &BotAdpter{Bot: bot}, nil
}

// SendMessage implements the port MessageSender
func (b *BotAdpter) SendMessage(chatID int64, text string) error {
	msg := tgbot.NewMessage(chatID, text)

	if _, err := b.Bot.Send(msg); err != nil {
		return err
	}
	return nil

}

// SendFile implements the port FileSender
func (b *BotAdpter) SendFile(chatID int64, path string) error {
	doc := tgbot.NewDocument(chatID, tgbot.FilePath(path))

	if _, err := b.Bot.Send(doc); err != nil {
		return err
	}
	return nil

}
