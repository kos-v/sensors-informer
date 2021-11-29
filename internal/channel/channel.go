package channel

import (
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/report"
)

type Channel interface {
	IsEnable() bool
	Send(r report.Report) error
}

func GetChannels(conf config.Config) []Channel {
	return []Channel{
		&TelegramBotChannel{conf},
	}
}
