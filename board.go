package gui

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
)

const (
	fieldWidth  = 3
	fieldHeight = 1
)

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "X"}

type tile struct {
	rec  *rectangle
	text *tl.Text
}

// Board represents a single board.
type Board struct {
	id     uuid.UUID
	config *BoardConfig
	tiles  []*tile
	ch     chan string

	x int
	y int
}

// BoardConfig holds configuration parameters for Board struct.
type BoardConfig struct {
	RulerColor Color
	TextColor  Color
	EmptyColor Color
	HitColor   Color
	MissColor  Color
	ShipColor  Color
	SunkColor  Color
	EmptyChar  byte
	HitChar    byte
	MissChar   byte
	ShipChar   byte
	SunkChar   byte
}

// NewBoardConfig returns a new config with default values.
func NewBoardConfig() *BoardConfig {
	return &BoardConfig{
		RulerColor: White,
		TextColor:  Black,
		EmptyColor: Blue,
		HitColor:   Red,
		MissColor:  Grey,
		ShipColor:  Green,
		SunkColor:  Magenta,
		EmptyChar:  '~',
		HitChar:    'H',
		MissChar:   'M',
		ShipChar:   'S',
		SunkChar:   'X',
	}
}

func (c *BoardConfig) getColor(state State) Color {
	switch state {
	case Hit:
		return c.HitColor
	case Miss:
		return c.MissColor
	case Ship:
		return c.ShipColor
	case Sunk:
		return c.SunkColor
	default:
		return c.EmptyColor
	}
}

func (c *BoardConfig) getChar(state State) byte {
	switch state {
	case Hit:
		return c.HitChar
	case Miss:
		return c.MissChar
	case Ship:
		return c.ShipChar
	case Sunk:
		return c.SunkChar
	default:
		return c.EmptyChar
	}
}

// NewBoard returns a new Board struct.
// X and Y are the coordinates of the top left corner of the board.
// If no config is provided, default values are used.
func NewBoard(x, y int, config *BoardConfig) *Board {
	if config == nil {
		config = NewBoardConfig()
	}

	b := &Board{
		id:     uuid.New(),
		config: config,
		ch:     make(chan string),
		x:      x,
		y:      y,
	}

	b.tiles = make([]*tile, 11*11)

	for n := 1; n <= 10; n++ {
		newX := n*fieldWidth + n
		horizontal := &tile{
			rec:  newRectangle(tl.NewRectangle(x+newX, y, fieldWidth, fieldHeight, b.config.RulerColor.toAttr())),
			text: tl.NewText(x+newX+(fieldWidth/2), y+(fieldHeight/2), letters[n-1], b.config.TextColor.toAttr(), b.config.RulerColor.toAttr()),
		}
		b.tiles[n] = horizontal

		newY := n*fieldHeight + n
		vertical := &tile{
			rec:  newRectangle(tl.NewRectangle(x, y+newY, fieldWidth, fieldHeight, b.config.RulerColor.toAttr())),
			text: tl.NewText(x+(fieldWidth/2), y+newY+(fieldHeight/2), fmt.Sprintf("%d", n), b.config.TextColor.toAttr(), b.config.RulerColor.toAttr()),
		}
		b.tiles[n*11] = vertical
	}

	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			newX := x*fieldWidth + x
			newY := y*fieldHeight + y
			tile := &tile{
				rec: newClickableRectangle(
					tl.NewRectangle(b.x+newX, b.y+newY, fieldWidth, fieldHeight, b.config.EmptyColor.toAttr()),
					fmt.Sprintf("%s%d", letters[x-1], y),
					b.ch,
				),
				text: tl.NewText(b.x+newX+(fieldWidth/2), b.y+newY+(fieldHeight/2), string(b.config.EmptyChar), b.config.TextColor.toAttr(), b.config.EmptyColor.toAttr()),
			}
			b.tiles[x+y*11] = tile
		}
	}

	return b
}

// SetStates sets the states of the board. The states are represented
// as a 10x10 matrix, where the first index is the X coordinate and
// the second index is the Y coordinate.
// Example: states[0][0] is the state of the field A1.
func (b *Board) SetStates(states [10][10]State) {
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			state := states[x-1][y-1]
			color := b.config.getColor(state).toAttr()
			b.tiles[x+y*11].rec.SetColor(color)
			b.tiles[x+y*11].text.SetColor(b.config.TextColor.toAttr(), color)
			b.tiles[x+y*11].text.SetText(string(b.config.getChar(state)))
		}
	}

}

func (b *Board) ID() uuid.UUID {
	return b.id
}

func (b *Board) Drawables() []tl.Drawable {
	d := []tl.Drawable{}
	for _, t := range b.tiles[1:] {
		d = append(d, t.rec, t.text)
	}
	return d
}

// Listen blocks until a field is clicked by the user and returns the
// field as a string containing coordinates. Use context to control
// cancelation and prevent listening indefinitely.
func (b *Board) Listen(ctx context.Context) string {
	select {
	case s := <-b.ch:
		return s
	case <-ctx.Done():
		return ""
	}
}
