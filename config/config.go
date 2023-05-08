package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	TelegramBot TelegramBotConfig
	Database    DatabaseConfig
}

type TelegramBotConfig struct {
	Token string `mapstructure:"TG_BOT_TOKEN"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int    `mapstructure:"POSTGRES_PORT"`
	Database string `mapstructure:"POSTGRES_DB"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
}

func LoadConfig(path string) (err error) {
	var tgBotConfig TelegramBotConfig
	var dbConfig DatabaseConfig

	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&tgBotConfig)
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		return err
	}

	Cfg = Config{TelegramBot: tgBotConfig, Database: dbConfig}
	return nil
}

func (dc *DatabaseConfig) GetURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
}
