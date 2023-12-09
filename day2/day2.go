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
	var power int
	var sumOfPower int

	for _, line := range sliceOfLine {
		//isGameValid := true
		subsets := GetSubsets(line)

		fewestNbOfDicesOfEachColor := map[string]int{
			"blue":  1,
			"green": 1,
			"red":   1,
		}

		for _, subset := range subsets {
			colorCubesOfSubset := strings.Split(subset, ",")

			diceOfColor := GetMapFromSubsets(colorCubesOfSubset)

			//invalidGame := diceOfColor["red"] > 12 || diceOfColor["green"] > 13 || diceOfColor["blue"] > 14
			//if invalidGame {
			//	isGameValid = false
			//}

			if diceOfColor["green"] > fewestNbOfDicesOfEachColor["green"] && diceOfColor["green"] != 0 {
				fewestNbOfDicesOfEachColor["green"] = diceOfColor["green"]
			}
			if diceOfColor["blue"] > fewestNbOfDicesOfEachColor["blue"] && diceOfColor["blue"] != 0 {
				fewestNbOfDicesOfEachColor["blue"] = diceOfColor["blue"]
			}
			if diceOfColor["red"] > fewestNbOfDicesOfEachColor["red"] && diceOfColor["red"] != 0 {
				fewestNbOfDicesOfEachColor["red"] = diceOfColor["red"]
			}
		}

		power = fewestNbOfDicesOfEachColor["green"] * fewestNbOfDicesOfEachColor["blue"] * fewestNbOfDicesOfEachColor["red"]
		sumOfPower += power
	}
	fmt.Println(sumOfPower)
}
