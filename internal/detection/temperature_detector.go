package detection

import (
	r "github.com/kos-v/sensors-informer/internal/report"
	s "github.com/kos-v/sensors-informer/internal/sensor"
	"github.com/kos-v/sensors-informer/internal/temperature"
	"github.com/kos-v/sensors-informer/internal/temperature/convert"
	"time"
)

type TemperatureDetector struct {
	Critical temperature.Value
	Reader   s.Reader
	Unit     temperature.Unit
}

func (d *TemperatureDetector) Detect() (*r.Report, error) {
	var report r.Report

	busList, err := d.Reader.ReadTemperatureSensors()
	if err != nil {
		return &report, err
	}

	for _, bus := range busList {
		for _, sensor := range bus.Sensors {
			if !sensor.HasInputInfo() {
				continue
			}

			tempVal := temperature.Value(sensor.GetInputInfo().GetValueAsFloat())
			if d.Unit == temperature.UnitFahrenheit {
				tempVal = convert.ToFahrenheit(temperature.UnitCelsius, tempVal)
			}
			if tempVal >= d.Critical {
				report.AddSensorReport(bus, sensor)
			}
		}
	}

	report.Time = time.Now()

	return &report, nil
}
