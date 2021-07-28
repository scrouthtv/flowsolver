package reader

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/scrouthtv/flowsolver/internal/game"
)

//go:embed tests/border3x3.txt
var borderTest string

func TestBorder(t *testing.T) {
	i := New(strings.NewReader(borderTest))
	g, err := i.Import()
	if err != nil {
		t.Fatal(err)
	}

	w, h := g.Size()

	t.Log("\n" + g.Pretty() + "---")

	if w != 3 || h != 3 {
		t.Errorf("Invalid size: %d/%d", w, h)
	}

	a, b, c, d := game.ColorA, game.ColorB, game.ColorC, game.ColorD
	expectColor(t, 0, 0, g.GetColor(0, 0), &a)
	expectColor(t, 2, 0, g.GetColor(2, 0), &b)
	expectColor(t, 0, 2, g.GetColor(0, 2), &c)
	expectColor(t, 2, 2, g.GetColor(2, 2), &d)

	none := game.ColorNone
	expectColor(t, 1, 0, g.GetColor(1, 0), &none)
	expectColor(t, 1, 1, g.GetColor(1, 1), &none)
	expectColor(t, 1, 2, g.GetColor(1, 2), &none)
	expectColor(t, 0, 1, g.GetColor(0, 1), &none)
	expectColor(t, 2, 1, g.GetColor(2, 1), &none)
}
