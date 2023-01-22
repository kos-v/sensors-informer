package channel

import (
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/contracts"
	"github.com/kos-v/sensors-informer/internal/message"
)

type Factory struct {
	channelListOpts config.ChannelListOpts
	msgFormatter    message.Formatter
}

func (f *Factory) CreateAvailableList() []contracts.Channel {
	var channels []contracts.Channel
	if f.channelListOpts.File.Enable {
		channels = append(channels, NewFileChannelService(f.channelListOpts.File, f.msgFormatter))
	}
	if f.channelListOpts.NotifySend.Enable {
		channels = append(channels, NewNotifySendChannelService(f.channelListOpts.NotifySend, f.msgFormatter))
	}
	if f.channelListOpts.Smtp.Enable {
		channels = append(channels, NewSmtpChannelService(f.channelListOpts.Smtp, f.msgFormatter))
	}
	if f.channelListOpts.TelegramBot.Enable {
		channels = append(channels, NewTelegramBotChannelService(f.channelListOpts.TelegramBot, f.msgFormatter))
	}

	return channels
}

func NewFactory(channelListOpts config.ChannelListOpts, msgFormatter message.Formatter) *Factory {
	return &Factory{channelListOpts: channelListOpts, msgFormatter: msgFormatter}
}
