package report

import (
	"github.com/kos-v/sensors-informer/internal/sensor"
	"time"
)

type Report struct {
	Time    time.Time
	Sensors []struct {
		BusName     string
		SensorName  string
		SensorValue float32
	}
}

func (r *Report) AddSensorReport(b sensor.Bus, s sensor.Sensor) {
	r.Sensors = append(r.Sensors, struct {
		BusName     string
		SensorName  string
		SensorValue float32
	}{BusName: b.Name, SensorName: s.Name, SensorValue: s.GetInputInfo().GetValueAsInt()})
}

func (r *Report) HasSensorReports() bool {
	return len(r.Sensors) > 0
}
