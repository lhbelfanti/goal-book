// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Example:
// go run main.go &
// http://localhost:8000/?width=2000;height=1000;cells=200;color=c1cF13

// Exercise 3.4
package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const angle = math.Pi / 6						// angle of x, y axes (=30°)
const xyrange = 30.0							// axis ranges (-xyrange..+xyrange)

var width = 600									// canvas width in pixels
var height = 320								// canvas height in pixels
var cells = 100									// number of grid cells
var xyscale = float64(width) / 2 / xyrange 		// pixels per x or y unit
var zscale = float64(height) * 0.4        		// pixels per z unit


var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		color := color.RGBA{R: 255, G: 255, B: 255, A: 1}

		if err := r.ParseForm(); err != nil {
			fmt.Print(err)
		}

		for k, v := range r.Form {
			switch k {
			case "width": {
				width, _ = strconv.Atoi(v[0])
			}
			case "height": {
				height, _ = strconv.Atoi(v[0])
			}
			case "cells": {
				cells, _ = strconv.Atoi(v[0])
			}
			case "color": {
				color, _ = parseHexColor("#" + v[0])
			}
			}
		}


		xyscale = float64(width) / 2 / xyrange 		// pixels per x or y unit
		zscale = float64(height) * 0.4        		// pixels per z unit

		fmt.Printf("Width: %d, height: %d, cells: %d, color: %v\n", width, height, cells, color)

		w.Header().Set("Content-Type", "image/svg+xml")

		plot(w, color)
	})



	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(out io.Writer, c color.RGBA) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
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
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgb(%d, %d, %d); stroke:#000; stroke-width: 0.8;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B)
		}
	}
	fmt.Fprintf(out, "</svg>")
}


func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)
	isValid := !math.IsNaN(z)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, isValid
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
