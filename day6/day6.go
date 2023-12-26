package day6

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile("[0-9]{1,4}")

func FormatRaceData(lines []string) [][]int {
	numbersOnALine := len(numberRegex.FindAllString(lines[0], -1))
	var raceData [][]int
	for i := 0; i < numbersOnALine; i++ {
		numbers := make([]int, 2)
		raceData = append(raceData, numbers)
	}
	for index, line := range lines {
		matchNumbers := numberRegex.FindAllString(line, -1)
		for numberIndex, matchNumber := range matchNumbers {
			number, _ := strconv.Atoi(matchNumber)
			raceData[numberIndex][index] = number
		}
	}
	return raceData
}

func FormatRaceDataPart2(lines []string) []int {
	var raceData []int
	for _, line := range lines {
		matchNumbers := numberRegex.FindAllString(line, -1)
		var concatLine string
		for _, matchNumber := range matchNumbers {
			concatLine += matchNumber
		}
		number, _ := strconv.Atoi(concatLine)
		raceData = append(raceData, number)
	}
	return raceData
}

func calcPart1(RacesTimesDistance [][]int) int {
	var multipliedNumbersOfWaysToBeatRecord int
	for _, raceData := range RacesTimesDistance {
		time := raceData[0]
		distance := raceData[1]

		numberOfWaysToBeatRecord := calcNumbersOfWaysToBeatRecord(time, distance)
		if multipliedNumbersOfWaysToBeatRecord == 0 {
			multipliedNumbersOfWaysToBeatRecord += numberOfWaysToBeatRecord
		} else {
			multipliedNumbersOfWaysToBeatRecord *= numberOfWaysToBeatRecord
		}
	}
	return multipliedNumbersOfWaysToBeatRecord
}

func calcPart2(RaceTimeDistance []int) int {
	time := RaceTimeDistance[0]
	distance := RaceTimeDistance[1]

	numberOfWaysToBeatRecord := calcNumbersOfWaysToBeatRecord(time, distance)
	fmt.Println(numberOfWaysToBeatRecord)
	return numberOfWaysToBeatRecord
}

func calcNumbersOfWaysToBeatRecord(time, distance int) int {
	numberOfWaysToBeatRecord := 0
	for speed := 0; speed <= time; speed++ {
		calculatedTime := time - speed
		calculatedDistance := speed * calculatedTime
		if calculatedDistance > distance {
			numberOfWaysToBeatRecord += 1
		}
	}
	return numberOfWaysToBeatRecord
}

func Day6(filePath string) {
	lines := utils.LinesInFile(filePath)
	part1Races := FormatRaceData(lines)
	part2Race := FormatRaceDataPart2(lines)

	part1Result := calcPart1(part1Races)
	part2Result := calcPart2(part2Race)

	fmt.Printf("Part 1: %d\n", part1Result)
	fmt.Printf("Part 2: %d\n", part2Result)
}
