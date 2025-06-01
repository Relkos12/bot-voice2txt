package handler

import (
	"log/slog"

	"github.com/Relkos12/bot-voice2txt/internal/entity"
	"github.com/Relkos12/bot-voice2txt/internal/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdates handle all updates from tg.
// The handler must receive a MessageService to send a message,
// a fileService to send a file and a logger to log errors
func HandleUpdates(bot tgbotapi.BotAPI, messageService *usecase.MessageService, fileService *usecase.FileService, logger *slog.Logger) {
	//set confing for getUpdatesChan
	u := setConfUpdates()

	//start polling
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		// start
		if update.Message.Text == "/start" {
			text := "Добро пожаловать, если вы хотите проверить свой английский, отправьте свое голосовое сообщение."
			msg := entity.NewMessage(update.Message.Chat.ID, text)

			err := messageService.SendMessage(msg)
			if err != nil {
				logger.Error("the message is not sent", "error", err)
			}
		}

	}
}

func setConfUpdates() tgbotapi.UpdateConfig {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 30

	return update
}
