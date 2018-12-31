package matrix

// Matrix embodies the data and methods to do
// matrix based calculations
type Matrix struct {
	data [][]float64
}

// NewMatrix constructs a new matrix
func NewMatrix(d [][]float64) Matrix {
	return Matrix{
		data: d,
	}
}

// At retrieves the values at a given row / column coordinate
func (m Matrix) At(r, c int) float64 {
	return m.data[r][c]
}
