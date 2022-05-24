package domain_test

import (
	"beat-challenge/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDayTime(test *testing.T) {
	test.Parallel()

	testCases := []struct {
		name     string
		location domain.Location
		expected bool
	}{
		{
			name:     "when day time is given",
			location: domain.Location{Timestamp: 1405594955},
			expected: true,
		},
		{
			name:     "when night time is given",
			location: domain.Location{Timestamp: 1624754398},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			got := testCase.location.IsDatyTime()

			assert.Equal(t, testCase.expected, got)
		})
	}
}
