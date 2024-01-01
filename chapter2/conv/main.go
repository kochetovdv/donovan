//Excercise 2.2. Convertations from argument in float to any metrics
// Напишите программу общего назначения для преобразования
// единиц, аналогичную cf , которая считывает числа из аргументов командной строки
// (или из стандартного ввода, если аргументы командной строки отсутствуют) и преобразует каждое число в
// другие единицы, как температуру — в градусы Цельсия и Фаренгейта, длину — в футы и метры, вес — в фунты и
// килограммы и т.д.

package main

import (
	distance "donovan/chapter2/conv/distanceconv"
	temperature "donovan/chapter2/conv/tempconv"
	weight "donovan/chapter2/conv/weightconv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}
		f := temperature.Fahrenheit(t)
		c := temperature.Celsius(t)
		m := distance.Meter(t)
		ft := distance.Feet(t)
		kg:= weight.Kilograms(t)
		p := weight.Pounds(t)
		fmt.Printf("%v = %v, %v = %v,\n %v = %v, %v = %v,\n %v = %v, %v = %v\n",
			f, temperature.FToC(f), c, temperature.CToF(c),
			m, distance.MToF(m), ft, distance.FToM(ft),
			kg, weight.KToP(kg), p, weight.PToK(p),
		)
	}
}
