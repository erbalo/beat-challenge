package ride_test

import (
	"beat-challenge/src/domain"
	"beat-challenge/src/services/ride"
	"reflect"
	"testing"
)

func TestCreate(test *testing.T) {
	test.Parallel()

	type args struct {
		row []string
	}

	testCases := []struct {
		name         string
		args         args
		expectedRide domain.Ride
		expectedErr  error
	}{
		{
			name: "happy path",
			args: args{row: []string{"1", "37.966660", "23.728308", "1405594957"}},
			expectedRide: domain.Ride{
				ID: 1,
				Locations: []domain.Location{
					{
						Lat:       37.966660,
						Lon:       23.728308,
						Timestamp: 1405594957,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			rideService := ride.NewService()

			rowChannel := make(chan []string)
			go func() {
				defer close(rowChannel)
				rowChannel <- testCase.args.row
			}()

			rideChannel := make(chan domain.Ride)
			go func() {
				defer close(rideChannel)
				rideChannel <- testCase.expectedRide
			}()

			errorChannel := make(chan error)
			go func() {
				defer close(errorChannel)
				errorChannel <- testCase.expectedErr
			}()

			gotRideChannel, gotErrorChannel := rideService.Create(rowChannel)
			if !reflect.DeepEqual(<-gotRideChannel, <-rideChannel) {
				t.Errorf("Create() ride: %v, want: %v", <-gotRideChannel, <-rideChannel)
			}

			if !reflect.DeepEqual(<-gotErrorChannel, <-errorChannel) {
				t.Errorf("Create() err: %v, want: %v", <-gotErrorChannel, <-errorChannel)
			}
		})
	}

}
