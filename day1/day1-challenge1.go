package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
Setup the Environment

go env -w GOROOT="/usr/lib/go"
go env -w GOPATH="/home/thepcn3rd/go/workspaces/AoC2023/day1"

Make the directories - src
Copy the commonFunctions folder into the src directory so that it can be referenced

// To cross compile for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o day1.bin -ldflags "-w -s" main.go


*/

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var sumInteger int
	sumInteger = 0
	// Read line by line
	for scanner.Scan() {
		var firstIntegerChar string
		var lastIntegerChar string
		var firstInteger int
		var lastInteger int
		var valueInteger int

		firstIntegerChar = ""
		lastIntegerChar = ""
		line := scanner.Text()
		line = strings.Replace(line, "\n", "", 1)
		//fmt.Println(line) // Process the line here, e.g., print it
		for _, char := range line {
			//fmt.Printf("%c\n", char)
			if unicode.IsDigit(char) {
				//fmt.Printf("%c is a number\n", char)
				if firstIntegerChar == "" {
					firstIntegerChar = fmt.Sprintf("%c", char)
				} else {
					lastIntegerChar = fmt.Sprintf("%c", char)
				}
			}
		}
		// After the line is scanned
		firstInteger, err = strconv.Atoi(firstIntegerChar)
		if err != nil {
			fmt.Println("Error converting first character to integer:", err)
			return
		}
		// Accomodates for only 1 integer being on a line
		if lastIntegerChar != "" {
			lastInteger, err = strconv.Atoi(lastIntegerChar)
			if err != nil {
				fmt.Println("Error converting first character to integer:", err)
				return
			}
		}
		if firstIntegerChar != "" && lastIntegerChar != "" {
			valueInteger = (firstInteger * 10) + lastInteger
		} else {
			valueInteger = (firstInteger * 10) + firstInteger
		}
		fmt.Printf("%d\n", valueInteger)
		sumInteger += valueInteger
	}
	fmt.Printf("%d", sumInteger)

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
