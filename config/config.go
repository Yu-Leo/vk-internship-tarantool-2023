package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	TelegramBot TelegramBotConfig `yaml:"telegramBot"`
	Database    DatabaseConfig    `yaml:"database"`
}

type TelegramBotConfig struct {
	Token string `yaml:"token"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
}

func LoadConfig(path string) error {
	var err error
	var config Config

	viper.SetConfigFile(path)

	err = viper.BindEnv("telegramBot.token", "TG_BOT_TOKEN")
	if err != nil {
		return err
	}
	err = viper.BindEnv("database.host", "DB_HOST")
	if err != nil {
		return err
	}
	err = viper.BindEnv("database.port", "DB_PORT")
	if err != nil {
		return err
	}
	err = viper.BindEnv("database.user", "DB_USER")
	if err != nil {
		return err
	}
	err = viper.BindEnv("database.password", "DB_PASSWORD")
	if err != nil {
		return err
	}
	err = viper.BindEnv("database.database", "DB_DATABASE")
	if err != nil {
		return err
	}
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	Cfg = config
	return nil
}

func (dc *DatabaseConfig) GetURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
}
