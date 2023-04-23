package gui

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/google/uuid"
)

// Drawable represents a collection of related objects 
// that can be drawn on the screen.
type Drawable interface {
	ID() uuid.UUID
	Drawables() []tl.Drawable
}
