package message

import (
	"fmt"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/report"
	"github.com/kos-v/sensors-informer/internal/temperature"
	tconvert "github.com/kos-v/sensors-informer/internal/temperature/convert"
	"strings"
)

var placeholderHandlers = []struct {
	name    string
	handler func(r *report.Report) string
}{
	{
		name: "timestamp",
		handler: func(r *report.Report) string {
			return r.Time.Format("2006-01-02 15:04:05 Z07")
		},
	},
}

type Formatter interface {
	FormatTitle(r *report.Report) string
	FormatBodyRows(r *report.Report) []string
}

type PredictableFormatter struct {
	Config config.Config
}

func (f *PredictableFormatter) FormatTitle(r *report.Report) string {
	return execPlaceholders(r, f.Config.Report.Format.Title.Text)
}

func (f *PredictableFormatter) FormatBodyRows(r *report.Report) []string {
	var rows []string

	for _, v := range r.Sensors {
		sensorValue := temperature.Value(v.SensorValue)
		unitLabel := "C"
		if f.Config.Report.Format.TemperatureUnit == temperature.UnitFahrenheit {
			sensorValue = tconvert.ToFahrenheit(temperature.UnitCelsius, sensorValue)
			unitLabel = "F"
		}

		rows = append(rows, fmt.Sprintf("%s::%s: %.1f°%s", v.BusName, v.SensorName, sensorValue, unitLabel))
	}

	return rows
}

func execPlaceholders(r *report.Report, target string) string {
	for _, item := range placeholderHandlers {
		ph := fmt.Sprintf("{%s}", item.name)
		if strings.Contains(target, ph) {
			target = strings.Replace(target, ph, item.handler(r), -1)
		}
	}

	return target
}
