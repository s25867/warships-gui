package gui

import (
	"fmt"
	"strconv"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

const (
	defaultTextBG = tl.ColorWhite
	defaultTextFG = tl.ColorBlack

	fieldWidth  = 3
	fieldHeight = 1
)

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

// BoardConfig holds configuration parameters for Board struct.
type BoardConfig struct {
	HitColor  string
	MissColor string
	ShipColor string

	HitChar  byte
	MissChar byte
	ShipChar byte
}

type Board struct {
	rectangles     []*tl.Rectangle
	texts          []*tl.Text
	clicableStates []*clickable
	states         []*tl.Rectangle
	statesTexts    []*tl.Text

	x int
	y int

	statesConf map[State]stateColorAndChar

	borderDrawed bool
}

func newBoard(x, y int, c *BoardConfig) (*Board, error) {
	statesConf, err := parseStatesFromConfig(c)
	if err != nil {
		return nil, err
	}
	b := &Board{
		x:          x,
		y:          y,
		statesConf: statesConf,
	}

	for i := 1; i < 11; i++ {
		newX := i*fieldWidth + i
		newY := i*fieldHeight + i

		b.rectangles = append(b.rectangles, tl.NewRectangle(x+newX, y, fieldWidth, fieldHeight, tl.ColorWhite))
		b.rectangles = append(b.rectangles, tl.NewRectangle(x, y+newY, fieldWidth, fieldHeight, tl.ColorWhite))

		b.texts = append(b.texts, tl.NewText(x+newX+(fieldWidth/2), y+(fieldHeight/2), letters[i-1], defaultTextFG, defaultTextBG))
		b.texts = append(b.texts, tl.NewText(x+(fieldWidth/2), y+newY+(fieldHeight/2), fmt.Sprintf("%d", i), defaultTextFG, defaultTextBG))

	}

	return b, nil
}

func (b *Board) setStates(states [10][10]State) *Board {
	b.states = make([]*tl.Rectangle, 0)
	b.statesTexts = make([]*tl.Text, 0)
	b.clicableStates = make([]*clickable, 0)

	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i*fieldWidth + i
			newY := j*fieldHeight + j
			color, text := b.statesConf[states[i-1][j-1]].color, b.statesConf[states[i-1][j-1]].char

			b.states = append(b.states, tl.NewRectangle(b.x+newX, b.y+newY, fieldWidth, fieldHeight, color))
			b.statesTexts = append(b.statesTexts, tl.NewText(b.x+newX+(fieldWidth/2), b.y+newY+(fieldHeight/2), text, tl.ColorBlack, color))
		}
	}

	return b
}

func (b *Board) setClicableStates(states [10][10]State) *Board {
	b.states = make([]*tl.Rectangle, 0)
	b.statesTexts = make([]*tl.Text, 0)
	b.clicableStates = make([]*clickable, 0)

	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			newX := i*fieldWidth + i
			newY := j*fieldHeight + j
			color, text := b.statesConf[states[i-1][j-1]].color, b.statesConf[states[i-1][j-1]].char

			rec := tl.NewRectangle(b.x+newX, b.y+newY, fieldWidth, fieldHeight, color)
			if states[i-1][j-1].clickAllowed() {
				b.clicableStates = append(b.clicableStates, newClickable(
					fmt.Sprintf("%s%d", letters[i-1], j), rec))
			} else {
				b.states = append(b.states, rec)
			}

			b.statesTexts = append(b.statesTexts, tl.NewText(b.x+newX+(fieldWidth/2), b.y+newY+(fieldHeight/2), text, tl.ColorBlack, color))
		}
	}

	return b
}

func rgbFromString(s string) (int, int, int, error) {
	var colors [3]int
	for i, e := range strings.Split(s, ",") {
		n, err := strconv.Atoi(e)
		if err != nil {
			return -1, -1, -1, err
		}
		colors[i] = n
	}
	return colors[0], colors[1], colors[2], nil
}

func stateFromConfOrDefault(s State, confColor, confChar string) (stateColorAndChar, error) {
	var scc stateColorAndChar
	if confColor != "" {
		r, g, b, err := rgbFromString(confColor)
		if err != nil {
			return stateColorAndChar{}, err
		}
		scc.color = tl.RgbTo256Color(r, g, b)
	} else {
		scc.color = defaultColorState[s].color
	}

	if confChar != "" {
		scc.char = confChar
	} else {
		scc.char = defaultColorState[s].char
	}

	return scc, nil
}

func parseStatesFromConfig(c *BoardConfig) (map[State]stateColorAndChar, error) {
	var err error
	s := make(map[State]stateColorAndChar, 0)

	s[Hit], err = stateFromConfOrDefault(Hit, c.HitColor, string(c.HitChar))
	if err != nil {
		return nil, err
	}
	s[Miss], err = stateFromConfOrDefault(Miss, c.MissColor, string(c.MissChar))
	if err != nil {
		return nil, err
	}
	s[Ship], err = stateFromConfOrDefault(Ship, c.ShipColor, string(c.ShipChar))
	if err != nil {
		return nil, err
	}

	s[None] = defaultColorState[None]

	return s, nil
}
