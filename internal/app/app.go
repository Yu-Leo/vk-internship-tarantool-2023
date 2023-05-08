package app

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/pkg/postgresql"
)

func Run(logger *logrus.Logger) {

	postgresConnection, err := postgresql.NewConnection(context.Background(), 2, &config.Cfg.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		postgresConnection.Release()
		logger.Info("Close Postgres connection")
	}()

	logger.Info("Open Postgres connection")

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
