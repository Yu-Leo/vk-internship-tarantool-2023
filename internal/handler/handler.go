package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/repositories"
)

type Handler struct {
	logger         *logrus.Logger
	bot            *tgbotapi.BotAPI
	noteRepository repositories.NoteRepository
}

func NewBotHandler(logger *logrus.Logger, bot *tgbotapi.BotAPI, repo repositories.NoteRepository) *Handler {
	h := &Handler{logger: logger, bot: bot, noteRepository: repo}
	return h
}

func (h *Handler) HandleUpdates(updates <-chan tgbotapi.Update) {
	for update := range updates {
		if update.Message != nil {
			h.logger.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			h.bot.Send(msg)
		}
	}
}