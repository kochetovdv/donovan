package main

import (
	"fmt"
	tempconv "donovan/chapter2/tempconv/tempconv0"
)

func main(){
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC) // ”100" °C
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC)) // "180й °F
//	fmt.Printf("%g\n", boilingFFreezingC)     // Ошибка компиляции: несоответствие типов
}