package gui

import tl "github.com/JoelOtter/termloop"

type Text struct {
	*tl.Text
}

func newText(t *tl.Text) *Text {
	return &Text{Text: t}
}

// SetText sets the text of the Text to be text and returns Text.
func (t *Text) SetText(text string) *Text {
	t.Text.SetText(text)
	return t
}
