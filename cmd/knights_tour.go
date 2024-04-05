package main

import (
	tour "cmd/knights_tour.go/pkg/tour"
	"cmd/knights_tour.go/pkg/board"
)

func main() {
	boardSize := 8

	chessBoard := board.GetBoard(boardSize)
	ok := tour.SolveKT(0, 5, &chessBoard)

	if ok {
		board.PrintBoard(&chessBoard)
	} else {
		println("No solution found")
	}
}
