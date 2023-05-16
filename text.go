package gui

import (
	tl "github.com/grupawp/termloop"
	"github.com/google/uuid"
)

type Text struct {
	id uuid.UUID
	t  *tl.Text
}

// TextConfig holds configuration parameters for Text struct.
type TextConfig struct {
	FgColor Color
	BgColor Color
}

// NewTextConfig returns a new config with default values.
func NewTextConfig() *TextConfig {
	return &TextConfig{
		FgColor: Black,
		BgColor: White,
	}
}

// NewText returns a new Text struct.
// X and Y are the coordinates of the top left corner of the text.
// If no config is provided, default values are used.
func NewText(x, y int, text string, cfg *TextConfig) *Text {
	if cfg == nil {
		cfg = NewTextConfig()
	}
	t := tl.NewText(x, y, text, cfg.FgColor.toAttr(), cfg.BgColor.toAttr())
	return &Text{
		id: uuid.New(),
		t:  t,
	}
}

// SetText sets the text of the Text struct.
func (t *Text) SetText(text string) {
	t.t.SetText(text)
}

// SetFgColor sets the foreground color of the Text struct.
func (t *Text) SetFgColor(c Color) {
	_, bg := t.t.Color()
	t.t.SetColor(tl.RgbTo256Color(int(c.Red), int(c.Green), int(c.Blue)), bg)
}

// SetBgColor sets the background color of the Text struct.
func (t *Text) SetBgColor(c Color) {
	fg, _ := t.t.Color()
	t.t.SetColor(fg, tl.RgbTo256Color(int(c.Red), int(c.Green), int(c.Blue)))
}

func (t *Text) ID() uuid.UUID {
	return t.id
}

func (t *Text) Drawables() []tl.Drawable {
	return []tl.Drawable{t.t}
}
