package fare_test

import (
	"beat-challenge/src/domain"
	"beat-challenge/src/services/fare"
	"beat-challenge/src/services/location"
	"reflect"
	"testing"
)

func TestEstimate(test *testing.T) {
	test.Parallel()

	type fields struct {
		locationSercice location.Service
	}

	type args struct {
		ride domain.Ride
	}

	testCases := []struct {
		name         string
		fields       fields
		args         args
		expectedFare []string
	}{
		{
			name: "happy path",
			fields: fields{
				locationSercice: nil,
			},
			args: args{
				ride: domain.Ride{
					ID: 1,
					Locations: []domain.Location{
						{
							Lat:       37.966660,
							Lon:       23.728308,
							Timestamp: 1405594957,
						},
					},
				},
			},
			expectedFare: []string{"1", "3.47"},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			fareService := fare.New(testCase.fields.locationSercice)

			rideChannel := make(chan domain.Ride)
			go func() {
				defer close(rideChannel)
				rideChannel <- testCase.args.ride
			}()

			fareChannel := make(chan []string)
			go func() {
				defer close(fareChannel)
				fareChannel <- testCase.expectedFare
			}()

			gotFareChannel := fareService.Estimate(rideChannel)
			if !reflect.DeepEqual(<-gotFareChannel, <-fareChannel) {
				t.Errorf("Estimate() amount: %v, want: %v", <-gotFareChannel, <-fareChannel)
			}
		})
	}
}
