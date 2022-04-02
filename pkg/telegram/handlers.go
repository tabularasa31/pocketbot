package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	commandStart = "start"
	commandAbout = "about"
	commandHelp  = "help"
)

// Обработка обычных сообщений
func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.SaveSuccessfully)
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		URL:         message.Text,
		AccessToken: accessToken,
	}); err != nil {
		return errUnableToSave
	}

	_, err = b.bot.Send(msg)
	return err
}

// Обработка команд
func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAbout:
		return b.handleAboutCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

// Обработка команды Start
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	defer b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "УСПЕШНО"))
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProccess(message)
	}
	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.messages.AlreadyAuthorized))
	return err
}

// Обработка команды About
func (b *Bot) handleAboutCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.About)
	_, err := b.bot.Send(msg)
	return err
}

// Обработка команды Help
func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Help)
	_, err := b.bot.Send(msg)
	return err
}

// Неизвестная команда
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknounnCommand)
	_, err := b.bot.Send(msg)
	return err
}
