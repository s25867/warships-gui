package gui

import (
	"context"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

type Drawer interface {
	DrawBoard(ctx context.Context, x, y int, status [10][10]State)
	DrawBoardAndCatchCoords(ctx context.Context, x, y int, status [10][10]State) string
	DrawText(ctx context.Context, x, y int, text string)

	//DrawTextAndCatchInput(ctx context.Context, x, y int, text string) string
	// GetCoords(ctx context.Context) string
	// SetState(ctx context.Context, coords string, state State)
}

type drawer struct {
	game *tl.Game
}

func NewDrawer(ctx context.Context) Drawer {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	go game.Start()

	return &drawer{game: game}
}

const (
	width  = 3
	height = 1
)

var (
	wantClick = false
	// setState  = false
	// lastState State
	// stateID   string
)

func (d drawer) DrawBoard(ctx context.Context, x, y int, status [10][10]State) {
	d.drawBoard(ctx, x, y, status, false)
}

func (d drawer) DrawBoardAndCatchCoords(ctx context.Context, x, y int, status [10][10]State) string {
	d.drawBoard(ctx, x, y, status, true)
	wantClick = true

	return <-boardChan
}

func (d drawer) drawBoard(ctx context.Context, x, y int, status [10][10]State, clickable bool) {
	var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	for i := 1; i < 11; i++ {
		newX := i * width
		d.game.Screen().AddEntity(tl.NewRectangle(x+newX+i, y, width, height, tl.ColorWhite))
		d.game.Screen().AddEntity(tl.NewText(x+newX+i+(width/2), y+(height/2), letters[i-1], tl.ColorBlack, tl.ColorWhite))

		newY := i * height
		d.game.Screen().AddEntity(tl.NewRectangle(x, y+newY+i, width, height, tl.ColorWhite))
		d.game.Screen().AddEntity(tl.NewText(x+(width/2), y+newY+i+(height/2), fmt.Sprintf("%d", i), tl.ColorBlack, tl.ColorWhite))
	}

	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i * width
			newY := j * height
			color, text := status[i-1][j-1].ColorAndText()

			if clickable {
				d.game.Screen().AddEntity(
					newClickable(
						fmt.Sprintf("%s%d", letters[i-1], j),
						tl.NewRectangle(x+newX+i, y+newY+j, width, height, color)))
			} else {
				d.game.Screen().AddEntity(tl.NewRectangle(x+newX+i, y+newY+j, width, height, color))
			}
			d.game.Screen().AddEntity(tl.NewText(x+newX+i+(width/2), y+newY+j+(height/2), text, tl.ColorBlack, color))
		}
	}
}

func (d *drawer) DrawText(ctx context.Context, x, y int, text string) {
	d.game.Screen().AddEntity(tl.NewText(x, y, text, tl.ColorWhite, tl.ColorBlack))
}

// func (d *drawer) DrawTextAndCatchInput(ctx context.Context, x, y int, text string) string {
// 	t := newInputText(tl.NewText(x, y, text, tl.ColorWhite, tl.ColorBlack))

// 	d.game.Screen().AddEntity(t)

// 	for <-t.done {
// 		// d.game.Screen().RemoveEntity(t)
// 		return t.readValue
// 	}

// 	return ""
// }

// func (d *drawer) GetCoords(ctx context.Context) string {
// 	wantClick = true
// 	return <-boardChan
// }

// func (d *drawer) SetState(ctx context.Context, coords string, state State) {
// 	switch state {
// 	case Hit:
// 		lastState = Hit
// 	case Miss:
// 		lastState = Miss
// 	}

// 	stateID = coords
// 	setState = true
// }
