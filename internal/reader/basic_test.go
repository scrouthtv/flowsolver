package reader

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/scrouthtv/flowsolver/internal/game"
)

func expectColor(t *testing.T, x int, y int, is *game.Color, should *game.Color) {
	t.Helper()

	if is == should {
		return
	}

	if is == nil {
		t.Errorf("%d/%d: Expected color %s, got nil", x, y, *should)
		return
	} else if should == nil {
		t.Errorf("%d/%d: Expected nil, got color %s", x, y, *is)
		return
	}

	if *is != *should {
		t.Errorf("%d/%d: Expected color %s, got %s", x, y, *should, *is)
	}
}

//go:embed tests/basic3x3.txt
var basic3x3 string

func TestBasic3x3(t *testing.T) {
	// ignore line break shenanigans for now
	mybasic := strings.ReplaceAll(basic3x3, "\r", "")
	i := New(strings.NewReader(mybasic))

	g, err := i.Import()
	if err != nil {
		t.Error(err)
	}

	t.Log("\n" + g.Pretty() + "---")

	w, h := g.Size()
	t.Log("Size: ", w, h)
	if w != 3 || h != 3 {
		t.Errorf("Expected 3x3, got %dx%d", w, h)
	}

	a := game.ColorA
	b := game.ColorB
	c := game.ColorC
	expectColor(t, 1, 0, g.GetColor(1, 0), &a)
	expectColor(t, 0, 1, g.GetColor(0, 1), &b)
	expectColor(t, 2, 1, g.GetColor(2, 1), &b)
	expectColor(t, 1, 2, g.GetColor(1, 2), &c)

	none := game.ColorNone
	expectColor(t, 1, 1, g.GetColor(1, 1), &none)
	expectColor(t, 2, 2, g.GetColor(2, 2), &none)
	expectColor(t, 0, 2, g.GetColor(0, 2), &none)
	expectColor(t, 2, 0, g.GetColor(2, 0), &none)
}

func TestLineBreaks(t *testing.T) {
	// make sure there are no windows line breaks:
	mybasic := strings.ReplaceAll(basic3x3, "\r", "")
	// add windows line breaks:
	mybasic = strings.ReplaceAll(mybasic, "\n", "\r\n")

	i := New(strings.NewReader(mybasic))

	g, err := i.Import()
	if err != nil {
		t.Error(err)
	}

	t.Log("\n" + g.Pretty() + "---")

	w, h := g.Size()
	t.Log("Size: ", w, h)
	if w != 3 || h != 3 {
		t.Errorf("Expected 3x3, got %dx%d", w, h)
	}

	a := game.ColorA
	b := game.ColorB
	c := game.ColorC
	expectColor(t, 1, 0, g.GetColor(1, 0), &a)
	expectColor(t, 0, 1, g.GetColor(0, 1), &b)
	expectColor(t, 2, 1, g.GetColor(2, 1), &b)
	expectColor(t, 1, 2, g.GetColor(1, 2), &c)

	none := game.ColorNone
	expectColor(t, 1, 1, g.GetColor(1, 1), &none)
	expectColor(t, 2, 2, g.GetColor(2, 2), &none)
	expectColor(t, 0, 2, g.GetColor(0, 2), &none)
	expectColor(t, 2, 0, g.GetColor(2, 0), &none)
}
