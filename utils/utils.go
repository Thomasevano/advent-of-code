package utils

import (
	"bufio"
	"fmt"
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
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func FileToString(filePath string) string {
	// Read the contents of the file
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
	}

	// Convert the byte slice to a string
	content := string(contentBytes)

	return content
}
