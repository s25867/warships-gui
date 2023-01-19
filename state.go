package gui

import (
	tl "github.com/JoelOtter/termloop"
)

type State string

var (
	// Hit means that some ship element is hit.
	Hit State = "Hit"
	// Miss means that shoot is missed.
	Miss State = "Miss"
	// Ship means that field represents whole ship element.
	Ship State = "Ship"
)

func (s State) colorAndText() (tl.Attr, string) {
	switch s {
	case Hit:
		return tl.ColorRed, "H"
	case Miss:
		return tl.RgbTo256Color(105, 105, 105), "M"
	case Ship:
		return tl.RgbTo256Color(139, 128, 0), "S"
	}
	return tl.ColorBlue, "~"
}

// clickAllowed returns true for states that are allowed to be clicked.
func (s State) clickAllowed() bool {
	switch s {
	case Hit:
		return false
	case Miss:
		return false
	case Ship:
		return true
	}

	return false
}
