// Package helping to transform coordinates from Lambert system into the WGS84 system
package lambertgo

import (
	"fmt"
	"math"
)


func latitudeISOFromLatitude(lat float64, e float64) float64{
	return math.Log10(math.Tan(math.Pi/4+lat/2)*math.Pow((1-e*math.Sin(lat))/(1+e*math.Sin(lat)),e/2));
}

func latitudeFromLatitudeISO( lat_iso float64, e float64, eps float64) float64{

	phi_0 := 2*math.Atan(math.Exp(lat_iso)) - math.Pi/2
	phi_i := 2*math.Atan(math.Pow((1+e*math.Sin(phi_0))/(1-e*math.Sin(phi_0)),e/2)*math.Exp(lat_iso)) - math.Pi/2

	delta := 100.0
	for  delta > eps {
		phi_0 = phi_i
		phi_i = 2*math.Atan(math.Pow((1+e*math.Sin(phi_0))/(1-e*math.Sin(phi_0)),e/2.0)*math.Exp(lat_iso)) - math.Pi/2
		delta = math.Abs(phi_i - phi_0)
	}
	return phi_i
}

func (pt * Point) lambertToGeographic(zone Zone, lon_merid float64,  e float64, eps float64) {

	n := lambertN[zone]
	C := lambertC[zone]
	x_s := lambertXs[zone]
	y_s := lambertYs[zone]
	x := pt.X
	y := pt.Y

	R := math.Sqrt(((x-x_s)*(x-x_s) + (y-y_s)*(y-y_s)))
	gamma := math.Atan((x-x_s)/(y_s-y))

	lon := lon_merid + gamma/n
	lat_iso := -1/n*math.Log(math.Abs(R/C))

	lat := latitudeFromLatitudeISO(lat_iso,e,eps)

	pt.X = lon
	pt.Y = lat
}

func lambertNormal(lat float64,  a float64, e float64) float64{

	sina := math.Sin(lat)
	return a/math.Sqrt(1-e*e*sina*sina)
}

func (pt *Point) geographicToCartesian(a float64, e float64){

	lat :=  pt.Y
	lon :=  pt.X
	he := pt.Z

	N := lambertNormal(lat,a,e)

	pt.X = (N+he)*math.Cos(lat)*math.Cos(lon)
	pt.Y = (N+he)*math.Cos(lat)*math.Sin(lon)
	pt.Z = (N*(1-e*e)+he)*math.Sin(lat);
}

func(pt *Point) cartesianToGeographic(meridien float64, a float64, e float64, eps float64){

	x := pt.X
	y := pt.Y
	z := pt.Z

	lon := meridien + math.Atan(y/x)

	module := math.Sqrt(x*x + y*y)

	phi_0 :=  math.Atan(z/(module*(1-(a*e*e)/math.Sqrt(x*x+y*y+z*z))))

	phi_i := math.Atan(z/module/(1-a*e*e*math.Cos(phi_0)/(module * math.Sqrt(1-e*e*math.Sin(phi_0)*math.Sin(phi_0)))));

	delta := 100.0
	for delta > eps {

		phi_0 = phi_i
		phi_i = math.Atan(z/module/(1-a*e*e*math.Cos(phi_0)/(module * math.Sqrt(1-e*e*math.Sin(phi_0)*math.Sin(phi_0)))))
		delta = math.Abs(phi_i - phi_0)
	}
	he := module/math.Cos(phi_i) - a/math.Sqrt(1-e*e*math.Sin(phi_i)*math.Sin(phi_i))

	pt.X = lon
	pt.Y = phi_i
	pt.Z = he
	pt.Unit = Radian
}

// ToWGS84 converts coordinates expressed in Meter in the lambert system to Radian in the WGS84 system.
// It takes the lambert Zone ine parameters
func (pt * Point) ToWGS84(zone Zone){

	if pt.Unit != Meter {
		fmt.Errorf("Could not transform Point which is not in METER\n")
		return
	}
	if Lambert93 == zone{
		pt.lambertToGeographic(zone,IERSLongitudeMeridian,EWGS84,DefaultEPS)
		pt.Unit = Radian
	} else {
		pt.lambertToGeographic(zone,ParisLongitudeMeridian,EClarkIGN,DefaultEPS)
		pt.Unit = Radian
		pt.geographicToCartesian(AClarkIGN,EClarkIGN)

		pt.X-=168
		pt.Y-=60
		pt.Z+=320

		pt.cartesianToGeographic(IERSLongitudeMeridian,AWGS84,EWGS84,DefaultEPS)

	}

}
