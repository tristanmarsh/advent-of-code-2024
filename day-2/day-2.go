package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	default:
		part1()
	}
}

// given a list of n integers across y lines
// - determine the number of valid lines where a line
// - contains numbers all increasing or all decreasing
// - increases or decreases only by 1, 2, or 3
func part1() {
	fmt.Println("Part 1")
	input := loadInputFile()
	valid := countValidLines(input)
	fmt.Println(valid)
}
func part2() {
	fmt.Println("Part 2")
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
	for _, line := range lines {
		if isValid(line) {
			valid++
		}
	}
	return
}

const (
	Ascending = iota
	Descending
)

func isValid(lineString string) bool {
	direction := Ascending
	line := lineToIntSlice(lineString)
	firstNumber := line[0]
	lastNumber := line[len(line)-1]
	if firstNumber > lastNumber {
		direction = Descending
	}

	for i, x := range line {
		// skip first iteration
		if i > 0 {
			previous := line[i-1]
			diff := math.Abs(float64(x) - float64(previous))
			if diff > 3 || diff == 0 {
				return false
			}
			if direction == Ascending && x < previous {
				return false
			}
			if direction == Descending && x > previous {
				return false
			}
		}
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
