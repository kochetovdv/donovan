// Упражнение 3.3. Окрасьте каждый многоугольник цветом, зависящим от его высоты, так,
// чтобы пики были красными (#ff0000), а низины — синими (#0000ff).

package main

import (
	"fmt"
	"math"
	"os"
)

const (
	white         = "#ffffff"
	red           = "#ff0000"
	blue          = "#0000ff"
	grey          = "#808080"
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange // Пикселей в единице х или у
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 3         // Углы осей х, у (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func main() {
	file, err := os.OpenFile("donovan.svg", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	svgContent := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:%v; fill:%v; stroke-width: 0.7' "+
		"width='%d' height='%d' >", grey, white, width, height)

	_, err = file.WriteString(svgContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			polygon := fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			_, err = file.WriteString(polygon)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	_, err = file.WriteString("</svg>")
	if err != nil {
		fmt.Println(err)
		return
	}
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
