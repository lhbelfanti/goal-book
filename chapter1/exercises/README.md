## Chapter 1 exercises

#### Exercise 1.1
Modify the echo program to also print `os.Args[0]`, the name of the command that invoked it.

#### Exercise 1.2
Modify the echo program to print the index and value of each of its arguments, one per line.

#### Exercise 1.3
Experiment to measure the difference in running time between our potentially inefficient version and the one that uses `strings.Join`.

#### Exercise 1.4
Modify dup2 to print the names of all files in which each duplicated line occurs.

#### Exercise 1.5
Change the Lissajous program's color palette to green on black, for added authenticity. To create the web color `#RRGGBB`, use `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.

#### Exercise 1.6
Modify the Lissajous program to produce images in multiple colors by adding more values to `palette` and then displaying them by changing the third argument of `SetColorIndex` in some interesting way.
