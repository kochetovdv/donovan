package main

import (
	"fmt"
	"os"
)

func main() {
	echo1()
	echo2()
	echo3()
}

// Аргументы
func echo1() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println("Аргументы: ", s)
}

// Команда и аргументы
func echo2() {
	s, sep := "", " "
	for _, arg := range os.Args {
		s += arg + sep
	}
	fmt.Println("Команда и аргументы: ", s)
}

// Индекс и аргумент в каждой стоке
func echo3() {
	fmt.Println("Введены следующие аргументы: ")
	for i, arg := range os.Args {
		fmt.Printf("№%d\t%s\n", i, arg)
	}
}
