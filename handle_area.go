package gui

import (
	"context"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

type HandleArea struct {
	id         uuid.UUID
	clickables []*rectangle
	ch         chan string
}

func NewHandleArea(objs map[string]Spatial) *HandleArea {
	ch := make(chan string)
	clickables := createClickables(objs, ch)
	return &HandleArea{
		id:         uuid.New(),
		clickables: clickables,
		ch:         ch,
	}
}

func (area *HandleArea) SetClickablesOn(objs map[string]Spatial) {
	clickables := createClickables(objs, area.ch)
	area.clickables = clickables
}

func createClickables(objs map[string]Spatial, ch chan<- string) (clickables []*rectangle) {
	for key, obj := range objs {
		x, y := obj.Position()
		w, h := obj.Size()
		rec := tl.NewRectangle(x, y, w, h, tl.ColorWhite)
		clickables = append(clickables, newClickableRectangle(rec, key, ch))
	}
	return
}

func (area *HandleArea) ID() uuid.UUID {
	return area.id
}

func (area *HandleArea) Drawables() []tl.Drawable {
	d := []tl.Drawable{}
	for _, c := range area.clickables {
		d = append(d, tl.Drawable(c))
	}
	return d
}

func (area *HandleArea) Listen(ctx context.Context) string {
	select {
	case s := <-area.ch:
		return s
	case <-ctx.Done():
		return ""
	}
}

func (area *HandleArea) GetClickables() []*rectangle {
	return area.clickables
}

func (area *HandleArea) GetChannel() chan string {
	return area.ch
}
