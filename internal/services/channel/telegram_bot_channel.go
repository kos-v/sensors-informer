package channel

import (
	"fmt"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/security"
	"strings"
)

type TelegramBotChannelService struct {
	opts             config.TelegramBotChannelOpts
	messageFormatter message.Formatter
}

func (ch *TelegramBotChannelService) GetName() string {
	return "telegramBot"
}

func (ch *TelegramBotChannelService) IsEnable() bool {
	return ch.opts.Enable
}

func (ch *TelegramBotChannelService) Send(r report.Report) error {
	bot, err := botapi.NewBotAPI(ch.opts.Token)
	if err != nil {
		return ch.hideErrorSecrets(err)
	}

	msg := botapi.NewMessage(ch.opts.ChatId, ch.format(r))
	_, err = bot.Send(msg)

	return ch.hideErrorSecrets(err)
}

func (ch *TelegramBotChannelService) format(r report.Report) string {
	msg := ""
	if t := ch.messageFormatter.FormatTitle(&r); t != "" {
		msg += t + "\n"
	}

	msg += strings.Join(ch.messageFormatter.FormatBodyRows(&r), "\n")

	return msg
}

func (ch *TelegramBotChannelService) hideErrorSecrets(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(security.HideSecrets([]string{ch.opts.Token}, err.Error()))
}

func NewTelegramBotChannelService(
	opts config.TelegramBotChannelOpts,
	msgFormatter message.Formatter,
) *TelegramBotChannelService {
	return &TelegramBotChannelService{opts: opts, messageFormatter: msgFormatter}
}
