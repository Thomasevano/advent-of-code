package day3_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasevano/advent-of-code/day3"
	"github.com/thomasevano/advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "it split line in a slice",
			args: "example.txt",
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day3.Main(tt.args), "they should be equal")
		})
	}
}

func TestGetSymbols(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want map[day3.Coordinates]string
	}{
		{
			name: "it get Symbols with coords of the input",
			args: utils.LinesInFile("example.txt"),
			want: map[day3.Coordinates]string{
				{1, 3}: "'*'",
				{3, 6}: "'#'",
				{4, 3}: "'*'",
				{5, 5}: "'+'",
				{8, 3}: "'$'",
				{8, 5}: "'*'",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day3.GetSymbols(tt.args), "they should be equal")
		})
	}
}

func TestGetNumbersAdjacentCells(t *testing.T) {
	tests := []struct {
		name                        string
		sliceIndex                  int
		line                        string
		numberCoords                []int
		numbersCoordsWithLineNumber day3.NumberCoords
		want                        map[day3.Coordinates]string
	}{
		{
			name:                        "it returns AdjacentCells when there is more than one number on line",
			sliceIndex:                  0,
			line:                        "467..114..",
			numberCoords:                []int{5, 8},
			numbersCoordsWithLineNumber: day3.NumberCoords{LineNumber: 0, Coords: day3.Coordinates{X: 5, Y: 8}},
			want: map[day3.Coordinates]string{
				{-1, 4}: "",
				{-1, 5}: "",
				{-1, 6}: "",
				{-1, 7}: "",
				{-1, 8}: "",
				{0, 4}:  "",
				{0, 5}:  "",
				{0, 6}:  "",
				{0, 7}:  "",
				{0, 8}:  "",
				{1, 4}:  "",
				{1, 5}:  "",
				{1, 6}:  "",
				{1, 7}:  "",
				{1, 8}:  "",
			},
		},
		{
			name:                        "it returns AdjacentCells when there is one number on line",
			sliceIndex:                  6,
			line:                        "..592.....",
			numberCoords:                []int{2, 5},
			numbersCoordsWithLineNumber: day3.NumberCoords{LineNumber: 0, Coords: day3.Coordinates{X: 2, Y: 5}},
			want: map[day3.Coordinates]string{
				{5, 1}: "",
				{5, 2}: "",
				{5, 3}: "",
				{5, 4}: "",
				{5, 5}: "",
				{6, 1}: "",
				{6, 2}: "",
				{6, 3}: "",
				{6, 4}: "",
				{6, 5}: "",
				{7, 1}: "",
				{7, 2}: "",
				{7, 3}: "",
				{7, 4}: "",
				{7, 5}: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day3.GetNumbersAdjacentCells(tt.sliceIndex, tt.line, tt.numberCoords,
				tt.numbersCoordsWithLineNumber), "they should be equal")
		})
	}
}
