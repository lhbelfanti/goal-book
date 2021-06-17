// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.5
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

type Color struct {
	Step  float64
	Color color.Color
}

var palette = []Color{
	{Color: color.RGBA{G: 0x04, B: 0x0f, A: 0xff}},
	{Color: color.RGBA{R: 0x03, G: 0x26, B: 0x28, A: 0xff}},
	{Color: color.RGBA{R: 0x07, G: 0x3e, B: 0x1e, A: 0xff}},
	{Color: color.RGBA{R: 0x18, G: 0x55, B: 0x08, A: 0xff}},
	{Color: color.RGBA{R: 0x5f, G: 0x6e, B: 0x0f, A: 0xff}},
	{Color: color.RGBA{R: 0x84, G: 0x50, B: 0x19, A: 0xff}},
	{Color: color.RGBA{R: 0x9b, G: 0x30, B: 0x22, A: 0xff}},
	{Color: color.RGBA{R: 0xb4, G: 0x92, B: 0x2f, A: 0xff}},
	{Color: color.RGBA{R: 0x94, G: 0xca, B: 0x3d, A: 0xff}},
	{Color: color.RGBA{R: 0x4f, G: 0xd5, B: 0x51, A: 0xff}},
	{Color: color.RGBA{R: 0x66, G: 0xff, B: 0xb3, A: 0xff}},
	{Color: color.RGBA{R: 0x82, G: 0xc9, B: 0xe5, A: 0xff}},
	{Color: color.RGBA{R: 0x9d, G: 0xa3, B: 0xeb, A: 0xff}},
	{Color: color.RGBA{R: 0xd7, G: 0xb5, B: 0xf3, A: 0xff}},
	{Color: color.RGBA{R: 0xfd, G: 0xd6, B: 0xf6, A: 0xff}},
	{Color: color.RGBA{R: 0xff, G: 0xf0, B: 0xf2, A: 0xff}},
}

const (
	xmin, ymin, xmax, ymax 	= -2, -2, +2, +2
	width, height          	= 2048, 2048
	maxIterations			= 128
)

func main() {
	filepath := os.Args[1]
	fi, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error while opening file: %s\n", err)
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fullPalette := interpolateColors()
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			m := mandelbrot(z)
			iteration := float64(maxIterations - m)

			var colorToPlot uint32 = 4294967295 // Max uint

			// The color depends on the number of iterations
			if int(math.Abs(iteration)) < len(fullPalette)-1 {
				color1 := fullPalette[int(math.Abs(iteration))]
				color2 := fullPalette[int(math.Abs(iteration))+1]
				colorToPlot = linearInterpolation(rgbaToUint(color1), rgbaToUint(color2), uint32(iteration))
			}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, uint32ToRgba(colorToPlot))
		}
	}
	png.Encode(fi, img) // NOTE: ignoring errors

	if err := fi.Close(); err != nil {
		fmt.Printf("Error while closing file: %s\n", err)
		os.Exit(1)
	}
}

func mandelbrot(c complex128) int {
	var z complex128
	n := 0
	for cmplx.Abs(z) <= 2 && n < maxIterations {
		z = z*z + c
		n++
	}

	return n
}

func linearInterpolation(c1, c2, mu uint32) uint32 {
	return c1*(1-mu) + c2*mu
}

func rgbaToUint(color color.RGBA) uint32 {
	r, g, b, a := color.RGBA()
	r /= 0xff
	g /= 0xff
	b /= 0xff
	a /= 0xff
	return uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
}

func uint32ToRgba(col uint32) color.RGBA {
	r := col >> 24 & 0xff
	g := col >> 16 & 0xff
	b := col >> 8 & 0xff
	a := 0xff
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}


func interpolateColors() []color.RGBA {
	var factor float64
	var steps []float64
	var cols []uint32
	var interpolated []uint32
	var interpolatedColors []color.RGBA

	factor = 1.0 / maxIterations
	for index, col := range palette {
		if col.Step == 0.0 && index != 0 {
			stepRatio := float64(index+1) / float64(len(palette))
			step := float64(int(stepRatio*100)) / 100 // truncate to 2 decimal precision
			steps = append(steps, step)
		} else {
			steps = append(steps, col.Step)
		}
		r, g, b, a := col.Color.RGBA()
		r /= 0xff
		g /= 0xff
		b /= 0xff
		a /= 0xff
		uintColor := uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
		cols = append(cols, uintColor)
	}

	var min, max, minColor, maxColor float64
	if len(palette) == len(steps) && len(palette) == len(cols) {
		for i := 0.0; i <= 1; i += factor {
			for j := 0; j < len(palette)-1; j++ {
				if i >= steps[j] && i < steps[j+1] {
					min = steps[j]
					max = steps[j+1]
					minColor = float64(cols[j])
					maxColor = float64(cols[j+1])
					uintColor := cosineInterpolation(maxColor, minColor, (i-min)/(max-min))
					interpolated = append(interpolated, uint32(uintColor))
				}
			}
		}
	}

	for _, pixelValue := range interpolated {
		r := pixelValue >> 24 & 0xff
		g := pixelValue >> 16 & 0xff
		b := pixelValue >> 8 & 0xff
		a := 0xff

		interpolatedColors = append(interpolatedColors, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
	}


	return interpolatedColors
}

func cosineInterpolation(c1, c2, mu float64) float64 {
	mu2 := (1 - math.Cos(mu*math.Pi)) / 2.0
	return c1*(1-mu2) + c2*mu2
}
