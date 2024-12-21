package main

import (
	"fmt"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/config"
	"github.com/lemavisaitov/telegram-pocketer-bot/internal/telegram"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"

	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

var logger = logging.GetLogger()

func main() {
	allLogsFile, err := os.OpenFile("logs/all.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Fatal()
	}
	log.SetOutput(allLogsFile)
	logger.Infof("Starting Telegram Pocketer Bot")

	cfg := config.GetConfig()
	fmt.Printf("token: %s\n", cfg.Token)
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		logger.Fatal(err)
	}
	bot.Debug = true

	pocketClient, err := pocket.NewClient(cfg.ConsumerKey)
	if err != nil {
		logger.Fatal(err)
	}

	tg := telegram.New(bot, pocketClient, "http://localhost/")
	err = tg.Start(logger)
	if err != nil {
		logger.Fatal(err)
	}
}
