package main

// This file shows how simple is to draw a single board,
// catch clicked field and draw it as a text inside a terminal.

import (
	"context"
	"fmt"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(ctx)

	t := d.DrawText(ctx, 2, 2, "") // initialize some text object

	coords := d.DrawBoardAndCatchCoords(ctx, 2, 4, [10][10]gui.State{}) // draw empty board at position (2,4)

	t.SetText(fmt.Sprintf("You clicked: %v ", coords))

	for {
		if d.IsClosed() { // wait until escape character has been pressed
			return
		}
	}
}
