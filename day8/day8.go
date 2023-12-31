package main

import (
	"fmt"
	"github.com/thomasevano/advent-of-code/utils"
	"regexp"
	"strings"
)

func main() {
	data := utils.LinesInFile("example1.txt")
	//data := utils.LinesInFile("example2.txt")
	//data := utils.LinesInFile("input.txt")

	directionString := data[0]
	fmt.Println(directionString)
	nodesSlice := data[2:]

	directionsSlice := strings.Split(directionString, "")
	fmt.Println(directionsSlice)

	var nodes = make(map[string][]string)
	lettersRegex := regexp.MustCompile(`[A-Z]{3}`)
	for _, node := range nodesSlice {
		foundLetter := lettersRegex.FindAllString(node, 3)
		nodes[foundLetter[0]] = []string{foundLetter[1], foundLetter[2]}
	}
	fmt.Println(nodes)

	var steps int
	var actualPosition = "AAA"
	var finalPosition = "ZZZ"
	for !(actualPosition == finalPosition) {
		for _, direction := range directionsSlice {
			if direction == "L" {
				actualPosition = nodes[actualPosition][0]
			} else {
				actualPosition = nodes[actualPosition][1]
			}
			steps++
			if actualPosition == finalPosition {
				break
			}
		}
	}

	fmt.Printf("steps: %d\n", steps)

}
