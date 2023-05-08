package handler

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/models"
)

func (h *Handler) Set(msg *tgbotapi.Message) (string, error) {
	input := strings.Split(msg.Text, " ")
	if len(input) < 4 {
		return invalidSetInputMessage, nil
	}

	note := getNoteFromSet(input)
	err := h.noteRepository.Set(fmt.Sprint(msg.Chat.ID), note)

	if err != nil {
		return "", err
	}

	text := fmt.Sprintf(dataMessage, note.ServiceName, note.Login, note.Password)
	return text, nil
}

func (h *Handler) Get(msg *tgbotapi.Message) (string, error) {
	input := strings.Split(msg.Text, " ")
	if len(input) < 2 {
		return invalidGetInputMessage, nil
	}

	serviceName := getServiceNameFromGetAndDel(input)

	note, err := h.noteRepository.Get(fmt.Sprint(msg.Chat.ID), serviceName)

	if err != nil {
		text := fmt.Sprintf(serviceNotFoundMessage, serviceName)
		return text, nil
	}

	text := fmt.Sprintf(dataMessage, note.ServiceName, note.Login, note.Password)

	return text, nil
}

func (h *Handler) Del(msg *tgbotapi.Message) (string, error) {
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