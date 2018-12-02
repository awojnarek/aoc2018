package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	m := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
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
			os.Exit(0)
		}

		// It's not - add it to our map
		m[count] = 1
	}
}
