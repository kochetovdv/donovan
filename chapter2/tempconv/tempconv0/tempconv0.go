// Пакет tempconv выполняет вычисления температур
// по Цельсию (Celsius) и по Фаренгейту (Fahrenheit),
package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + AbsoluteZeroC) }
func KToC(k Kelvin) Celsius     { return Celsius(k + Kelvin(AbsoluteZeroC))}
func FToK(f Fahrenheit) Kelvin  { return CToK(FToC(f)) }
func KToF(k Kelvin) Fahrenheit  { return CToF(KToC(k)) }