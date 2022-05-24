package location_test

import (
	"beat-challenge/src/domain"
	"beat-challenge/src/services/location"
	"testing"
)

func TestDistance(test *testing.T) {
	test.Parallel()

	type args struct {
		startLocation domain.Location
		endLocation   domain.Location
	}

	testCases := []struct {
		name             string
		args             args
		expectedDistance float64
	}{
		{
			name: "happy path",
			args: args{
				startLocation: domain.Location{
					Lat: 37.966660,
					Lon: 23.728308,
				},
				endLocation: domain.Location{
					Lat: 37.966627,
					Lon: 23.728263,
				},
			},
			expectedDistance: 0.005387608950290441,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			service := location.New()

			if got := service.Distance(testCase.args.startLocation, testCase.args.endLocation); got != testCase.expectedDistance {
				t.Errorf("Distance() distance: %v, want: %v", got, testCase.expectedDistance)
			}
		})
	}
}
