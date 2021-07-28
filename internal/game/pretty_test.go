package game

import "testing"

func TestPretty(t *testing.T) {
	g := New(3, 3)
	g.SetColor(0, 0, ColorA)
	g.SetColor(1, 0, ColorB)
	g.SetColor(0, 2, ColorC)
	pretty := g.Pretty()
	expected := "AB \n   \nC  \n"
	if pretty != expected {
		t.Errorf("Expected \n'%s' got \n'%s'", expected, pretty)
	}

	g.SetColor(2, 2, ColorD)
	pretty = g.Pretty()
	expected = "AB \n   \nC D\n"
	if pretty != expected {
		t.Errorf("Expected \n'%s' got \n'%s'", expected, pretty)
	}
}
