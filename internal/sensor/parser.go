package sensor

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	SensorTypeFan         = "fan"
	SensorTypeTemperature = "temperature"
	SensorTypeUnknown     = "unknown"

	busNamePattern    = `^(\w+-+)+\d+$`
	sensorInfoPattern = `^  (\w*)(:\s*)(\d+\.\d+)$`
	sensorNamePattern = `^([^ ].+):$`
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

func (i SensorInfo) GetValueAsInt() float32 {
	f, err := strconv.ParseFloat(i.Value, 32)
	if err != nil {
		return 0
	}

	return float32(f)
}

type Parser struct {
	SensorsOut string
}

func (p *Parser) Parse() []Bus {
	var busList []Bus

	for _, line := range strings.Split(p.SensorsOut, "\n") {
		if p.match(busNamePattern, line) {
			busList = append(busList, Bus{Name: line})
			continue
		}

		lastBus := len(busList) - 1
		if lastBus < 0 {
			continue
		}

		if p.match(sensorNamePattern, line) {
			busList[lastBus].Sensors = append(busList[lastBus].Sensors, Sensor{
				Name: p.parseSensorName(line),
				Type: SensorTypeUnknown,
			})
			continue
		}

		lastSensor := len(busList[lastBus].Sensors) - 1
		if lastSensor < 0 {
			continue
		}

		if p.match(sensorInfoPattern, line) {
			info := p.parseSensorInfo(line)
			if info.FullName != "" {
				busList[lastBus].Sensors[lastSensor].Type = p.parseSensorType(strings.TrimSpace(line))
				busList[lastBus].Sensors[lastSensor].Info = append(busList[lastBus].Sensors[lastSensor].Info, info)
			}
		}
	}

	return busList
}

func (p *Parser) match(pattern, data string) bool {
	matched, _ := regexp.MatchString(pattern, data)
	return matched
}

func (p *Parser) parseSensorInfo(data string) SensorInfo {
	info := SensorInfo{}
	re, _ := regexp.Compile(sensorInfoPattern)

	matches := re.FindStringSubmatch(data)
	if len(matches) < 4 {
		return info
	}

	info.FullName = matches[1]
	info.Value = matches[3]

	nameParts := strings.Split(info.FullName, "_")
	if len(nameParts) > 1 {
		info.Name = strings.Join(nameParts[1:], "_")
	}

	return info
}

func (p *Parser) parseSensorName(data string) string {
	re, _ := regexp.Compile(sensorNamePattern)
	matches := re.FindStringSubmatch(data)
	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}
func (p *Parser) parseSensorType(data string) string {
	if strings.HasPrefix(data, "temp") {
		return SensorTypeTemperature
	} else if strings.HasPrefix(data, "fan") {
		return SensorTypeFan
	}

	return SensorTypeUnknown
}
