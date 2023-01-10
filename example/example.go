package main

import (
	"context"
	"time"

	gui "github.com/grupawp/warships-gui"
)

const (
	username = "Kopytko I"
	opponent = "Stary wyjadacz"
)

func main() {
	ctx := context.Background()

	d := gui.NewDrawer(ctx)

	// username := d.DrawTextAndCatchInput(ctx, 5, 1, fmt.Sprintf("Podaj nazwę gracza: "))

	// d.DrawText(ctx, 5, 2, fmt.Sprintf("Nazwa gracza: %s", username))

	_ = d.DrawBoardAndCatchCoords(ctx, 5, 3, [10][10]gui.State{
		{gui.Hit, gui.Miss},
		{gui.Hit},
		{gui.Hit},
		{gui.Hit},
		{gui.Hit},
		{gui.Miss},
		{gui.Hit},
		{gui.Hit},
		{gui.Hit},
		{gui.Hit},
	})

	// d.DrawText(ctx, 60, 1, fmt.Sprintf("Kliknieta koordynata: %s", coord))

	// d.DrawText(ctx, 60, 1, fmt.Sprintf("Nazwa gracza: %s", opponent))

	// d.DrawBoard(ctx, 60, 3, [10][10]gui.State{
	// 	{gui.Hit, gui.Miss},
	// 	{gui.Hit},
	// 	{gui.Hit},
	// 	{gui.Miss},
	// 	{gui.Hit},
	// 	{gui.Miss},
	// 	{gui.Hit},
	// 	{gui.Hit},
	// 	{gui.Hit},
	// 	{gui.Miss},
	// }, false)

	// for {
	// 	coord := d.GetCoords(ctx)

	// 	d.DrawText(ctx, 60, 1, fmt.Sprintf("Kliknieta koordynata: %s", coord))

	// 	switch string(coord[0]) {
	// 	case "A", "B", "E":
	// 		d.SetState(ctx, coord, gui.Hit)

	// 	default:
	// 		d.SetState(ctx, coord, gui.Miss)
	// 	}

	// }
	time.Sleep(time.Second * 5)
}
