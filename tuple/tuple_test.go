package tuple

import "testing"

func TestTupleWhenIsAPoint(t *testing.T) {
	tup := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
	if tup.X != 4.3 {
		t.Error("Expected tup.X to equal 4.3, got ", tup.X)
	}
	if tup.Y != -4.2 {
		t.Error("Expected tup.Y to equal -4.2, got ", tup.X)
	}
	if tup.Z != 3.1 {
		t.Error("Expected tup.Z to equal 3.1, got ", tup.X)
	}
	if tup.W != 1.0 {
		t.Error("Expected tup.W to equal 1.0, got ", tup.X)
	}
	if tup.IsPoint() != true {
		t.Error("Expected tup to be a point")
	}
	if tup.IsVector() != false {
		t.Error("Expected tup to not be a vector")
	}
}
func TestTupleWhenIsAVector(t *testing.T) {
	tup := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
	if tup.X != 4.3 {
		t.Error("Expected tup.X to equal 4.3, got ", tup.X)
	}
	if tup.Y != -4.2 {
		t.Error("Expected tup.Y to equal -4.2, got ", tup.X)
	}
	if tup.Z != 3.1 {
		t.Error("Expected tup.Z to equal 3.1, got ", tup.X)
	}
	if tup.W != 0.0 {
		t.Error("Expected tup.W to equal 0.0, got ", tup.X)
	}
	if tup.IsPoint() != false {
		t.Error("Expected tup to not be a point")
	}
	if tup.IsVector() != true {
		t.Error("Expected tup to be a vector")
	}
}

func TestNewPoint(t *testing.T) {
	point := NewPoint(4, -4, 3)
	if point.IsPoint() != true {
		t.Error("Expected to have created a point")
	}
	if point.IsVector() == true {
		t.Error("Expected to not have created a vector")
	}
}

func TestNewVector(t *testing.T) {
	vec := NewVector(4, -4, 3)
	if vec.IsPoint() == true {
		t.Error("Expected to not have created a point")
	}
	if vec.IsVector() != true {
		t.Error("Expected to have created a vector")
	}
}
