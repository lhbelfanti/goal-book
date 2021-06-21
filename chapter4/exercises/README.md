## Chapter 4 exercises

#### Exercise 4.1
Write a function that counts the number of bits that are different in two SHA256 hashes. (See `PopCount` from Section 2.6.2.)

#### Exercise 4.2
Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

#### Exercise 4.3
Rewrite `reverse` to use an array pointer instead of a slice.

#### Exercise 4.4
Write a version of `rotate` that operates in a single pass.

#### Exercise 4.5
Write an in-place function to eliminate adjacent duplicates in a `[]string` slice.

#### Exercise 4.6
Write an in-place function that squashes each run of adjacent Unicode spaces (see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.

#### Exercise 4.7
Modify `reverse` to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

#### Exercise 4.8
Modify `charcount` to count letters, digits, and so on in their Unicode categories, using functions like `unicode.IsLetter`.

#### Exercise 4.9
Write a program `wordfreq` to report the frequency of each word in an input text file. Call `input.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines.
