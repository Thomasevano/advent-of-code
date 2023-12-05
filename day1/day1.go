package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var letterDigitsToNumeric = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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

func GetDigits(line string) []string {
	r := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
	correctedLine := r.Replace(line)

	digitsRegex := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
	matches := digitsRegex.FindAllString(correctedLine, -1)

	result := GetFirstAndLastDigit(matches)

	return result
}

func ConvertStringDigitToNumeric(digit string) string {
	var numericDigit string
	isDigitNumber := regexp.MustCompile(`\d`).MatchString(digit)

	if !isDigitNumber {
		numericDigit += letterDigitsToNumeric[digit]
	} else {
		numericDigit += digit
	}

	return numericDigit
}

func GetFirstAndLastDigit(digits []string) []string {
	var slice []string
	slice = append(slice, digits[0], digits[len(digits)-1])
	return slice
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func Main(filePath string) {
	sliceOfLine := LinesInFile(filePath)
	var array []int
	var numericDigits string
	var intDigits int

	for _, element := range sliceOfLine {
		digitsArray := make([]string, 0, 2)

		for _, digit := range GetDigits(element) {
			digitsArray = append(digitsArray, ConvertStringDigitToNumeric(digit))
			numericDigits = strings.Join(digitsArray, "")
			number, _ := strconv.Atoi(numericDigits)
			intDigits = number
		}
		array = append(array, intDigits)
	}
	result := sum(array)
	fmt.Println(result)
}
