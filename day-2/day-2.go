package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var partFlag = flag.Int("part", 1, "Part 1 or 2")

func main() {
	flag.Parse()

	switch *partFlag {
	case 1:
		part1()
	case 2:
		part2()
	}
}

// Given a list of n levels (integers) across y reports (lines)
// - determine the number of valid reports where a line
// - contains reports (numbers) all
// - increasing or decreasing by 1, 2, or 3
func part1() {
	fmt.Println("Part 1")
	input := loadInputFile()
	valid := countValidLines(input)
	fmt.Println(valid)
}

// Given the same list of n levels (integers) across y reports (lines)
// - revealuate with the same criteria if each report is valid
// - but allowing for a single level (integer) to be pruned
func part2() {
	fmt.Println("Part 2")
	input := loadInputFile()
	valid := countValidLinesPart2(input)
	fmt.Println(valid)
}

func loadInputFile() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func countValidLines(input string) (valid int) {
	lines := strings.Split(input, "\n")
	for _, stringLine := range lines {
		line := lineToIntSlice(stringLine)
		if isLineValid(line) {
			valid++
		}
	}
	return
}

func countValidLinesPart2(input string) (valid int) {
	lines := strings.Split(input, "\n")
	for _, stringLine := range lines {
		line := lineToIntSlice(stringLine)
		if isLineValid(line) || hasValidSubset(line) {
			valid++
		}
	}
	return
}

func hasValidSubset(line []int) bool {
	for i := 0; i < len(line); i++ {
		subset := copySliceWithoutIndex(line, i)
		isSubsetValid := isLineValid(subset)
		if isSubsetValid {
			return true
		}
	}
	return false
}

type Direction int

const (
	Ascending  Direction = 0
	Descending Direction = 1
)

func isLineValid(line []int) bool {
	direction := Ascending
	if line[0] > line[len(line)-1] {
		direction = Descending
	}

	for i, x := range line {
		// skip first iteration
		if i > 0 {
			previous := line[i-1]
			isPairValid := isPairValid(x, previous, direction)
			if !isPairValid {
				return false
			}
		}
	}
	return true
}

func isPairValid(x int, prev int, direction Direction) bool {
	diff := x - prev
	if diff < 0 {
		diff = -diff
	}
	if diff > 3 || diff == 0 {
		return false
	}

	if direction == Ascending && x < prev {
		return false
	}

	if direction == Descending && x > prev {
		return false
	}

	return true
}

func lineToIntSlice(lineString string) (line []int) {
	numberRegex := regexp.MustCompile(`(\d+)`)
	matches := numberRegex.FindAllString(lineString, -1)
	for _, match := range matches {
		i, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err)
		}
		line = append(line, i)
	}
	return
}

// The source of all my pain
// copy on write does not prevent the underlying slice entires from being modified
// an explicit new slice must be created with make and copy to prevent modification of the passed slice
func copySliceWithoutIndex(s []int, index int) []int {
	// Create a new slice to ensure the original is unmodified
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return append(newSlice[:index], newSlice[index+1:]...)
}
