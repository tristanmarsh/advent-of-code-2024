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

func masMatchesFromCoords(searchLocations []Coords, board Board) (count int) {
	for _, location := range searchLocations {
		if checkMasNeighbours(location, board) {
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

// determine if 9 squares around "A" form one of 4 valid arrangments where both "M" chars are on one side.
// Top       Right     Bottom   Left
// [M]-[M]   S -[M]    S - S    [M]- S
//  - A -    - A -     - A -     - A -
//  S - S    S -[M]   [M]-[M]   [M]- S

// Then check for the corresponding "S" on the adjacent side
func checkMasNeighbours(location Coords, board Board) bool {
	searchChars := []string{"M", "S"}
	for _, point := range PointNeighbourMap {
		neighbour := addCoords(location, point)
		if !coordsAreValid(neighbour, board) {
			return false
		}
		if !slices.Contains(searchChars, string(getBoardSquare(neighbour, board))) {
			return false
		}
	}

	nw := getBoardSquare(addCoords(location, PointNeighbourMap[NW]), board)
	ne := getBoardSquare(addCoords(location, PointNeighbourMap[NE]), board)
	se := getBoardSquare(addCoords(location, PointNeighbourMap[SE]), board)
	sw := getBoardSquare(addCoords(location, PointNeighbourMap[SW]), board)

	// top M bottom S
	if nw == "M" && ne == "M" {
		if sw == "S" && se == "S" {
			return true
		}
	}

	// right M left S
	if nw == "S" && ne == "M" {
		if sw == "S" && se == "M" {
			return true
		}
	}

	// bottom M top S
	if nw == "S" && ne == "S" {
		if sw == "M" && se == "M" {
			return true
		}
	}

	// left M right S
	if nw == "M" && ne == "S" {
		if sw == "M" && se == "S" {
			return true
		}
	}

	return false
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
