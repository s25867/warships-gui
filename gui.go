package gui

import (
	"context"

	"github.com/google/uuid"
	"github.com/grupawp/termloop"
	tl "github.com/grupawp/termloop"
)

type GUI struct {
	game      *tl.Game
	drawables map[uuid.UUID][]tl.Drawable
}

type Button struct {
	tl.Rectangle
	text     *tl.Text
	callback func()
	id       uuid.UUID
}

func NewButton(x, y int, width, height int, text string, callback func()) *Button {
	button := &Button{
		Rectangle: *tl.NewRectangle(x, y, width, height, tl.ColorWhite),
		text:      tl.NewText(x, y, text, tl.ColorBlack, tl.ColorDefault),
		callback:  callback,
		id:        uuid.New(),
	}
	return button
}

func (b *Button) ID() uuid.UUID {
	return b.id
}

func (b *Button) Contains(x, y int) bool {
	bx, by := b.Position()
	bw, bh := b.Size()
	return x >= bx && x < bx+bw && y >= by && y < by+bh
}

func (b *Button) Draw(s *termloop.Screen) {
	b.Rectangle.Draw(s)
}

func (b *Button) Tick(ev termloop.Event) {
	// todo
}

func (b *Button) Drawables() []tl.Drawable {
	return []tl.Drawable{&b.Rectangle, b.text}
}

func (g *GUI) DrawButton(x, y int, width, height int, text string, callback func()) {
	button := NewButton(x, y, width, height, text, callback)
	g.Draw(button)
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
