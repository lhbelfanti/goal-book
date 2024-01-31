## Chapter 10 exercises

#### Exercise 10.1

Extend the `jpeg` program so that it converts any supported input format to any output format, using `image.Decode` to detect the input format and a flag to select the output format.

#### Exercise 10.2

Define a generic archive file-reading function capable of reading ZIP files (archive/zip) and POSIX tar filex (archive/tar).
Use a registration mechanism sililar to the one described above so that support for each file format can be plugged in using blank imports.

#### Exercise 10.3

Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service hosts the code samples for this book.
(HTTP requests from `go get` include the `go-get` parameter so that servers can disginguish them from ordinary browser requests.)

## Result

```shell
go run fetch.go http://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
```

#### Exercise 10.4

Construct a tool that reports the set of all packages in the worksapce that transitively depend on the packages specified by the arguments.
Hint: you will need to run `to list` twice, once for the initial packages and once for all packages.
You may want to parse its JSON output using the `encodin g/json` package (§4.5)

