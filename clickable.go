package gui

import tl "github.com/JoelOtter/termloop"

type clickable struct {
	*tl.Rectangle
	id string
}

func newClickable(id string, rec *tl.Rectangle) *clickable {
	return &clickable{id: id, Rectangle: rec}
}

func (c *clickable) Tick(e tl.Event) {

	switch e.Key {
	case tl.MouseLeft:
		c.processClick(e)
	}

	if setState && c.id == stateID {
		c.stateChanged()
	}
}

func (c *clickable) processClick(e tl.Event) {
	x, y := c.Position()
	if e.MouseX >= x && e.MouseY >= y && e.MouseX <= (x+width*scale) && e.MouseY <= (y+height*scale) {
		if wantClick {
			defer func() { wantClick = false }()
			boardChan <- c.id
		}
	}
}

func (c *clickable) stateChanged() {

	defer func() { setState = false }()
	switch lastState {
	case Hit:
		c.SetColor(tl.ColorRed)
	case Miss:
		c.SetColor(tl.ColorBlack)
	}

}
