// To execute Go code, please declare a func main() in a package "main"
// .      0          1.   2. ....
// [["snowflake", "is", "one", "of", "the", "best", "companies"], <- 0
//
//	["weather", "is", "good", "today"], <- 1
//	["snowflake", "stock", "is", "up", "today"]] <- 2
//
// func index(x [][]input)
//
// func search(word string)
// snowflake -> [[0, 0], [2, 0]]
// is -> [[0, 1], [1, 1], [2, 2]
// invalid -> []
//
// func search2(word1 string, word2 string)
// snowflake, is -> [[0, 0]]
// snowflake -> results [[,],[,],...]
// is -> adj + 1
// good, today -> [[1, 2]]
// invalid invalid -> []
//
// func searchN(words []string)
// one of the best
// stock is up
// ...
package main

import "fmt"

type Point struct {
	row int
	col int
}

// Index1 simple implementation
type Index1 struct {
	data map[string][]Point
}

func Index1Make() *Index1 {
	return &Index1{data: make(map[string][]Point)}
}

func (i Index1) Index(input [][]string) {
	for row, words := range input {
		for col, word := range words {
			points, found := i.data[word]
			if !found {
				points = make([]Point, 0)
				i.data[word] = points
			}
			i.data[word] = append(points, Point{row: row, col: col})
		}
	}
}

func (i Index1) Search(word string) []Point {
	return i.data[word]
}

// Index2 faster implementation for multi-word searches
type Index2 struct {
	data map[string]map[int]map[int]bool
}

func Index2Make() *Index2 {
	return &Index2{data: make(map[string]map[int]map[int]bool)}
}

func (i Index2) Index(input [][]string) {
	for row, words := range input {
		for col, word := range words {
			points, found := i.data[word]
			if !found {
				points = make(map[int]map[int]bool)
				i.data[word] = points
			}
			cols, foundCols := points[row]
			if !foundCols {
				cols = make(map[int]bool)
				points[row] = cols
			}

			cols[col] = true
		}
	}
}

func (i Index2) Search(word1, word2 string) []Point {
	results := make([]Point, 0)
	word1Points, found1 := i.data[word1]
	if !found1 {
		return nil
	}
	for row, cols := range word1Points {
		for col := range cols {
			if word2 == "" {
				results = append(results, Point{row: row, col: col})
				continue
			}

			word2found := i.data[word2][row][col+1]
			if word2found {
				results = append(results, Point{row: row, col: col})
			}
		}
	}

	return results
}

func main() {
	var input [][]string = [][]string{
		{"snowflake", "is", "one", "of", "the", "best", "companies"},
		{"weather", "is", "good", "today"},
		{"snowflake", "stock", "is", "up"}}

	fmt.Printf("\n\nSearch1:\n")
	i1 := Index1Make()
	i1.Index(input)
	fmt.Printf("snowflake: %v\n", i1.Search("snowflake"))
	fmt.Printf("is       : %v\n", i1.Search("is"))
	fmt.Printf("nadda       : %v\n", i1.Search("nadda"))

	fmt.Printf("\n\nSearch2:\n")
	i2 := Index2Make()
	i2.Index(input)
	fmt.Printf("snowflake is: %v\n", i2.Search("snowflake", "is"))
	fmt.Printf("snowflake   : %v\n", i2.Search("snowflake", ""))
	fmt.Printf("is          : %v\n", i2.Search("is", ""))
	fmt.Printf("good today  : %v\n", i2.Search("good", "today"))
	fmt.Printf("nadda       : %v\n", i2.Search("nada", "today"))
}

// Your previous Plain Text content is preserved below:

// This is just a simple shared plaintext pad, with no execution capabilities.

// When you know what language you'd like to use for your interview,
// simply choose it from the dots menu on the tab, or add a new language
// tab using the Languages button on the left.

// You can also change the default language your pads are created with
// in your account settings: https://app.coderpad.io/settings

// Enjoy your interview!
