package raytracer_test

import (
	"testing"

	"alqahtani.io/raytracer"
)

func TestNewColor(t *testing.T) {
	got := raytracer.NewColor(-0.5, 0.4, 1.7).Tuple()
	red, green, blue := got.X, got.Y, got.Z

	if red != -0.5 && green != 0.4 && blue != 1.7 {
		t.Errorf("got %v want {-0.5, 0.4, 1.7}", got)
	}
}

func TestColorAdd(t *testing.T) {
	got1, got2 := raytracer.NewColor(0.9, 0.6, 0.75), raytracer.NewColor(0.7, 0.1, 0.25)
	got := got1.Add(got2)
	want := raytracer.NewColor(1.6, 0.7, 1.0)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
