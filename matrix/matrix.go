package matrix

import (
	"errors"

	"github.com/muzfuz/ray/float"
)

// Matrix embodies the data and methods to do
// matrix based calculations
type Matrix [][]float64

// NewMatrix constructs a new matrix
func NewMatrix(rows, columns int) Matrix {
	rc := make([][]float64, rows)
	for i := range rc {
		rc[i] = make([]float64, columns)
	}
	return Matrix(rc)
}

// At retrieves the values at a given row / column coordinate
func (m Matrix) At(r, c int) float64 {
	return m[r][c]
}

// Equal will compare two instances and return true if they are the same
func (m Matrix) Equal(m2 Matrix) bool {
	if m.rows() != m2.rows() || m.cols() != m2.cols() {
		return false
	}
	for r := range m {
		for c := range m[r] {
			if !float.Equal(m[r][c], m2[r][c]) {
				return false
			}
		}
	}
	return true
}

// Multiply will take two matrices and multiply them
func (m Matrix) Multiply(m2 Matrix) (Matrix, error) {
	if m.rows() != m2.cols() {
		return Matrix{}, errors.New("cannot multiply two matrices with differing row + column sizes")
	}
	newMat := NewMatrix(m.rows(), m2.cols())
	for r := range m {
		for c := range m2[r] {
			for i := 0; i < m.rows(); i++ { // uses m.rows() as delimiter, but could also use m2.cols()
				newMat[r][c] += m[r][i] * m2[i][c]
			}
		}
	}
	return newMat, nil
}

func (m Matrix) rows() int {
	return len(m)
}

func (m Matrix) cols() int {
	return len(m[0])
}
