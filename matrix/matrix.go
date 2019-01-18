package matrix

import (
	"errors"
	"fmt"
	"math"

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

// Translation returns a 4x4 translation matrix
func Translation(x, y, z float64) Matrix {
	return Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

// Scaling returns a 4x4 scaling matrix
func Scaling(x, y, z float64) Matrix {
	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

// RotationX returns a 4x4 rotation matrix on the X axis
func RotationX(r float64) Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, math.Cos(r), -math.Sin(r), 0},
		{0, math.Sin(r), math.Cos(r), 0},
		{0, 0, 0, 1},
	}
}

// RotationY returns a 4x4 rotation matrix on the Y axis
func RotationY(r float64) Matrix {
	return Matrix{
		{math.Cos(r), 0, math.Sin(r), 0},
		{0, 1, 0, 0},
		{-math.Sin(r), 0, math.Cos(r), 0},
		{0, 0, 0, 1},
	}
}

// RotationZ returns a 4x4 rotation matrix on the Z axis
func RotationZ(r float64) Matrix {
	return Matrix{
		{math.Cos(r), -math.Sin(r), 0, 0},
		{math.Sin(r), math.Cos(r), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// Shearing returns a 4x4 shearing matrix
func Shearing(xy, xx, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		{1, xy, xx, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
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

// Determinant returns the determinant of either
// a 2x2 submatrix, or a larger one. The calculations
// are slightly different depending on which is found.
func (m Matrix) Determinant() float64 {
	if m.rows() == 2 && m.cols() == 2 {
		return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	}
	det := 0.0
	if len(m) > 0 {
		for i, val := range m[0] {
			det += val * m.Cofactor(0, i)
		}
	}
	return det
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
func (m Matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

// Cofactor first calculates the Minor and then
// determines wether its appropriate to negate the result.
func (m Matrix) Cofactor(row, col int) float64 {
	min := m.Minor(row, col)
	if (row+col)%2 == 0 {
		return min
	}
	return -min
}

// Invertible uses the determinant to determine wether
// the matrix is invertible.
func (m Matrix) Invertible() bool {
	if float.Equal(0.0, m.Determinant()) {
		return false
	}
	return true
}

// Inverse determines the inverse of a matrix
func (m Matrix) Inverse() (Matrix, error) {
	if !m.Invertible() {
		return Matrix{}, fmt.Errorf("matrix %#v is not invertible", m)
	}
	m2 := NewMatrix(m.rows(), m.cols())
	det := m.Determinant()
	for r := range m {
		for c := range m[r] {
			cof := m.Cofactor(r, c)
			m2[c][r] = cof / det
		}
	}
	return m2, nil
}

func (m Matrix) rows() int {
	return len(m)
}

func (m Matrix) cols() int {
	return len(m[0])
}
