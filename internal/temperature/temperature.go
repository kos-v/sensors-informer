package temperature

type Unit string
type Value float32

const (
	UnitCelsius    Unit = "c"
	UnitFahrenheit Unit = "f"
)

func IsSupportedUnit(v Unit) bool {
	return v == UnitCelsius || v == UnitFahrenheit
}
