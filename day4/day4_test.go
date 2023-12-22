package day4_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasevano/advent-of-code/day4"
	"testing"
)

func TestSplitNumbers(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]map[string][]int
	}{
		{
			name: "it split line in winning numbers and numbers i have",
			args: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: map[string]map[string][]int{
				"Card 1": {
					"winning numbers":  {41, 48, 83, 86, 17},
					"numbers you have": {83, 86, 6, 31, 17, 9, 48, 53},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day4.SplitNumbers(tt.args), "they should be equal")
		})
	}
}

func TestDay4(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "it return sum of match points for all the cards",
			args: "example.txt",
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day4.Day4(tt.args), "they should be equal")
		})
	}
}
