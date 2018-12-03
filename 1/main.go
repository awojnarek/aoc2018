package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open our file
	file, err := os.Open("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create our variables to house things
	var frequencies []int
	var trackerFreqs []int
	tracker := 0

	// Populate our slice
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, _ := strconv.Atoi(scan.Text())
		frequencies = append(frequencies, i)
	}

	// Loop over our slice and do work, keep going until we find what we need
	for count := 0; count <= len(frequencies); count++ {

		// If we reached the end of the slice, repeat by setting count to 0
		if count == len(frequencies) {
			count = 0
		}

		// Define our integer
		integer := frequencies[count]

		// Add the last with our running tracker
		tracker = integer + tracker

		// Let's loop through the slice and see if we have it
		for _, s := range trackerFreqs {
			if s == tracker {
				fmt.Println("Found first dupe:", tracker)
				os.Exit(0)
			}
		}

		// If we looped and didn't find it, record our math and keep looping baby.
		trackerFreqs = append(trackerFreqs, tracker)
	}

	/*
		for scanner.Scan() {
			fmt.Println("shit")
			// Convert to integer
			input := scanner.Text()
			integer, _ := strconv.Atoi(input)

			// Add to our existing count
			count = count + integer

			// Check to see if this new count is already in our map
			_, present := m[count]

			// If it's in our map already, break
			if present == true {
				fmt.Println(count)
				fmt.Println("Exiting")
				os.Exit(0)
			}

			// It's not - add it to our map
			m[count] = 1
		}
	*/
}
