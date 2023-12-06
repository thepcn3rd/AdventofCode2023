package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Advent of Code 2023

Setup the Environment

go env -w GOROOT="/usr/lib/go"
go env -w GOPATH="/home/thepcn3rd/go/workspaces/AoC2023/day4"

Make the directories - src
Copy the commonFunctions folder into the src directory so that it can be referenced

// To cross compile for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o day2.bin -ldflags "-w -s" main.go


*/

// CheckError checks for errors
func checkError(reasonString string, err error, exitBool bool) {
	if err != nil && exitBool == true {
		fmt.Printf("%s\n", reasonString)
		//fmt.Printf("%s\n\n", err)
		os.Exit(0)
	} else if err != nil && exitBool == false {
		fmt.Printf("%s\n", reasonString)
		//fmt.Printf("%s\n", err)
		return
	}
}

type CardStruct struct {
	CardName       string
	WinningNumbers string
	CardNumbers    string
	WinningInt     []int
	CardInt        []int
	TotalMatches   int
	CardPoints     int64
}

func convertStringToIntList(numberString string) []int {
	var listInt []int
	listNumbers := strings.Split(numberString, " ")
	//fmt.Println(arrayWinningNumbers)
	for _, val := range listNumbers {
		//fmt.Printf("Value: '%s'\n", val)
		if val != "" {
			//fmt.Println(strings.Replace(val, " ", "", -1))
			i, err := strconv.Atoi(strings.Replace(val, " ", "", -1))
			//fmt.Printf("%d\n", i)
			checkError("Unable to convert the integer", err, true)
			listInt = append(listInt, i)
		}
		//fmt.Println(Card.WinningInt)
	}
	return listInt
}

func main() {
	var Cards []CardStruct
	var TotalPoints int
	file, err := os.Open("input.txt")
	checkError("Unable to open file", err, true)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var Card CardStruct
		line := scanner.Text()
		//fmt.Println(line)
		getCardName := strings.Split(line, ":")
		Card.CardName = getCardName[0]
		getCards := strings.Split(getCardName[1], "|")
		Card.WinningNumbers = getCards[0]
		Card.CardNumbers = getCards[1]
		Card.WinningInt = convertStringToIntList(Card.WinningNumbers)
		Card.CardInt = convertStringToIntList(Card.CardNumbers)
		for _, W := range Card.WinningInt {
			for _, C := range Card.CardInt {
				if W == C {
					Card.TotalMatches += 1
				}
			}
		}
		/*
			switch Card.TotalMatches {
			case 1:
				Card.CardPoints = 1
			case 2:
				Card.CardPoints = 2
			case 3:
				Card.CardPoints = 4
			}
		*/
		// Convert the above switch statement to a binary number for faster calculation
		var binaryNumber string
		if Card.TotalMatches > 1 {
			binaryNumber = "1" + strings.Repeat("0", (Card.TotalMatches-1))
		} else if Card.TotalMatches == 1 {
			binaryNumber = "1"
		}
		Card.CardPoints, _ = strconv.ParseInt(binaryNumber, 2, 64)
		fmt.Printf("%s - Card Points: %d\n", Card.CardName, Card.CardPoints)
		TotalPoints += int(Card.CardPoints)
		Cards = append(Cards, Card)
	}
	fmt.Println(TotalPoints)
}
