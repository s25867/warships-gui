package main

import (
	"context"

	gui "github.com/grupawp/warships-gui"
)

type MyStr string

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(ctx)

	go d.DrawBoard(ctx)

	for {
		coord := d.GetCoords(ctx)

		switch string(coord[0]) {
		case "A", "B", "E":
			d.SetState(ctx, coord, gui.Hit)

		default:
			d.SetState(ctx, coord, gui.Miss)
		}
	}
}
