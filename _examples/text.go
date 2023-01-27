package main

// This file shows how simple is to draw and edit some text inside a terminal.

import (
	"context"
	"time"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(ctx)

	t := d.DrawText(ctx, 2, 2, "Example text.")
	defer d.RemoveText(ctx, t) // it's not necessary, but it's a good practice to cleanup when have that possibility

	time.Sleep(5 * time.Second) // sleep to see a difference

	t.SetText("Example text after 5 seconds of sleeping")

	for {
		if d.IsClosed() { // wait until escape character has been pressed
			return
		}
	}
}
