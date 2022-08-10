package raytracer_test

import (
	"math"
	"testing"

	"alqahtani.io/raytracer"
)

func TestTupleIsAPoint(t *testing.T) {
	got := raytracer.Tuple{X: 4.3, Y: -4.2, Z: 1.0, W: 1.0}
	want := 1.0

	if got.W != want {
		t.Errorf("got %f want %f", got.W, want)
	}
}

func TestTupleIsAVector(t *testing.T) {
	got := raytracer.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
	want := 0.0

	if got.W != want {
		t.Errorf("got %f want %f", got.W, want)
	}
}

func TestNewPoint(t *testing.T) {
	got := raytracer.NewPoint(4, -4, 3)
	want := raytracer.Tuple{4, -4, 3, 1.0}

	if !got.Equals(&want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewVector(t *testing.T) {
	got := raytracer.NewVector(4, -4, 3)
	want := raytracer.Tuple{4, -4, 3, 0.0}

	if !got.Equals(&want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAddTuples(t *testing.T) {
	a := &raytracer.Tuple{3, -2, 5, 1}
	b := &raytracer.Tuple{-2, 3, 1, 0}
	got := a.Add(b)
	want := raytracer.Tuple{1, 1, 6, 1}

	if !got.Equals(&want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubTwoPoints(t *testing.T) {
	a := raytracer.NewPoint(3, 2, 1)
	b := raytracer.NewPoint(5, 6, 7)
	got := a.Sub(b)
	want := raytracer.NewVector(-2, -4, -6)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubVecFromPoint(t *testing.T) {
	p := raytracer.NewPoint(3, 2, 1)
	v := raytracer.NewVector(5, 6, 7)
	got := p.Sub(v)
	want := raytracer.NewPoint(-2, -4, -6)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubTwoVecs(t *testing.T) {
	v1 := raytracer.NewVector(3, 2, 1)
	v2 := raytracer.NewVector(5, 6, 7)
	got := v1.Sub(v2)
	want := raytracer.NewVector(-2, -4, -6)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSubVecFromZeroVec(t *testing.T) {
	zv := raytracer.ZeroVector()
	v := raytracer.NewVector(1, -2, 3)
	got := zv.Sub(v)
	want := raytracer.NewVector(-1, 2, -3)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNegateTuple(t *testing.T) {
	v := &raytracer.Tuple{1, -2, 3, -4}
	got := v.Negate()
	want := &raytracer.Tuple{-1, 2, -3, 4}

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScaleUp(t *testing.T) {
	v := &raytracer.Tuple{1, -2, 3, -4}
	got := v.Scale(3.5)
	want := &raytracer.Tuple{3.5, -7, 10.5, -14}

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestScaleDown(t *testing.T) {
	v := &raytracer.Tuple{1, -2, 3, -4}
	got := v.Scale(0.5)
	want := &raytracer.Tuple{0.5, -1, 1.5, -2}

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMagnitude1(t *testing.T) {
	v := raytracer.NewVector(1, 0, 0)
	got := v.Magnitude()
	want := 1.0

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMagnitude2(t *testing.T) {
	v := raytracer.NewVector(0, 1, 0)
	got := v.Magnitude()
	want := 1.0

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMagnitude3(t *testing.T) {
	v := raytracer.NewVector(0, 0, 1)
	got := v.Magnitude()
	want := 1.0

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMagnitude4(t *testing.T) {
	v := raytracer.NewVector(1, 2, 3)
	got := v.Magnitude()
	want := math.Sqrt(14.0)

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestMagnitude5(t *testing.T) {
	v := raytracer.NewVector(-1, -2, -3)
	got := v.Magnitude()
	want := math.Sqrt(14.0)

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestNormalize1(t *testing.T) {
	v := raytracer.NewVector(4, 0, 0)
	got := v.Normalize()
	want := raytracer.NewVector(1, 0, 0)

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNormalize2(t *testing.T) {
	v := raytracer.NewVector(1, 2, 3)
	got := v.Normalize()
	want := raytracer.NewVector(1.0/math.Sqrt(14), 2.0/math.Sqrt(14), 3.0/math.Sqrt(14))

	if !got.Equals(want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNormalize3(t *testing.T) {
	v := raytracer.NewVector(1, 2, 3)
	got := v.Normalize().Magnitude()
	want := 1.0

	if math.Abs(got-want) > raytracer.Epsilon {
		t.Errorf("got %v want %v", got, want)
	}
}
