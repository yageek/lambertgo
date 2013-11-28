package lambertgo

import (
	"math"
)

const (
	LambertI = iota
	LambertII = iota
	LambertIII = iota
	LambertIV = iota
	Lambert93 =iota
)

const (
	Degree = iota
	Grad = iota
	Radian = iota
)

const (
	degreeToradian = 180.0/math.Pi
	radianTodegree = math.Pi/180

	gradTodegree = 180.0/200.0
	degreeTograd = 200.0/180.0

	gradToradian = 200.0/math.Pi
	radiantTograd = math.Pi/200.0
)

type Transform Point

type Point struct {
	X float32
	Y float32
	Z float32
	Unit int32
}

func (pt *Point) scale(s float32){

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
}
func (pt *Point) ToGrad(){
	switch pt.Unit {
	case Radian:
		pt.scale(radiantTograd)
	case Degree:
		pt.scale(degreeTograd)
	default:

	}
}

func (pt *Point) ToRadian(){
	switch pt.Unit {
	case Grad:
		pt.scale(gradToradian)
	case Degree:
		pt.scale(degreeToradian)
	default:

	}
}
