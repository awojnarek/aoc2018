package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func offByOne(a string, b string) bool {
	arr1 := strings.Split(a, "")
	arr2 := strings.Split(b, "")
	count := 0

	for index, string := range arr1 {
		if count > 1 {
			return false
		}
		if arr2[index] == string {
			continue
		} else {
			count++
		}
	}

	if count == 1 {
		return true
	}

	return false
}

func findAlikeString(x []string) string {
	str1 := x[0]
	str2 := x[1]
	fmt.Println(str1)

	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")

	for index, val := range arr1 {
		if arr2[index] != val {
			fmt.Println("found:", arr2[index], val)
			arr1 = arr1[:index+copy(arr1[index:], arr1[index+1:])]
			return strings.Join(arr1[:], "")
		}
	}
	return "false"
}

func main() {

	// Open our file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create our variables to house things
	var inputSlice []string
	var outputSlice []string

	// Populate our slice
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inputSlice = append(inputSlice, scan.Text())
	}

	// Loop over our slice
	for _, i := range inputSlice {
		if len(outputSlice) != 0 {
			break
		}
		for _, n := range inputSlice {
			if i == n {
				continue
			}

			x := offByOne(i, n)
			if x == true {
				outputSlice = append(outputSlice, i, n)
			}
		}
	}

	stringAlike := findAlikeString(outputSlice)
	fmt.Println(stringAlike)
}
