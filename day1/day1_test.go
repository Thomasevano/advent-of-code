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
		want []string
	}{
		{
			name: "it returns numeric digits of each element",
			args: "11abc2",
			want: []string{"1", "2"},
		},
		{
			name: "it returns numeric and written digits of each element",
			args: "two1nine",
			want: []string{"2", "9"},
		},
		{
			name: "it returns numeric and written digits of each element",
			args: "eightwothree",
			want: []string{"8", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.GetDigits(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstAndLastDigit(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "it returns first and last digits of each element",
			args: []string{"1", "1", "2"},
			want: []string{"1", "2"},
		},
		{
			name: "it returns first and last digits of each element",
			args: []string{"two", "1", "nine"},
			want: []string{"two", "nine"},
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

func TestConvertStringDigitToNumeric(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "it returns first and last digits of each element",
			args: "two",
			want: "2",
		},
		{
			name: "it returns first and last digits of each element",
			args: "3",
			want: "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day1.ConvertStringDigitToNumeric(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringDigitToNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
