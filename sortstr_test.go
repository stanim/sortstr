package sortstr_test

import (
	"fmt"

	"github.com/stanim/sortstr"
)

// ExampleBy demonstrates multisorting of string slices.
func ExampleBy() {
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

	// Output:
	//
	// By author, title, year:
	// Ringo Star
	// Ringo Star, 1965, Untitled
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, 12-Bar Original
	// John Lennon, 1965, Let It Be
	// John Lennon, 1968, Let It Be
	// George Harrison, 1968, While My Guitar Gently Weeps
	//
	// By year, author, title:
	// Ringo Star
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, 12-Bar Original
	// John Lennon, 1965, Let It Be
	// Ringo Star, 1965, Untitled
	// George Harrison, 1968, While My Guitar Gently Weeps
	// John Lennon, 1968, Let It Be
	//
	// By title, author, year:
	// Ringo Star
	// John Lennon, 1965, 12-Bar Original
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, Let It Be
	// John Lennon, 1968, Let It Be
	// Ringo Star, 1965, Untitled
	// George Harrison, 1968, While My Guitar Gently Weeps
}

// ExampleByHeaders demonstrates multisorting of string
// slices with column headers instead of indices.
func ExampleByHeaders() {
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

	// Output:
	//
	// By -author, title, year:
	// Ringo Star, 1965, Untitled
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, 12-Bar Original
	// John Lennon, 1965, Let It Be
	// John Lennon, 1968, Let It Be
	// George Harrison, 1968, While My Guitar Gently Weeps
	//
	// By year, author, title:
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, 12-Bar Original
	// John Lennon, 1965, Let It Be
	// Ringo Star, 1965, Untitled
	// George Harrison, 1968, While My Guitar Gently Weeps
	// John Lennon, 1968, Let It Be
	//
	// By title, author, year:
	// John Lennon, 1965, 12-Bar Original
	// Paul McCartney, 1963, All My Loving
	// John Lennon, 1965, Let It Be
	// John Lennon, 1968, Let It Be
	// Ringo Star, 1965, Untitled
	// George Harrison, 1968, While My Guitar Gently Weeps
	// Unknown column header
}
