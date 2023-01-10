package gui

import (
	tl "github.com/JoelOtter/termloop"
)

type State string

var (
	Hit  State = "Hit"
	Miss State = "Miss"
	Ship State = "Ship"
)

func (s State) ColorAndText() (tl.Attr, string) {
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
