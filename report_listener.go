package main

import (
	"github.com/kos-v/sensors-informer/channel"
	"github.com/kos-v/sensors-informer/config"
	r "github.com/kos-v/sensors-informer/report"
	"log"
	"time"
)

type ReportListener struct {
	Ch       chan r.Report
	Channels []channel.Channel
	Config   config.Config

	lastSendTime int64
}

func (l *ReportListener) Listen() {
	go func() {
		for report := range l.Ch {
			l.handle(report)
		}
	}()
}

func (l *ReportListener) handle(report r.Report) {
	send := false
	for _, ch := range l.Channels {
		if ch.IsEnable() && l.canSendByTimeout() {
			err := ch.Send(report)
			if err != nil {
				log.Printf("Error: %s\n", err.Error())
				continue
			}
			send = true
		}
	}

	if send {
		l.lastSendTime = time.Now().Unix()
	}
}

func (l *ReportListener) canSendByTimeout() bool {
	return time.Now().Unix() >= l.lastSendTime+int64(l.Config.Report.RepeatTimeout)
}
