package channel

import (
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
)

type Channel interface {
	IsEnable() bool
	Send(r report.Report) error
}

func GetChannels(conf config.Config, formatter message.Formatter) []Channel {
	return []Channel{
		&TelegramBotChannel{Config: conf, MessageFormatter: formatter},
		&FileChannel{Config: conf, MessageFormatter: formatter},
	}
}
