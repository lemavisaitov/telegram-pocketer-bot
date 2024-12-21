package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lemavisaitov/telegram-pocketer-bot/pkg/logging"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message, logger *logging.Logger) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know such a command")
	switch message.Command() {
	case commandStart:
		msg.Text = "You entered the command /start"
		msg.ReplyToMessageID = message.MessageID
		_, err := b.bot.Send(msg)
		return err
	default:
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
