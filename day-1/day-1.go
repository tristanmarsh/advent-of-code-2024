package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Given two lists of integers, sort ascending and calculate the difference between corresponding elements in each list
func main() {
	listsFile := loadListFile()
	listA, listB := extractListsFromFile(listsFile)

	intListA := stringListToIntList(listA)
	intListB := stringListToIntList(listB)
	fmt.Println("Lists")
	fmt.Println("A: ", intListA)
	fmt.Println("B: ", intListB)

	sort.Ints(intListA)
	sort.Ints(intListB)
	fmt.Println("\nSorted")
	fmt.Println("A: ", intListA)
	fmt.Println("B: ", intListB)

	distanceList := getDistanceList(intListA, intListB)

	fmt.Println("\nDistances\n")
	fmt.Println("D: ", distanceList)

	// fmt.Println("\nRaw:")
	// printIntSlice(distanceList)

	sum := sumIntegers(distanceList)
	fmt.Println("\nSum:")
	fmt.Println(sum)

}

func loadListFile() (list string) {
	file, err := os.ReadFile("lists.txt")
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

func printIntSlice(list []int) {
	for _, x := range list {
		fmt.Println(x)
	}
}
