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

// Equal will compare two instances and return true if they are the same
func (m Matrix) Equal(m2 Matrix) bool {
	if len(m.data) != len(m2.data) || len(m.data[0]) != len(m2.data[0]) {
		return false
	}
	for r := range m.data {
		for c := range m.data[r] {
			if m.data[r][c] != m2.data[r][c] {
				return false
			}
		}
	}
	return true
}
