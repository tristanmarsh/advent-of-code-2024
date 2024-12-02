package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

// Given two lists of integers:
// - sort ascending
// - calculate the difference between corresponding elements in each list
// - sum the list of differences
func part1() {
	fmt.Println("Part 1")
	listsFile := loadInputFile()
	listA, listB := extractListsFromFile(listsFile)

	intListA := stringListToIntList(listA)
	intListB := stringListToIntList(listB)

	sort.Ints(intListA)
	sort.Ints(intListB)

	distanceList := getDistanceList(intListA, intListB)

	sum := sumIntegers(distanceList)
	fmt.Println(sum)
}

// Given two lists of integers:
// derive a "similarity score" by summing the product of
// each id in the left list, by it's number of occurances in the right list
func part2() {
	listsFile := loadInputFile()
	listA, listB := extractListsFromFile(listsFile)
	intListA := stringListToIntList(listA)
	intListB := stringListToIntList(listB)

	listASimilarity := getListSimilarity(intListA, intListB)
	fmt.Println("Part 2")
	fmt.Println(listASimilarity)
}

func loadInputFile() (list string) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func extractListsFromFile(listsFile string) (listA []string, listB []string) {
	lines := strings.Split(listsFile, "\n")
	numberRegex := regexp.MustCompile(`(\d+)`)

	for _, line := range lines {
		matches := numberRegex.FindAllString(line, -1)
		if len(matches) != 2 {
			log.Fatal("Invalid input file format. Expected 2 integers per line.")
		}
		listA = append(listA, matches[0])
		listB = append(listB, matches[1])
	}
	return listA, listB
}

func stringListToIntList(list []string) (result []int) {
	for _, string := range list {
		i, err := strconv.Atoi(string)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}
	return
}

func getDistanceList(sliceA []int, sliceB []int) (distances []int) {
	for i := 0; i < len(sliceA); i++ {
		distances = append(distances, int(math.Abs(float64(sliceA[i])-float64(sliceB[i]))))
	}
	return
}

func sumIntegers(list []int) (sum int) {
	for _, x := range list {
		sum = sum + x
	}
	return
}

func getListSimilarity(a []int, b []int) (similarityScore int) {
	for _, id := range a {
		count := getIdCount(id, b)
		idScore := count * id
		similarityScore = similarityScore + idScore
	}
	return
}

func getIdCount(id int, ids []int) (similarity int) {
	for _, i := range ids {
		if i == id {
			similarity++
		}
	}
	return
}
