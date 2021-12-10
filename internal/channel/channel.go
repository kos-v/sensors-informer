package channel

import (
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
)

type Channel interface {
	IsEnable() bool
	Send(r report.Report) error
}

func GetChannels(config conf.Config, formatter message.Formatter) []Channel {
	return []Channel{
		&TelegramBotChannel{Opts: config.Channels.TelegramBot, MessageFormatter: formatter},
		&FileChannel{Opts: config.Channels.File, MessageFormatter: formatter},
		&NotifySendChannel{Opts: config.Channels.NotifySend, MessageFormatter: formatter},
		&SmtpChannel{Opts: config.Channels.Smtp, MessageFormatter: formatter},
	}
}
