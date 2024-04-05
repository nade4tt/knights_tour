package tour

import (
	"cmd/knights_tour.go/pkg/board"
	"testing"
)

func TestTour(t *testing.T) {
	for boardSize := 6; boardSize <= 100; boardSize++ {
		chessBoard := board.GetBoard(boardSize)

		validMoves := [][2]int{
			{0, 0},
			{boardSize / 2, boardSize / 2},
			{boardSize - 1, boardSize - 1},
			{0, boardSize - 1},
			{boardSize - 1, 0},
			{0, boardSize / 2},
			{boardSize / 2, 0},
		}

		for _, move := range validMoves {
			chessBoard[move[1]][move[0]] = 1
		}

		for i := range boardSize {
			for j := range boardSize {
				found := false
				for _, subSlice := range validMoves {
					if subSlice[0] == j && subSlice[1] == i {
						found = true
						break
					}
				}

				if found {
					if isValid(j, i, &chessBoard, boardSize) {
						t.Errorf("Expected invalid move for grid size %d at (%d, %d)", boardSize, j, i)
					}
				} else {
					if !isValid(j, i, &chessBoard, boardSize) {
						t.Errorf("Expected valid move for grid size %d at (%d, %d)", boardSize, i, j)
					}
				}
			}
		}
	}
}
