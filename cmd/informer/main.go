package main

import (
	"flag"
	"github.com/kos-v/sensors-informer/internal"
	"github.com/kos-v/sensors-informer/internal/channel"
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/sensor"
	"log"
)

func main() {
	specificConfig := flag.String("config", "", "The path to a specific configuration file")
	flag.Parse()

	config, err := conf.LoadConfig(*specificConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	rch := make(chan report.Report)

	listener := internal.ReportListener{
		Ch:       rch,
		Channels: channel.GetChannels(*config, &message.PredictableFormatter{Config: *config}),
		Config:   *config,
	}
	listener.Listen()

	detector := internal.Detector{
		Config: *config,
		Reader: &sensor.CommandReader{config.LmSensors.Command},
		Rch:    rch,
	}
	detector.Run()
}
