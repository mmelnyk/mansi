package mansi

import "testing"

func TestCursorUpBy(t *testing.T) {
	var data = []struct {
		pos      int
		expected string
	}{
		{1, CSI + "1A"},
		{2, CSI + "2A"},
		{3, CSI + "3A"},
		{0, ""},
		{-1, ""},
		{-10, ""},
	}

	for _, i := range data {
		if CursorUpBy(i.pos) != i.expected {
			t.Errorf("Return does not match expectation for %d", i.pos)
		}
	}
}

func TestCursorDownBy(t *testing.T) {
	var data = []struct {
		pos      int
		expected string
	}{
		{1, CSI + "1B"},
		{2, CSI + "2B"},
		{3, CSI + "3B"},
		{0, ""},
		{-1, ""},
		{-10, ""},
	}

	for _, i := range data {
		if CursorDownBy(i.pos) != i.expected {
			t.Errorf("Return does not match expectation for %d", i.pos)
		}
	}
}

func TestCursorForwardBy(t *testing.T) {
	var data = []struct {
		pos      int
		expected string
	}{
		{1, CSI + "1C"},
		{2, CSI + "2C"},
		{3, CSI + "3C"},
		{0, ""},
		{-1, ""},
		{-10, ""},
	}

	for _, i := range data {
		if CursorForwardBy(i.pos) != i.expected {
			t.Errorf("Return does not match expectation for %d", i.pos)
		}
	}
}

func TestCursorBackBy(t *testing.T) {
	var data = []struct {
		pos      int
		expected string
	}{
		{1, CSI + "1D"},
		{2, CSI + "2D"},
		{3, CSI + "3D"},
		{0, ""},
		{-1, ""},
		{-10, ""},
	}

	for _, i := range data {
		if CursorBackBy(i.pos) != i.expected {
			t.Errorf("Return does not match expectation for %d", i.pos)
		}
	}
}

func TestSetWindowTitle(t *testing.T) {
	var data = []struct {
		title    string
		expected string
	}{
		{"", OSC + "0;" + BEL},
		{" ", OSC + "0; " + BEL},
		{"Just Title", OSC + "0;Just Title" + BEL},
	}

	for _, i := range data {
		if SetWindowTitle(i.title) != i.expected {
			t.Errorf("Return does not match expectation for %s", i.title)
		}
	}
}

func TestCursorPosition(t *testing.T) {
	var data = []struct {
		line     int
		column   int
		expected string
	}{
		{1, 1, CSI + "1;1H"},
		{0, 1, CSI + "1;1H"},
		{1, 0, CSI + "1;1H"},
		{0, 0, CSI + "1;1H"},
		{1, -1, CSI + "1;1H"},
		{-1, 1, CSI + "1;1H"},
		{-1, -1, CSI + "1;1H"},
		{2, 1, CSI + "2;1H"},
		{1, 2, CSI + "1;2H"},
		{2, 2, CSI + "2;2H"},
		{3, 10, CSI + "3;10H"},
		{100, 100, CSI + "100;100H"},
	}

	for _, i := range data {
		if CursorPosition(i.column, i.line) != i.expected {
			t.Errorf("Return does not match expectation for %d:%d", i.line, i.column)
		}
	}
}
