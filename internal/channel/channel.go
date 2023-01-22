package channel

import (
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
)

type Channel interface {
	GetName() string
	IsEnable() bool
	Send(r report.Report) error
}

func GetChannels(config conf.Config, formatter message.Formatter) []Channel {
	return []Channel{
	}
}
