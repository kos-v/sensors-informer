package main

import (
	"github.com/kos-v/sensors-informer/internal"
	"github.com/kos-v/sensors-informer/internal/channel"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/sensor"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	rch := make(chan report.Report)

	listener := internal.ReportListener{
		Ch:       rch,
		Channels: channel.GetChannels(*conf),
		Config:   *conf,
	}
	listener.Listen()

	detector := internal.Detector{
		Config: *conf,
		Reader: &sensor.CommandReader{conf.LmSensors.Command},
		Rch:    rch,
	}
	detector.Run()
}
