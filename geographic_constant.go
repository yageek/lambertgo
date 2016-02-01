package lambertgo

import (
	"math"
)

var lambertN = [...]float64{0.7604059656, 0.7289686274, 0.6959127966, 0.6712679322, 0.7289686274, 0.7256077650}
var lambertC = [...]float64{11603796.98, 11745793.39, 11947992.52, 12136281.99, 11745793.39, 11754255.426}
var lambertXs = [...]float64{600000.0, 600000.0, 600000.0, 234.358, 600000.0, 700000.0}
var lambertYs = [...]float64{5657616.674, 6199695.768, 6791905.085, 7239161.542, 8199695.768, 12655612.050}

//Zone represents the different lambert zones.
type Zone int32

const (
	// LambertI represents the first Lambert zone.
	LambertI = 0
	// LambertII represents the second Lambert zone.
	LambertII = 1
	// LambertIII represents the third Lambert zone.
	LambertIII = 2
	// LambertIV represents the fourth Lambert zone.
	LambertIV = 3
	// LambertIIE represents the extended lambert zone.
	LambertIIE = 4
	// Lambert93 represents the lambert 93 zone.
	Lambert93 = 5
)
const (
	// Degree represents standart degrees unit.
	Degree = iota
	// Grad represents standart grad unit.
	Grad = iota
	// Radian represents standart radian unit.
	Radian = iota
	// Meter represents standart meter unit.
	Meter = iota

	//EClarkIGN represents the E value from Clark ellipsoïd.
	EClarkIGN float64 = 0.08248325676
	//EWGS84 represents the E value from WGS84 ellipsoïd.
	EWGS84 float64 = 0.08181919106

	//AClarkIGN represents the A value from Clark ellipsoïd.
	AClarkIGN float64 = 6378249.2

	//AWGS84 represents the A value from WGS84 ellipsoïd.
	AWGS84 float64 = 6378137.0

	//ParisLongitudeMeridian is the Paris's meridian longitude in radian.
	ParisLongitudeMeridian float64 = 0
	//GreenwichLongitudeMeridian is the Greenwich's meridian longitude in radian.
	GreenwichLongitudeMeridian float64 = 0.04079234433
	//IERSLongitudeMeridian is the IERS's meridian longitude in radian.
	IERSLongitudeMeridian float64 = 3 * math.Pi / 180.0
)

// DefaultEPS is the comparison margin
// used in the different algorithms.
var DefaultEPS float64 = 1e-10
