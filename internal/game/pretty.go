package game

import "strings"

func (g *Game) Pretty() string {
	var out strings.Builder

	out.Grow(g.width * g.height)

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			out.WriteRune(g.getColor(x, y).AsRune())
		}
		out.WriteRune('\n')
	}

	return out.String()
}

func (c *Color) AsRune() rune {
	if c == nil {
		return 'x'
	}

	if *c == ColorNone {
		return ' '
	}

	if *c <= ColorA+26 {
		return rune(*c + 'A' - ColorA)
	}

	return '#'
}

type ErrInvalidColor struct {
	r rune
}

func (e *ErrInvalidColor) Error() string {
	return "invalid color: " + string(e.r)
}

func ColorFromRune(r rune) (*Color, error) {
	if r == 'x' {
		return nil, nil
	}

	if r == ' ' {
		none := ColorNone
		return &none, nil
	}

	if r >= 'A' && r <= 'Z' {
		c := Color(int(r-'A') + int(ColorA))
		return &c, nil
	}

	return nil, &ErrInvalidColor{r}
}
