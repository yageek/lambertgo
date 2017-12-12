package lambertgo

import "math"

// Point represents a generic point with three space coordinates and unit.
type Point struct {
	X    float64
	Y    float64
	Z    float64
	Unit int32
}

const (
	degreeToradian float64 = math.Pi / 180.0
	radianTodegree float64 = 180.0 / math.Pi

	gradTodegree float64 = 180.0 / 200.0
	degreeTograd float64 = 200.0 / 180.0

	gradToradian  float64 = math.Pi / 200.0
	radiantTograd float64 = 200.0 / math.Pi
)

func (pt *Point) scale(s float64) {

	pt.X *= s
	pt.Y *= s
	pt.Z *= s
}

// ToDegree converts the coordinates from Radian or Grad to Degree.
func (pt *Point) ToDegree() {
	switch pt.Unit {
	case Radian:
		pt.scale(radianTodegree)
	case Grad:
		pt.scale(gradTodegree)
	default:

	}
	pt.Unit = Degree
}

// ToGrad converts the coordinates from Degree or Radian to Gradian.
func (pt *Point) ToGrad() {
	switch pt.Unit {
	case Radian:
		pt.scale(radiantTograd)
	case Degree:
		pt.scale(degreeTograd)
	default:

	}
	pt.Unit = Grad
}

// ToRadian converts the coordinates from Degree or Grad to Radian.
func (pt *Point) ToRadian() {
	switch pt.Unit {
	case Grad:
		pt.scale(gradToradian)
	case Degree:
		pt.scale(degreeToradian)
	default:

	}
	pt.Unit = Radian
}
