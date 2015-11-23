package handlers

import (
	"testing"
)

func TestFlash(t *testing.T) {
	r := RGB{}
	var tests = []struct {
		Color    string
		Alt      string
		Time     string
		Expected string
	}{
		{"0,255,0", "255,0,0", "5", "set_rgb 0 255 0 255 0 0 flash 5"},
		{"", "255,0,0", "0", ""},
		{"0,255,0", "", "0", ""},
	}

	for _, v := range tests {
		r.LastRun = ""
		r.Flash(v.Color, v.Alt, v.Time)

		if r.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", r.LastRun, v.Expected)
		}
	}
}

func TestUpdate(t *testing.T) {
	r := RGB{}
	var tests = []struct {
		Color    string
		Expected string
	}{
		{"0,255,0", "set_rgb 0 255 0"},
		{"", ""},
		{"0,255,0", ""},
	}

	for _, v := range tests {
		r.LastRun = ""
		r.Update(v.Color)

		if r.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", r.LastRun, v.Expected)
		}
	}
}
