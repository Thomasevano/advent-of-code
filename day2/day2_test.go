package day2_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasevano/2023-advent-of-code/day2"
	"testing"
)

// RULES: Determine which games would have been possible with only 12 red cubes, 13 green cubes, and 14 blue cubes

func TestGetSubsets(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "it split game line in subsets",
			args: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: []string{" 3 blue, 4 red;", " 1 red, 2 green, 6 blue;", " 2 green"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, day2.GetSubsets(tt.args), "they should be equal")
		})
	}
}

func TestGetMapFromSubsets(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want map[string]int
	}{
		{
			name: "it return map of number of dices for each color for each subset",
			args: []string{" 3 blue", "4 red;"},
			want: map[string]int{
				"blue": 3,
				"red":  4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day2.GetMapFromSubsets(tt.args), "they should be equal")
		})
	}
}

func TestCheckIfGameValid(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want map[string]int
	}{
		{
			name: "it return map of number of dices for each color for each subset",
			args: []string{" 3 blue", "4 red;"},
			want: map[string]int{
				"blue": 3,
				"red":  4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day2.GetMapFromSubsets(tt.args), "they should be equal")
		})
	}
}
