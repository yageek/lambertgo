package  lambertgo

import (
	"fmt"
	"math"
	"testing"
)


func TestGeographicToCartesian(t* testing.T){

	lon := [...]float64{0.01745329248 ,0.00290888212 ,0.00581776423}
	lat := [...]float64{0.02036217457,0.00000000000 ,-0.03199770300}
	he  := [...]float64{100.0000,10.0000 ,2000.0000}
	a := [...]float64{6378249.2000 ,6378249.2000 ,6378249.2000}
	e := [...]float64{0.08248325679 ,0.08248325679 ,0.08248325679}
	expected := [...]*Point{

		&Point{6376064.6955,111294.6230,128984.7250, 0 },
		&Point{6378232.2149,18553.5780,0, 0 },
		&Point{6376897.5369,37099.7050,-202730.9070, 0 }}
		for index := range lon{
			ptExpected  := expected[index]
			ptOrg  := &Point{lon[index],lat[index],he[index],0}

			ptOrg.geographicToCartesian(a[index],e[index])

			if math.Abs(ptOrg.X - ptExpected.X) > 1e-4{
				t.Errorf("Longitude too far from expected  - Expected : %f - Computed :%f",ptExpected.X, ptOrg.X)
			}
		}
	}

	func TestLambertNormal(t *testing.T){

		n := 6393174.9755
		lat := 0.97738438100
		a := 6378388.0000
		e := 0.081991890

		calc := lambertNormal(lat,a,e)
		if math.Abs(n - calc) > 1e-4{
			t.Errorf("Lambert Normal too far from expected - Expected : %f - Computed : %f", n, calc)

		}
	}


	func TestLatitudeFromLatitudeISO(t *testing.T){

		lat_iso :=[...]float64{1.00552653648,-0.30261690060 ,0.2000000000};
		e :=[...]float64{0.08199188998,0.08199188998,0.08199188998};
		eps:=[...]float64{1.0e-11,1.0e-11,1.0e-11};

		phi:=[...]float64{0.87266462600, -0.29999999997 ,0.19998903369};


		for index :=  range lat_iso{
			result := latitudeFromLatitudeISO(lat_iso[index], e[index], eps[index])

			if math.Abs(result - phi[index])  > 1e-4{
				t.Errorf("Latitude too far from expected - Expected : %f - Computed : %f",phi[index],result)
			}

		}


	}

	func TestCartesianToGeographic(t *testing.T){

		a:= [...]float64{6378249.2000, 6378249.2000 ,6378249.2000}
		e:= [...]float64{0.08248325679, 0.08248325679, 0.08248325679}
		x:= [...]float64{6376064.6950, 6378232.2150, 6376897.5370}
		y:= [...]float64{111294.6230, 18553.5780, 37099.7050}
		z:= [...]float64{128984.7250, 0.0000, -202730.9070}
		eps:= [...]float64{1e-11,1e-11,1e-11}

		lon:= [...]float64{0.01745329248, 0.00290888212, 0.00581776423}
		lat:= [...]float64{0.02036217457, 0.00000000000, -0.03199770301}
		he:= [...]float64{99.9995, 10.0001, 2000.0001}

		ign_eps := 1e-11;
		for i := range a{

			pt := &Point{x[i],y[i],z[i],Meter}
			pt.cartesianToGeographic(ParisLongitudeMeridian,a[i],e[i],eps[i])

			if math.Abs(pt.X - lon[i]) > ign_eps{
				t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f",lon[i],pt.X)
			}
			if math.Abs(pt.Y - lat[i]) > ign_eps{
				t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f",lat[i],pt.Y)
			}
			if math.Abs(pt.Z - he[i]) > 1e-4{
				t.Errorf("Height too far from expected - Expected : %.11f - Computed : %.11f",he[i],pt.Z)
			}
		}

	}
	func TestLambertToGeographic(t *testing.T){
		ptOrigin := &Point{1029705.083,272723.849,0,0}
		ptExpected := &Point{0.145512099,0.872664626,0,0}

		ptOrigin.lambertToGeographic(LambertI,GreenwichLongitudeMeridian,EClarkIGN,1e-9)

		if math.Abs(ptOrigin.X - ptExpected.X) > 1e-9{
			t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f",ptExpected.X,ptOrigin.X)
		}


		if  math.Abs(ptOrigin.Y - ptExpected.Y) > 1e-9{
			t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f",ptExpected.Y,ptOrigin.Y)
		}




	}
	func TestDegree(t * testing.T){

		ptOrigin := &Point{180.0,180.0,180.0,Degree}
		ptExpected := *ptOrigin

		ptOrigin.ToDegree()
		if  ptExpected.X != ptOrigin.X && ptExpected.Y != ptOrigin.Y && ptExpected.Z != ptOrigin.Z && ptExpected.Unit != ptOrigin.Unit {
			t.Error("Point should have the same value")
		}
	}
	func TestLambert93(t * testing.T){

		ptOrigin := &Point{668832.5384,6950138.7285,0,Meter}
		ptExpected := &Point{2.56865, 49.64961, 0,Degree}

		ptOrigin.ToWGS84(Lambert93)
		ptOrigin.ToDegree()

		if  math.Abs(ptOrigin.X - ptExpected.X) > 1e-5{
			t.Errorf("Longitude too far from expected - Expected : %.11f - Computed : %.11f",ptExpected.X,ptOrigin.X)
		}
		if  math.Abs(ptOrigin.Y - ptExpected.Y) > 1e-5{
			t.Errorf("Latitude too far from expected - Expected : %.11f - Computed : %.11f",ptExpected.Y,ptOrigin.Y)
		}


	}


func ExampleToWGS84(){

	var point *Point = &Point{668832.5384,6950138.7285,0,Meter}
	point.ToWGS84(Lambert93)
	fmt.Printf("Latitude:%5f - Longitude:%.5f", point.Y, point.X)
	// Output:
	// Latitude:0.866549 - Longitude:0.04483
}
