package lambertgo

import (
	"fmt"
	"math"
	"testing"
)

func TestGeographicToCartesian(t *testing.T) {

	lon := [...]float64{0.01745329248, 0.00290888212, 0.00581776423}
	lat := [...]float64{0.02036217457, 0.00000000000, -0.03199770300}
	he := [...]float64{100.0000, 10.0000, 2000.0000}
	a := [...]float64{6378249.2000, 6378249.2000, 6378249.2000}
	e := [...]float64{0.08248325679, 0.08248325679, 0.08248325679}
	expected := [...]*Point{

		{6376064.6955, 111294.6230, 128984.7250, 0},
		{6378232.2149, 18553.5780, 0, 0},
		{6376897.5369, 37099.7050, -202730.9070, 0}}
	for index := range lon {
		ptExpected := expected[index]
		ptOrg := &Point{lon[index], lat[index], he[index], 0}

		ptOrg.geographicToCartesian(a[index], e[index])

		if math.Abs(ptOrg.X-ptExpected.X) > 1e-4 {
			t.Errorf("Longitude too far from expected  - Expected : %f - Computed :%f", ptExpected.X, ptOrg.X)
		}
	}
}

func TestLambertNormal(t *testing.T) {

	n := 6393174.9755
	lat := 0.97738438100
	a := 6378388.0000
	e := 0.081991890

	calc := lambertNormal(lat, a, e)
	if math.Abs(n-calc) > 1e-4 {
		t.Errorf("Lambert Normal too far from expected - Expected : %f - Computed : %f", n, calc)

	}
}

func TestLatitudeFromLatitudeISO(t *testing.T) {

	latISO := [...]float64{1.00552653648, -0.30261690060, 0.2000000000}
	e := [...]float64{0.08199188998, 0.08199188998, 0.08199188998}
	eps := [...]float64{1.0e-11, 1.0e-11, 1.0e-11}

	phi := [...]float64{0.87266462600, -0.29999999997, 0.19998903369}

	for index := range latISO {
		result := latitudeFromLatitudeISO(latISO[index], e[index], eps[index])

		if math.Abs(result-phi[index]) > 1e-4 {
			t.Errorf("Latitude too far from expected - Expected : %f - Computed : %f", phi[index], result)
		}

	}

}

func TestCartesianToGeographic(t *testing.T) {

	a := [...]float64{6378249.2000, 6378249.2000, 6378249.2000}
	e := [...]float64{0.08248325679, 0.08248325679, 0.08248325679}
	x := [...]float64{6376064.6950, 6378232.2150, 6376897.5370}
	y := [...]float64{111294.6230, 18553.5780, 37099.7050}
	z := [...]float64{128984.7250, 0.0000, -202730.9070}
	eps := [...]float64{1e-11, 1e-11, 1e-11}

	lon := [...]float64{0.01745329248, 0.00290888212, 0.00581776423}
	lat := [...]float64{0.02036217457, 0.00000000000, -0.03199770301}
	he := [...]float64{99.9995, 10.0001, 2000.0001}

	ignEPS := 1e-11
	for i := range a {

		pt := &Point{x[i], y[i], z[i], Meter}
		pt.cartesianToGeographic(ParisLongitudeMeridian, a[i], e[i], eps[i])

		if math.Abs(pt.X-lon[i]) > ignEPS {
			t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", lon[i], pt.X)
		}
		if math.Abs(pt.Y-lat[i]) > ignEPS {
			t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", lat[i], pt.Y)
		}
		if math.Abs(pt.Z-he[i]) > 1e-4 {
			t.Errorf("Height too far from expected - Expected : %.11f - Computed : %.11f", he[i], pt.Z)
		}
	}

}
func TestLambertToGeographic(t *testing.T) {
	ptOrigin := &Point{1029705.083, 272723.849, 0, 0}
	ptExpected := &Point{0.145512099, 0.872664626, 0, 0}

	ptOrigin.lambertToGeographic(LambertI, GreenwichLongitudeMeridian, EClarkIGN, 1e-9)

	if math.Abs(ptOrigin.X-ptExpected.X) > 1e-9 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.X, ptOrigin.X)
	}

	if math.Abs(ptOrigin.Y-ptExpected.Y) > 1e-9 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.Y, ptOrigin.Y)
	}

}
func TestDegree(t *testing.T) {

	ptOrigin := &Point{180.0, 180.0, 180.0, Degree}
	ptExpected := *ptOrigin

	ptOrigin.ToDegree()
	if ptExpected.X != ptOrigin.X && ptExpected.Y != ptOrigin.Y && ptExpected.Z != ptOrigin.Z && ptExpected.Unit != ptOrigin.Unit {
		t.Error("Point should have the same value")
	}
}
func TestLambert93(t *testing.T) {

	ptOrigin := &Point{668832.5384, 6950138.7285, 0, Meter}
	ptExpected := &Point{2.56865, 49.64961, 0, Degree}

	ptOrigin.ToWGS84(Lambert93)
	ptOrigin.ToDegree()

	if math.Abs(ptOrigin.X-ptExpected.X) > 1e-5 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.X, ptOrigin.X)
	}
	if math.Abs(ptOrigin.Y-ptExpected.Y) > 1e-5 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.Y, ptOrigin.Y)
	}

}

func TestZenithStrasbourg(t *testing.T) {

	ptOrigin := &Point{994300.623, 113409.981, 0, Meter}
	ptExpected := &Point{7.68639475277068, 48.5953456709144, 0, Degree}

	ptOrigin.ToWGS84(LambertI)
	ptOrigin.ToDegree()

	if math.Abs(ptOrigin.X-ptExpected.X) > 1e-5 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.X, ptOrigin.X)
	}
	if math.Abs(ptOrigin.Y-ptExpected.Y) > 1e-5 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.Y, ptOrigin.Y)
	}
}

func TestBugLambertIIE(t *testing.T) {

	ptOrigin := &Point{369419, 1986498, 0, Meter}
	ptExpected := &Point{-0.579117201473994, 44.84071560809383, 0, Degree}

	ptOrigin.ToWGS84(LambertIIE)
	ptOrigin.ToDegree()

	if math.Abs(ptOrigin.X-ptExpected.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.X, ptOrigin.X)
	}
	if math.Abs(ptOrigin.Y-ptExpected.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.Y, ptOrigin.Y)
	}
}

func TestZenithStrasbourgToLambertI(t *testing.T) {

	ptExpected := &Point{994300.623, 113409.981, 0, Meter}
	ptOrigin := &Point{7.68639475277068, 48.5953456709144, 0, Degree}

	ptOrigin.ToRadian()
	ptOrigin.ToLambert(LambertI)

	if math.Abs(ptOrigin.X-ptExpected.X) > 1 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.X, ptOrigin.X)
	}
	if math.Abs(ptOrigin.Y-ptExpected.Y) > 1 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", ptExpected.Y, ptOrigin.Y)
	}
}

func TestToLambertIToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(LambertI)
	point.ToWGS84(LambertI)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func TestToLambertIIToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(LambertII)
	point.ToWGS84(LambertII)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func TestToLambertIIIToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(LambertIII)
	point.ToWGS84(LambertIII)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func TestToLambertIVToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(LambertIV)
	point.ToWGS84(LambertIV)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func TestToLambertIIEToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(LambertIIE)
	point.ToWGS84(LambertIIE)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func TestToLambert93ToWGS84(t *testing.T) {

	var pointTest = &Point{2.228389, 48.824054, 0, Degree}
	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(Lambert93)
	point.ToWGS84(Lambert93)
	point.ToDegree()
	if math.Abs(point.X-pointTest.X) > 1e-3 {
		t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.X, pointTest.X)
	}
	if math.Abs(point.Y-pointTest.Y) > 1e-3 {
		t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f", pointTest.Y, pointTest.Y)
	}
}

func ExamplePoint_ToWGS84() {

	var point = &Point{668832.5384, 6950138.7285, 0, Meter}
	point.ToWGS84(Lambert93)
	fmt.Printf("Latitude:%5f - Longitude:%.5f", point.Y, point.X)
	// Output:
	// Latitude:0.866549 - Longitude:0.04483
}

func ExamplePoint_ToLambert93() {

	var point = &Point{2.228389, 48.824054, 0, Degree}
	point.ToRadian()
	point.ToLambert(Lambert93)
	fmt.Printf("X:%5f - Y:%.5f", point.X, point.Y)
	// Output:
	// X:643349.551264 - Y:6858498.18110
}

func ExamplePoint_ToLambertI() {

	var point = &Point{0.14551209900, 0.87266462600, 0, Radian}
	point.ToLambert(LambertI)
	fmt.Printf("X:%3f - Y:%.3f", point.X, point.Y)
	// Output:
	// X:1029705.081876 - Y:272723.847
}

func ExamplePoint_ToLambertIIE() {

	// var point = &Point{2.228389, 48.824054, 0, Degree}
	// point.ToRadian()
	// point.ToLambert(LambertIIE)
	// fmt.Printf("X:%5f - Y:%.5f, Z: %5.f", point.X, point.Y, point.Z)
	// Output:
	// X:592055,69 - Y:2425079,29
}
