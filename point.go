package lambertgo

type Transform Point

type Zone int32

type Point struct {
	X float64
	Y float64
	Z float64
	Unit int32
}

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
