[![Build Status](https://travis-ci.org/YaGeek/lambertgo.png?branch=develop)](https://travis-ci.org/YaGeek/lambertgo)

#Description
A simple Go package to convert coordinates in Lambert projections system to GPS WGS84 coordinates. It is based on the [IGN alorithms and methods](http://geodesie.ign.fr/contenu/fichiers/documentation/algorithmes/notice/NTG_71.pdf)

# Install
	go get github.com/YaGeek/lambertgo
# Usage

```go
import lgo "github.com/YaGeek/lambertgo"
import "fmt"

//Declares origin point and translated point
var point *lgo.Point= &lgo.Point{994272.661,113467.422,0,lgo.Meter}

//Converts pointOrg in Lambert Zone 1 to WGS84 - Results in Radian
point.ToWGS84(lgo.LambertI)

//Converts to Degree
point.ToDegree()

fmt.Printf("Latitude:%.5f - Longitude:%.5f",point.Y,point.X)
```

#License
Copyright (c) 2013 Yannick Heinrich - Released under the GPLv2 License.

