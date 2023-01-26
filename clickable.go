package gui

import tl "github.com/JoelOtter/termloop"

var (
	boardChan = make(chan string)
	wantClick = false
)

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
}

func (c *clickable) processClick(e tl.Event) {
	x, y := c.Position()
	if e.MouseX >= x && e.MouseY >= y && e.MouseX <= (x+fieldWidth) && e.MouseY <= (y+fieldHeight) {
		if wantClick {
			defer func() { wantClick = false }()
			boardChan <- c.id
		}
	}
}
