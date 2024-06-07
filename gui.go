package gui

import (
	"context"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type GUI struct {
	game      *tl.Game
	drawables map[uuid.UUID][]tl.Drawable
	buttons   map[uuid.UUID]*Button
}

// NewGUI returns a new GUI instance.
// If debug is true, the GUI will print logs to the terminal
// after exiting.
func NewGUI(debug bool) *GUI {
	game := tl.NewGame()
	game.Screen().SetFps(60)
	game.SetDebugOn(debug)

	d := &GUI{
		game:      game,
		drawables: make(map[uuid.UUID][]tl.Drawable),
		buttons:   make(map[uuid.UUID]*Button),
	}

	return d
}

// Start displays the GUI and blocks until endKey is pressed or
// context is done. If endKey is nil it defaults to Ctrl+C.
func (g *GUI) Start(ctx context.Context, endKey *tl.Key) {
	if endKey != nil {
		g.game.SetEndKey(*endKey)
	}
	g.game.Start(ctx)
}

// Draw draws the given Drawable on the screen.
func (g *GUI) Draw(d Drawable) {
	g.drawables[d.ID()] = d.Drawables()
	for _, di := range g.drawables[d.ID()] {
		g.game.Screen().AddEntity(di)
	}
}

// DrawButton adds a button to the GUI at the specified position with the given label and callback function.
func (g *GUI) DrawButton(x, y, width, height int, label string, callback func()) {
	button := NewButton(x, y, width, height, label, callback)
	g.buttons[button.ID()] = button
	for _, drawable := range button.Drawables() {
		g.game.Screen().AddEntity(drawable)
	}
}

// Remove removes the given Drawable from the screen.
func (g *GUI) Remove(d Drawable) {
	for _, drawable := range g.drawables[d.ID()] {
		g.game.Screen().RemoveEntity(drawable)
	}
	delete(g.drawables, d.ID())
}

// Log takes a log string and additional parameters, which can be
// substituted into the string using standard fmt.Printf rules.
// If debug mode is on, the formatted log will be printed to the
// terminal when GUI exits.
func (g *GUI) Log(format string, a ...any) {
	g.game.Log(format, a...)
}
