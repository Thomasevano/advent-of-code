package main

import (
	"cmp"
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"slices"
	"strconv"
	"strings"
)

var (
	fiveOfAKind  = []int{5}
	fourOfAKind  = []int{4, 1}
	fullHouse    = []int{3, 2}
	threeOfAKind = []int{3, 1, 1}
	twoPair      = []int{2, 2, 1}
	onePair      = []int{2, 1, 1, 1}
	highCard     = []int{1, 1, 1, 1, 1}
)

var CharToIntValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Hand struct {
	Cards         string
	Bid           int
	NumberOfChars []int
	Type          string
	CardsValue    []int
}

func main() {
	// should return 6440
	//data := utils.LinesInFile("aoc-example.txt")
	// should return part 1: 1343
	//data := utils.LinesInFile("example.txt")
	// should return part 1: 6592
	// should return Part 2: 6839
	//data := utils.LinesInFile("example1.txt")
	// should return 248559379
	data := utils.LinesInFile("input.txt")
	hands, handsWithSameType := getHandData(data)

	orderedHandsByType := orderHandsByType(handsWithSameType)
	handsSortedByType := getHandsFromIndex(hands, orderedHandsByType)

	sortedHands := sortByHandWeight(handsSortedByType)

	var sortedHandsBid []int
	for _, handsByType := range sortedHands {
		for _, hand := range handsByType {
			sortedHandsBid = append(sortedHandsBid, hand.Bid)
		}
	}

	totalWinning := 0
	for index, bid := range sortedHandsBid {
		//fmt.Printf("Rank: %d -> Bid: %d\n", index+1, bid)
		totalWinning += (index + 1) * bid
	}
	fmt.Println(totalWinning)
}

func getHandData(data []string) ([]Hand, map[string][]int) {
	var handsWithBids []Hand
	handsWithSameType := make(map[string][]int)
	for handIndex, handWithBid := range data {
		var hand Hand
		hand.Cards, hand.Bid = getHandCardAndBid(handWithBid)
		hand.NumberOfChars = []int{}

		countForEachChar := make(map[rune]int)
		for _, card := range hand.Cards {
			hand.CardsValue = append(hand.CardsValue, CharToIntValue[string(card)])
			countForEachChar[card]++
		}
		for _, count := range countForEachChar {
			hand.NumberOfChars = append(hand.NumberOfChars, count)
		}
		hand.Type = determineType(hand)

		handsWithSameType[hand.Type] = append(handsWithSameType[hand.Type], handIndex)
		handsWithBids = append(handsWithBids, hand)
	}
	return handsWithBids, handsWithSameType
}

func getHandCardAndBid(handWithBid string) (string, int) {
	separateCardsAndBid := strings.Split(handWithBid, " ")
	cards := separateCardsAndBid[0]
	bid, _ := strconv.Atoi(separateCardsAndBid[1])

	return cards, bid
}

func determineType(hand Hand) string {
	switch {
	case utils.ElementsMatch(fiveOfAKind, hand.NumberOfChars):
		return "fiveOfAKind"
	case utils.ElementsMatch(fourOfAKind, hand.NumberOfChars):
		return "fourOfAKind"
	case utils.ElementsMatch(fullHouse, hand.NumberOfChars):
		return "fullHouse"
	case utils.ElementsMatch(threeOfAKind, hand.NumberOfChars):
		return "threeOfAKind"
	case utils.ElementsMatch(twoPair, hand.NumberOfChars):
		return "twoPair"
	case utils.ElementsMatch(onePair, hand.NumberOfChars):
		return "onePair"
	case utils.ElementsMatch(highCard, hand.NumberOfChars):
		return "highCard"
	}
	return ""
}

func orderHandsByType(handsWithSameType map[string][]int) [][]int {
	indexesOfHandsByType := make([][]int, 7)
	for handType, handIndexes := range handsWithSameType {
		switch {
		case handType == "highCard":
			indexesOfHandsByType[0] = handIndexes
			break
		case handType == "onePair":
			indexesOfHandsByType[1] = handIndexes
			break
		case handType == "twoPair":
			indexesOfHandsByType[2] = handIndexes
			break
		case handType == "threeOfAKind":
			indexesOfHandsByType[3] = handIndexes
			break
		case handType == "fullHouse":
			indexesOfHandsByType[4] = handIndexes
			break
		case handType == "fourOfAKind":
			indexesOfHandsByType[5] = handIndexes
			break
		case handType == "fiveOfAKind":
			indexesOfHandsByType[6] = handIndexes
			break
		}
	}
	indexesOfHandsByType = deleteEmpty(indexesOfHandsByType)
	return indexesOfHandsByType
}

func deleteEmpty(s [][]int) [][]int {
	var r [][]int
	for _, str := range s {
		if len(str) > 0 {
			r = append(r, str)
		}
	}
	return r
}

func getHandsFromIndex(hands []Hand, handsWithSameType [][]int) [][]Hand {
	var handsOrderedByType = make([][]Hand, len(handsWithSameType))
	for i, indexesOfHands := range handsWithSameType {
		handsOrderedByType[i] = make([]Hand, len(indexesOfHands))
		for index, handIndex := range indexesOfHands {
			handsOrderedByType[i][index] = hands[handIndex]
		}
	}
	return handsOrderedByType
}

func sortByHandWeight(handsSortedByType [][]Hand) [][]Hand {
	for _, handsByType := range handsSortedByType {
		slices.SortStableFunc(handsByType, func(a, b Hand) int {
			result := 0
			for index := 0; index < len(a.CardsValue); index++ {
				result = cmp.Compare(a.CardsValue[index], b.CardsValue[index])
				if result == -1 || result == 1 {
					break
				}
			}
			return result
		})
	}
	return handsSortedByType
}
