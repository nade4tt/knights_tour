package main

import (
	board "cmd/knights_tour.go/pkg/board"
	tour "cmd/knights_tour.go/pkg/tour"
	"testing"
)

func TestMain(t *testing.T) {
	boardSize := 8

	// Test board
	chessBoard := board.GetBoard(boardSize)
	N := len(chessBoard)
	if N != boardSize {
		t.Errorf("Invalid board size. Expected %d, got %d\n", boardSize, N)
	}

	N = len(chessBoard[0])
	if N != boardSize {
		t.Errorf("Invalid board size. Expected %d, got %d\n", boardSize, N)
	}

	for i := range boardSize {
		for j := range boardSize {
			if chessBoard[i][j] != -1 {
				t.Errorf("Unexpected value in the board at x: %d y: %d", j, i)
			}
		}
	}

	// Test tour
	chessBoard[0][0] = 1
	chessBoard[6][5] = 2

	if !tour.IsValid(0, 2, &chessBoard, boardSize) {
		t.Error("Move shall be valid")
	}

	if !tour.IsValid(7, 7, &chessBoard, boardSize) {
		t.Error("Move shall be valid")
	}

	if tour.IsValid(0, 0, &chessBoard, boardSize) {
		t.Error("Move shall be invalid")
	}

	if tour.IsValid(5, 6, &chessBoard, boardSize) {
		t.Error("Move shall be invalid")
	}
}
