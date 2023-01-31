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
	None State = ""
)

var defaultColorState map[State]stateColorAndChar = map[State]stateColorAndChar{
	Hit:  {tl.ColorRed, "H"},
	Miss: {tl.RgbTo256Color(105, 105, 105), "M"},
	Ship: {tl.RgbTo256Color(139, 128, 0), "S"},
	None: {tl.ColorBlue, "~"},
}

type stateColorAndChar struct {
	color tl.Attr
	char  string
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
