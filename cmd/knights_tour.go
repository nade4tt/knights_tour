package main

import (
	board "cmd/knights_tour.go/pkg/board"
)

func main() {
	chessBoard := board.GetBoard(8)
	board.PrintBoard(&chessBoard)
}
