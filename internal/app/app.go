package app

import (
	"github.com/kos-v/sensors-informer/internal"
	"github.com/kos-v/sensors-informer/internal/channel"
	conf "github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/detection"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/sensor"
)

type app struct {
	configPath string
}

func (a *app) Run() error {
	config, err := conf.LoadConfig(a.configPath)
	if err != nil {
		return err
	}

	rch := make(chan report.Report)

	listener := internal.NewReportListener(
		config.Report.RepeatTimeout,
		channel.GetChannels(*config, &message.PredictableFormatter{Opts: config.Report.Format}),
		rch,
	)
	listener.Listen()

	sensorsReader := &sensor.CommandReader{Command: config.LmSensors.Command}
	watcher := detection.Watcher{
		Detectors: []detection.Detector{
			&detection.TemperatureDetector{
				Critical: config.Sensors.Temperature.Critical,
				Reader:   sensorsReader,
				Unit:     config.Sensors.Temperature.Unit,
			},
		},
		Interval:      config.Sensors.PollingInterval,
		ReportChannel: rch,
	}
	watcher.Run()

	return nil
}

func NewApp(configPath string) *app {
	return &app{configPath: configPath}
}
