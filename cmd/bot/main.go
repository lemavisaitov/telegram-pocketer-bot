package main

import (
	"github.com/boltdb/bolt"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/config"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/repository"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/repository/boltdb"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/server"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/telegram"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"

	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

var logger = logging.GetLogger()

func main() {
	var err error
	allLogsFile, err := os.OpenFile("logs/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Fatal()
	}
	log.SetOutput(allLogsFile)
	logger.Infof("Starting Telegram Pocketer Bot")

	cfg := config.GetConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		logger.Fatal(err)
	}
	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.ConsumerKey)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := initDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	tokenRepository := boltdb.NewTokenRepository(db)

	tg := telegram.New(bot, pocketClient, tokenRepository, "https://t.me/mygetpocket_bot")

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/mygetpocket_bot")

	go func() {
		if err = tg.Start(logger); err != nil {
			logger.Fatal(err)
		}
	}()
	if err = authorizationServer.Start(); err != nil {
		logger.Fatal(err)
	}
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)

	if err != nil {
		return nil, err
	}
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
