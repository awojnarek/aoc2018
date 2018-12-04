package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Open our file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create our variables to house things
	var inputSlice []string

	// Create some ints to store overall counts.
	// These should only increment by 1.
	var doubleAmount int
	var tripleAmount int

	// Populate our slice
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inputSlice = append(inputSlice, scan.Text())
	}

	// Loop over our slice
	for _, i := range inputSlice {

		// These are our ints that house the number of doubles, inside each string, we only care about discrete entries.
		// I.E if the string has two p's and two o's then only count double as one, and increment doubleAmount once.
		var doubleInt int
		var tripleInt int

		// Loop over our string one character at a time
		for _, n := range strings.Split(i, "") {

			// If the string contains exactly two counts, then add it to our double count
			if strings.Count(i, n) == 2 {
				doubleInt++
				continue
			}

			// If the string contains exactly three counts, then add it to our triple count
			if strings.Count(i, n) == 3 {
				tripleInt++
			}
		}

		if doubleInt >= 1 {
			doubleAmount++
		}

		if tripleInt >= 1 {
			tripleAmount++
		}
	}
	fmt.Println(doubleAmount, tripleAmount)
}
