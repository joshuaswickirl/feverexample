package fever_test

import (
	"testing"

	"github.com/joshuaswickirl/feverexample/fever"
)

func TestFever(t *testing.T) {
	var temp float64 = 100
	got := fever.Determine(temp)

	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNoFever(t *testing.T) {
	var temp float64 = 99.5
	got := fever.Determine(temp)

	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
