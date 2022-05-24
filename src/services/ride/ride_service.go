package ride

import (
	"beat-challenge/src/domain"
	"fmt"
	"strconv"
)

const (
	BUFFER_SIZE = 50
)

type Service interface {
	Create(dataChannel <-chan []string) (<-chan domain.Ride, <-chan error)
}

type serviceImpl struct{}

func NewService() Service {
	return &serviceImpl{}
}

func (service *serviceImpl) Create(rowChannel <-chan []string) (<-chan domain.Ride, <-chan error) {
	rideChannel := make(chan domain.Ride, BUFFER_SIZE)
	errorChannel := make(chan error)

	go func() {
		var rideId uint64
		var rideLocations []domain.Location

		for row := range rowChannel {
			id, location, err := parseRow(row)
			if err != nil {
				errorChannel <- fmt.Errorf("failed to parse row, with error %v", err)
			} else {
				if len(rideLocations) != 0 && rideId != id {
					rideChannel <- domain.Ride{
						ID:        rideId,
						Locations: rideLocations,
					}

					rideLocations = []domain.Location{}
				}

				rideId = id
				rideLocations = append(rideLocations, *location)
			}

		}

		rideChannel <- domain.Ride{
			ID:        rideId,
			Locations: rideLocations,
		}

		close(rideChannel)
		close(errorChannel)
	}()

	return rideChannel, errorChannel
}

func parseRow(row []string) (uint64, *domain.Location, error) {
	if len(row) < 4 {
		return 0, nil, fmt.Errorf("row size hasn't the required structure, with error %v", row)
	}

	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert ride id")
	}

	latitude, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert latitude")
	}

	longitude, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert longitude")
	}

	timestamp, err := strconv.ParseInt(row[3], 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to convert timestamp")
	}

	return id, &domain.Location{
		Lat:       latitude,
		Lon:       longitude,
		Timestamp: int64(timestamp),
	}, nil
}
