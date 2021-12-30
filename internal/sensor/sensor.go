package sensor

import "strconv"

const (
	SensorTypeFan         = "fan"
	SensorTypeTemperature = "temperature"
	SensorTypeUnknown     = "unknown"
)

type Bus struct {
	Name    string
	Sensors []Sensor
}

type Sensor struct {
	Name string
	Info []SensorInfo
	Type string
}

func (s Sensor) HasInputInfo() bool {
	return s.GetInputInfo().Name != ""
}

func (s Sensor) GetInputInfo() SensorInfo {
	for _, v := range s.Info {
		if v.Name == "input" {
			return v
		}
	}

	return SensorInfo{}
}

type SensorInfo struct {
	FullName string
	Name     string
	Value    string
}

func (i SensorInfo) GetValueAsFloat() float32 {
	f, err := strconv.ParseFloat(i.Value, 32)
	if err != nil {
		return 0
	}

	return float32(f)
}
