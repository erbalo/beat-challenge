package fare

import (
	"beat-challenge/src/domain"
	"beat-challenge/src/services/location"
	"fmt"
	"math"
)

const (
	BUFFER_SIZE           = 50
	MIN_SPEED             = float64(10)
	MAX_SPEED             = float64(100)
	STANDARD_FLAG         = float64(1.3)
	IDLE_HOUR_RATE        = float64(11.9)
	MIN_TOTAL             = float64(3.47)
	ON_TRANSIT_RATE_DAY   = float64(0.74)
	ON_TRANSIT_RATE_NIGHT = float64(1.3)
)

type Service interface {
	Estimate(rideChannel <-chan domain.Ride) <-chan []string
}

type serviceImpl struct {
	locationService location.Service
}

func New(locationService location.Service) Service {
	return &serviceImpl{
		locationService: locationService,
	}
}

func (service *serviceImpl) Estimate(rideChannel <-chan domain.Ride) <-chan []string {
	fareChannel := make(chan []string, BUFFER_SIZE)
	go func() {
		for ride := range rideChannel {
			estimatedFare := service.estimate(&ride)

			fareChannel <- []string{
				fmt.Sprintf("%v", estimatedFare.RideID),
				fmt.Sprintf("%.2f", estimatedFare.Amount),
			}
		}

		close(fareChannel)
	}()

	return fareChannel
}

func (service *serviceImpl) estimate(ride *domain.Ride) *domain.Fare {
	total := STANDARD_FLAG

	for i := 0; i < len(ride.Locations)-1; i++ {
		for j := i + 1; j < len(ride.Locations); j++ {
			startLocation := ride.Locations[i]
			endLocation := ride.Locations[i+1]

			deltaTime := float64(endLocation.Timestamp - startLocation.Timestamp)

			deltaDistance := service.locationService.Distance(startLocation, endLocation)
			speed := (deltaDistance / deltaTime) * 3600

			if speed > MAX_SPEED {
				i++
				continue
			}

			if speed <= MIN_SPEED {
				total += (deltaTime / 3600) * IDLE_HOUR_RATE
				break
			}

			if startLocation.IsDatyTime() {
				total += deltaDistance * ON_TRANSIT_RATE_DAY
			} else {
				total += deltaDistance * ON_TRANSIT_RATE_NIGHT
			}

			break
		}
	}

	total = math.Max(total, MIN_TOTAL)

	return &domain.Fare{
		RideID: ride.ID,
		Amount: total,
	}
}
