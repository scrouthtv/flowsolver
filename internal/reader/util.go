package reader

import "github.com/scrouthtv/flowsolver/internal/game"

func shouldSetColor(c *game.Color) bool {
	if c == nil {
		return false
	}

	if *c == game.ColorNone {
		return false
	}

	return true
}
