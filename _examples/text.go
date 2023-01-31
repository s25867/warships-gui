package main

// This file shows how simple is to draw and edit some text inside a terminal.

import (
	"context"
	"log"
	"time"

	gui "github.com/grupawp/warships-gui"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(&gui.Config{})

	t, err := d.DrawText(ctx, 2, 2, "Example text.", nil)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second) // sleep to see a difference

	t.SetText("Example text after 5 seconds of sleeping")
	time.Sleep(2 * time.Second)
	d.RemoveText(ctx, t)

	t, err = d.DrawText(ctx, 2, 2, "Example text with custom background color.", &gui.TextConfig{BackgroundColor: "153,255,51"})
	if err != nil {
		log.Fatal(err)
	}
	defer d.RemoveText(ctx, t) // it's not necessary, but it's a good practice to cleanup when have that possibility

	for {
		if !d.IsGameRunning() { // wait until escape character has been pressed
			return
		}
	}
}
