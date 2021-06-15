// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.2
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
	mogulHeight	  = 0.1
	mogulLength	  = 2 * (math.Pi/12)	// moguls1.svg -- (math.Pi/6) // moguls2.svg and moguls3.svg
	mogulWidth	  = 2 * (math.Pi/4)		// moguls1.svg -- (math.Pi/6) // moguls2.svg and moguls3.svg
	eggBoxHeight  = 0.2
	eggBoxF		  = 2 * (math.Pi/6)
	saddleHeight  = 1.0
	saddleF	  	  = 2 * (math.Pi/(cells/3))
)


var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	filepath := os.Args[1]
	fi, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error while opening file: %s\n", err)
		os.Exit(1)
	}

	fi.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, valid := corner(i+1, j)
			if !valid {
				continue
			}
			bx, by, valid := corner(i, j)
			if !valid {
				continue
			}
			cx, cy, valid := corner(i, j+1)
			if !valid {
				continue
			}
			dx, dy, valid := corner(i+1, j+1)
			if !valid {
				continue
			}
			fi.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	fi.WriteString("</svg>")

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := moguls(x, y)
	// z := eggBox(x, y)
	// z := saddle(x, y)

	isValid := !math.IsNaN(z)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isValid
}

func eggBox(x, y float64) float64 {
	r := math.Sin(eggBoxF*x) * math.Sin(eggBoxF*y)
	return eggBoxHeight * r
}

func moguls(x, y float64) float64 {
	r := math.Cos(mogulLength*x) * math.Cos(mogulWidth*y)
	// r := math.Sin(mogulLength*x) * math.Sin(mogulWidth*y) // -> moguls3.svg
	return mogulHeight * math.Abs(r)
}

func saddle(x, y float64) float64 {
	r := math.Sin(saddleF*x) * math.Sin(saddleF*y)
	return saddleHeight * r
}
