package gui

import (
	"context"
	"sync"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type GUI struct {
	game      *tl.Game
	drawables map[uuid.UUID][]tl.Drawable
	screens   map[string]*tl.Screen
	mutex     sync.Mutex
}

// NewGUI returns a new GUI instance.
// If debug is true, the GUI will print logs to the terminal
// after exiting.
func NewGUI(debug bool) *GUI {
	game := tl.NewGame()
	game.Screen().SetFps(60)
	game.SetDebugOn(debug)
	screens := make(map[string]*tl.Screen)
	screens["init"] = game.Screen()
	d := &GUI{
		game:      game,
		drawables: make(map[uuid.UUID][]tl.Drawable),
		screens:   screens,
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
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.drawables[d.ID()] = d.Drawables()
	for _, di := range g.drawables[d.ID()] {
		g.game.Screen().AddEntity(di)
	}
}

func (g *GUI) NewScreen(name string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	screen := tl.NewScreen()
	screen.SetFps(60)
	g.screens[name] = screen
}

func (g *GUI) RemoveScreen(name string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	delete(g.screens, name)
}

func (g *GUI) SetScreen(name string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	screen, ok := g.screens[name]
	if !ok {
		return false
	}
	g.game.SetScreen(screen)
	return true
}

func (g *GUI) ScreenNames() []string {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	names := make([]string, 0)
	for name := range g.screens {
		names = append(names, name)
	}
	return names
}

// Remove removes the given Drawable from the screen.
func (g *GUI) Remove(d Drawable) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

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
