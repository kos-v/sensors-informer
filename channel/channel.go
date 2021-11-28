package channel

import (
	"github.com/kos-v/sensors-informer/config"
	"github.com/kos-v/sensors-informer/report"
)

type Channel interface {
	IsEnable() bool
	Send(r report.Report) error
}

func GetChannels(conf config.Config) []Channel {
	return []Channel{
		&TelegramBot{conf},
	}
}
