package game

import "testing"

func TestColorStringer(t *testing.T) {
	if ColorNone.String() != "none" {
		t.Errorf("ColorNone.String() should be none, but is %s", ColorNone.String())
	}

	if ColorA.String() != "A" {
		t.Errorf("ColorA.String() should be A, but is %s", ColorA.String())
	}

	if ColorJ.String() != "J" {
		t.Errorf("ColorJ.String() should be J, but is %s", ColorJ.String())
	}

	cz := Color('Z' - 'A' + ColorA)
	if cz.String() != "Z" {
		t.Errorf("ColorZ.String() should be Z, but is %s", cz.String())
	}

	c35 := Color(35 + ColorA)
	if c35.String() != "35" {
		t.Errorf("Color35.String() should be Color(35), but is %s", c35.String())
	}

}
