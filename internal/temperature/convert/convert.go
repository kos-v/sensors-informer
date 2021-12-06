package convert

import "github.com/kos-v/sensors-informer/internal/temperature"

func ToFahrenheit(unit temperature.Unit, val temperature.Value) temperature.Value {
	if unit == temperature.UnitCelsius {
		return val*1.8 + 32
	}
	return val
}
