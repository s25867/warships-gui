package gui

import (
	"context"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

type Drawer interface {
	DrawBoard(ctx context.Context)
	GetCoords(ctx context.Context) string
	SetState(ctx context.Context, coords string, state State)
}

type drawer struct {
	game *tl.Game
}

func NewDrawer(ctx context.Context) Drawer {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	return &drawer{game: game}
}

const (
	scale  = 2
	width  = 2
	height = 1
)

func (d drawer) DrawBoard(ctx context.Context) {

	var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

	for i := 1; i < 11; i++ {
		x := i * width * scale
		d.game.Screen().AddEntity(tl.NewRectangle(1+x+(i), 1, width*scale, height*scale, tl.ColorWhite))
		d.game.Screen().AddEntity(tl.NewText(1+x+(1*i)+2, 1, letters[i-1], tl.ColorBlack, tl.ColorWhite))

		y := i * scale * height
		d.game.Screen().AddEntity(tl.NewRectangle(1, 1+y+i, width*scale, height*scale, tl.ColorWhite))
		d.game.Screen().AddEntity(tl.NewText(2, 2+y+i, fmt.Sprintf("%d", i), tl.ColorBlack, tl.ColorWhite))
	}

	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			x := i * width * scale
			y := j * scale * height
			d.game.Screen().AddEntity(
				newClickable(
					fmt.Sprintf("%s%d", letters[i-1], j),
					tl.NewRectangle(1+x+i, 1+y+j, width*scale, height*scale, tl.ColorBlue)))
		}
	}

	d.game.Start()
}

var (
	wantClick = false
	setState  = false
	lastState State
	stateID   string
)

func (d *drawer) GetCoords(ctx context.Context) string {
	wantClick = true
	return <-boardChan
}

func (d *drawer) SetState(ctx context.Context, coords string, state State) {
	switch state {
	case Hit:
		lastState = Hit
	case Miss:
		lastState = Miss
	}

	stateID = coords
	setState = true
}
