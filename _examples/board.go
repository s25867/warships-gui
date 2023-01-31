package main

// This file shows how simple is to draw a single board inside a terminal.

import (
	"context"
	"log"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(&gui.Config{})
	b, err := d.NewBoard(2, 2, &gui.BoardConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer d.RemoveBoard(ctx, b)

	d.DrawBoard(ctx, b, [10][10]gui.State{}) // draw empty board at position (2,2)

	for {
		if !d.IsGameRunning() { // wait until escape character has been pressed
			return
		}
	}
}
_
