package day5

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"strconv"
	"strings"
)

func ExtractSeeds(mapSlice []string) []int {
	seeds := strings.SplitAfter(mapSlice[0], ":")
	seeds = strings.Fields(seeds[1])

	return ConvertStringsToInt(seeds)
}

func ConvertStringsToInt(stringsSlice []string) []int {
	var numbers []int
	for _, element := range stringsSlice {
		number, _ := strconv.Atoi(element)
		numbers = append(numbers, number)
	}

	return numbers
}

func GetNumbersForEachMap(mapSlice []string) [][]string {
	var mapConverter [][]string

	for _, mapNumbers := range mapSlice {
		mapConverter = append(mapConverter, strings.Split(mapNumbers, "\n")[1:])
	}

	return mapConverter
}

func ConvertMapDataToInt(maps [][]string) [][][]int {
	var mapsData = make([][][]int, len(maps))
	for mapIndex, mapsElements := range maps {
		mapsData[mapIndex] = make([][]int, len(mapsElements))
		for elementIndex, element := range mapsElements {
			digits := strings.Fields(element)
			numbers := ConvertStringsToInt(digits)
			mapsData[mapIndex][elementIndex] = numbers
		}
	}
	return mapsData
}

func GetSeedLocationNumber(seed int, mapsNumbers [][][]int) int {
	actualNumber := seed
	for _, mapSequence := range mapsNumbers {
		actualNumber = GetSeedCategoryNumber(mapSequence, actualNumber)
	}
	return actualNumber
}

func GetSeedCategoryNumber(mapSequence [][]int, actualNumber int) int {
	for _, sequence := range mapSequence {
		categoryNumber := CalcSeedMapping(actualNumber, sequence)
		if categoryNumber != actualNumber {
			return categoryNumber
		}
	}
	return actualNumber
}

func CalcSeedMapping(actualNumber int, mapNumbers []int) int {
	destinationRangeStart := mapNumbers[0]
	sourceRangeStart := mapNumbers[1]
	rangeLength := mapNumbers[2]

	if actualNumber >= sourceRangeStart && actualNumber <= sourceRangeStart+rangeLength-1 {
		if actualNumber == sourceRangeStart {
			return destinationRangeStart
		}
		return destinationRangeStart + actualNumber - sourceRangeStart
	}
	return actualNumber
}

func Day5(filePath string) {
	fileContent := utils.FileToString(filePath)
	lowestLocationNumber := 0

	mapSlice := strings.Split(fileContent, "\n\n")
	seeds := ExtractSeeds(mapSlice)

	mapSlice = mapSlice[1:]
	maps := GetNumbersForEachMap(mapSlice)
	data := ConvertMapDataToInt(maps)

	for _, seed := range seeds {
		locationNumber := GetSeedLocationNumber(seed, data)

		if lowestLocationNumber == 0 || locationNumber < lowestLocationNumber {
			lowestLocationNumber = locationNumber
		}
	}
	fmt.Println(lowestLocationNumber)
}
