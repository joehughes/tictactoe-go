package main

import (
	"errors"
	"fmt"
	"strconv"
)

const redColor = "\033[31;1m"
const grayColor = "\033[38;5;239m"
const resetColor = "\033[0m"
const reverseColor = "\033[7m"

var lines = [8][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
	{1, 5, 9},
	{3, 5, 7},
}

type player struct {
	name  string
	piece rune
}

type lineEval struct {
	line         [3]int
	winningPiece rune
}

func printBoard(board [3][3]rune) {
	printBoardWithHighlight(board, []int{})
}

func printBoardWithHighlight(board [3][3]rune, line []int) {
	fmt.Println()
	for rowIndex, row := range board {
		if rowIndex > 0 {
			fmt.Println("━╋━╋━")
		}
		for colIndex, val := range row {
			if colIndex > 0 {
				fmt.Print("┃")
			}
			squareIndex, _ := indicesToSquareNumber(rowIndex, colIndex)
			for _, lineSquareIndex := range line {
				if lineSquareIndex == squareIndex {
					fmt.Print(reverseColor)
				}
			}
			if val == 0 {
				fmt.Print(grayColor + fmt.Sprint(squareIndex) + resetColor)
			} else {
				var boardVal string
				if val == '❌' {
					boardVal = "X"
				} else {
					boardVal = "O"
				}
				fmt.Print(redColor + boardVal + resetColor)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func squareNumberToIndices(number int) (rowIndex int, colIndex int, err error) {
	if number < 1 || number > 9 {
		return 0, 0, errors.New("Number is out of range 1 - 9")
	}

	rowIndex = (number - 1) / 3
	colIndex = (number - 1) % 3
	return
}

func indicesToSquareNumber(rowIndex int, colIndex int) (squareNumber int, err error) {
	if rowIndex < 0 || rowIndex >= 3 {
		return 0, errors.New("rowIndex out of range")
	}
	if colIndex < 0 || colIndex >= 3 {
		return 0, errors.New("colIndex out of range")
	}

	return rowIndex*3 + colIndex + 1, nil
}

func evaluateLines(board [3][3]rune) [8]lineEval {
	result := [8]lineEval{}
	for lineIndex, line := range lines {
		var xCount, oCount int
		for _, squareNumber := range line {
			rowIndex, colIndex, _ := squareNumberToIndices(squareNumber)
			if board[rowIndex][colIndex] == '❌' {
				xCount += 1
			} else if board[rowIndex][colIndex] == '⭕' {
				oCount += 1
			}
		}
		var winningPiece rune
		if xCount == 3 {
			winningPiece = '❌'
		} else if oCount == 3 {
			winningPiece = '⭕'
		}
		result[lineIndex] = lineEval{line: line, winningPiece: winningPiece}
	}
	return result
}

func main() {
	var board [3][3]rune

	player1 := player{name: "Player 1", piece: '❌'}
	player2 := player{name: "Player 2", piece: '⭕'}
	players := [2]player{player1, player2}
	turn := 0
	var hasWinner bool
	var hasDraw bool

	fmt.Print("\n\nWelcome to TIC ❌ TAC ⭕ TOE\n\n")

	for !(hasWinner || hasDraw) {

		printBoard(board)

		currentPlayer := players[turn%2]
		fmt.Print(string(currentPlayer.piece) + " " + currentPlayer.name + ", please enter your move: ")
		var moveInput string
		fmt.Scanln(&moveInput)

		number, inputError := strconv.Atoi(moveInput)
		rowIndex, colIndex, inputError2 := squareNumberToIndices(number)

		if inputError != nil || inputError2 != nil {
			fmt.Println("Please enter a square number from the board")
			continue
		}

		if board[rowIndex][colIndex] != 0 {
			fmt.Println("That square is taken, please choose an empty one")
			continue
		}

		board[rowIndex][colIndex] = currentPlayer.piece

		for _, eval := range evaluateLines(board) {
			if eval.winningPiece != 0 {
				hasWinner = true
				printBoardWithHighlight(board, eval.line[:])
				//				printBoard(board)
				fmt.Println("We have a winner: " + string(eval.winningPiece))
			}
		}

		turn += 1
	}
}
