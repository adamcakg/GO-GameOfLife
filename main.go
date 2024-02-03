package main

import (
	"fmt"
	"time"
)

func main() {

	var rows = 3
	var cols = 3

	fmt.Print("Select grid size: ")
	fmt.Scan(&rows, &cols)

	game := initializeGame(rows, cols)

	game.randomizeBoard()

	for {
		game.print()

		game.tick()

		time.Sleep(2 * time.Second)
	}

}
