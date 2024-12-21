package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"
)

const (
	commandStart           = "start"
	replyStartTemplate     = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне доступ. Для этого переходи по ссылке:\n%s"
	replyAlreadyAuthorized = "Ты уже авторизирован"
)

func (b *Bot) handleCommand(message *tgbotapi.Message, logger *logging.Logger) error {
	logger.Info("handle command")
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message, logger)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know such a command")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message, logger *logging.Logger) error {
	logger.Infof("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message, logger *logging.Logger) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message, logger)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}
