package writer_test

import (
	"beat-challenge/src/util/writer"
	"reflect"
	"testing"
)

const (
	TEST_FOLDER = "__test__/"
)

func TestWrite(test *testing.T) {

	test.Parallel()

	type args struct {
		fare []string
	}

	testCases := []struct {
		name             string
		args             args
		expected         int
		expectedErr      error
		expectedErrState bool
	}{
		{
			name: "happy path",
			args: args{
				fare: []string{"1", "9.12"},
			},
			expected:         0,
			expectedErr:      nil,
			expectedErrState: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			writer := writer.New(TEST_FOLDER + "test-writer.csv")

			fareChannel := make(chan []string)
			go func() {
				defer close(fareChannel)
				fareChannel <- testCase.args.fare
			}()

			successChannel := make(chan int)
			go func() {
				defer close(successChannel)
				successChannel <- testCase.expected
			}()

			errorChannel := make(chan error)
			go func() {
				defer close(errorChannel)
				errorChannel <- testCase.expectedErr
			}()

			gotSuccessChannel, gotErrorChannel, err := writer.Write(fareChannel)
			if (err != nil) != testCase.expectedErrState {
				t.Errorf("Write() error: %v, want error: %v", err, testCase.expectedErrState)
				return
			}

			if !reflect.DeepEqual(<-gotSuccessChannel, <-successChannel) {
				t.Errorf("Write() got success: %v, want done: %v", <-gotSuccessChannel, <-successChannel)
			}

			if !reflect.DeepEqual(<-gotErrorChannel, <-errorChannel) {
				t.Errorf("Write() got error: %v, want error: %v", <-gotErrorChannel, <-errorChannel)
			}
		})
	}

}
