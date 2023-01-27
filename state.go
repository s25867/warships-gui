package gui

import (
	tl "github.com/JoelOtter/termloop"
)

// State represents available object that can be draw on the Board.
type State string

var (
	// Hit means that some ship element is hit.
	Hit State = "Hit"
	// Miss means that shoot is missed.
	Miss State = "Miss"
	// Ship means that field represents whole ship element.
	Ship State = "Ship"
	// None means that field has no defined State.
	None State = "None"
)

var colorState map[State]tl.Attr = map[State]tl.Attr{
	Hit:  tl.ColorRed,
	Miss: tl.RgbTo256Color(105, 105, 105),
	Ship: tl.RgbTo256Color(139, 128, 0),
	None: tl.ColorBlue,
}

type StateColors struct {
	Hit  tl.Attr
	Miss tl.Attr
	Ship tl.Attr
	None tl.Attr
}

func (s State) colorAndText() (tl.Attr, string) {
	switch s {
	case Hit:
		return colorState[Hit], "H"
	case Miss:
		return colorState[Miss], "M"
	case Ship:
		return colorState[Ship], "S"
	}
	return colorState[None], "~"
}

// clickAllowed returns true for states that are allowed to be clicked.
func (s State) clickAllowed() bool {
	switch s {
	case Hit:
		return false
	case Miss:
		return false
	case Ship:
		return false
	}

	return true
}
