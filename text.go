package gui

import (
	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
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
func NewText(x, y int, text string, config *TextConfig) *Text {
	if config == nil {
		config = NewTextConfig()
	}
	t := tl.NewText(x, y, text, config.FgColor.toAttr(), config.BgColor.toAttr())
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
	t.t.SetColor(c.toAttr(), bg)
}

// SetBgColor sets the background color of the Text struct.
func (t *Text) SetBgColor(c Color) {
	fg, _ := t.t.Color()
	t.t.SetColor(fg, c.toAttr())
}

func (t *Text) ID() uuid.UUID {
	return t.id
}

func (t *Text) Drawables() []tl.Drawable {
	return []tl.Drawable{t.t}
}
