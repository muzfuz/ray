package matrix

import (
	"math"
	"testing"

	"github.com/muzfuz/raytrace/tuple"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	is := assert.New(t)

	m := NewMatrix(4, 4)
	m = Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	is.Equal(1.0, m[0][0])
	is.Equal(4.0, m[0][3])
	is.Equal(5.5, m[1][0])
	is.Equal(7.5, m[1][2])
	is.Equal(11.0, m[2][2])
	is.Equal(13.5, m[3][0])
	is.Equal(15.5, m[3][2])

	m2 := NewMatrix(2, 2)
	m2 = Matrix{
		{-3, 5},
		{1, -2},
	}
	is.Equal(-3.0, m2[0][0])
	is.Equal(5.0, m2[0][1])
	is.Equal(1.0, m2[1][0])
	is.Equal(-2.0, m2[1][1])

	m3 := NewMatrix(3, 3)
	m3 = Matrix{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}
	is.Equal(-3.0, m3[0][0])
	is.Equal(-2.0, m3[1][1])
	is.Equal(1.0, m3[2][2])
}

func TestEqual(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	b := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	c := Matrix{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}
	is.True(a.Equal(b))
	is.False(a.Equal(c))
}

func TestMultiply(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	b := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}

	expected := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	mult := a.Multiply(b)
	is.True(mult.Equal(expected))
}

func TestMultiplyTuple(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	b := tuple.NewPoint(1, 2, 3)
	expected := tuple.NewPoint(18, 24, 33)

	res := a.MultiplyTuple(b)
	is.True(res.Equal(expected))
}

func TestIdentityMatrix(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}

	res := a.Multiply(Identity())

	is.Equal(res, a)
}

func TestTranspose(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}

	tnsps := Matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}

	is.Equal(tnsps, a.Transpose())
	is.Equal(Identity(), Identity().Transpose())
}

func TestDeterminant(t *testing.T) {
	is := assert.New(t)

	mat := Matrix{
		{1, 5},
		{-3, 2},
	}

	d := mat.Determinant()
	is.Equal(17.0, d)
}

func TestDeterminant3by3(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	c := a.Cofactor(0, 0)
	is.Equal(56.0, c)
	c = a.Cofactor(0, 1)
	is.Equal(12.0, c)
	c = a.Cofactor(0, 2)
	is.Equal(-46.0, c)
	d := a.Determinant()
	is.Equal(-196.0, d)
}

func TestDeterminant4by4(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}

	c := a.Cofactor(0, 0)
	is.Equal(690.0, c)
	c = a.Cofactor(0, 1)
	is.Equal(447.0, c)
	c = a.Cofactor(0, 2)
	is.Equal(210.0, c)
	c = a.Cofactor(0, 3)
	is.Equal(51.0, c)
	d := a.Determinant()
	is.Equal(-4071.0, d)
}

func TestSubmatrix(t *testing.T) {
	is := assert.New(t)

	mat3 := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}
	e := Matrix{
		{-3, 2},
		{0, 6},
	}
	is.Equal(e, mat3.Submatrix(0, 2))

	mat4 := Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}
	e2 := Matrix{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}
	is.Equal(e2, mat4.Submatrix(2, 1))
}

func TestMinor(t *testing.T) {
	is := assert.New(t)

	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	b := m.Submatrix(1, 0).Determinant()
	is.Equal(25.0, b)

	min := m.Minor(1, 0)
	is.Equal(25.0, min)
}

func TestCofactor(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	min := a.Minor(0, 0)
	is.Equal(-12.0, min)
	cof := a.Cofactor(0, 0)
	is.Equal(-12.0, cof)

	min = a.Minor(1, 0)
	is.Equal(25.0, min)
	cof = a.Cofactor(1, 0)
	is.Equal(-25.0, cof)
}

func TestInvertible(t *testing.T) {
	is := assert.New(t)
	a := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}
	is.Equal(-2120.0, a.Determinant())
	is.True(a.Invertible())

	b := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}
	is.Equal(0.0, b.Determinant())
	is.False(b.Invertible())

	c := Matrix{}
	is.Equal(0.0, c.Determinant())
	is.False(c.Invertible())
}

func TestInverse(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}

	e := Matrix{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	}

	is.Equal(532.0, a.Determinant())
	is.Equal(-160.0, a.Cofactor(2, 3))
	is.Equal(105.0, a.Cofactor(3, 2))

	b, err := a.Inverse()
	is.NoError(err)
	is.Equal(-160.0/532.0, b[3][2])
	is.Equal(105.0/532.0, b[2][3])
	is.True(b.Equal(e))
}

func TestInverseAnother(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}
	e1 := Matrix{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	}
	res1, err := a.Inverse()
	is.NoError(err)
	is.True(res1.Equal(e1))

	b := Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}
	e2 := Matrix{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	}
	res2, err := b.Inverse()
	is.NoError(err)
	is.True(res2.Equal(e2))
}

func TestMultiplyProductByInverse(t *testing.T) {
	is := assert.New(t)

	a := Matrix{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	}
	b := Matrix{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	}
	c := a.Multiply(b)

	inv, err := b.Inverse()
	is.NoError(err)

	res := c.Multiply(inv)
	is.True(a.Equal(res))
}

func TestTranslation(t *testing.T) {
	is := assert.New(t)

	transform := Translation(5.0, -3.0, 2.0)
	p := tuple.NewPoint(-3.0, 4.0, 5.0)

	// Multiply by translation matrix
	e := tuple.NewPoint(2.0, 1.0, 7.0)
	res := transform.MultiplyTuple(p)
	is.Equal(e, res)

	// Multiply by the inverse of a translation matrix
	inv, _ := transform.Inverse()
	e2 := tuple.NewPoint(-8.0, 7.0, 3.0)
	res = inv.MultiplyTuple(p)
	is.Equal(e2, res)

	// Translation does not affect vectors
	v := tuple.NewVector(-3.0, 4.0, 5.0)
	res = transform.MultiplyTuple(v)
	is.Equal(v, res)
}

func TestScaling(t *testing.T) {
	is := assert.New(t)

	scaling := Scaling(2.0, 3.0, 4.0)
	p := tuple.NewPoint(-4.0, 6.0, 8.0)
	e := tuple.NewPoint(-8.0, 18.0, 32.0)

	// Scaling applies to points
	res := scaling.MultiplyTuple(p)
	is.Equal(e, res)

	// Scaling applies to vectors
	v := tuple.NewVector(-4.0, 6.0, 8.0)
	e2 := tuple.NewVector(-8.0, 18.0, 32.0)
	res = scaling.MultiplyTuple(v)
	is.Equal(e2, res)

	// Multiply by the inverse
	inv, _ := scaling.Inverse()
	e3 := tuple.NewVector(-2.0, 2.0, 2.0)
	res = inv.MultiplyTuple(v)
	is.Equal(e3, res)
}

func TestReflection(t *testing.T) {
	is := assert.New(t)

	scaling := Scaling(-1.0, 1.0, 1.0)
	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(-2.0, 3.0, 4.0)
	res := scaling.MultiplyTuple(p)
	is.Equal(e, res)
}

func TestRotationX(t *testing.T) {
	is := assert.New(t)

	// Rotating a point around the x axis
	p := tuple.NewPoint(0, 1, 0)
	// by a half quarter turn
	halfQuarter := RotationX(math.Pi / 4)
	hq := halfQuarter.MultiplyTuple(p)
	e := tuple.NewPoint(0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0)
	is.True(hq.Equal(e))
	// by a full quarter turn
	fullQuarter := RotationX(math.Pi / 2)
	fq := fullQuarter.MultiplyTuple(p)
	e2 := tuple.NewPoint(0, 0, 1)
	is.True(fq.Equal(e2))

	// Rotating a point in the inverse of the x axis
	inv, err := halfQuarter.Inverse()
	is.NoError(err)
	i := inv.MultiplyTuple(p)
	e3 := tuple.NewPoint(0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0)
	is.True(i.Equal(e3))
}

func TestRotationY(t *testing.T) {
	is := assert.New(t)

	p := tuple.NewPoint(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4.0)
	fullQuarter := RotationY(math.Pi / 2.0)

	//Half quarter turn
	hq := halfQuarter.MultiplyTuple(p)
	e := tuple.NewPoint(math.Sqrt(2.0)/2.0, 0, math.Sqrt(2.0)/2.0)
	is.True(hq.Equal(e))

	// Full quarter turn
	fq := fullQuarter.MultiplyTuple(p)
	e2 := tuple.NewPoint(1.0, 0.0, 0.0)
	is.True(fq.Equal(e2))
}

func TestRotationZ(t *testing.T) {
	is := assert.New(t)

	p := tuple.NewPoint(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4.0)
	fullQuarter := RotationZ(math.Pi / 2.0)

	// Half quarter turn
	hq := halfQuarter.MultiplyTuple(p)
	e := tuple.NewPoint(-math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0, 0)
	is.True(hq.Equal(e))

	// Full quarter turn
	fq := fullQuarter.MultiplyTuple(p)
	e2 := tuple.NewPoint(-1.0, 0.0, 0.0)
	is.True(fq.Equal(e2))
}

func TestShearingXY(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(1.0, 0.0, 0.0, 0.0, 0.0, 0.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(5.0, 3.0, 4.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestShearingXZ(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(0.0, 1.0, 0.0, 0.0, 0.0, 0.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(6.0, 3.0, 4.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestShearingYX(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(0.0, 0.0, 1.0, 0.0, 0.0, 0.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(2.0, 5.0, 4.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestShearingYZ(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(0.0, 0.0, 0.0, 1.0, 0.0, 0.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(2.0, 7.0, 4.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestShearingZX(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(2.0, 3.0, 6.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestShearingZY(t *testing.T) {
	is := assert.New(t)

	transform := Shearing(0.0, 0.0, 0.0, 0.0, 0.0, 1.0)

	p := tuple.NewPoint(2.0, 3.0, 4.0)
	e := tuple.NewPoint(2.0, 3.0, 7.0)

	res := transform.MultiplyTuple(p)
	is.Equal(res, e)
}

func TestIndividualTransformations(t *testing.T) {
	is := assert.New(t)

	p := tuple.NewPoint(1.0, 0.0, 1.0)

	a := RotationX(math.Pi / 2.0)
	b := Scaling(5.0, 5.0, 5.0)
	c := Translation(10.0, 5.0, 7.0)

	p2 := a.MultiplyTuple(p)
	e2 := tuple.NewPoint(1.0, -1.0, 0)
	is.True(p2.Equal(e2))

	p3 := b.MultiplyTuple(p2)
	e3 := tuple.NewPoint(5.0, -5.0, 0.0)
	is.True(p3.Equal(e3))

	p4 := c.MultiplyTuple(p3)
	e4 := tuple.NewPoint(15.0, 0.0, 7.0)
	is.True(p4.Equal(e4))
}

func TestChainedTransformations(t *testing.T) {
	is := assert.New(t)

	p := tuple.NewPoint(1.0, 0.0, 1.0)
	a := RotationX(math.Pi / 2.0)
	b := Scaling(5.0, 5.0, 5.0)
	c := Translation(10.0, 5.0, 7.0)

	tr := c.Multiply(b).Multiply(a)

	e := tuple.NewPoint(15.0, 0.0, 7.0)
	res := tr.MultiplyTuple(p)
	is.Equal(e, res)
}
