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

#### Exercise 4.10
Modify `issues` to report the results in age categories, say less than a month old, less than a year old, and more than a year old.

#### Exercise 4.11
Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.

#### Exercise 4.12
The popular web comic `xkcd` has a `JSON` interface. For example, a request to `http://xkcd.com/571/info.0.json` produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool `xkcd` that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.

#### Exercise 4.13
The JSON-based web service of the Open Movie Database lets you search `https://omdbapi.com/` for a movie by name and download its poster image. Write a tool `poster` that downloads the poster image for the movie named on the command line.
