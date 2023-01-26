package gui

import (
	"context"
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

const (
	defaultTextBG = tl.ColorWhite
	defaultTextFG = tl.ColorBlack

	fieldWidth  = 3
	fieldHeight = 1
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
		newX := i*fieldWidth + i
		newY := i*fieldHeight + i

		b.Rectangles = append(b.Rectangles, tl.NewRectangle(x+newX, y, fieldWidth, fieldHeight, tl.ColorWhite))
		b.Rectangles = append(b.Rectangles, tl.NewRectangle(x, y+newY, fieldWidth, fieldHeight, tl.ColorWhite))

		b.Texts = append(b.Texts, tl.NewText(x+newX+(fieldWidth/2), y+(fieldHeight/2), letters[i-1], b.textFG, b.textBG))
		b.Texts = append(b.Texts, tl.NewText(x+(fieldWidth/2), y+newY+(fieldHeight/2), fmt.Sprintf("%d", i), b.textFG, b.textBG))

	}

	return b
}

func (b *Board) drawStates(ctx context.Context, x, y int, states [10][10]State) *Board {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i*fieldWidth + i
			newY := j*fieldHeight + j
			color, text := states[i-1][j-1].colorAndText()

			b.States = append(b.States, tl.NewRectangle(x+newX, y+newY, fieldWidth, fieldHeight, color))
			b.StatesTexts = append(b.StatesTexts, tl.NewText(x+newX+(fieldWidth/2), y+newY+(fieldHeight/2), text, tl.ColorBlack, color))
		}
	}

	return b
}

func (b *Board) drawClicableStates(ctx context.Context, x, y int, states [10][10]State) *Board {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i*fieldWidth + i
			newY := j*fieldHeight + j
			color, text := states[i-1][j-1].colorAndText()

			rec := tl.NewRectangle(x+newX, y+newY, fieldWidth, fieldHeight, color)
			if states[i-1][j-1].clickAllowed() {
				b.ClicableStates = append(b.ClicableStates, newClickable(
					fmt.Sprintf("%s%d", letters[i-1], j), rec))
			} else {
				b.States = append(b.States, rec)
			}

			b.StatesTexts = append(b.StatesTexts, tl.NewText(x+newX+(fieldWidth/2), y+newY+(fieldHeight/2), text, tl.ColorBlack, color))
		}
	}

	return b
}
