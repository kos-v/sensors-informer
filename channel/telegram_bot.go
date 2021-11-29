package channel

import (
	"fmt"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kos-v/sensors-informer/config"
	"github.com/kos-v/sensors-informer/report"
	"strings"
	"unicode/utf8"
)

type TelegramBotChannel struct {
	Config config.Config
}

func (ch *TelegramBotChannel) IsEnable() bool {
	return ch.Config.Channels.TelegramBot.Enable
}

func (ch *TelegramBotChannel) Send(r report.Report) error {
	bot, err := botapi.NewBotAPI(ch.Config.Channels.TelegramBot.Token)
	if err != nil {
		return ch.hideSecrets(err)
	}

	msg := botapi.NewMessage(ch.Config.Channels.TelegramBot.ChatId, ch.format(r))
	_, err = bot.Send(msg)

	return ch.hideSecrets(err)
}

func (ch *TelegramBotChannel) format(r report.Report) string {
	msg := fmt.Sprintf("Critical temperature readings:\n")
	for _, v := range r.Sensors {
		msg += fmt.Sprintf("\"%s::%s\": %.1f°С\n", v.BusName, v.SensorName, v.SensorValue)
	}
	return msg
}

func (ch *TelegramBotChannel) hideSecrets(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(strings.Replace(
		err.Error(),
		ch.Config.Channels.TelegramBot.Token,
		strings.Repeat("*", utf8.RuneCountInString(ch.Config.Channels.TelegramBot.Token)),
		-1,
	))
}
