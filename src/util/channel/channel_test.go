package channel_test

import (
	"beat-challenge/src/util/channel"
	"reflect"
	"testing"
)

func TestMergeErrorChannels(test *testing.T) {
	test.Parallel()

	type args struct {
		channels error
	}

	testCases := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "happy path",
			args: args{
				channels: nil,
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			channels := make(chan error)
			go func() {
				defer close(channels)
				channels <- testCase.args.channels
			}()

			errorChannel := make(chan error)
			go func() {
				defer close(errorChannel)
				errorChannel <- testCase.expectedErr
			}()

			gotErrorChannel := channel.MergeErrorChannels(channels)
			if !reflect.DeepEqual(<-gotErrorChannel, <-errorChannel) {
				t.Errorf("MergeErrorChannels() %v, want: %v", <-gotErrorChannel, <-errorChannel)
			}
		})
	}
}
