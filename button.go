package gui

import (
	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type Button struct {
	id       uuid.UUID
	gui      *GUI
	x        int
	y        int
	width    int
	height   int
	label    string
	callback func()
	rect     *tl.Rectangle
	txt      *tl.Text
}

// NewButton creates a new Button at the specified position with the given dimensions, label, and callback function.
func NewButton(gui *GUI, x, y, width, height int, label string, callback func()) *Button {
	btn := &Button{
		id:       uuid.New(),
		gui:      gui,
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		label:    label,
		callback: callback,
		rect:     tl.NewRectangle(x, y, width, height, tl.ColorBlue),
		txt:      tl.NewText(x+width/2-len(label)/2, y+height/2, label, tl.ColorWhite, tl.ColorBlue),
	}

	return btn
}

func (b *Button) ID() uuid.UUID {
	return b.id
}

func (b *Button) Drawables() []tl.Drawable {
	return []tl.Drawable{b.rect, b.txt}
}

// Tick handles mouse click events to trigger the button's callback.
func (b *Button) Tick(e tl.Event) {
	if e.Type == tl.EventMouse {
		b.gui.Log("Mouse event detected at (%d, %d)", e.MouseX, e.MouseY)
		if e.Key == tl.MouseLeft && e.MouseX >= b.x && e.MouseX < b.x+b.width && e.MouseY >= b.y && e.MouseY < b.y+b.height {
			b.gui.Log("Button clicked")
			b.callback()
		}
	}
}
