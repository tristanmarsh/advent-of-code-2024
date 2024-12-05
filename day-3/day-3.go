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

// Given a corrupted string with junk data, find the `mul(2,3)` function calls accepting 2 integers between 1 and 3 characters long
// Sum the integer result of all mul calls
func part1() {
	fmt.Println("Part 1")
	input := loadInput()
	mulCalls := extractMulCommands(input)
	mulArguments := extractArgumentPairs(mulCalls)
	mulResults := processMulCalls(mulArguments)
	result := sumIntArray(mulResults)
	fmt.Println(result)
}

func part2() {
	fmt.Println("Part 2")
	input := loadInput()
	extractMulCommands(input)
	result := 0
	fmt.Println(result)
}

func loadInput() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}
	return string(file)
}

func extractMulCommands(input string) []string {
	mulFunctionCalls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	return mulFunctionCalls.FindAllString(input, -1)
}

func extractArgumentPairs(input []string) []string {
	var mulPairs []string
	argumentPair := regexp.MustCompile(`[0-9]{1,3},[0-9]{1,3}`)
	for _, call := range input {
		x := argumentPair.FindString(call)
		mulPairs = append(mulPairs, x)
	}
	return mulPairs
}

func processMulCalls(input []string) (result []int) {
	for _, pair := range input {
		result = append(result, processMulPair(pair))
	}
	return
}

func processMulPair(input string) int {
	pair := strings.Split(input, ",")
	a, b := pair[0], pair[1]
	aInt, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	return aInt * bInt
}

func sumIntArray(input []int) (result int) {
	for _, x := range input {
		result = result + x
	}
	return
}
