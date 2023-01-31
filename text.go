package gui

import tl "github.com/JoelOtter/termloop"

const (
	defaultTextBG = tl.ColorWhite
	defaultTextFG = tl.ColorBlack
)

type Text struct {
	*tl.Text
}

// TextConfig holds configuration parameters for Text struct.
type TextConfig struct {
	ForegroundColor string
	BackgroundColor string
}

func newText(x, y int, text string, tc *TextConfig) (*Text, error) {
	fgColor, bgColor := defaultTextFG, defaultTextBG
	if tc != nil && tc.ForegroundColor != "" {
		r, b, g, err := rgbFromString(tc.ForegroundColor)
		if err != nil {
			return nil, err
		}
		fgColor = tl.RgbTo256Color(r, g, b)
	}
	if tc != nil && tc.BackgroundColor != "" {
		r, b, g, err := rgbFromString(tc.BackgroundColor)
		if err != nil {
			return nil, err
		}
		bgColor = tl.RgbTo256Color(r, g, b)
	}
	t := tl.NewText(x, y, text, fgColor, bgColor)

	return &Text{Text: t}, nil
}

// SetText sets the text of the Text to be text and returns Text.
func (t *Text) SetText(text string) *Text {
	t.Text.SetText(text)
	return t
}
