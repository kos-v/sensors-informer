package channel

import (
	"fmt"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/security"
)

type TelegramBotChannel struct {
	Config           config.Config
	MessageFormatter message.Formatter
}

func (ch *TelegramBotChannel) IsEnable() bool {
	return ch.Config.Channels.TelegramBot.Enable
}

func (ch *TelegramBotChannel) Send(r report.Report) error {
	bot, err := botapi.NewBotAPI(ch.Config.Channels.TelegramBot.Token)
	if err != nil {
		return ch.hideErrorSecrets(err)
	}

	msg := botapi.NewMessage(ch.Config.Channels.TelegramBot.ChatId, ch.format(r))
	_, err = bot.Send(msg)

	return ch.hideErrorSecrets(err)
}

func (ch *TelegramBotChannel) format(r report.Report) string {
	msg := ""
	if t := ch.MessageFormatter.FormatTitle(&r); t != "" {
		msg += t + "\n"
	}
	for _, v := range ch.MessageFormatter.FormatBodyRows(&r) {
		msg += v + "\n"
	}
	return msg
}

func (ch *TelegramBotChannel) hideErrorSecrets(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(security.HideSecrets([]string{ch.Config.Channels.TelegramBot.Token}, err.Error()))
}
