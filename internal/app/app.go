package app

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/handler"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/repositories"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/pkg/postgresql"
)

func Run(logger *logrus.Logger) {
	postgresConnection, err := postgresql.NewConnection(context.Background(), 3, &config.Cfg.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		postgresConnection.Release()
		logger.Info("Close Postgres connection")
	}()

	logger.Info("Open Postgres connection")

	noteRepo := repositories.NewPostgresNoteRepository(postgresConnection)

	bot, err := tgbotapi.NewBotAPI(config.Cfg.TelegramBot.Token)
	if err != nil {
		logger.Panic(err)
	}

	bot.Debug = false

	logger.Infof("Authorized on account https://t.me/%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)

	botHandler := handler.NewBotHandler(logger, bot, noteRepo)
	botHandler.HandleUpdates(updates)
}
