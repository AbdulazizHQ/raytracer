package raytracer

type Matrix struct {
	matrix [][]float64
}

func NewMatrix(width int, height int) *Matrix {
	m := make([][]float64, width)
	for i := 0; i < width; i++ {
		m[i] = make([]float64, height)
	}

	return &Matrix{m}
}

func (m *Matrix) Set(i int, j int, value float64) {
	m.matrix[i][j] = value
}

func (m *Matrix) Get(i int, j int) float64 {
	return m.matrix[i][j]
}
