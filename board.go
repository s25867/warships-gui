package gui

import (
	"context"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

const (
	defaultTextBG = tl.ColorWhite
	defaultTextFG = tl.ColorBlack
)

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

type Board struct {
	Rectangles     []*tl.Rectangle
	Texts          []*tl.Text
	ClicableStates []*clickable
	States         []*tl.Rectangle
	StatesTexts    []*tl.Text

	textBG tl.Attr
	textFG tl.Attr
}

func newBoard(x, y int) *Board {

	b := &Board{
		textBG: defaultTextBG,
		textFG: defaultTextFG,
	}

	for i := 1; i < 11; i++ {
		newX := i*width + i
		newY := i*height + i

		b.Rectangles = append(b.Rectangles, tl.NewRectangle(x+newX, y, width, height, tl.ColorWhite))
		b.Rectangles = append(b.Rectangles, tl.NewRectangle(x, y+newY, width, height, tl.ColorWhite))

		b.Texts = append(b.Texts, tl.NewText(x+newX+(width/2), y+(height/2), letters[i-1], b.textFG, b.textBG))
		b.Texts = append(b.Texts, tl.NewText(x+(width/2), y+newY+(height/2), fmt.Sprintf("%d", i), b.textFG, b.textBG))

	}

	return b
}

func (b *Board) drawStates(ctx context.Context, x, y int, states [10][10]State) *Board {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i * width
			newY := j * height
			color, text := states[i-1][j-1].colorAndText()

			b.States = append(b.States, tl.NewRectangle(x+newX, y+newY, width, height, color))
			b.StatesTexts = append(b.StatesTexts, tl.NewText(x+newX+(width/2), y+newY+(height/2), text, tl.ColorBlack, color))
		}
	}

	return b
}

func (b *Board) drawClicableStates(ctx context.Context, x, y int, states [10][10]State) *Board {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i*width + i
			newY := j*height + j
			color, text := states[i-1][j-1].colorAndText()

			if states[i-1][j-1].clickAllowed() {
				b.ClicableStates = append(b.ClicableStates, newClickable(
					fmt.Sprintf("%s%d", letters[i-1], j),
					tl.NewRectangle(x+newX+i, y+newY+j, width, height, color)))
			} else {
				b.States = append(b.States, tl.NewRectangle(x+newX, y+newY, width, height, color))
			}

			b.StatesTexts = append(b.StatesTexts, tl.NewText(x+newX+(width/2), y+newY+(height/2), text, tl.ColorBlack, color))
		}
	}

	return b
}
