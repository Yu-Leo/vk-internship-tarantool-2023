package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/app"
)

const configPath = "./dev.env"

func main() {
	err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	app.Run(logger)
}
