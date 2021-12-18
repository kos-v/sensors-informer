package internal

import (
	"github.com/kos-v/sensors-informer/internal/channel"
	r "github.com/kos-v/sensors-informer/internal/report"
	"log"
	"sync"
	"time"
)

func NewReportListener(repeatTimeout uint, channels []channel.Channel, rch chan r.Report) *ReportListener {
	return &ReportListener{
		Ch:            rch,
		Channels:      channels,
		RepeatTimeout: repeatTimeout,
		LastSendTime:  &lastSendTimeStorage{storage: make(map[string]int64)},
	}
}

type ReportListener struct {
	Ch            chan r.Report
	Channels      []channel.Channel
	RepeatTimeout uint
	LastSendTime  *lastSendTimeStorage
}

func (l *ReportListener) Listen() {
	go func() {
		for report := range l.Ch {
			l.handle(report)
		}
	}()
}

func (l *ReportListener) handle(report r.Report) {
	for _, item := range l.Channels {
		tch := item
		if !tch.IsEnable() {
			continue
		}

		currTime := time.Now().Unix()
		if !l.LastSendTime.Has(tch.GetName()) || currTime >= l.nexSendTime(tch.GetName()) {
			go func() {
				err := tch.Send(report)
				if err != nil {
					log.Printf("Error: %s\n", err.Error())
				}
				l.LastSendTime.Set(tch.GetName(), time.Now().Unix())
			}()
		}
	}
}

func (l *ReportListener) nexSendTime(channel string) int64 {
	return l.LastSendTime.Get(channel) + int64(l.RepeatTimeout)
}

type lastSendTimeStorage struct {
	mx      sync.Mutex
	storage map[string]int64
}

func (s *lastSendTimeStorage) Get(k string) int64 {
	s.mx.Lock()
	defer s.mx.Unlock()

	if !s.has(k) {
		return 0
	}
	return s.storage[k]
}

func (s *lastSendTimeStorage) Has(k string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.has(k)
}

func (s *lastSendTimeStorage) Set(k string, v int64) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.storage[k] = v
}

func (s *lastSendTimeStorage) has(k string) bool {
	_, ok := s.storage[k]
	return ok
}
