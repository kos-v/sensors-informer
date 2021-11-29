package sensor

import (
	"io/ioutil"
	"os/exec"
)

type Reader interface {
	Read() ([]Bus, error)
	ReadTemperatureSensors() ([]Bus, error)
}

type FileReader struct {
	Path string
}

func (r *FileReader) Read() ([]Bus, error) {
	data, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}

	parser := Parser{
		SensorsOut: string(data),
	}

	return parser.Parse(), nil
}

func (r *FileReader) ReadTemperatureSensors() ([]Bus, error) {
	return readTemperatureSensors(r)
}

type CommandReader struct {
	Command string
}

func (r *CommandReader) Read() ([]Bus, error) {
	cmd := exec.Command(r.Command, "-uA")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	parser := Parser{
		SensorsOut: string(stdout),
	}

	return parser.Parse(), nil
}

func (r *CommandReader) ReadTemperatureSensors() ([]Bus, error) {
	return readTemperatureSensors(r)
}

func readTemperatureSensors(reader Reader) ([]Bus, error) {
	busList, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var newBusList []Bus
	for _, bus := range busList {
		newBus := Bus{}
		for _, sensor := range bus.Sensors {
			if sensor.Type == SensorTypeTemperature {
				newBus.Sensors = append(newBus.Sensors, sensor)
			}
		}
		if len(newBus.Sensors) > 0 {
			newBus.Name = bus.Name
			newBusList = append(newBusList, newBus)
		}
	}

	return newBusList, nil
}
