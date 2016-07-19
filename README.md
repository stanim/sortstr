# sortstr

[![Travis CI](https://img.shields.io/travis/stanim/sortstr/master.svg?style=flat-square)](https://travis-ci.org/stanim/sortstr)
[![Coverage Status](http://codecov.io/github/stanim/sortstr/coverage.svg?branch=master)](http://codecov.io/github/stanim/sortstr?branch=master)
[![Documentation and Examples](https://godoc.org/github.com/stanim/sortstr?status.svg)](https://godoc.org/github.com/stanim/sortstr)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/stanim/sortstr/blob/master/LICENSE)

Sort slices of strings based on multiple indices or column
headers.
Reverse sort is possible by providing negative index
values or prefixing column header titles with `-`.
Note that indices are not zero based, but start with 1.
If an index is out of range, an empty string will be
used for comparing, rather than throwing a runtime
panic.

Normally you only need to use this function:

    By(rows, indices)

or these two together:

    NewHeaders(titles)
    ByHeaders(headers, rows, titles)

Only if you need to sort the same rows multiple times
in different orders, you'll need to use the `Multi` type
directly.

This was developed with the
[xlsx](https://github.com/tealeg/xlsx)
package in mind, but can be used independently.

### Examples

Using indices:

```go
var rows = [][]string{
	{"John Lennon", "1968", "Let It Be"},
	{"John Lennon", "1965", "Let It Be"},
	{"John Lennon", "1965", "12-Bar Original"},
	{"Paul McCartney", "1963", "All My Loving"},
	{"George Harrison", "1968",
		"While My Guitar Gently Weeps"},
	{"Ringo Star", "1965", "Untitled"},
	{"Ringo Star"},
}
sortstr.By(rows, -1, 3, 2) // reverse order for first column
sortstr.Print("By author, title, year", rows, ", ")
sortstr.By(rows, 2, 1, 3)
sortstr.Print("By year, author, title", rows, ", ")
sortstr.By(rows, 3, 1, 2)
sortstr.Print("By title, author, year", rows, ", ")
```

Output:

	By author, title, year:
	Ringo Star
	Ringo Star, 1965, Untitled
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, 12-Bar Original
	John Lennon, 1965, Let It Be
	John Lennon, 1968, Let It Be
	George Harrison, 1968, While My Guitar Gently Weeps

	By year, author, title:
	Ringo Star
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, 12-Bar Original
	John Lennon, 1965, Let It Be
	Ringo Star, 1965, Untitled
	George Harrison, 1968, While My Guitar Gently Weeps
	John Lennon, 1968, Let It Be

	By title, author, year:
	Ringo Star
	John Lennon, 1965, 12-Bar Original
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, Let It Be
	John Lennon, 1968, Let It Be
	Ringo Star, 1965, Untitled
	George Harrison, 1968, While My Guitar Gently Weeps

Using column headers:

```go
titles := []string{"author", "year", "title"}
headers := sortstr.NewHeaders(titles)
rows := [][]string{
	{"John Lennon", "1968", "Let It Be"},
	{"John Lennon", "1965", "Let It Be"},
	{"John Lennon", "1965", "12-Bar Original"},
	{"Paul McCartney", "1963", "All My Loving"},
	{"George Harrison", "1968",
		"While My Guitar Gently Weeps"},
	{"Ringo Star", "1965", "Untitled"},
}
err := sortstr.ByHeaders(headers, rows,
	"-author", "title", "year") // reverse order for author
if err != nil {
	fmt.Println("Unknown column header")
}
sortstr.Print("By -author, title, year", rows, ", ")
err = sortstr.ByHeaders(headers, rows,
	"year", "author", "title")
if err != nil {
	fmt.Println("Unknown column header")
} else {
	sortstr.Print("By year, author, title", rows, ", ")
}
err = sortstr.ByHeaders(headers, rows,
	"title", "author", "year")
if err != nil {
	fmt.Println("Unknown column header")
} else {
	sortstr.Print("By title, author, year", rows, ", ")
}
err = sortstr.ByHeaders(headers, rows,
	"title", "author", "disc")
if err != nil {
	fmt.Println("Unknown column header")
} else {
	sortstr.Print("By title, author, year", rows, ", ")
}
```

Output:

	By -author, title, year:
	Ringo Star, 1965, Untitled
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, 12-Bar Original
	John Lennon, 1965, Let It Be
	John Lennon, 1968, Let It Be
	George Harrison, 1968, While My Guitar Gently Weeps

	By year, author, title:
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, 12-Bar Original
	John Lennon, 1965, Let It Be
	Ringo Star, 1965, Untitled
	George Harrison, 1968, While My Guitar Gently Weeps
	John Lennon, 1968, Let It Be

	By title, author, year:
	John Lennon, 1965, 12-Bar Original
	Paul McCartney, 1963, All My Loving
	John Lennon, 1965, Let It Be
	John Lennon, 1968, Let It Be
	Ringo Star, 1965, Untitled
	George Harrison, 1968, While My Guitar Gently Weeps
	Unknown column header

### Documentation

See [godoc](https://godoc.org/github.com/stanim/sortstr) for more documentation and examples.

### License

Released under the [MIT License](https://github.com/stanim/sortstr/blob/master/LICENSE).
