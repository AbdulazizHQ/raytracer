package raytracer

import "math"

const Epsilon float64 = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func floatEquals(a float64, b float64) bool {
	return math.Abs(b-a) < Epsilon
}

func (a *Tuple) Equals(b *Tuple) bool {
	return floatEquals(a.X, b.X) &&
		floatEquals(a.Y, b.Y) &&
		floatEquals(a.Z, b.Z) &&
		floatEquals(a.W, b.W)
}

func (a *Tuple) Add(b *Tuple) *Tuple {
	a.W += b.W
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z

	return a
}

func (a *Tuple) Sub(b *Tuple) *Tuple {
	a.W -= b.W
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z

	return a
}

func (a *Tuple) Scale(alpha float64) *Tuple {
	a.W *= alpha
	a.X *= alpha
	a.Y *= alpha
	a.Z *= alpha

	return a
}

func (a *Tuple) Negate() *Tuple {
	a.W *= -1
	a.X *= -1
	a.Y *= -1
	a.Z *= -1

	return a
}

func NewPoint(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func NewVector(x float64, y float64, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func ZeroVector() *Tuple {
	return NewVector(0, 0, 0)
}

func ZeroPoint() *Tuple {
	return NewPoint(0, 0, 0)
}

func (a *Tuple) Magnitude() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

func (a *Tuple) Normalize() *Tuple {
	m := a.Magnitude()
	a.W /= m
	a.X /= m
	a.Y /= m
	a.Z /= m

	return a
}
