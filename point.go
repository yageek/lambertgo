package lambertgo

import "math"

type Transform Point

type Zone int32

type Point struct {
	X float64
	Y float64
	Z float64
	Unit int32
}

const (
	degreeToradian float64 = math.Pi/180.0
	radianTodegree float64 = 180.0/math.Pi

	gradTodegree float64 = 180.0/200.0
	degreeTograd float64 = 200.0/180.0

	gradToradian float64 = math.Pi/200.0
	radiantTograd float64 = 200.0/math.Pi
)
func (pt *Point) scale(s float64){

	pt.X*= s
	pt.Y*= s
	pt.Z*= s
}

func (pt *Point) ToDegree(){
	switch pt.Unit {
	case Radian:
		pt.scale(radianTodegree)
	case Grad:
		pt.scale(gradTodegree)
	default:

	}
	pt.Unit = Degree
}
func (pt *Point) ToGrad(){
	switch pt.Unit {
	case Radian:
		pt.scale(radiantTograd)
	case Degree:
		pt.scale(degreeTograd)
	default:

	}
	pt.Unit = Grad
}

func (pt *Point) ToRadian(){
	switch pt.Unit {
	case Grad:
		pt.scale(gradToradian)
	case Degree:
		pt.scale(degreeToradian)
	default:

	}
	pt.Unit = Radian
}
