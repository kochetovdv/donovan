// Пакет distanceconv выполняет конвертацию метров в футы и обратно.

package distanceconv

type Meter float64
type Feet float64

const (
	FeetCoefficient float64 = 3.28084
)

func MToF(m Meter) Feet { return Feet(m*Meter(FeetCoefficient)) }
func FToM(f Feet) Meter { return Meter(f/Feet(FeetCoefficient)) }