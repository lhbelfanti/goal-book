## Chapter 3 exercises

#### Exercise 3.1
If the function `f` returns a non-finite `float64` value, the SVG file will contain invalid `<polygon>` element (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.

#### Exercise 3.2
Experiment with visualizations of other functions from the `math` package. Can you produce an egg box, moguls or a saddle?

#### Exercise 3.3
Color each polygon based on its height, so that the peaks are colored red `(#ff0000)` and the valleys blue `(#0000ff)`.

#### Exercise 3.4
Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the `Content-Type` header like this:

`w.Header().Set("Content-Type", "image/svg+xml")`

(This step was not required in the Lissajous example because the server uses standard heuristics to recognize common formats like PNG from the first 512 bytes of the response, and generates the proper header.) Allow the client to specify values like height, width, and color as HTTP request parameters.

#### Exercise 3.5
Implement a full-color Mandelbrot set using the function `image.NewRGBA` and the type `color.RGBA` or `color.YCbCr`.

#### Exercise 3.6
Supersampling is a technique to reduce the effect of pixelation by computing the color value at seveal points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels". Implement it.

#### Exercise 3.7
Another simple fractal uses Newton's method to find complex solutions to a function such as `z⁴-1 = 0`. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.

#### Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float` and `big.Rat`. (The latter two types are found in the `math/big` package. `Float` uses arbitrary but bounded-precision floating-point; `Rat` uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom level do rendering artifacts become visible?

Benchmark:
```
BenchmarkMandelbrotComplex64
BenchmarkMandelbrotComplex64-4   	42029980	        29.3 ns/op

BenchmarkMandelbrotComplex128
BenchmarkMandelbrotComplex128-4   	50313486	        23.5 ns/op

BenchmarkMandelbrotBigFloat
BenchmarkMandelbrotBigFloat-4   	1984020	       		505 ns/op

BenchmarkMandelbrotBigRat
BenchmarkMandelbrotBigRat-4   	  	519090	      		2056 ns/op
```
#### Exercise 3.9
Write a web server that renders fractals and writes the image data to the client. Allow the client to specify the `x`, `y` and `zoom` values as parameters to the HTTP request.

#### Exercise 3.10
Write a non-recursive version of `comma`, using `bytes.Buffer` instead of string concatenation.

#### Exercise 3.11
Enhance `comma` so that it deals correctly with floating-point numbers and an optional sign.

#### Exercise 3.12
Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.

#### Exercise 3.13
Write `const` declarations for KB, MB, up through YB as compactly as you can.
