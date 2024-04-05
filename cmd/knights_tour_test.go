package main

import (
	board "cmd/knights_tour.go/pkg/board"
	tour "cmd/knights_tour.go/pkg/tour"
	"testing"
)

func TestMain(t *testing.T) {
  boardSizes := []int{6, 8, 10, 12}

  for _, boardSize := range boardSizes {
    chessBoard := board.GetBoard(boardSize)

    // Test every position
    for i := range boardSize {
      for j := range boardSize {
        chessBoard = board.GetBoard(boardSize)
        ok := tour.SolveKT(i, j, &chessBoard)
        if !ok {
          t.Errorf("No solution found for board size %d at (%d, %d)", boardSize, i, j)
        }
      }
    }
  }
}
