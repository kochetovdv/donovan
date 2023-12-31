// Упражнение 3.1. Если функция f возвращает значение float64, не являющееся 
// конечным, SVG-файл содержит неверные элементы <polygon> (хотя многие визуализаторы SVG успешно 
// обрабатывают эту ситуацию). Измените программу так, чтобы некорректные многоугольники были опущены

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange // Пикселей в единице х или у
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей х, у (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style=' stroke: grey; fill: white; stroke-width: 0.7' "+
		"width=,%d' height='%d' >", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z, ok := f(x, y)
	if !ok {
		return 0, 0
	}
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, у float64) (float64, bool) {
	r := math.Hypot(x, у) // Расстояние от (0,0)
	if math.IsInf(r, 0) || math.IsNaN(r) {
		return 0, false
	}
	return math.Sin(r) / r, true
}
