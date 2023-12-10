package utils

import (
	"bufio"
	"log"
	"os"
)

func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	var result []string
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
