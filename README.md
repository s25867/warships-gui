# warships-gui

Warships-gui is an advance tool that allows you to draw your 'Warships Online' game into a terminal.
If you are familiar with goroutines and channels that library is excellent for you. 
Otherwise see simpler solution: 
- https://github.com/grupawp/warships-lightgui

## Quick Start
To start a new game and take control under terminal use **NewDrawer()** method as below:
```go
d := gui.NewDrawer(&gui.Config{})
```
After that call library starts a new gouroutine and takes control above current terminal, so you can easily write there any content. 

Whats more, type **gui.Config** allows to overwrite default parameters and looks like:
```go
// Config holds customizable parameters for drawer.
type Config struct {
	// EndKey is an end character using to exit game, 'space' is default.
	EndKey tl.Key
}
```

## Drawing a board
To draw a board you need to create a new **Board** object. In a single game you can create as many boards as you need, so if you want to create one you need to call one of **Drawer** methods:
```go
// NewBoard creates new instance of Board at given (x,y) position.
func (d drawer) NewBoard(x, y int, c *BoardConfig) (*Board, error)
```
It also allows to overwrite default parameters but that fields are not required, when none is set then it takes a default value. **BoardConfig** looks like below:
```go
// BoardConfig holds configuration parameters for Board struct.
type BoardConfig struct {
	HitColor  string
	MissColor string
	ShipColor string

	HitChar  byte
	MissChar byte
	ShipChar byte
}
```
A string value of color has to be a RGB representation in given format: "r,g,b" so it expect three integer values in range(0,255) separated by comma. Otherwise **NewBoard()** methods returns an error.

When we have a **Board** instance then we can start drawing using one of the methods below: 
```go
// DrawBoard draws 10x10 board with left upper corner begins at (x,y) point.
// It fills fields as it's given in 'states' argument.
func (d *drawer) DrawBoard(ctx context.Context, b *Board, states [10][10]State)

// DrawBoardAndCatchCoords does same as 'DrawBoard' method.
// But after drawing it waits for mouse action that returns clicked field, e.g. "B6".
// This allows to click only on the "see state".
func (d *drawer) DrawBoardAndCatchCoords(ctx context.Context, b *Board, states [10][10]State) string
```

## States
State is internal representation of a single field on the board. It's one of a constrant values:
```go
const (
	// Hit means that some ship element is hit.
	Hit 
	// Miss means that shoot is missed.
	Miss 
	// Ship means that field represents whole ship element.
	Ship 
	// None means that field has no defined State.
	None 
)
```
So if you want to draw a **Ship** state at point (2,5) and a **Miss** state at point (2,6) you can call **DrawBoard** like below:
```go
d := gui.NewDrawer(&gui.Config{})
b := d.NewBoard(1, 1, &gui.BoardConfig{})

states := [10][10]gui.State{}{}
states[2][5] = gui.Ship
states[2][6] = gui.Miss

d.DrawBoard(ctx, b, states)
```

## Drawing a text
Drawing a text is similar like drawing a board. Everything that you need is to create a new **Text** instance:
```go
// DrawText creates a new Text at position (x, y).
// It sets the Text's text to be text.
// Returns a pointer to the new Text.
func (d *drawer) DrawText(ctx context.Context, x, y int, text string) *Text
```
You can change value of the **Text** at any time by calling:
```go
// SetText sets the text of the Text to be text and returns Text.
func (t *Text) SetText(text string) *Text
```

## Removing elements 
Our GUI supports removing objects from **Drawer**. To do it you should call one of the methods below:
```go
// RemoveBoard removes existing Board from screen.
func (d *drawer) RemoveBoard(ctx context.Context, b *Board)

// RemoveText removes existing Text from screen.
func (d *drawer) RemoveText(ctx context.Context, t *Text)
```

## Examples
Look inside **_examples/** to see some example files. Enjoy! 