package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	var result []string
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetSubsets(line string) []string {
	subsets := strings.SplitAfter(line, ":")
	subsets = subsets[1:]
	subsets = strings.SplitAfter(subsets[0], ";")

	return subsets
}

func GetMapFromSubsets(colorCubesOfSubset []string) map[string]int {
	diceOfColor := make(map[string]int)
	for _, colorCube := range colorCubesOfSubset {
		colorRegex := regexp.MustCompile(`blue|green|red`)
		numberRegex := regexp.MustCompile("[0-9]{1,2}")

		matchColor := colorRegex.FindAllString(colorCube, -1)
		matchNumber := numberRegex.FindAllString(colorCube, -1)

		number, _ := strconv.Atoi(matchNumber[0])

		diceOfColor[matchColor[0]] = number
	}
	return diceOfColor
}

func Main(filePath string) {
	sliceOfLine := LinesInFile(filePath)
	var gameIdSum int

	for index, line := range sliceOfLine {
		isGameValid := true
		subsets := GetSubsets(line)

		for _, subset := range subsets {
			colorCubesOfSubset := strings.Split(subset, ",")

			diceOfColor := GetMapFromSubsets(colorCubesOfSubset)

			invalidGame := diceOfColor["red"] > 12 || diceOfColor["green"] > 13 || diceOfColor["blue"] > 14
			if invalidGame {
				isGameValid = false
			}
		}
		if isGameValid {
			gameIdSum += index + 1
		}
	}
	fmt.Println(gameIdSum)
}
