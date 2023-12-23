package day4

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"slices"
	"strconv"
	"strings"
)

var cardsSlice []map[string][]int
var part1, part2 = 0, 0
var cardsCount = map[int]int{}

func SplitNumbers(line string) []map[string][]int {
	subset := strings.SplitAfter(line, ":")
	cardNumber := subset[0]
	cardNumber = cardNumber[0 : len(cardNumber)-1]

	subset = strings.SplitAfter(subset[1], "|")

	winningNumbers := StringToInt(TrimStringNumbers(subset[0]))

	numbersIHave := StringToInt(TrimStringNumbers(subset[1]))

	numbers := map[string][]int{
		"winning numbers":  winningNumbers,
		"numbers you have": numbersIHave,
	}

	cardsSlice = append(cardsSlice, numbers)
	return cardsSlice
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

	for cardIndex, card := range cardsSlice {
		matchPoints := 0
		matchedNumbers := 0
		cardsCount[cardIndex+1] += 1
		for _, number := range card["numbers you have"] {
			if slices.Contains(card["winning numbers"], number) {
				if matchPoints == 0 {
					matchPoints += 1
				} else {
					matchPoints = matchPoints * 2
				}
				matchedNumbers += 1
			}
		}
		fmt.Printf("Card: %d with Matched numbers: %d \n", cardIndex+1, matchedNumbers)
		if matchedNumbers > 0 {
			for i := 0; i < matchedNumbers; i++ {
				if cardIndex == 0 {
					cardsCount[cardIndex+1+i+1] += 1
				} else {
					cardsCount[cardIndex+1+i+1] += cardsCount[cardIndex+1]
				}
			}
		}
		part2 += cardsCount[cardIndex+1]
		part1 += matchPoints
	}

	fmt.Println(part1)
	fmt.Println(part2)
	return part1
}
