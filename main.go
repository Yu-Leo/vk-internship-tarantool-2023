package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
)

const configPath = "config/config.yaml"

func main() {
	err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()

	bot, err := tgbotapi.NewBotAPI(config.Cfg.TelegramBot.Token)
	if err != nil {
		logger.Panic(err)
	}

	bot.Debug = true

	logger.Info("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			logger.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
