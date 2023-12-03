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
			args: "example.txt",
			want: []string{"11abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.LinesInFile(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinesInFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDigits(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "it returns first and last digits of each element",
			args: "11abc2",
			want: "112",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.GetDigits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildDockerComposeNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstAndLastDigit(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "it returns first and last digits of each element",
			args: "112",
			want: 12,
		},
		{
			name: "it returns first and last digits of each element",
			args: "38",
			want: 38,
		},
		{
			name: "it returns first and last digits of each element",
			args: "12345",
			want: 15,
		},
		{
			name: "it returns first and last digits of each element",
			args: "7",
			want: 77,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.GetFirstAndLastDigit(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstAndLastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
