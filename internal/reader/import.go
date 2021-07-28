package reader

import (
	"io"
	"strings"

	"github.com/scrouthtv/flowsolver/internal/game"
)

type Position struct {
	x int
	y int
}

// Importer imports a game.
// It first reads the first line to determine the width of the game.
// Afterwards, the rest of the stream is read to determine set colors
// as well as the height.
// The caller has to take care of stream buffering.
type Importer struct {
	in     io.Reader
	colors map[Position]game.Color
	y      int
}

func New(in io.Reader) *Importer {
	return &Importer{in: in, colors: make(map[Position]game.Color)}
}

func (i *Importer) Import() (*game.Game, error) {
	line, width, err := i.firstLine()
	if err != nil {
		return nil, err
	}

	i.y = 0
	i.parseLine(line)

	// don't discard line breaks, as firstLine() already did:
	line, err = i.subseqLine(width, false)
	if err != nil {
		return i.finalize(width), nil
	}

	i.y = 1
	i.parseLine(line)

	for {
		line, err = i.subseqLine(width, true)
		if err != nil {
			return i.finalize(width), nil
		}

		i.y++
		i.parseLine(line)
	}
}

func (i *Importer) finalize(width int) *game.Game {
	g := game.New(width, i.y+1)

	for pos, color := range i.colors {
		g.SetColor(pos.x, pos.y, color)
	}

	return g
}

func (i *Importer) subseqLine(w int, discardbreaks bool) (string, error) {
	if discardbreaks {
		r, err := i.read()
		if err != nil {
			return "", err
		}

		for r != '\n' {
			r, err = i.read()
			if err != nil {
				return "", err
			}
		}
	}

	buf := make([]byte, w)
	_, err := i.in.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// parseLine extracts all the required colors from the specified line
// and adds them to the importer's memory.
// The current line is taken from the importer.
func (i *Importer) parseLine(line string) {
	for x, r := range line {
		c, err := game.ColorFromRune(r)
		if err == nil && shouldSetColor(c) {
			i.colors[Position{x: x, y: i.y}] = *c
		}
	}
}

// read reads from the stream skipping \r characters.
func (i *Importer) read() (byte, error) {
	buf := make([]byte, 1)
	_, err := i.in.Read(buf)
	if err != nil {
		return 0, err
	}

	for buf[0] == '\r' {
		_, err := i.in.Read(buf)
		if err != nil {
			return 0, err
		}
	}

	return buf[0], nil
}

func (i *Importer) firstLine() (string, int, error) {
	var line strings.Builder

	read, err := i.read()
	if err != nil {
		return "", 0, err
	}

	for read != '\n' {
		line.WriteByte(read)
		read, err = i.read()
		if err != nil {
			return "", 0, err
		}
	}

	return line.String(), line.Len(), nil
}
