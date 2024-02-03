package main

import (
	"reflect"
	"testing"
)

func TestBoard3x3WithI(t *testing.T) {

	game := initializeGame(3, 3)

	game.generateBoard(I)

	t.Log(game.board)

	if !reflect.DeepEqual(game.board, [][]bool{{false, false, false}, {true, true, true}, {false, false, false}}) {
		t.Fatalf("Generation failed")
	}

	game.tick()

	if !reflect.DeepEqual(game.board, [][]bool{{false, true, false}, {false, true, false}, {false, true, false}}) {
		t.Fatalf("First tick failed")
	}

	game.tick()

	if !reflect.DeepEqual(game.board, [][]bool{{false, false, false}, {true, true, true}, {false, false, false}}) {
		t.Fatalf("Second tick failed")
	}
}
