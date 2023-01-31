package main

// This file shows how simple is to draw a single board,
// catch clicked field and draw it as a text inside a terminal.

import (
	"context"
	"fmt"
	"log"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(&gui.Config{})

	b, err := d.NewBoard(2, 4, &gui.BoardConfig{})
	if err != nil {
		log.Fatal(err)
	}
	defer d.RemoveBoard(ctx, b)

	coords := d.DrawBoardAndCatchCoords(ctx, b, [10][10]gui.State{}) // draw empty board at position (2,4)

	t := d.DrawText(ctx, 2, 2, "") // initialize some text object
	t.SetText(fmt.Sprintf("You clicked: %v ", coords))

	for {
		if !d.IsGameRunning() { // wait until escape character has been pressed
			return
		}
	}
}
