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
	d := make([][]float64, rows)
	for i := range d {
		d[i] = make([]float64, columns)
	}
	return Matrix(d)
}

// At retrieves the values at a given row / column coordinate
func (m Matrix) At(r, c int) float64 {
	return m[r][c]
}

// Equal will compare two instances and return true if they are the same
func (m Matrix) Equal(m2 Matrix) bool {
	if len(m) != len(m2) || len(m[0]) != len(m2[0]) {
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
	rows, columns := len(m), len(m2[0])
	if rows != columns {
		return Matrix{}, errors.New("cannot multiply two matrices with differing row + column sizes")
	}
	newMat := NewMatrix(rows, columns)
	for r := range m {
		for c := range m2[r] {
			for i := 0; i < rows; i++ {
				newMat[r][c] += m[r][i] * m2[i][c]
			}
		}
	}
	return newMat, nil
}
