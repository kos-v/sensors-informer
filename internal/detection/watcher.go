package detection

import (
	r "github.com/kos-v/sensors-informer/internal/report"
	"log"
	"sync"
	"time"
)

type Watcher struct {
	Detectors     []Detector
	Interval      uint
	ReportChannel chan r.Report
}

func (w *Watcher) Run() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			for _, d := range w.Detectors {
				report, err := d.Detect()
				if err == nil {
					if report.HasSensorReports() {
						w.ReportChannel <- *report
					}
				} else {
					log.Printf("Error: %s", err.Error())
				}
			}

			time.Sleep(time.Duration(w.Interval) * time.Second)
		}
	}()
}
