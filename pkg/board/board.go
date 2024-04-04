package board

import (
	"fmt"
	"strconv"
)

func GetBoard(size int) [][]int {
	board := make([][]int, size)
	for i := range size {
		board[i] = make([]int, size)
		for j := range size {
			board[i][j] = -1
		}
	}
	return board
}

func PrintBoard(board *[][]int) {
	N := len((*board))

	padding := len(strconv.Itoa(N*N)) + 1
	for i := range N {
		for j := range N {
			fmt.Printf("%*d", padding, (*board)[i][j])
		}
		fmt.Print("\n")
	}
}
