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
Advent of Code 2023

Setup the Environment

go env -w GOROOT="/usr/lib/go"
go env -w GOPATH="/home/thepcn3rd/go/workspaces/AoC2023/day3"

Make the directories - src
Copy the commonFunctions folder into the src directory so that it can be referenced

// To cross compile for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o day3.bin -ldflags "-w -s" main.go


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

type PuzzleStruct struct {
	Pieces []PieceStruct
}

type PieceStruct struct {
	CoordX      int
	CoordY      int
	Value       string
	IntValue    int // If the value is a number convert to an integer for math
	SpecialChar bool
	Adjacent    bool // True if adjacent to a symbol, False if not adjacent
}

func initPiece(p PieceStruct) PieceStruct {
	p.CoordX = 0
	p.CoordY = 0
	p.Value = "-"
	p.IntValue = 0
	p.SpecialChar = false
	p.Adjacent = false
	return p
}

func main() {
	// Created testCaseInput.txt to include a number that is a single digit and a 3 digit number at the end of a line
	file, err := os.Open("input.txt")
	checkError("Unable to open file", err, true)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var puzzle PuzzleStruct
	var x, maxX int
	var y, maxY int
	x = 0
	y = 0
	maxX = 0
	maxY = 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\n", "", -1)
		// Read one character at a time
		for _, char := range line {
			fmt.Printf("%c", char)
			var piece PieceStruct
			piece = initPiece(piece)
			piece.CoordX = x
			piece.CoordY = y
			piece.Value = string(char)
			// Determine if the value is an integer and then convert it to an intValue
			if unicode.IsDigit(char) {
				piece.IntValue, _ = strconv.Atoi(string(char))
			} else {
				piece.IntValue = -1 // Place a -1 for the value if it does not exist
			}
			// Determine if the value is a special character
			if unicode.IsDigit(char) || string(char) == "." {
				piece.SpecialChar = false // This is not needed but left for clarity
			} else {
				piece.SpecialChar = true
			}
			// Get the maxX prior to calculating adjacency
			if x > maxX {
				maxX = x
			}

			puzzle.Pieces = append(puzzle.Pieces, piece)

			// Increment x on every new character read unless it is the newline char
			x += 1
		}
		// Get the maxY
		if y > maxY {
			maxY = y
		}
		// Increment y on every new line read
		y += 1
		x = 0 // Reset x on a line that is read
		fmt.Printf("\n")
	}
	fmt.Printf("\nDimensions of Puzzle: \n")
	fmt.Printf("\tminX - 0, maxX - %d\n", maxX)
	fmt.Printf("\tminY - 0, maxY - %d\n\n", maxY)

	// Determine if the piece is adjacent to a special character
	// Determine Algorithm to Determine if a Special Character is adjacent to numbers
	for p := range puzzle.Pieces {
		for q := range puzzle.Pieces { // Current X, Y
			// Priority: 1 -- x - 1 unless x = 0 AND y - 1 unless y = 0 (Left and Up)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX - 1)) && (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY - 1)) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 2 -- y - 1 unless y = 0 (Up) AND X = X
			if (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY - 1)) && puzzle.Pieces[q].CoordX == puzzle.Pieces[p].CoordX && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 3 -- x + 1 unless x = maxX AND y - 1 unless y = 0 (Right and Up)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX + 1)) && (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY - 1)) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 4 -- x - 1 unless x = 0 (Left)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX - 1)) && (puzzle.Pieces[q].CoordY == puzzle.Pieces[p].CoordY) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 5 -- x + 1 unless x = maxX (Right)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX + 1)) && (puzzle.Pieces[q].CoordY == puzzle.Pieces[p].CoordY) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 6 -- x - 1 unless x = 0 AND y + 1 unless y = maxY (Left and Down)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX - 1)) && (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY + 1)) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 7 -- y + 1 unless y - maxY (Down)
			if (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY + 1)) && puzzle.Pieces[q].CoordX == puzzle.Pieces[p].CoordX && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
			// Priority: 8 -- x + 1 unless x = maxX AND y + 1 unless y = maxY (Right and Down)
			if (puzzle.Pieces[q].CoordX == (puzzle.Pieces[p].CoordX + 1)) && (puzzle.Pieces[q].CoordY == (puzzle.Pieces[p].CoordY + 1)) && puzzle.Pieces[p].SpecialChar == true {
				puzzle.Pieces[q].Adjacent = true
			}
		}
	}

	// Calculate the sum of the pieces that are not adjacent
	var sumPieces int
	sumPieces = 0

	hundredsPlaceInt := 0
	hundredsPlaceAdjacent := false
	tensPlaceInt := 0
	tensPlaceAdjacent := false
	onesPlaceInt := 0
	onesPlaceAdjacent := false
	lastCoordY := 0
	sameNumberFactor := 0
	for p := range puzzle.Pieces {
		// Is the number a 3 digit number
		// Use case is 467. on the first line
		if onesPlaceInt >= 0 && tensPlaceInt >= 0 && hundredsPlaceInt > 0 && sameNumberFactor == 0 {
			calculateNumber := (onesPlaceInt * 1) + (tensPlaceInt * 10) + (hundredsPlaceInt * 100)
			if onesPlaceAdjacent == true || tensPlaceAdjacent == true || hundredsPlaceAdjacent == true {
				// Sum the pieces that are adjacent
				sumPieces += calculateNumber
				fmt.Printf("Calculated Number: %d - Adjacent - Running Total: %d\n", calculateNumber, sumPieces)
			} else {
				fmt.Printf("Calculated Number: %d\n", calculateNumber)

			}
			sameNumberFactor = 3
			// Is the number a 2 digit number
		} else if onesPlaceInt < 0 && tensPlaceInt >= 0 && hundredsPlaceInt > 0 && sameNumberFactor == 0 {
			calculateNumber := (tensPlaceInt * 1) + (hundredsPlaceInt * 10)
			if tensPlaceAdjacent == true || hundredsPlaceAdjacent == true {
				// Sum the pieces that are adjacent
				sumPieces += calculateNumber
				fmt.Printf("Calculated Number: %d - Adjacent - Running Total: %d\n", calculateNumber, sumPieces)
			} else {
				fmt.Printf("Calculated Number: %d\n", calculateNumber)
			}
			sameNumberFactor = 2
		} else if onesPlaceInt < 0 && tensPlaceInt < 0 && hundredsPlaceInt > 0 && sameNumberFactor == 0 {
			calculateNumber := (hundredsPlaceInt * 1)
			if hundredsPlaceAdjacent == true {
				// Sum the pieces that are adjacent
				sumPieces += calculateNumber
				fmt.Printf("Calculated Number: %d - Adjacent - Running Total: %d\n", calculateNumber, sumPieces)
			} else {
				fmt.Printf("Calculated Number: %d\n", calculateNumber)
			}
			sameNumberFactor = 1
		}

		// Is the number a single digit number
		// Calculate the lastCoordY to determine end of line
		hundredsPlaceInt = tensPlaceInt
		hundredsPlaceAdjacent = tensPlaceAdjacent
		tensPlaceInt = onesPlaceInt
		tensPlaceAdjacent = onesPlaceAdjacent
		onesPlaceInt = puzzle.Pieces[p].IntValue
		onesPlaceAdjacent = puzzle.Pieces[p].Adjacent
		//lastCoordY = puzzle.Pieces[p].CoordY
		if sameNumberFactor > 0 {
			sameNumberFactor -= 1
		}
		// Use case 879..................54.100..355 - The 54 is not recorded
		if puzzle.Pieces[p].IntValue > 0 {
			sameNumberFactor = 0
		}
		// Use case if a number is at the end of a line and a new number starts at the beginning of the next line
		// Calculated Number: 219 - Adjacent - Running Total: 436773
		//** Calculated Number: 635
		//Calculated Number: 334

		//.................%........*.216.831....................509.....392............508...........&...........780.........219.....*......&....*635
		//334............&........862................945.....162*.........*.......529.......164.......627.381*........591........*83...........155....
		// Resolved the above due to missing 635 being adjacent, however it lists in the debug 353 and 533 erronously
		if puzzle.Pieces[p].CoordY == (lastCoordY + 1) {
			sameNumberFactor = 0
		}
		lastCoordY = puzzle.Pieces[p].CoordY
		// Another issue is it does not list the last number of the puzzle...
	}
	fmt.Printf("\nSum of Pieces that are Adjacent: %d\n\n", sumPieces)
	//fmt.Println(puzzle)
}
