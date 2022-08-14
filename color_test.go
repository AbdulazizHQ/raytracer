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
	a, b := raytracer.NewColor(0.9, 0.6, 0.75), raytracer.NewColor(0.7, 0.1, 0.25)
	got := a.Add(b)
	want := raytracer.NewColor(1.6, 0.7, 1.0)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestColorSubtract(t *testing.T) {
	a, b := raytracer.NewColor(0.9, 0.6, 0.75), raytracer.NewColor(0.7, 0.1, 0.25)
	got := a.Subtract(b)
	want := raytracer.NewColor(0.2, 0.5, 0.5)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestColorScale(t *testing.T) {
	a := raytracer.NewColor(0.2, 0.3, 0.4)
	scaler := 2.0
	got := a.Scale(scaler)
	want := raytracer.NewColor(0.4, 0.6, 0.8)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestColorMultiply(t *testing.T) {
	a, b := raytracer.NewColor(1, 0.2, 0.4), raytracer.NewColor(0.9, 1, 0.1)
	got := a.Multiply(b)
	want := raytracer.NewColor(0.9, 0.2, 0.04)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}
