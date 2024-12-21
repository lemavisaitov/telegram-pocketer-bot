package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	redirectURL  string
}

func New(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectURL string) *Bot {
	return &Bot{
		bot:          bot,
		pocketClient: pocketClient,
		redirectURL:  redirectURL,
	}
}

func (b *Bot) Start(logger *logging.Logger) error {
	logger.Infof("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates, logger)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, logger *logging.Logger) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message, logger)
			logger.Error(err)
			continue
		}
		err := b.handleMessage(update.Message, logger)
		logger.Error(err)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	return updates
}
