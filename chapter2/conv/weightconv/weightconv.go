// Пакет weightconv выполняет конвертацию килограмм в фунты и обратно.

package weightconv

type Kilograms float64
type Pounds float64

const (
	PoundsСoefficient float64 = 2.20462
)

func KToP(k Kilograms) Pounds { return Pounds(k*Kilograms(PoundsСoefficient)) }
func PToK(p Pounds) Kilograms { return Kilograms(p/Pounds(PoundsСoefficient)) }