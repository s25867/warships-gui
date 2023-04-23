package gui

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/google/uuid"
)

type GUI struct {
	game      *tl.Game
	drawables map[uuid.UUID][]tl.Drawable
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
	}

	return d
}

// Start displays the GUI and blocks until the it exits.
// If endKey is not nil, the GUI will exit when the key is pressed.
// Default if Ctrl + C.
func (g GUI) Start(endKey *tl.Key) {
	if endKey != nil {
		g.game.SetEndKey(*endKey)
	}
	g.game.Start()
}

// Draw draws the given Drawable on the screen.
func (g *GUI) Draw(d Drawable) {
	g.drawables[d.ID()] = d.Drawables()
	for _, di := range g.drawables[d.ID()] {
		g.game.Screen().AddEntity(di)
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
