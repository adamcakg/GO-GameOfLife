package main

import (
	"fmt"
)

type Game struct {
	board [][]bool
}

func initializeGame(rows, cols int) *Game {

	var board = make([][]bool, rows)

	for i := range board {
		board[i] = make([]bool, cols)
	}

	return &Game{board: board}
}

func (game *Game) randomizeBoard() {
	game.board[1][2] = true
	game.board[1][1] = true
	game.board[1][0] = true
}

func (game *Game) print() {
	//Cleares the console
	fmt.Print("\033[H\033[2J")

	for i := range game.board {
		for j := range game.board[i] {
			fmt.Print(game.board[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()
}

func (game *Game) tick() {

	var board = make([][]bool, len(game.board))

	for i := range board {
		board[i] = make([]bool, len(game.board[0]))
	}

	for i := range game.board {
		for j := range game.board[i] {
			count := 0
			width := len(game.board[i])
			height := len(game.board)

			if isIndexValid(i-1, j-1, width, height) && game.board[i-1][j-1] {
				count++
			}

			if isIndexValid(i, j-1, width, height) && game.board[i][j-1] {
				count++
			}

			if isIndexValid(i+1, j-1, width, height) && game.board[i+1][j-1] {
				count++
			}

			if isIndexValid(i-1, j, width, height) && game.board[i-1][j] {
				count++
			}

			if isIndexValid(i+1, j, width, height) && game.board[i+1][j] {
				count++
			}

			if isIndexValid(i-1, j+1, width, height) && game.board[i-1][j+1] {
				count++
			}

			if isIndexValid(i, j+1, width, height) && game.board[i][j+1] {
				count++
			}

			if isIndexValid(i+1, j+1, width, height) && game.board[i+1][j+1] {
				count++
			}

			switch {
			case count < 2:
				board[i][j] = false
			case count >= 2 && count <= 3:
				board[i][j] = game.board[i][j]

				if !game.board[i][j] && count == 3 {
					board[i][j] = true
				}
			case count > 3:
				board[i][j] = false
			}
		}
	}

	game.board = board
}

func isIndexValid(x, y, width, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}
