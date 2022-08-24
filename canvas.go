package raytracer

import (
	"bytes"
	"fmt"
	"math"
)

type Canvas struct {
	canvas [][]Color
}

func NewCanvas(width int, height int) *Canvas {
	canvas := make([][]Color, width)
	for i := range canvas {
		canvas[i] = make([]Color, height)
		for j := range canvas[i] {
			canvas[i][j] = *NewColor(0, 0, 0)
		}
	}

	return &Canvas{canvas}
}

func (c *Canvas) Width() int {
	return len(c.canvas)
}

func (c *Canvas) Height() int {
	if len(c.canvas) == 0 {
		return 0
	}

	return len(c.canvas[0])
}

func (c *Canvas) ColorAt(width int, height int) Color {
	return c.canvas[width][height]
}

func (c *Canvas) SetColorAt(width int, height int, color Color) {
	c.canvas[width][height] = color
}

func (c *Canvas) ToPPM() string {
	var ppm bytes.Buffer
	ppm.WriteString("P3\n")
	ppm.WriteString(fmt.Sprintf("%v %v\n", c.Width(), c.Height()))
	ppm.WriteString("255\n")

	for j := 0; j < c.Height(); j++ {
		var line bytes.Buffer
		for i := 0; i < c.Width(); i++ {
			c := c.canvas[i][j]
			line.WriteString(fmt.Sprintf("%v %v %v ", int(math.Round(clamp(c.t.X*255))), int(math.Round(clamp(c.t.Y*255))), int(math.Round(clamp(c.t.Z*255)))))
		}
		line.Truncate(line.Len() - 1)
		line = splitSegments(line, 70)
		line.WriteRune('\n')
		ppm.Write(line.Bytes())
	}
	return ppm.String()
}

func splitSegments(b bytes.Buffer, segmentSize int) bytes.Buffer {
	bs := b.Bytes()
	segments := b.Len() / segmentSize
	for i := 1; i <= segments; i++ {
		var pos int
		for pos = i * (segmentSize - 1); bs[pos] != ' '; pos-- {
		}
		bs[pos] = '\n'
	}

	return b
}

func clamp(value float64) float64 {
	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return value

}
