package game

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Game represents a game area.
// The coordinates are laid out like this:
// 0 -> x
// |
// V y
// .
type Game struct {
	colors [][]Color
	width  int
	height int
}

// NewGame creates an empty game.
func NewGame(w int, h int) *Game {
	g := make([][]Color, w)
	for i := range g {
		g[i] = make([]Color, h)
	}

	return &Game{colors: g, width: w, height: h}
}

// SetColor sets the color at the specified position.
func (g *Game) SetColor(x int, y int, c Color) {
	g.colors[x][y] = c
}

// GetColor returns the color at the given position.
// If x or y are out of bounds, it returns nil.
// If the color at x/y hasn't been set yet, it returns ColorNone.
func (g *Game) GetColor(x int, y int) *Color {
	if x < 0 || y < 0 {
		return nil
	}

	if x >= g.width || y >= g.height {
		return nil
	}

	return g.getColor(x, y)
}

// getColor returns the color at the given position.
// If x or y are out of bounds, it panics.
func (g *Game) getColor(x int, y int) *Color {
	return &g.colors[x][y]
}

func (g *Game) GetNeighbor(x int, y int, d Direction) *Color {
	switch d {
	case Up:
		return g.GetColor(x, y-1)
	case Down:
		return g.GetColor(x, y+1)
	case Left:
		return g.GetColor(x-1, y)
	case Right:
		return g.GetColor(x+1, y)
	}

	panic("unknown direction")
}
