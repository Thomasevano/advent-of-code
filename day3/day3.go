package day3

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"regexp"
	"strconv"
	"unicode"
)

type Coordinates struct {
	X, Y int
}

type NumberCoords struct {
	LineNumber int
	Coords     Coordinates
}

var partNumbers = map[Coordinates][]int{}
var numbers = map[NumberCoords]string{}
var symbolsCoords = map[Coordinates]string{}
var part1, part2 = 0, 0

func GetSymbols(sliceOfLine []string) map[Coordinates]string {
	for sliceIndex, line := range sliceOfLine {
		for charIndex, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				symbolsCoords[Coordinates{sliceIndex, charIndex}] = strconv.QuoteRune(char)
			}
		}
	}
	return symbolsCoords
}

func GetNumbersAdjacentCells(sliceIndex int, line string, numberCoords []int, numbersCoordsWithLineNumber NumberCoords,
) map[Coordinates]string {
	adjacentCells := map[Coordinates]string{}
	for lineIndex := numberCoords[0]; lineIndex < numberCoords[1]; lineIndex++ {
		numbers[numbersCoordsWithLineNumber] += string(line[lineIndex])
		for _, d := range []Coordinates{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		} {
			adjacentCellCoords := Coordinates{sliceIndex + d.X, lineIndex + d.Y}
			adjacentCells[adjacentCellCoords] = ""
		}
	}
	return adjacentCells
}

func Main(filePath string) int {
	sliceOfLine := utils.LinesInFile(filePath)
	symbolsCoords = GetSymbols(sliceOfLine)

	for sliceIndex, line := range sliceOfLine {
		numbersCoords := regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1)

		for _, numberCoords := range numbersCoords {
			numbersCoordsWithLineNumber := NumberCoords{sliceIndex, Coordinates{numberCoords[0], numberCoords[1]}}

			adjacentCells := GetNumbersAdjacentCells(sliceIndex, line, numberCoords, numbersCoordsWithLineNumber)

			number, _ := strconv.Atoi(numbers[numbersCoordsWithLineNumber])
			for cellCoord := range adjacentCells {
				if _, ok := symbolsCoords[cellCoord]; ok {
					partNumbers[cellCoord] = append(partNumbers[cellCoord], number)
					part1 += number
				}
			}
		}
	}
	for symbolCoords, digits := range partNumbers {
		if symbolsCoords[symbolCoords] == "'*'" && len(digits) == 2 {
			part2 += digits[0] * digits[1]
		}
	}
	fmt.Println(partNumbers)
	fmt.Printf("Part 1: %d \n", part1)
	fmt.Printf("Part 2: %d \n", part2)
	//return part1, part2
	return part2
}
