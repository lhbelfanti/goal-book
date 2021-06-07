## Chapter 2 exercises

#### Exercise 2.1
Add types, constants, and functions to `tempconv` for processing temperatures in the Kelvin scale, where zero Kelvin is -273.15°C and a difference of 1K has the same magnitude as 1°C.

#### Exercise 2.2
Write a general-purpose unit-conversion program analogous to `cf` that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.

#### Exercise 2.3
Rewrite `PopCount` to use a loop instead of a single expression. Compare the performance of the two versions.

#### Exercise 2.4
Write a version of `PopCount` that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.

#### Exercise 2.5
The expression `x&(x-1)` clears the rightmost non-zero bit of `x`. Write a version of `PopCount` that counts bits by using this fact, and assess its performance.
