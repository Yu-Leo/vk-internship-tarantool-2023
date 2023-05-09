package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/config"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/app"
)

func getConfigPath() string {
	args := os.Args
	if len(args) == 1 {
		logrus.Fatal("There are no command line arg with path to config file")
	}
	return args[1]
}

func main() {
	err := config.LoadConfig(getConfigPath())
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()
	app.Run(logger)
}
