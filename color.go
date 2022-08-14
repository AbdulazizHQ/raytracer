package raytracer

type Color struct {
	t *Tuple
}

func NewColor(x float64, y float64, z float64) *Color {
	return &Color{&Tuple{x, y, z, 0.0}}
}

func (c *Color) Tuple() *Tuple {
	return c.t
}

func (c *Color) Equals(c2 *Color) bool {
	return c.t.Equals(c2.t)
}

func (c *Color) Add(c2 *Color) *Color {
	c.t = c.t.Add(c2.t)

	return c
}

func (c *Color) Subtract(c2 *Color) *Color {
	c.t = c.t.Sub(c2.t)

	return c
}

func (c *Color) Scale(s float64) *Color {
	c.t = c.t.Scale(s)

	return c
}

func (c *Color) Multiply(c2 *Color) *Color {
	c.t.X *= c2.t.X
	c.t.Y *= c2.t.Y
	c.t.Z *= c2.t.Z

	return c
}
