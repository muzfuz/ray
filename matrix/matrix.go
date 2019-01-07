package matrix

import (
	"errors"

	"github.com/muzfuz/ray/tuple"

	"github.com/muzfuz/ray/float"
)

// Matrix embodies the data and methods to do
// matrix based calculations
type Matrix [][]float64

// NewMatrix constructs a new matrix
func NewMatrix(rows, columns int) Matrix {
	mat := make(Matrix, rows)
	for c := range mat {
		mat[c] = make([]float64, columns)
	}
	return mat
}

// Identity returns the identity matrix
func Identity() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
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
	if m.cols() != m2.rows() {
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

// MultiplyTuple multiplies the matrix by a tuple
func (m Matrix) MultiplyTuple(v tuple.Tuple) (tuple.Tuple, error) {
	colMat := NewMatrix(4, 1)
	colMat = Matrix{
		{v.X},
		{v.Y},
		{v.Z},
		{v.W},
	}
	mat, err := m.Multiply(colMat)
	if err != nil {
		return tuple.Tuple{}, err
	}
	return tuple.Tuple{
		X: mat[0][0],
		Y: mat[1][0],
		Z: mat[2][0],
		W: mat[3][0],
	}, nil
}

// Transpose returns the transpose of the current Matrix
func (m Matrix) Transpose() Matrix {
	mat := NewMatrix(m.cols(), m.rows())
	for r := range m {
		for c := range m[r] {
			mat[c][r] = m[r][c]
		}
	}
	return mat
}

// Determinant returns the determinant of a 2x2 submatrix
func (m Matrix) Determinant() (float64, error) {
	if m.rows() != 2 || m.cols() != 2 {
		return 0.0, errors.New("can only find the determinant of a 2x2 matrix")
	}
	return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0]), nil
}

// Submatrix returns a new matrix with
// a given column and row sliced off the original matrix
func (m Matrix) Submatrix(row, col int) Matrix {
	mat := NewMatrix(m.rows(), m.cols())
	for c := range m {
		copy(mat[c], m[c])
	}
	mat = append(mat[:row], mat[row+1:]...)
	for c := range mat {
		mat[c] = append(mat[c][:col], mat[c][col+1:]...)
	}
	return mat
}

// Minor calculates the determinant of a submatrix
func (m Matrix) Minor(row, col int) (float64, error) {
	return m.Submatrix(row, col).Determinant()
}

func (m Matrix) rows() int {
	return len(m)
}

func (m Matrix) cols() int {
	return len(m[0])
}
