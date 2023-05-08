package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/repositories"
)

type commandHandler = func(msg *tgbotapi.Message) (string, error)

type Handler struct {
	logger          *logrus.Logger
	bot             *tgbotapi.BotAPI
	noteRepository  repositories.NoteRepository
	commandHandlers map[string]commandHandler
}

func NewBotHandler(logger *logrus.Logger, bot *tgbotapi.BotAPI, repo repositories.NoteRepository) *Handler {
	h := &Handler{logger: logger, bot: bot, noteRepository: repo}
	h.commandHandlers = map[string]commandHandler{
		"set": h.Set,
		"get": h.Get,
		"del": h.Del,
	}
	return h
}

func (h *Handler) HandleUpdates(updates <-chan tgbotapi.Update) {
	for update := range updates {
		msg := update.Message
		if msg == nil || msg.Chat == nil {
			continue
		}
		h.handleMessage(msg)
	}
}

func (h *Handler) handleMessage(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "")
	reply.ReplyToMessageID = msg.MessageID

	handlerFunc, found := h.commandHandlers[msg.Command()]

	if !found {
		reply.Text = helpMessage
		h.send(reply)
		return
	}

	go func(msg *tgbotapi.Message) {
		text, err := handlerFunc(msg)

		if err != nil {
			text = internalErrorMessage
			h.logger.Error(err)
		}

		reply.Text = text
		reply.ParseMode = "MarkDown"
		h.send(reply)
	}(msg)
}

func (h *Handler) send(msg tgbotapi.MessageConfig) {
	_, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
