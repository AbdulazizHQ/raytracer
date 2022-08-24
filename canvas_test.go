package raytracer_test

import (
	"strings"
	"testing"

	"alqahtani.io/raytracer"
)

func TestNewCanvas(t *testing.T) {
	got := raytracer.NewCanvas(10, 20)

	if got.Width() != 10 {
		t.Errorf("width: got %v want %v", got.Width(), 10)
	}
	if got.Height() != 20 {
		t.Errorf("height: got %v want %v", got.Height(), 20)
	}
	want := raytracer.NewColor(0, 0, 0)
	for i := 0; i < got.Width(); i++ {
		for j := 0; j < got.Height(); j++ {
			c := got.ColorAt(i, j)
			if !c.Equals(want) {
				t.Errorf("color (%v, %v), got %v want %v", i, j, got, want)
			}
		}
	}
}

func TestSetColorAt(t *testing.T) {
	want := raytracer.NewColor(1, 2, 3)
	cnvs := raytracer.NewCanvas(1, 1)
	cnvs.SetColorAt(0, 0, *want)
	got := cnvs.ColorAt(0, 0)
	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPPM1(t *testing.T) {
	want := `P3
5 3
255`
	cnvs := raytracer.NewCanvas(5, 3)
	got := strings.Join(strings.Split(cnvs.ToPPM(), "\n")[0:3], "\n")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestPPM2(t *testing.T) {
	want := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	cnvs := raytracer.NewCanvas(5, 3)
	cnvs.SetColorAt(0, 0, *raytracer.NewColor(1.5, 0, 0))
	cnvs.SetColorAt(2, 1, *raytracer.NewColor(0, 0.5, 0))
	cnvs.SetColorAt(4, 2, *raytracer.NewColor(-0.5, 0, 1))

	got := strings.Join(strings.Split(cnvs.ToPPM(), "\n")[3:7], "\n")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestPPM3(t *testing.T) {
	want := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`
	cnvs := raytracer.NewCanvas(10, 2)
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			cnvs.SetColorAt(i, j, *raytracer.NewColor(1, 0.8, 0.6))
		}
	}

	got := strings.Join(strings.Split(cnvs.ToPPM(), "\n")[3:7], "\n")

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
