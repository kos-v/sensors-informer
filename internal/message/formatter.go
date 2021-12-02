package message

import (
	"fmt"
	"github.com/kos-v/sensors-informer/internal/report"
)

type Formatter interface {
	FormatTitle(r *report.Report, withTimestamp bool) string
	FormatBodyRows(r *report.Report) []string
}

type PredictableFormatter struct{}

func (f *PredictableFormatter) FormatTitle(r *report.Report, withTimestamp bool) string {
	tt := ""
	if withTimestamp {
		tt = fmt.Sprintf("[%s]\n", r.Time.Format("2006-01-02 15:04:05 Z07"))
	}

	return tt + "Critical temperature readings"
}

func (f *PredictableFormatter) FormatBodyRows(r *report.Report) []string {
	var rows []string

	for _, v := range r.Sensors {
		rows = append(rows, fmt.Sprintf("%s::%s: %.1f°С", v.BusName, v.SensorName, v.SensorValue))
	}

	return rows
}
