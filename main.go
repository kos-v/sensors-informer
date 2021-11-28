package main

import (
	"github.com/kos-v/sensors-informer/channel"
	"github.com/kos-v/sensors-informer/config"
	"github.com/kos-v/sensors-informer/report"
	"github.com/kos-v/sensors-informer/sensor"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	rch := make(chan report.Report)

	listener := ReportListener{
		Ch:       rch,
		Channels: channel.GetChannels(*conf),
		Config:   *conf,
	}
	listener.Listen()

	detector := Detector{
		Config: *conf,
		Reader: &sensor.CommandReader{conf.LmSensors.Command},
		Rch:    rch,
	}
	detector.Run()
}
