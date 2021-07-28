package game

import "strings"

func (g *Game) Pretty() string {
	var out strings.Builder

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			out.WriteRune(g.getColor(x, y).AsRune())
		}
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

func ColorFromRune(r rune) *Color {
	if r == 'x' {
		return nil
	}

	if r == ' ' {
		none := ColorNone
		return &none
	}

	if r >= 'A' && r <= 'Z' {
		c := Color(int(r-'A') + int(ColorA))
		return &c
	}

	panic("can't read colors > Z")
}
