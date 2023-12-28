package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Shim for 2 param return values
func M(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

//func TestGetHandData(t *testing.T) {
//	tests := []struct {
//		name              string
//		args              []string
//		hands             []Hand
//		handsWithSameType map[string][]int
//	}{
//		{
//			name: "it convert input data to wanted data to be processesd",
//			args: utils.LinesInFile("example.txt"),
//			hands: []Hand{
//				{
//					Cards:         "32T3K",
//					Bid:           765,
//					NumberOfChars: onePair,
//					Type:          "onePair",
//					CardsValue:    []int{3, 2, 10, 3, 13},
//				},
//				{
//					"T55J5",
//					684,
//					[]int{1, 3, 1},
//					"threeOfAKind",
//					[]int{10, 5, 5, 11, 5},
//				},
//				{
//					"KK677",
//					28,
//					[]int{1, 2, 2},
//					"twoPair",
//					[]int{13, 13, 6, 7, 7},
//				},
//				{
//					"KTJJT",
//					220,
//					[]int{2, 1, 2},
//					"twoPair",
//					[]int{13, 10, 11, 11, 10},
//				},
//				{
//					"QQQJA",
//					483,
//					[]int{1, 3, 1},
//					"threeOfAKind",
//					[]int{12, 12, 12, 11, 14},
//				},
//			},
//			handsWithSameType: map[string][]int{
//				"onePair":      {0},
//				"twoPair":      {2, 3},
//				"threeOfAKind": {1, 4},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			assert.ElementsMatch(t, tt.hands, M(getHandData(tt.args)), "they should be equal")
//		})
//	}
//}

func TestDetermineType(t *testing.T) {
	tests := []struct {
		name string
		hand Hand
		want string
	}{
		{
			name: "return type fiveOfAKind when 5 same cards",
			hand: Hand{
				Cards:         "AAAAA",
				Bid:           0,
				NumberOfChars: fiveOfAKind,
				Type:          "",
				CardsValue:    []int{14, 14, 14, 14, 14},
			},
			want: "fiveOfAKind",
		}, {
			name: "return type fourOfAKind when 4 same cards",
			hand: Hand{
				Cards:         "AA8AA",
				Bid:           0,
				NumberOfChars: fourOfAKind,
				Type:          "",
				CardsValue:    []int{14, 14, 8, 14, 14},
			},
			want: "fourOfAKind",
		}, {
			name: "return type threeOfAKind when 3 same cards",
			hand: Hand{
				Cards:         "TTT98JA",
				Bid:           0,
				NumberOfChars: threeOfAKind,
				Type:          "",
				CardsValue:    []int{10, 10, 10, 9, 8, 11, 14},
			},
			want: "threeOfAKind",
		}, {
			name: "return type twoPair when 2 times 2 same cards",
			hand: Hand{
				Cards:         "23432",
				Bid:           0,
				NumberOfChars: twoPair,
				Type:          "",
				CardsValue:    []int{2, 3, 4, 3, 2},
			},
			want: "twoPair",
		}, {
			name: "return type onePair when 2 same cards",
			hand: Hand{
				Cards:         "A23A4",
				Bid:           0,
				NumberOfChars: onePair,
				Type:          "",
				CardsValue:    []int{14, 2, 3, 14, 4},
			},
			want: "onePair",
		}, {
			name: "return type highCard when all cards are distinct",
			hand: Hand{
				Cards:         "23456",
				Bid:           0,
				NumberOfChars: highCard,
				Type:          "",
				CardsValue:    []int{2, 3, 4, 5, 6},
			},
			want: "highCard",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, determineType(tt.hand), "they should be equal")
		})
	}
}

func TestOrderHandsByType(t *testing.T) {
	tests := []struct {
		name              string
		handsWithSameType map[string][]int
		want              [][]int
	}{
		{
			name: "transform map in an ordered by type slice",
			handsWithSameType: map[string][]int{
				"onePair":      {0},
				"twoPair":      {2, 3},
				"threeOfAKind": {1, 4},
			},
			want: [][]int{{0}, {2, 3}, {1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, orderHandsByType(tt.handsWithSameType), "they should be equal")
		})
	}
}

func TestGetHandsFromIndex(t *testing.T) {
	tests := []struct {
		name               string
		hands              []Hand
		orderedHandsByType [][]int
		want               [][]Hand
	}{
		{
			name: "return hands coressponding to Index",
			hands: []Hand{
				{
					Cards:         "32T3K",
					Bid:           765,
					NumberOfChars: onePair,
					Type:          "onePair",
					CardsValue:    []int{3, 2, 10, 3, 13},
				},
				{
					"T55J5",
					684,
					[]int{1, 3, 1},
					"threeOfAKind",
					[]int{10, 5, 5, 11, 5},
				},
				{
					"KK677",
					28,
					[]int{1, 2, 2},
					"twoPair",
					[]int{13, 13, 6, 7, 7},
				},
				{
					"KTJJT",
					220,
					[]int{2, 1, 2},
					"twoPair",
					[]int{13, 10, 11, 11, 10},
				},
				{
					"QQQJA",
					483,
					[]int{1, 3, 1},
					"threeOfAKind",
					[]int{12, 12, 12, 11, 14},
				},
			},
			orderedHandsByType: [][]int{{0}, {2, 3}, {1, 4}},
			want: [][]Hand{
				{

					{
						Cards:         "32T3K",
						Bid:           765,
						NumberOfChars: onePair,
						Type:          "onePair",
						CardsValue:    []int{3, 2, 10, 3, 13},
					},
				},
				{

					{
						"KK677",
						28,
						[]int{1, 2, 2},
						"twoPair",
						[]int{13, 13, 6, 7, 7},
					},
					{
						"KTJJT",
						220,
						[]int{2, 1, 2},
						"twoPair",
						[]int{13, 10, 11, 11, 10},
					},
				},
				{
					{
						"T55J5",
						684,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{10, 5, 5, 11, 5},
					},
					{
						"QQQJA",
						483,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{12, 12, 12, 11, 14},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, getHandsFromIndex(tt.hands, tt.orderedHandsByType), "they should be equal")
		})
	}
}

func TestSortByHandWeight(t *testing.T) {
	tests := []struct {
		name  string
		hands [][]Hand
		want  [][]Hand
	}{
		{
			name: "",
			hands: [][]Hand{
				{

					{
						Cards:         "32T3K",
						Bid:           765,
						NumberOfChars: onePair,
						Type:          "onePair",
						CardsValue:    []int{3, 2, 10, 3, 13},
					},
				},
				{

					{
						"KK677",
						28,
						[]int{1, 2, 2},
						"twoPair",
						[]int{13, 13, 6, 7, 7},
					},
					{
						"KTJJT",
						220,
						[]int{2, 1, 2},
						"twoPair",
						[]int{13, 10, 11, 11, 10},
					},
				},
				{
					{
						"T55J5",
						684,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{10, 5, 5, 11, 5},
					},
					{
						"QQQJA",
						483,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{12, 12, 12, 11, 14},
					},
				},
			},
			want: [][]Hand{
				{

					{
						Cards:         "32T3K",
						Bid:           765,
						NumberOfChars: onePair,
						Type:          "onePair",
						CardsValue:    []int{3, 2, 10, 3, 13},
					},
				},
				{

					{
						"KTJJT",
						220,
						[]int{2, 1, 2},
						"twoPair",
						[]int{13, 10, 11, 11, 10},
					},
					{
						"KK677",
						28,
						[]int{1, 2, 2},
						"twoPair",
						[]int{13, 13, 6, 7, 7},
					},
				},
				{
					{
						"T55J5",
						684,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{10, 5, 5, 11, 5},
					},
					{
						"QQQJA",
						483,
						[]int{1, 3, 1},
						"threeOfAKind",
						[]int{12, 12, 12, 11, 14},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, sortByHandWeight(tt.hands), "they should be equal")
		})
	}
}
