package day6_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasevano/advent-of-code/day6"
	"github.com/thomasevano/advent-of-code/utils"
	"testing"
)

func TestFormatRaceData(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]int
	}{
		{
			name: "it returns race data in tuple",
			args: utils.LinesInFile("example.txt"),
			want: [][]int{
				{7, 9},
				{15, 40},
				{30, 200},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, day6.FormatRaceData(tt.args), "they should be equal")
		})
	}
}
