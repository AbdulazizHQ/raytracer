package raytracer_test

import (
	"alqahtani.io/raytracer"
	"testing"
)

type values struct {
	i     int
	j     int
	value float64
}

func TestNewMatrix(t *testing.T) {
	m := raytracer.NewMatrix(4, 4)
	vs := []values{
		{0, 0, 1.0},
		{1, 2, 7.5},
		{3, 2, 15.5},
	}
	for _, v := range vs {
		m.Set(v.i, v.j, v.value)
	}

	for _, v := range vs {
		got := m.Get(v.i, v.j)
		wants := v.value
		if got != wants {
			t.Fatalf("got m[%v][%v] = %v, wants m[%v][%v] = %v", v.i, v.j, got, v.i, v.j, wants)
		}
	}

}
