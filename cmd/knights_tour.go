package main

import (
	"cmd/knights_tour.go/pkg/board"
	tour "cmd/knights_tour.go/pkg/tour"
	"flag"
	"fmt"
)

func main() {
	boardSize := flag.Int("n", 8, "Size of the board")
	startX := flag.Int("x", 0, "Starting x position")
	startY := flag.Int("y", 0, "Starting x position")
	help := flag.Bool("h", false, "Display help information")
	flag.Parse()

	if *help {
		fmt.Printf("Usage:\n-n: board size\n-x: starting x position\n-y: starting y position\n\n")
		return
	}

	if *boardSize <= 0 || *boardSize > 12 {
		fmt.Printf("Board size must be between 1 and 12\n")
		return
	}

	chessBoard := board.GetBoard(*boardSize)
	ok := tour.SolveKT(*startX, *startY, &chessBoard)

	fmt.Printf("Board size: %dx%d\n", *boardSize, *boardSize)
	fmt.Printf("Starting position: (%d, %d)\n\n", *startX, *startY)

	if ok {
		board.PrintBoard(&chessBoard)
	} else {
		println("No solution found")
	}
}
