package gui

import (
	"unicode/utf8"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

// TextInput represents a text input Input.
type TextInput struct {
	*tl.Text
	content   string
	maxLength int
	x, y      int
	id        uuid.UUID
}

// NewTextInput creates a new TextInput.
func NewTextInput(x, y, maxLength int) *TextInput {
	return &TextInput{
		Text:      tl.NewText(x, y, "", tl.ColorWhite, tl.ColorBlack),
		maxLength: maxLength,
		x:         x,
		y:         y,
		id:        uuid.New(),
	}
}

// Draw updates the TextInput text and draws it on the screen.
func (ta *TextInput) Draw(screen *tl.Screen) {
	ta.Text.SetText(ta.content)
	ta.Text.Draw(screen)
}

// Tick handles key press events for the TextInput.
func (ta *TextInput) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyBackspace || event.Key == tl.KeyBackspace2 {
			if len(ta.content) > 0 {
				_, size := utf8.DecodeLastRuneInString(ta.content)
				ta.content = ta.content[:len(ta.content)-size]
			}
		} else if utf8.RuneCountInString(ta.content) < ta.maxLength {
			ta.content += string(event.Ch) // Convert the rune to a string before adding it
		}
	}
}

// GetContent returns the current content of the TextInput.
func (ta *TextInput) GetContent() string {
	return ta.content
}

// ID returns the unique identifier of the TextInput.
func (ta *TextInput) ID() uuid.UUID {
	return ta.id
}

// Drawables returns the TextInput as a slice of tl.Drawable.
func (ta *TextInput) Drawables() []tl.Drawable {
	return []tl.Drawable{ta}
}
