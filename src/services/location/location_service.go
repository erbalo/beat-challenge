package location

import (
	"beat-challenge/src/domain"
	"math"
)

const (
	EARTH_RADIUS = float64(6371)
)

type Service interface {
	Distance(startLocation, endLocation domain.Location) float64
}

type serviceImpl struct{}

func New() Service {
	return &serviceImpl{}
}

/**
* 	Haversine
* 	formula: a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
*	c = 2 ⋅ atan2( √a, √(1−a) )
*	d = R ⋅ c
 */
func (service *serviceImpl) Distance(startLocation, endLocation domain.Location) float64 {
	var deltaLat = (endLocation.Lat - startLocation.Lat) * (math.Pi / 180)
	var deltaLon = (endLocation.Lon - startLocation.Lon) * (math.Pi / 180)

	startLatRadian := startLocation.Lat * (math.Pi / 180)
	endLatRadian := endLocation.Lat * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(startLatRadian)*math.Cos(endLatRadian)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}
