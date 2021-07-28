package game

import "strconv"

type Color uint8

const (
	ColorNone Color = iota
	ColorA
	ColorB
	ColorC
	ColorD
	ColorE
	ColorF
	ColorG
	ColorH
	ColorI
	ColorJ
)

func (c Color) String() string {
	if c == ColorNone {
		return "none"
	}

	if c < 27 {
		return string('A' + c - 1)
	}

	return "Color(" + strconv.Itoa(int(c)-1) + ")"
}
