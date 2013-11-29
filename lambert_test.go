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
func TestDegree(t * testing.T){

	ptOrigin := &Point{180.0,180.0,180.0,Degree}
	ptExpected := *ptOrigin

	ptOrigin.ToDegree()
	if  ptExpected.X != ptOrigin.X && ptExpected.Y != ptOrigin.Y && ptExpected.Z != ptOrigin.Z && ptExpected.Unit != ptOrigin.Unit {
		t.Error("Point should have the same value")
	}
}
func TestLambert93(t * testing.T){

	ptOrigin := &Point{668832.5384,6950138.7285,0,0};
	ptOrigin.ToWGS84(Lambert93)
	fmt.Printf("Lat:%.9f - Lon:%.9f",ptOrigin.Y, ptOrigin.X)

}
