package main

import (
	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()

	app.Run(logger)
}
