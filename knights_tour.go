package main

import (
	"fmt"
)

const GridSize = 8
const StartX = 0
const StartY = 0

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

func Walk(x int, y int, pointCount int, solution *[GridSize][GridSize]int) bool {
	// Base case
	if pointCount == GridSize*GridSize {
		return true
	}

	// Off the grid
	if x < 0 || x >= GridSize || y < 0 || y >= GridSize {
		return false
	}

	// Already visited
	if solution[y][x] != -1 {
		return false
	}

	// PRE condition
	solution[y][x] = pointCount
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
	solution[y][x] = -1
	return false
}

func SolveKnightsTour(x int, y int) [GridSize][GridSize]int {
	var solution [GridSize][GridSize]int
	for i := range solution {
		for j := range solution {
			solution[j][i] = -1
		}
	}

	Walk(x, y, 1, &solution)
	return solution
}

func GetPadding() int {
	maxVal := GridSize * GridSize
	count := 0
	for maxVal > 0 {
		maxVal /= 10
		count++
	}
	return count
}

func PrintGrid(grid [GridSize][GridSize]int) {
	padding := GetPadding()

	for i := range grid {
		for j := range grid {
			fmt.Printf("%*d  ", padding, grid[j][i])
		}
		fmt.Print("\n")
	}
}

func main() {
	grid := SolveKnightsTour(StartX, StartY)
	PrintGrid(grid)
}
