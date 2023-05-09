package handler

import (
	"fmt"
	"strings"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/models"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/pkg/encoding"
)

const maxStringLen = 80

func (h *Handler) set(msg *tgbotapi.Message) (string, bool, error) {
	input := strings.Split(msg.Text, " ")
	if len(input) < 4 {
		return invalidSetInputMessage, false, nil
	}

	note := getNoteFromSet(input)

	tooLong := utf8.RuneCountInString(note.ServiceName) > maxStringLen ||
		utf8.RuneCountInString(note.Login) > maxStringLen || utf8.RuneCountInString(note.Password) > maxStringLen
	if tooLong {
		return fmt.Sprintf(tooLongMessage, maxStringLen), false, nil
	}

	rawPassword := note.Password
	note.Password = encoding.Encode(note.Password)

	err := h.noteRepository.Set(fmt.Sprint(msg.Chat.ID), note)

	if err != nil {
		return "", false, err
	}

	text := fmt.Sprintf(dataMessage, note.ServiceName, note.Login, rawPassword)
	return text, true, nil
}

func (h *Handler) get(msg *tgbotapi.Message) (string, bool, error) {
	input := strings.Split(msg.Text, " ")
	if len(input) < 2 {
		return invalidGetInputMessage, false, nil
	}

	serviceName := getServiceNameFromGetAndDel(input)

	note, err := h.noteRepository.Get(fmt.Sprint(msg.Chat.ID), serviceName)
	if err != nil {
		text := fmt.Sprintf(serviceNotFoundMessage, serviceName)
		return text, false, nil
	}

	note.Password, err = encoding.Decode(note.Password)
	if err != nil {
		text := fmt.Sprintf(serviceNotFoundMessage, serviceName)
		return text, false, nil
	}

	text := fmt.Sprintf(dataMessage, note.ServiceName, note.Login, note.Password)

	return text, true, nil
}

func (h *Handler) del(msg *tgbotapi.Message) (string, error) {
	input := strings.Split(msg.Text, " ")
	if len(input) < 2 {
		return invalidDelInputMessage, nil
	}

	serviceName := getServiceNameFromGetAndDel(input)

	err := h.noteRepository.Del(fmt.Sprint(msg.Chat.ID), serviceName)
	text := fmt.Sprintf(serviceDeletedMessage, serviceName)

	if err != nil {
		text = fmt.Sprintf(serviceNotFoundMessage, serviceName)
	}

	return text, nil
}

func getNoteFromSet(input []string) *models.Note {
	serviceName := input[1]

	for i := 2; i < len(input)-2; i++ {
		serviceName += " " + input[i]
	}

	note := &models.Note{
		ServiceName: serviceName,
		Login:       input[len(input)-2],
		Password:    input[len(input)-1],
	}
	return note
}

func getServiceNameFromGetAndDel(input []string) string {
	serviceName := input[1]

	for i := 2; i < len(input); i++ {
		serviceName += " " + input[i]
	}
	return serviceName
}
