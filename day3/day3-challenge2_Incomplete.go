package main

/*
INCOMPLETE SOLUTION

Advent of Code 2023

Setup the Environment

go env -w GOROOT="/usr/lib/go"
go env -w GOPATH="/home/thepcn3rd/go/workspaces/AoC2023/day3"

Make the directories - src
Copy the commonFunctions folder into the src directory so that it can be referenced

// To cross compile for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o day3.bin -ldflags "-w -s" main.go


*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type mapsPieceStruct struct {
	CoordX      int
	CoordY      int
	Value       string
	IntValue    int // If the value is a number convert to an integer for math
	SpecialChar bool
	Adjacent    bool // True if adjacent to a symbol, False if not adjacent
	Gear        bool // If the Value an asterisk
	GearValue   int
	Gear1Value  int // If Gear what is the closest Gear1Value
	Gear2Value  int // If Gear what is the closest Gear2Value
}

func mapsInitPiece(p mapsPieceStruct) mapsPieceStruct {
	p.CoordX = -1
	p.CoordY = -1
	p.Value = "-"
	p.IntValue = 0
	p.SpecialChar = false
	p.Adjacent = false
	p.Gear = false
	p.GearValue = -1
	p.Gear1Value = -1
	p.Gear2Value = -1
	return p
}

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

func main() {
	var xyCoordMap [140][140]mapsPieceStruct
	//xyCoordMap[0][0] = PieceStruct{0, 0, ".", -1, false, false}
	//xyCoordMap[1][0] = PieceStruct{0, 1, ".", -1, false, false}
	//fmt.Println(xyCoordMap[0])

	file, err := os.Open("input.txt")
	checkError("Unable to open file", err, true)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var y, maxX, maxY int
	y = 0 // Y Coordinate
	maxX = 0
	maxY = 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\n", "", -1)
		// Build the xy Coorindate Map reading one character at a time
		for x, char := range line {
			fmt.Printf("%c", char)
			var piece mapsPieceStruct
			piece.CoordX = x
			piece.CoordY = y
			piece = mapsInitPiece(piece)
			piece.Value = string(char)
			//fmt.Printf(piece.Value)
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
			if piece.Value == "*" {
				piece.Gear = true
			}
			// Get the maxX prior to calculating adjacency
			if x > maxX {
				maxX = x
			}
			// Don't make the map unless the line count changes or the Y cooridnate increases
			//if maxY < y || maxY == 0 {
			//	xyCoordMap[y] = make(map[int]mapsPieceStruct)
			//}
			//fmt.Println(x)
			//xyCoordMap[y][x] = mapsPieceStruct{piece.Value, piece.IntValue, piece.SpecialChar, piece.Adjacent, piece.Gear, piece.Gear1Value, piece.Gear2Value}
			xyCoordMap[y][x] = piece
			//xyCoordMap[y] = make(map[int]mapsPieceStruct)
			//fmt.Println(piece)
			//xyCoordMap[y][x].Value = piece.Value
			//fmt.Println(xyCoordMap[y][x].Value)
		}
		// Get the maxY
		if y > maxY {
			maxY = y
		}
		// Increment y on every new line read
		y += 1
		fmt.Printf("\n")
	}
	fmt.Printf("\nDimensions of Puzzle: \n")
	fmt.Printf("\tminX - 0, maxX - %d\n", maxX)
	fmt.Printf("\tminY - 0, maxY - %d\n\n", maxY)
	// Testing the output
	//fmt.Println(xyCoordMap[1][5])

	// Determine if the piece is adjacent to a special character
	// Determine Algorithm to Determine if a Special Character is adjacent to numbers
	for y = 0; y < maxY; y++ {
		//fmt.Println(y)
		for x := 0; x < maxX; x++ {
			// Verifying the output of the values in the struct...
			fmt.Printf("%s", xyCoordMap[y][x].Value)
			// Priority: 1 -- x - 1 unless x = 0 AND y - 1 unless y = 0 (Left and Up)
			// If the current character is (an Integer or an asterisk) and it is inbounds of the puzzle
			if (xyCoordMap[y][x].IntValue >= 0 || xyCoordMap[y][x].Gear == true) && x-1 >= 0 && y-1 >= 0 {
				// If the adjacent character is a special character then change adjacent to true
				if xyCoordMap[y-1][x-1].SpecialChar == true {
					xyCoordMap[y][x].Adjacent = true
				}
				// Identify if the numerical value is adjacent to a gear
				gearValue := 0
				if xyCoordMap[y][x].Gear == true && xyCoordMap[y-1][x-1].IntValue >= 0 {
					// Possible 3 digit number
					//  x x x C _ _ _ 	C - Verify this is not a number for a 2 digit number
					//	_ _ _ * _ _ _
					//  _ _ _ _ _ _ _

					// New Idea...

					/*
						if x-3 >= 0 && xyCoordMap[y-1][x].IntValue < 0 {
							//  x x x C _ _ _ 	C - Verify this is not a number for a 2 digit number
							//	_ _ _ * _ _ _
							//  _ _ _ _ _ _ _
							if xyCoordMap[y-1][x-3].IntValue >= 0 && xyCoordMap[y-1][x-2].IntValue >= 0 {
								gearValue = (xyCoordMap[y-1][x-3].IntValue * 100) + (xyCoordMap[y-1][x-2].IntValue * 10) + (xyCoordMap[y-1][x-1].IntValue * 1)
								//  N x x C _ _ _   N - Verify this is not a number
								//	_ _ _ * _ _ _
								//  _ _ _ _ _ _ _
							} else if xyCoordMap[y-1][x-2].IntValue >= 0 {
								gearValue = (xyCoordMap[y-1][x-2].IntValue * 10) + (xyCoordMap[y-1][x-1].IntValue * 1)
								//  _ N x C _ _ _   N - Verify this is not a number
								//	_ _ _ * _ _ _
								//  _ _ _ _ _ _ _
							} else {
								gearValue = (xyCoordMap[y-1][x-1].IntValue * 1)
							}
						}
						if x-2 >= 0 && x == maxX {
							//  E x x x E _ _   E - Edge of the Puzzle (What if C is the Edge?)
							//	_ _ _ * _ _ _
							//  _ _ _ _ _ _ _
							if xyCoordMap[y-1][x-2].IntValue >= 0 && xyCoordMap[y-1][x].IntValue >= 0 {
								gearValue = (xyCoordMap[y-1][x-2].IntValue * 100) + (xyCoordMap[y-1][x-1].IntValue * 10) + (xyCoordMap[y-1][x].IntValue * 1)
							//  E N x x E _ _
							//	_ _ _ * _ _ _
							//  _ _ _ _ _ _ _
							} else if xyCoordMap[y-1][x-2].IntValue < 0 && xyCoordMap[y-1][x].IntValue >= 0 {
								gearValue = (xyCoordMap[y-1][x-1].IntValue * 10) + (xyCoordMap[y-1][x].IntValue * 1)
							}
						}
						if x+1 <= maxX && x+2 <= maxX {
							//  _ C x x x C _
							//	_ _ _ * _ _ _
							//  _ _ _ _ _ _ _
							if xyCoordMap[y-1][x].IntValue >= 0 && xyCoordMap[y-1][x+1].IntValue >= 0 {
								gearValue = (xyCoordMap[y-1][x-1].IntValue * 100) + (xyCoordMap[y-1][x].IntValue * 10) + (xyCoordMap[y-1][x+1].IntValue * 1)
							}
						}
					*/
				}
				if gearValue > 0 {
					fmt.Printf("\nGear: %d\n", gearValue)
				}
			}
			// Priority: 2 -- y - 1 unless y = 0 (Up) AND X = X

			// Priority: 3 -- x + 1 unless x = maxX AND y - 1 unless y = 0 (Right and Up)

			// Priority: 4 -- x - 1 unless x = 0 (Left)

			// Priority: 5 -- x + 1 unless x = maxX (Right)

			// Priority: 6 -- x - 1 unless x = 0 AND y + 1 unless y = maxY (Left and Down)

			// Priority: 7 -- y + 1 unless y - maxY (Down)

			// Priority: 8 -- x + 1 unless x = maxX AND y + 1 unless y = maxY (Right and Down)
		}
		fmt.Printf("\n")
	}

}
