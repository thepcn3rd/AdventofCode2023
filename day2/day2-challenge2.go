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
go env -w GOPATH="/home/thepcn3rd/go/workspaces/AoC2023/day2"

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

type gameStruct struct {
	Number        int
	CubeSets      setsStruct
	Possible      bool
	RedCount      int
	BlueCount     int
	GreenCount    int
	MinRedCount   int
	MinBlueCount  int
	MinGreenCount int
	CubePower     int
}

type setsStruct struct {
	Sets []string
}

func main() {
	file, err := os.Open("input.txt")
	checkError("Unable to open file", err, true)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	gameIDSum := 0
	gameCubePowerSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var gs gameStruct
		gs.Possible = true
		// Split Game with Number with Cubes
		lineParts := strings.Split(line, ":")
		fmt.Printf("****************************\nLine: %s\n", line)
		gs.Number, err = strconv.Atoi((strings.Split(lineParts[0], " ")[1]))
		gs.CubeSets.Sets = append(gs.CubeSets.Sets, lineParts[1])
		checkError("Unable to convert game number to an integer", err, true)
		// Minimum number of cubes of each color for the game
		gs.MinBlueCount = 0
		gs.MinGreenCount = 0
		gs.MinRedCount = 0
		setParts := strings.Split(lineParts[1], ";")
		for i := range setParts {
			fmt.Printf("\tSet: %s\n", setParts[i])
			cubeParts := strings.Split(setParts[i], ",")
			for j := range cubeParts {
				// Looks at each set individually
				gs.RedCount = 0
				gs.BlueCount = 0
				gs.GreenCount = 0
				fmt.Printf("\t\tCube Part: %s\n", cubeParts[j])
				// When splitting the string the number has a space in the front
				cubeKV := strings.Split(cubeParts[j], " ")
				//fmt.Println(cubeKV[2])
				cubeValue, err := strconv.Atoi(cubeKV[1])
				checkError("Unable to convert the count to an integer", err, true)

				// Total Counts of the Colors
				switch cubeKV[2] {
				case "red":
					gs.RedCount += cubeValue
					// Total Count
					if gs.RedCount > 12 {
						fmt.Printf("\t\t\tTotal Red Count: %d\n", gs.RedCount)
						gs.Possible = false
					}
					// Calculate the mimumun cubes needed
					if gs.MinRedCount < cubeValue {
						gs.MinRedCount = cubeValue
					}
				case "blue":
					gs.BlueCount += cubeValue
					if gs.BlueCount > 14 {
						fmt.Printf("\t\t\tTotal Blue Count: %d\n", gs.BlueCount)
						gs.Possible = false
					}
					// Calculate the mimumun cubes needed
					if gs.MinBlueCount < cubeValue {
						gs.MinBlueCount = cubeValue
					}
				case "green":
					gs.GreenCount += cubeValue
					if gs.GreenCount > 13 {
						fmt.Printf("\t\t\tTotal Green Count: %d\n", gs.GreenCount)
						gs.Possible = false
					}
					// Calculate the mimumun cubes needed
					if gs.MinGreenCount < cubeValue {
						gs.MinGreenCount = cubeValue
					}
				}
			}
		}

		if gs.Possible == true {
			fmt.Printf("Game: %d - Possible\n\n", gs.Number)
			gameIDSum += gs.Number
		} else {
			fmt.Printf("Game: %d - NOT Possible\n\n", gs.Number)
		}

		// Calculate Cube Power
		gs.CubePower = gs.MinBlueCount * gs.MinGreenCount * gs.MinRedCount
		fmt.Printf("Game Cube Power: %d\n", gs.CubePower)
		gameCubePowerSum += gs.CubePower
	}
	fmt.Printf("\n******************************\nGame ID Sum: %d\n", gameIDSum)
	fmt.Printf("Game Cube Power Sum: %d\n", gameCubePowerSum)
}
