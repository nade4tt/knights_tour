package board

import "testing"

func TestBoard(t *testing.T) {

	for boardSize := 2; boardSize < 100; boardSize++ {
		b := GetBoard(boardSize)
		if len(b) != boardSize {
			t.Errorf("Expected board size %d, but got %d", boardSize, len(b))
		}

		for i := 0; i < boardSize; i++ {
			if len(b[i]) != boardSize {
				t.Errorf("Expected board size %d, but got %d", boardSize, len(b[i]))
			}

			for j := 0; j < boardSize; j++ {
        if b[i][j] != -1 {
          t.Errorf("Expected board cell to be -1, but got %d", b[i][j])
        }
			}
		}
	}
}
