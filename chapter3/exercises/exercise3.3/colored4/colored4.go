// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.3 (fourth approach)
package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
)

const (
	width, height = 1000, 600            // canvas size in pixels
	cells         = 100                // number of grid cells
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

	fi.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1, valid := corner(i+1, j)
			if !valid {
				continue
			}
			bx, by, z2, valid := corner(i, j)
			if !valid {
				continue
			}
			cx, cy, z3, valid := corner(i, j+1)
			if !valid {
				continue
			}
			dx, dy, z4, valid := corner(i+1, j+1)
			if !valid {
				continue
			}

			depth := [4]float64{z1, z2, z3, z4}
			c := colorized(depth)

			//fmt.Printf("(%d,%d) -- A: %g - B: %g - C: %g - D: %g -> color: %v\n", i, j, z1, z2, z3, z4, c)

			fi.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgb(%d, %d, %d); stroke:#000; stroke-width: 0.8;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B))
		}
	}
	fi.WriteString("</svg>")

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	isValid := !math.IsNaN(z)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, isValid
}

func colorized(z [4]float64) color.NRGBA {
	v := chooseDepth(z)

	var maxRange float64 = 1
	var maxColors float64 = 510
	var r float64 = 255
	var g float64 = 0
	var b float64 = 255

	t := (v * maxColors) / maxRange

	if t > 255 { // Higher peaks
		b -= t - r
	} else if t <= 255 && t > 0 { // Lower peaks
		r -= t
		b = t - r
	} else if t < 0 { // Valley
		r = math.Abs(t)
	}

	c := color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 1}
	return c
}

func chooseDepth(array [4]float64) (float64) {
	maxPositive := -10.0
	minNegative := -10.0
	for _, value := range array {
		if value > 0 && value > maxPositive {
			maxPositive = value
		}

		if value < 0 && value > minNegative {
			minNegative = value
		}
	}

	if maxPositive != -10.0 {
		return maxPositive
	}

	return minNegative
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
