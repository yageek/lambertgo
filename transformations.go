// Package lambertgo helps to transform coordinates from Lambert system into the WGS84 system.
package lambertgo

import (
	"fmt"
	"math"
)

func latitudeISOFromLatitude(lat float64, e float64) float64 {
	return math.Log(math.Tan(math.Pi/4+lat/2) * math.Pow((1-e*math.Sin(lat))/(1+e*math.Sin(lat)), e/2))
}

func latitudeFromLatitudeISO(latISO float64, e float64, eps float64) float64 {

	phi0 := 2*math.Atan(math.Exp(latISO)) - math.Pi/2
	phii := 2*math.Atan(math.Pow((1+e*math.Sin(phi0))/(1-e*math.Sin(phi0)), e/2)*math.Exp(latISO)) - math.Pi/2

	delta := 100.0
	for delta > eps {
		phi0 = phii
		phii = 2*math.Atan(math.Pow((1+e*math.Sin(phi0))/(1-e*math.Sin(phi0)), e/2.0)*math.Exp(latISO)) - math.Pi/2
		delta = math.Abs(phii - phi0)
	}
	return phii
}

func (pt *Point) lambertToGeographic(zone Zone, lonMerid float64, e float64, eps float64) {

	n := lambertN[zone]
	C := lambertC[zone]
	xs := lambertXs[zone]
	ys := lambertYs[zone]
	x := pt.X
	y := pt.Y

	R := math.Sqrt(((x-xs)*(x-xs) + (y-ys)*(y-ys)))
	gamma := math.Atan((x - xs) / (ys - y))

	lon := lonMerid + gamma/n
	latISO := -1 / n * math.Log(math.Abs(R/C))

	lat := latitudeFromLatitudeISO(latISO, e, eps)

	pt.X = lon
	pt.Y = lat
}

func (pt *Point) geographicToLambert(zone Zone, lonMerid float64, e float64, eps float64) {

	n := lambertN[zone]
	C := lambertC[zone]
	xs := lambertXs[zone]
	ys := lambertYs[zone]
	lon := pt.X
	lat := pt.Y

	latIso := latitudeISOFromLatitude(lat, e)
	x := xs + C*math.Exp(-n*latIso)*math.Sin(n*(lon-lonMerid))
	y := ys - C*math.Exp(-n*latIso)*math.Cos(n*(lon-lonMerid))

	pt.X = x
	pt.Y = y
}

func lambertNormal(lat float64, a float64, e float64) float64 {

	sina := math.Sin(lat)
	return a / math.Sqrt(1-e*e*sina*sina)
}

func (pt *Point) geographicToCartesian(a float64, e float64) {

	lat := pt.Y
	lon := pt.X
	he := pt.Z

	N := lambertNormal(lat, a, e)

	pt.X = (N + he) * math.Cos(lat) * math.Cos(lon)
	pt.Y = (N + he) * math.Cos(lat) * math.Sin(lon)
	pt.Z = (N*(1-e*e) + he) * math.Sin(lat)
}

func (pt *Point) cartesianToGeographic(meridien float64, a float64, e float64, eps float64) {

	x := pt.X
	y := pt.Y
	z := pt.Z

	lon := meridien + math.Atan(y/x)

	module := math.Sqrt(x*x + y*y)

	phi0 := math.Atan(z / (module * (1 - (a*e*e)/math.Sqrt(x*x+y*y+z*z))))

	phii := math.Atan(z / module / (1 - a*e*e*math.Cos(phi0)/(module*math.Sqrt(1-e*e*math.Sin(phi0)*math.Sin(phi0)))))

	delta := 100.0
	for delta > eps {

		phi0 = phii
		phii = math.Atan(z / module / (1 - a*e*e*math.Cos(phi0)/(module*math.Sqrt(1-e*e*math.Sin(phi0)*math.Sin(phi0)))))
		delta = math.Abs(phii - phi0)
	}
	he := module/math.Cos(phii) - a/math.Sqrt(1-e*e*math.Sin(phii)*math.Sin(phii))

	pt.X = lon
	pt.Y = phii
	pt.Z = he
	pt.Unit = Radian
}

// ToWGS84 converts coordinates expressed in Meter in the lambert system to Radian in the WGS84 system.
// It takes the lambert Zone ine parameters.
func (pt *Point) ToWGS84(zone Zone) {

	if pt.Unit != Meter {
		fmt.Println("Could not transform Point which is not in METER")
		return
	}
	if Lambert93 == zone {
		pt.lambertToGeographic(zone, IERSLongitudeMeridian, EWGS84, DefaultEPS)
		pt.Unit = Radian
	} else {
		pt.lambertToGeographic(zone, ParisLongitudeMeridian, EClarkIGN, DefaultEPS)
		pt.Unit = Radian
		pt.geographicToCartesian(AClarkIGN, EClarkIGN)

		pt.X -= 168
		pt.Y -= 60
		pt.Z += 320

		pt.cartesianToGeographic(GreenwichLongitudeMeridian, AWGS84, EWGS84, DefaultEPS)

	}
}

// ToLambert converts coordinates expressed in Radian in the WGS84 system to Meter in the lambert system.
// It takes the lambert Zone in parameters.
func (pt *Point) ToLambert(zone Zone) {

	if pt.Unit != Radian {
		fmt.Println("Could not transform Point which is not in Radian")
		return
	}
	if Lambert93 == zone {
		pt.geographicToLambert(zone, IERSLongitudeMeridian, EWGS84, DefaultEPS)
		pt.Unit = Meter
	} else {
		pt.geographicToLambert(zone, GreenwichLongitudeMeridian, EClarkIGN, DefaultEPS)
		pt.Unit = Meter
	}
}
