package main

// This file shows how simple is to draw a single board inside a terminal.

import (
	"context"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(ctx)

	d.DrawBoard(ctx, 2, 2, [10][10]gui.State{}) // draw empty board at position (2,2)

	for {
		if d.IsClosed() { // wait until escape character has been pressed
			return
		}
	}
}
