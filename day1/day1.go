package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetDigits(line string) string {
	var result string
	for _, char := range line {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}
	return result
}

func GetFirstAndLastDigit(stringsOfDigits string) int {
	var finalString string
	finalString = stringsOfDigits[0:1] + stringsOfDigits[len(stringsOfDigits)-1:]
	digits, _ := strconv.Atoi(finalString)
	return digits
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func Main() int {
	sliceOfLine := LinesInFile("day1/input.txt")
	var array []int
	for _, element := range sliceOfLine {
		array = append(array, GetFirstAndLastDigit(GetDigits(element)))
	}
	result := sum(array)
	fmt.Println(result)
	return result
}
