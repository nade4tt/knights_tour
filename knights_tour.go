package main

import (
	"flag"
	"fmt"
	"os"
)

var directions = [8][2]int{
	{2, 1},
	{1, 2},
	{-1, 2},
	{-2, 1},
	{-2, -1},
	{-1, -2},
	{1, -2},
	{2, -1},
}

func Walk(x int, y int, pointCount int, solution *[][]int) bool {
	// Base case
	GridSize := len((*solution)[0])
	if pointCount == GridSize*GridSize {
		return true
	}

	// Off the grid
	if x < 0 || x >= GridSize || y < 0 || y >= GridSize {
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
	for i := range len(directions) {
		newX := x + directions[i][0]
		newY := y + directions[i][1]

		if Walk(newX, newY, pointCount, solution) {
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

	Walk(x, y, 1, &solution)
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
