// func createBoard()
//     1. Initialize board as a 2D array of size N x N
//     2. Set all values in board to 0
//     3. Return board

// Function isValid(x, y, board)
//     1. If x and y are within the bounds of the board and board[x][y] is 0
//         a. Return True
//     2. Else
//         a. Return False

// Function nextMove(x, y, board)
//     1. Initialize minDeg to N+1 and minMoveIndex to -1
//     2. For each move in moves
//         a. Calculate newX, newY based on the current move
//         b. If newX, newY is a valid position
//             i. Count valid onward moves from newX, newY
//             ii. If the count is less than minDeg
//                 - Set minDeg to this count
//                 - Update minMoveIndex to the current move index
//     3. If minMoveIndex is not -1
//         a. Calculate the next move using minMoveIndex
//         b. Return the next position (nx, ny) and True
//     4. Else
//         a. Return 0, 0, False

// Function solveKT(startX, startY)
//     1. Initialize board using createBoard()
//     2. Set the starting position on board to 1 (board[startX][startY] = 1)
//     3. For each move from 2 to N*N
//         a. Get the next move using nextMove()
//         b. If no valid move is found, return False
//         c. Else, mark the next move on the board and update the current position
//     4. Print the board
//     5. Return True

// Main
//  1. Set startX, startY to the desired starting position
//  2. If solveKT(startX, startY) is False
//     a. Print "Solution does not exist"
//  3. Else
//     a. Print the solution
package main

import "fmt"

func createBoard() {

}

func isValid() {

}

func nextMove(x int, y int, board *[][]int) {

}

func solveKT(startX int, startY int) {

}

func main() {
	fmt.Println("Knight's Tour Problem")
}
