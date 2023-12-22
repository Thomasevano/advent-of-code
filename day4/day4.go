package day4

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"slices"
	"strconv"
	"strings"
)

var cards = map[string]map[string][]int{}
var part1 = 0

func SplitNumbers(line string) map[string]map[string][]int {

	subset := strings.SplitAfter(line, ":")
	gameNumber := subset[0]
	gameNumber = gameNumber[0 : len(gameNumber)-1]

	subset = strings.SplitAfter(subset[1], "|")

	winningNumbers := StringToInt(TrimStringNumbers(subset[0]))

	numbersIHave := StringToInt(TrimStringNumbers(subset[1]))

	numbers := map[string][]int{
		"winning numbers":  winningNumbers,
		"numbers you have": numbersIHave,
	}

	cards[gameNumber] = numbers
	return cards
}

func TrimStringNumbers(numbers string) []string {
	numbers = strings.Trim(numbers, " ")
	numbers = strings.Replace(numbers, "|", "", 1)

	stringNumbersSuite := strings.Split(numbers, " ")

	if stringNumbersSuite[len(stringNumbersSuite)-1] == "" {
		stringNumbersSuite = stringNumbersSuite[0 : len(stringNumbersSuite)-1]
	}

	return stringNumbersSuite
}

func StringToInt(stringNumbers []string) []int {
	intNumbers := []int{}
	for _, stringNumber := range stringNumbers {
		number, _ := strconv.Atoi(stringNumber)
		intNumbers = append(intNumbers, number)
	}
	intNumbers = slices.DeleteFunc(intNumbers, func(n int) bool {
		return n == 0 // delete the odd numbers
	})
	return intNumbers
}

func Day4(filePath string) int {
	sliceOfLine := utils.LinesInFile(filePath)

	for _, line := range sliceOfLine {
		SplitNumbers(line)
	}

	for _, card := range cards {
		matchPoints := 0
		for _, number := range card["numbers you have"] {
			if slices.Contains(card["winning numbers"], number) {
				if matchPoints == 0 {
					matchPoints += 1
				} else {
					matchPoints = matchPoints * 2
				}
			}
		}
		part1 += matchPoints
		fmt.Println(part1)
	}
	return part1
}
