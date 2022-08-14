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
