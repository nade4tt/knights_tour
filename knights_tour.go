package main

import (
	"flag"
	"fmt"
	"os"
)

type Move struct {
	x int
	y int
}

var directions = [8]Move{
	{x: 2, y: 1},
	{x: 1, y: 2},
	{x: -1, y: 2},
	{x: -2, y: 1},
	{x: -2, y: -1},
	{x: -1, y: -2},
	{x: 1, y: -2},
	{x: 2, y: -1},
}

func IsMoveValid(x int, y int, gridSize int, solution *[][]int) bool {
	// Off the grid
	if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
		return false
	}

	// Already visited
	if (*solution)[y][x] != -1 {
		return false
	}

	return true
}

func GetDirections(x int, y int) []Move {
	result := make([]Move, 0)
	for _, move := range directions {
		newX, newY := move.x+x, move.y+y
		result = append(result, Move{x: newX, y: newY})
	}
	return result
}

func Walk(x int, y int, pointCount int, gridSize int, solution *[][]int) bool {
	// Base case
	if pointCount == gridSize*gridSize+1 {
		return true
	}

	// Off the grid
	if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
		return false
	}

	// Already visited
	if (*solution)[y][x] != -1 {
		return false
	}

	// PRE condition
	(*solution)[y][x] = pointCount
	pointCount++

	// Recurse
	// for i := range directions {
	for _, move := range GetDirections(x, y) {
		if Walk(move.x, move.y, pointCount, gridSize, solution) {
			return true
		}
	}

	// POST condition
	(*solution)[y][x] = -1
	return false
}

func SolveKnightsTour(x int, y int, gridSize int) [][]int {
	solution := make([][]int, gridSize)
	for i := range gridSize {
		solution[i] = make([]int, gridSize)
		for j := range gridSize {
			solution[i][j] = -1
		}
	}

	Walk(x, y, 1, gridSize, &solution)
	return solution
}

func GetPadding(gridSize int) int {
	maxVal := gridSize * gridSize
	count := 0
	for maxVal > 0 {
		maxVal /= 10
		count++
	}
	return count
}

func PrintGrid(grid *[][]int) {
	gridSize := len((*grid))

	padding := GetPadding(gridSize)

	for i := range len((*grid)) {
		for j := range len((*grid)[0]) {
			fmt.Printf("%*d  ", padding, (*grid)[j][i])
		}
		fmt.Print("\n")
	}
}

func validateInputArguments(x *int, y *int, gridSize int) {
	if *x < 0 || *x >= gridSize {
		fmt.Println("Invalid X coordinate")
		os.Exit(1)
	}
	if *y < 0 || *y >= gridSize {
		fmt.Println("Invalid Y coordinate")
		os.Exit(1)
	}
}

func main() {
	startX := flag.Int("x", 0, "Starting X coordinate")
	startY := flag.Int("y", 0, "Starting Y coordinate")
	gridSize := flag.Int("size", 8, "Size of a grid")
	flag.Parse()

	validateInputArguments(startX, startY, *gridSize)

	fmt.Printf("X: %d, Y: %d\n", *startX, *startY)

	grid := SolveKnightsTour(*startX, *startY, *gridSize)
	PrintGrid(&grid)
}
