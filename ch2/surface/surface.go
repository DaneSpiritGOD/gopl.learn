// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	file, err := os.Create("wow.svg")
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			writeCorner(file, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(file, "</svg>")
}

func writeCorner(file *os.File, ax, ay, bx, by, cx, cy, dx, dy float64) {
	if isValid(ax) &&
		isValid(ay) &&
		isValid(ay) &&
		isValid(ay) &&
		isValid(ay) &&
		isValid(ay) &&
		isValid(ay) {
		fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy)
	}
}

func isValid(num float64) bool {
	return !math.IsInf(num, 1) && !math.IsInf(num, -1) && !math.IsNaN(num)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func fEggbox(x, y float64) float64 { //鸡蛋盒
	r := 0.2 * (math.Cos(x) + math.Cos(y))
	return r
}

func fSaddle(x, y float64) float64 { //马鞍
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	r := y*y/a2 - x*x/b2
	return r
}
