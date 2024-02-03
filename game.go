package main

import (
	"fmt"
	"math/rand"
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

func (game *Game) generateBoard(mode GameMode) {

	switch mode {
	case I:
		game.board[1][2] = true
		game.board[1][1] = true
		game.board[1][0] = true
	case Random:
		for i := range game.board {
			for j := range game.board[i] {
				if rand.Int()%2 == 0 {
					game.board[i][j] = true
				}
			}
		}
	case Toad:
		game.board[0][4] = true
		game.board[1][4] = true
		game.board[1][3] = true
		game.board[2][4] = true
		game.board[2][3] = true
		game.board[3][3] = true
	}
}

func (game *Game) print() {
	//Cleares the console
	fmt.Print("\033[H\033[2J")

	for i := range game.board {
		for j := range game.board[i] {
			if game.board[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print("  ")
			}

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
