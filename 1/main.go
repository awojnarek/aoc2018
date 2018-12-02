package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		stringOperator := strings.Split(string(input), "")[0]
		stringNumber := input[1:]
		integer, _ := strconv.Atoi(stringNumber)

		if stringOperator == "+" {
			count = count + integer
		}

		if stringOperator == "-" {
			count = count - integer
		}
	}
	fmt.Println(count)
}
