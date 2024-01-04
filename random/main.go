package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerlD int
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000 // Зарплата снижена, пишет мало строк кода

	position := &dilbert.Position
	*position = "Senior " + *position // Повышен в должности

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (активный участник команды)"

	// (*employeeOfTheMonth).Position += " (активный участник команды)"
	fmt.Println(EmployeeByID(dilbert.ManagerlD).Position) // Босс

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // Уволить...
}

func EmployeeByID(id int) *Employee{
	return &Employee{}
}