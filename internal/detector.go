package internal

import (
	"github.com/kos-v/sensors-informer/internal/config"
	r "github.com/kos-v/sensors-informer/internal/report"
	s "github.com/kos-v/sensors-informer/internal/sensor"
	"github.com/kos-v/sensors-informer/internal/temperature"
	"log"
	"sync"
	"time"
)

type Detector struct {
	Config config.Config
	Reader s.Reader
	Rch    chan r.Report
}

func (d *Detector) Run() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			report, err := d.detect()
			if err == nil {
				if report.HasSensorReports() {
					d.Rch <- report
				}
			} else {
				log.Printf("Error: %s", err.Error())
			}

			time.Sleep(time.Duration(d.Config.Sensors.PollingInterval) * time.Second)
		}
	}()
}

func (d *Detector) detect() (r.Report, error) {
	var report r.Report

	busList, err := d.Reader.ReadTemperatureSensors()
	if err != nil {
		return report, err
	}

	for _, bus := range busList {
		for _, sensor := range bus.Sensors {
			if !sensor.HasInputInfo() {
				continue
			}

			if temperature.Value(sensor.GetInputInfo().GetValueAsInt()) > d.Config.Sensors.CriticalTemperature {
				report.AddSensorReport(bus, sensor)
			}
		}
	}

	report.Time = time.Now()

	return report, nil
}
