package handler

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/repositories"
)

const msgDeleteDelay = 10 * time.Second

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
		msg := update.Message
		if msg == nil || msg.Chat == nil {
			continue
		}
		reply := tgbotapi.NewMessage(msg.Chat.ID, "")
		reply.ReplyToMessageID = msg.MessageID

		if !h.validateMessage(reply, msg) {
			continue
		}

		go h.handleMessage(reply, msg)
	}
}

func (h *Handler) handleMessage(reply tgbotapi.MessageConfig, msg *tgbotapi.Message) {
	var isMessageDeleted bool
	var text string
	var err error

	switch msg.Command() {
	case "start":
		text = startMessage
	case "set":
		text, isMessageDeleted, err = h.set(msg)
		reply.ParseMode = "markdown"
	case "get":
		text, isMessageDeleted, err = h.get(msg)
		reply.ParseMode = "markdown"
	case "del":
		text, err = h.del(msg)
		reply.ParseMode = "markdown"
	default:
		text = helpMessage
	}

	if err != nil {
		text = internalErrorMessage
		h.logger.Error(err)
	}

	reply.Text = text
	err = h.sendMessage(reply, isMessageDeleted)
	if err != nil {
		h.logger.Error(err)
	}
}

func (h *Handler) validateMessage(reply tgbotapi.MessageConfig, msg *tgbotapi.Message) bool {
	symbol := findInvalidSymbol(msg.Text[1:])

	if symbol == "" {
		return true
	}

	reply.Text = fmt.Sprintf(invalidSymbolMessage, symbol)

	err := h.sendMessage(reply, false)
	if err != nil {
		h.logger.Error(err)
	}

	return false
}

func findInvalidSymbol(str string) string {
	const availableSymbols = "abcdefghijklmnopqrstuvwxyz" +
		"абвгдеёжзийклмнопрстуфхцчшщъыьэюя" +
		"!#$%&()*+,-./0123456789:;<=>?@[]^_{|}~; "

	for _, symbol := range str {
		if !strings.Contains(availableSymbols, strings.ToLower(string(symbol))) {
			return string(symbol)
		}
	}
	return ""
}

func (h *Handler) sendMessage(msg tgbotapi.MessageConfig, delete bool) error {
	botMsg, err := h.bot.Send(msg)
	if err != nil {
		return err
	}

	if !delete {
		return nil
	}

	time.Sleep(msgDeleteDelay)

	delMsg := tgbotapi.NewDeleteMessage(msg.ChatID, msg.ReplyToMessageID)
	_, err = h.bot.Request(delMsg)
	if err != nil {
		return err
	}

	_, err = h.bot.Request(tgbotapi.NewDeleteMessage(msg.ChatID, botMsg.MessageID))
	return err
}
