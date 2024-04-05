package tour

func IsValid(x int, y int, board *[][]int, boardSize int) bool {
	return x >= 0 && x <= boardSize-1 && y >= 0 && y <= boardSize-1 && (*board)[y][x] == -1
}

var moves = [8][2]int{
	{1, 2},
	{1, -2},
	{2, 1},
	{2, -1},
	{-1, 2},
	{-1, -2},
	{-2, 1},
	{-2, -1},
}

func nextMove(x int, y int, chessBoard *[][]int, boardSize int) (int, int, bool) {
	minDeg := boardSize + 1
	minMoveIndex := -1

	for moveIndex, move := range moves {
		newX := x + move[0]
		newY := y + move[1]

		count := 0
		if IsValid(newX, newY, chessBoard, boardSize) {
			for _, move := range moves {
				if IsValid(newX+move[0], newY+move[1], chessBoard, boardSize) {
					count++
				} else {
					continue
				}
			}

			if count < minDeg {
				minDeg = count
				minMoveIndex = moveIndex
			}
		}
	}

	if minMoveIndex == -1 {
		return 0, 0, false
	}

	nx := x + moves[minMoveIndex][0]
	ny := y + moves[minMoveIndex][1]

	return nx, ny, true

}

func SolveKT(startX int, startY int, chessBoard *[][]int) bool {
	(*chessBoard)[startY][startX] = 1

	boardSize := len(*chessBoard)
	x, y := startX, startY
	for i := 2; i <= boardSize*boardSize; i++ {
		nextX, nextY, ok := nextMove(x, y, chessBoard, boardSize)
		if !ok {
			return false
		}

		(*chessBoard)[nextY][nextX] = i
		x = nextX
		y = nextY
	}

	return true
}
