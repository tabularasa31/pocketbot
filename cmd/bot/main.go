package main

import (
	"PocketBot/pkg/config"
	"PocketBot/pkg/repository"
	"PocketBot/pkg/repository/boltdb"
	"PocketBot/pkg/server"
	"PocketBot/pkg/telegram"
	"github.com/boltdb/bolt"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("КОНФИГ:  ", cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	//bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, cfg.AuthServerURL, cfg.Messages)

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, cfg.TelegramBotURL)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}

}

// Инициализация базы данных
func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
