package lambertgo

import (
	"math"
)
var lambertN = [...]float64{0.7604059656, 0.7289686274, 0.6959127966, 0.6712679322, 0.7289686274, 0.7256077650}
var lambertC = [...]float64{11603796.98, 11745793.39, 11947992.52, 12136281.99, 11745793.39, 11754255.426}
var lambertXs = [...]float64{600000.0, 600000.0, 600000.0, 234.358, 600000.0, 700000.0}
var lambertYs = [...]float64{5657616.674, 6199695.768, 6791905.085, 7239161.542, 8199695.768, 12655612.050}

const (
	LambertI = 0
	LambertII = 1
	LambertIII = 2
	LambertIV = 3
	LambertII_e = 4
	Lambert93 = 5
)
const (
	Degree = iota
	Grad = iota
	Radian = iota
	Meter = iota

	EClarkIGN float64 = 0.08248325676
	EWGS84 float64 = 0.08181919106

	AClarkIGN float64 = 6378249.2
	AWGS84 float64 = 6378137.0

	ParisLongitudeMeridian  float64 = 0
	GreenwichLongitudeMeridian float64 = 0.04079234433
	IERSLongitudeMeridian float64 = 3*math.Pi/180.0

)

var DefaultEPS float64 = 1e-10
