package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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

// Given a crossword supporting all 8 directions (horizontally, vertically, diagonally, and reversed). How many times is `XMAS` appear?
func part1() {
	fmt.Println("Part 1")
	data := loadInputFile()
	board := makeBoard(data)
	// begin search from "X" squares
	xCoords := getAllCharCoords("X", board)
	count := wordMatchesFromCoords("XMAS", xCoords, board)
	fmt.Println(count)
}

// Search for and count the number of diagonal "MAS" which cross intersect at A forming an X
func part2() {
	fmt.Println("Part 2")
	data := loadInputFile()
	board := makeBoard(data)
	// begin search from "A" squares
	aCoords := getAllCharCoords("A", board)
	count := masMatchesFromCoords(aCoords, board)
	fmt.Println(count)
}

func loadInputFile() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

type Square string
type Line []Square

type Board struct {
	lines  []Line
	height int
	width  int
}

type Coords struct {
	x int
	y int
}

func makeBoard(input string) (board Board) {
	lines := strings.Split(input, "\n")
	board.height = len(lines)
	board.width = len(lines[0])
	for _, line := range lines {
		var realLine Line
		if len(line) != board.width {
			log.Fatal("Invalid word search input. All lines must be equal length")
		}

		squares := strings.Split(line, "")
		for _, square := range squares {
			realLine = append(realLine, Square(square))
		}
		board.lines = append(board.lines, realLine)
	}
	return
}

// find coordinates of a given character
func getAllCharCoords(target string, board Board) (results []Coords) {
	for y, line := range board.lines {
		for x, square := range line {
			if string(square) == target {
				results = append(results, Coords{x, y})
			}
		}
	}
	return
}

func wordMatchesFromCoords(word string, searchLocations []Coords, board Board) (count int) {
	for _, head := range searchLocations {
		count += checkAllDirections(word, head, board)
	}
	return
}

type Direction = int

// Directions for word detection
const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

// Directions translated to cartesian XY offset
var DirectionMap = map[Direction]Coords{
	N:  {0, -1},
	NE: {1, -1},
	E:  {1, 0},
	SE: {1, 1},
	S:  {0, 1},
	SW: {-1, 1},
	W:  {-1, 0},
	NW: {-1, -1},
}

func checkAllDirections(word string, head Coords, board Board) (count int) {
	for _, coords := range DirectionMap {
		if directionContainsWord(word, head, coords, board) {
			count++
		}
	}
	return count
}

func directionContainsWord(word string, head Coords, directionMapCoords Coords, board Board) bool {
	var coords = head
	for _, char := range strings.Split(word, "") {
		if !coordsAreValid(coords, board) || string(getBoardSquare(coords, board)) != char {
			return false
		}
		coords = addCoords(coords, directionMapCoords)
	}
	return true
}

func masMatchesFromCoords(searchLocations []Coords, board Board) (count int) {
	for _, location := range searchLocations {
		if hasMasCrossWord(location, board) {
			count++
		}
	}
	return count
}

var PointNeighbourMap = map[Direction]Coords{
	NW: DirectionMap[NW],
	NE: DirectionMap[NE],
	SE: DirectionMap[SE],
	SW: DirectionMap[SW],
}

// An "A" square is a valid X-"MAS" if it is one of the 4 valid arrangements:
//
// Top       Right     Bottom   Left
// [M]-[M]   S -[M]    S - S    [M]- S
//  - A -    - A -     - A -     - A -
//  S - S    S -[M]   [M]-[M]   [M]- S
//
// 2 "M" and 2 "S" chars found in opposing cardinal directions.

// Combing 4 squares sequentially from any point in either direction results in 1 of 4 valid strings
// MMSS, MSSM, SMMS, SSMM
func hasMasCrossWord(center Coords, board Board) bool {
	validStrings := []string{"MMSS", "MSSM", "SSMM", "SMMS"}
	var points []string

	for _, point := range PointNeighbourMap {
		pointCoords := addCoords(center, point)
		if coordsAreValid(pointCoords, board) {
			points = append(points, string(getBoardSquare(pointCoords, board)))
		}
	}

	return slices.Contains(validStrings, strings.Join(points, ""))
}

func addCoords(current Coords, directionCoords Coords) (next Coords) {
	next.x = current.x + directionCoords.x
	next.y = current.y + directionCoords.y

	return
}

func coordsAreValid(coords Coords, board Board) bool {
	return coords.x >= 0 && coords.x < board.width && coords.y >= 0 && coords.y < board.height
}

func getBoardSquare(coords Coords, board Board) (square Square) {
	return board.lines[coords.y][coords.x]
}
