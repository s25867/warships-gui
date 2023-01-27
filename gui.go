package gui

import (
	"context"
	"strconv"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

// Drawer exposes methods available when new game starts.
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

// NewDrawer returns new instance of Drawer interface.
// It starts game in a new gouritine and takes control above the terminal.
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

func NewDrawerWithOpts(o *Opts) (Drawer, error) {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	endKey := tl.KeySpace
	if o.EndKey != 0 {
		endKey = o.EndKey
	}
	game.SetEndKey(endKey)

	for k, v := range o.BasicStateColors {
		r, b, g, err := rgbFromString(v)
		if err != nil {
			return nil, err
		}
		colorState[k] = tl.RgbTo256Color(r, g, b)
	}

	d := &drawer{game: game}

	go func() {
		game.Start()
		d.done = true
	}()

	return d, nil
}

type Opts struct {
	// EndKey is an end character using to exit game, 'space' is default.
	EndKey tl.Key

	// BasicStateColors allows overwriting of any State to own RGB representation of color.
	// It has to be separated by '.' e.g. "120,0,19".
	// You don't need to set each of State but only this ones with should have got different color than default.
	BasicStateColors map[State]string
}

// IsClosed returns information about current game.
// When game is already closed then it returns 'true'.
func (d drawer) IsClosed() bool {
	return d.done
}

// DrawBoard draws 10x10 board with left upper corner begins at (x,y) point.
// It fills fields as it's given in 'states' argument.
func (d drawer) DrawBoard(ctx context.Context, x, y int, states [10][10]State) {
	d.drawBoard(ctx, x, y, states, false)
}

// DrawBoardAndCatchCoords does same as 'DrawBoard' method.
// But after drawing it waits for mouse action that returns clicked field, e.g. "B6".
// This allows to click only on the "see state".
func (d drawer) DrawBoardAndCatchCoords(ctx context.Context, x, y int, states [10][10]State) string {
	d.drawBoard(ctx, x, y, states, true)
	wantClick = true

	return <-boardChan
}

// DrawText creates a new Text at position (x, y).
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

func rgbFromString(s string) (int, int, int, error) {
	var colors [3]int
	for i, e := range strings.Split(s, ",") {
		n, err := strconv.Atoi(e)
		if err != nil {
			return -1, -1, -1, err
		}
		colors[i] = n
	}
	return colors[0], colors[1], colors[2], nil
}
