package detection

import r "github.com/kos-v/sensors-informer/internal/report"

type Detector interface {
	Detect() (*r.Report, error)
}
