package game

import "testing"

func expectColor(t *testing.T, is *Color, should *Color) {
	t.Helper()

	if is == should {
		return
	}

	if is == nil {
		t.Errorf("Expected color %s, got nil", *should)
		return
	} else if should == nil {
		t.Errorf("Expected nil, got color %s", *is)
		return
	}

	if *is != *should {
		t.Errorf("Expected color %s, got %s", *should, *is)
	}
}

func TestGetSetColor(t *testing.T) {
	g := New(3, 2)

	g.SetColor(0, 0, ColorA)
	g.SetColor(1, 1, ColorB)

	a := ColorA
	expectColor(t, g.GetColor(0, 0), &a)

	b := ColorB
	expectColor(t, g.GetColor(1, 1), &b)

	none := ColorNone
	expectColor(t, g.GetColor(1, 0), &none)
	expectColor(t, g.GetColor(2, 1), &none)

	expectColor(t, g.GetColor(-1, 0), nil)
	expectColor(t, g.GetColor(0, -1), nil)
	expectColor(t, g.GetColor(3, 0), nil)
	expectColor(t, g.GetColor(0, 2), nil)
}

func TestGetNeighbor(t *testing.T) {
	g := New(3, 3)
	g.SetColor(1, 1, ColorA)

	//  0 1 2 -> x
	// 0  D
	// 1R o L
	// 2  U
	// |
	// V y

	a := ColorA
	expectColor(t, g.GetNeighbor(1, 2, Up), &a)
	expectColor(t, g.GetNeighbor(1, 0, Down), &a)
	expectColor(t, g.GetNeighbor(2, 1, Left), &a)
	expectColor(t, g.GetNeighbor(0, 1, Right), &a)

	none := ColorNone
	expectColor(t, g.GetNeighbor(2, 2, Up), &none)
	expectColor(t, g.GetNeighbor(0, 0, Down), &none)
	expectColor(t, g.GetNeighbor(2, 2, Left), &none)
	expectColor(t, g.GetNeighbor(0, 0, Right), &none)

	expectColor(t, g.GetNeighbor(0, 0, Up), nil)
	expectColor(t, g.GetNeighbor(2, 2, Down), nil)
	expectColor(t, g.GetNeighbor(0, 0, Left), nil)
	expectColor(t, g.GetNeighbor(2, 2, Right), nil)
}
