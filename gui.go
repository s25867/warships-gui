package gui

import (
	"context"

	tl "github.com/JoelOtter/termloop"
)

type Drawer interface {
	IsClosed() bool
	DrawBoard(ctx context.Context, x, y int, states [10][10]State)
	DrawBoardAndCatchCoords(ctx context.Context, x, y int, states [10][10]State) string
	DrawText(ctx context.Context, x, y int, text string) *Text
	RemoveText(ctx context.Context, t *Text)
}

type drawer struct {
	game *tl.Game
	done bool
}

func NewDrawer(ctx context.Context) Drawer {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	game.SetEndKey(tl.KeySpace)

	d := &drawer{game: game}

	go func() {
		game.Start()
		d.done = true
	}()

	return d
}

// IsClosed returns information about current game.
// When game should be quitted then it returns 'true'.
func (d drawer) IsClosed() bool {
	return d.done
}

const (
	width  = 3
	height = 1
)

var (
	wantClick = false
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
	board := newBoard(x, y)

	for _, rec := range board.Rectangles {
		d.game.Screen().AddEntity(rec)
	}
	for _, t := range board.Texts {
		d.game.Screen().AddEntity(t)
	}

	if clickable {
		board.drawClicableStates(ctx, x, y, status)
	} else {
		board.drawStates(ctx, x, y, status)
	}

	for _, cs := range board.ClicableStates {
		d.game.Screen().AddEntity(cs)
	}
	for _, ss := range board.States {
		d.game.Screen().AddEntity(ss)
	}
	for _, st := range board.StatesTexts {
		d.game.Screen().AddEntity(st)
	}
}

// DrawText creates a new Text, at position (x, y).
// It sets the Text's text to be text.
// Returns a pointer to the new Text.
func (d *drawer) DrawText(ctx context.Context, x, y int, text string) *Text {
	t := newText(tl.NewText(x, y, text, defaultTextFG, defaultTextBG))
	d.game.Screen().AddEntity(t)
	return t
}

// RemoveText removes existing Text from screen.
func (d *drawer) RemoveText(ctx context.Context, t *Text) {
	d.game.Screen().RemoveEntity(t)
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
