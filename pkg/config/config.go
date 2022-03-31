package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBPath            string `mapstructure:"db_file"`
	Messages          Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	About             string `mapstructure:"about"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SaveSuccessfully  string `mapstructure:"saved_successfully"`
	UnknounnCommand   string `mapstructure:"unknown_command"`
}

// Парсим файл yml
func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.AddConfigPath("/cmd/configs")
	viper.AddConfigPath(".")
	viper.SetConfigName("main")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &c.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &c.Messages.Errors); err != nil {
		return nil, err
	}

	if err := ParseEnv(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Парсим переменные окружения
func ParseEnv(c *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url", "AUTH_SERVER_URL"); err != nil {
		return err
	}

	c.TelegramToken = viper.GetString("token")
	c.PocketConsumerKey = viper.GetString("consumer_key")
	c.AuthServerURL = viper.GetString("auth_server_url")
	log.Println("c.TelegramBotURL :: ", c.TelegramBotURL)

	return nil
}
