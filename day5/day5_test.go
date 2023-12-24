package day5_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasevano/advent-of-code/day5"
	"github.com/thomasevano/advent-of-code/utils"
	"strings"
	"testing"
)

func TestExtractSeeds(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []int
	}{
		{
			name: "it extract seeds from input data",
			args: strings.Split(utils.FileToString("example.txt"), "\n\n"),
			want: []int{79, 14, 55, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day5.ExtractSeeds(tt.args), "they should be equal")
		})
	}
}

func TestGetNumbersForEachMap(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{
			name: "it format input data to processable data",
			args: strings.Split(utils.FileToString("example.txt"), "\n\n")[1:],
			want: [][]string{
				{"50 98 2", "52 50 48"},
				{"0 15 37", "37 52 2", "39 0 15"},
				{"49 53 8", "0 11 42", "42 0 7", "57 7 4"},
				{"88 18 7", "18 25 70"},
				{"45 77 23", "81 45 19", "68 64 13"},
				{"0 69 1", "1 0 69"},
				{"60 56 37", "56 93 4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day5.GetNumbersForEachMap(tt.args), "they should be equal")
		})
	}
}

func TestConvertMapDataToInt(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want [][][]int
	}{
		{
			name: "It return maps digits into numbers",
			args: [][]string{
				{"50 98 2", "52 50 48"},
				{"0 15 37", "37 52 2", "39 0 15"},
				{"49 53 8", "0 11 42", "42 0 7", "57 7 4"},
				{"88 18 7", "18 25 70"},
				{"45 77 23", "81 45 19", "68 64 13"},
				{"0 69 1", "1 0 69"},
				{"60 56 37", "56 93 4"},
			},
			want: [][][]int{
				{{50, 98, 2}, {52, 50, 48}},
				{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
				{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
				{{88, 18, 7}, {18, 25, 70}},
				{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}},
				{{0, 69, 1}, {1, 0, 69}},
				{{60, 56, 37}, {56, 93, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day5.ConvertMapDataToInt(tt.args), "they should be equal")
		})
	}
}

func TestGetSeedLocationNumber(t *testing.T) {
	mapsNumbers := [][][]int{
		{{50, 98, 2}, {52, 50, 48}},
		{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
		{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
		{{88, 18, 7}, {18, 25, 70}},
		{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}},
		{{0, 69, 1}, {1, 0, 69}},
		{{60, 56, 37}, {56, 93, 4}},
	}
	tests := []struct {
		name        string
		seeds       int
		mapsNumbers [][][]int
		want        int
	}{
		{
			name:        "It location number for each seed",
			seeds:       79,
			mapsNumbers: mapsNumbers,
			want:        82,
		}, {
			name:        "It location number for each seed",
			seeds:       14,
			mapsNumbers: mapsNumbers,
			want:        43,
		}, {
			name:        "It location number for each seed",
			seeds:       55,
			mapsNumbers: mapsNumbers,
			want:        86,
		}, {
			name:        "It location number for each seed",
			seeds:       13,
			mapsNumbers: mapsNumbers,
			want:        35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day5.GetSeedLocationNumber(tt.seeds, tt.mapsNumbers), "they should be equal")
		})
	}
}

func TestCalcSeedMapping(t *testing.T) {
	tests := []struct {
		name         string
		actualNumber int
		mapNumbers   []int
		want         int
	}{
		{
			name:         "it return corresponding numbers of actualNumber throught map",
			actualNumber: 79,
			mapNumbers:   []int{50, 98, 2},
			want:         79,
		},
		{
			name:         "it return corresponding numbers of actualNumber throught map",
			actualNumber: 79,
			mapNumbers:   []int{52, 50, 48},
			want:         81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day5.CalcSeedMapping(tt.actualNumber, tt.mapNumbers), "they should be equal")
		})
	}
}
