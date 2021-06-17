// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 3.8
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
	"strconv"
)

const (
	cmpl64 = iota
	cmpl128
	bigFloat
	bigRat // extremely slow rendering speed...
	typeSize
)

func main() {
	filepath := os.Args[1]

	for i := 0; i < typeSize; i++ {
		fi, err := os.Create(filepath + strconv.Itoa(i) + ".png")
		if err != nil {
			fmt.Printf("Error while opening file: %s\n", err)
			os.Exit(1)
		}

		draw(fi, i)

		if err := fi.Close(); err != nil {
			fmt.Printf("Error while closing file: %s\n", err)
			os.Exit(1)
		}
	}
}

func draw(out io.Writer, typ int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			switch typ {
			case cmpl64:
				img.Set(px, py, mandelbrot64(z))
			case cmpl128:
				img.Set(px, py, mandelbrot128(z))
			case bigFloat:
				img.Set(px, py, mandelbrotFloat(z))
			case bigRat:
				img.Set(px, py, mandelbrotRat(z))
			}
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	vR := &big.Float{}
	vI := &big.Float{}
	for n := uint8(0); n < iterations; n++ {
		// (r+i)^2=r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2

		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) > 0 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	// Problems with high resolution images and slow 'escape' speed.
	// Because multiplication of arbitrary-precision arithmetic
	// achieve O(N log(N) log(log(N))) complexity
	// https://en.wikipedia.org/wiki/Arbitrary-precision_arithmetic#Implementation_issues
	// To create the example image that is uploaded I had to reduce this number to 10 and reduce the canvas size to the half
	const iterations = 200
	const contrast = 15

	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	vR := &big.Rat{}
	vI := &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		// (r+i)^2=r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2

		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) > 0 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
