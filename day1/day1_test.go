package day1_test

import (
	"github.com/thomasevano/2023-advent-of-code/day1"
	"reflect"
	"testing"
)

func TestLinesInFile(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "it returns each line in a string inside a slice",
			args: "./example.txt",
			want: []string{"11abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.LinesInFile(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildDockerComposeNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}
