package gui

import (
	"context"

	tl "github.com/JoelOtter/termloop"
)

// Config holds customizable parameters for drawer.
type Config struct {
	// EndKey is an end character using to exit game, 'space' is default.
	EndKey tl.Key
}

type drawer struct {
	game *tl.Game
	done bool
}

// NewDrawer returns new instance of drawer struct.
// It starts game in a new gouritine and takes control above the terminal.
func NewDrawer(c *Config) *drawer {
	game := tl.NewGame()
	game.Screen().SetFps(60)

	endKey := tl.KeySpace
	if c.EndKey != 0 {
		endKey = c.EndKey
	}
	game.SetEndKey(endKey)

	d := &drawer{game: game}

	go func() {
		game.Start()
		d.done = true
	}()

	return d
}

// IsGameRunning returns information about current game.
// It returns 'true' when game is in progress
// or returns 'false' when game is finished.
func (d drawer) IsGameRunning() bool {
	return !d.done
}

// NewBoard creates new instance of Board at given (x,y) position.
func (d drawer) NewBoard(x, y int, c *BoardConfig) (*Board, error) {
	return newBoard(x, y, c)
}

// DrawBoard draws 10x10 board with left upper corner begins at (x,y) point.
// It fills fields as it's given in 'states' argument.
func (d *drawer) DrawBoard(ctx context.Context, b *Board, states [10][10]State) {
	for _, v := range b.clicableStates {
		d.game.Screen().RemoveEntity(v)
	}
	b.clicableStates = nil
	for _, v := range b.states {
		d.game.Screen().RemoveEntity(v)
	}
	b.states = nil
	for _, v := range b.statesTexts {
		d.game.Screen().RemoveEntity(v)
	}
	b.statesTexts = nil

	b.setStates(states)

	d.drawBoard(ctx, b)
}

// DrawBoardAndCatchCoords does same as 'DrawBoard' method.
// But after drawing it waits for mouse action that returns clicked field, e.g. "B6".
// This allows to click only on the "see state".
func (d *drawer) DrawBoardAndCatchCoords(ctx context.Context, b *Board, states [10][10]State) string {
	for _, v := range b.clicableStates {
		d.game.Screen().RemoveEntity(v)
	}
	b.clicableStates = nil
	for _, v := range b.states {
		d.game.Screen().RemoveEntity(v)
	}
	b.states = nil
	for _, v := range b.statesTexts {
		d.game.Screen().RemoveEntity(v)
	}
	b.statesTexts = nil

	b.setClicableStates(states)

	d.drawBoard(ctx, b)
	wantClick = true

	return <-boardChan
}

// RemoveBoard removes existing Board from screen.
func (d *drawer) RemoveBoard(ctx context.Context, b *Board) {
	for _, v := range b.rectangles {
		d.game.Screen().RemoveEntity(v)
	}
	for _, v := range b.texts {
		d.game.Screen().RemoveEntity(v)
	}
	for _, v := range b.clicableStates {
		d.game.Screen().RemoveEntity(v)
	}
	for _, v := range b.states {
		d.game.Screen().RemoveEntity(v)
	}
	for _, v := range b.statesTexts {
		d.game.Screen().RemoveEntity(v)
	}
}

// DrawText creates a new Text at position (x, y).
// It sets the Text's text to be text.
// Returns a pointer to the new Text.
func (d *drawer) DrawText(ctx context.Context, x, y int, text string, tc *TextConfig) (*Text, error) {
	t, err := newText(x, y, text, tc)
	if err != nil {
		return nil, err
	}
	d.game.Screen().AddEntity(t)
	return t, nil
}

// RemoveText removes existing Text from screen.
func (d *drawer) RemoveText(ctx context.Context, t *Text) {
	d.game.Screen().RemoveEntity(t)
}

func (d *drawer) drawBoard(ctx context.Context, b *Board) {
	if !b.borderDrawed {
		for _, v := range b.rectangles {
			d.game.Screen().AddEntity(v)
		}
		for _, v := range b.texts {
			d.game.Screen().AddEntity(v)
		}
		b.borderDrawed = true
	}

	for _, cs := range b.clicableStates {
		d.game.Screen().AddEntity(cs)
	}
	for _, ss := range b.states {
		d.game.Screen().AddEntity(ss)
	}
	for _, st := range b.statesTexts {
		d.game.Screen().AddEntity(st)
	}
}
