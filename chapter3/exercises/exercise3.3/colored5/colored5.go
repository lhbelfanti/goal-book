// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.3 (fifth approach)
package main

import (
	"fmt"
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
	filepath := os.Args[1]
	fi, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error while opening file: %s\n", err)
		os.Exit(1)
	}

	fi.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: green; fill: black; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if z > 0 {
				fi.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:black;stroke:red;stroke-width:1'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy))
			} else {
				fi.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:black;stroke:blue;stroke-width:1'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy))

			}
		}
	}
	fi.WriteString("</svg>")

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
