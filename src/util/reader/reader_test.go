package reader_test

import (
	"beat-challenge/src/util/reader"
	"reflect"
	"testing"
)

const (
	TEST_FOLDER = "__test__/"
)

func TestRead(test *testing.T) {
	test.Parallel()

	type args struct {
		filePath string
	}

	testCases := []struct {
		name               string
		args               args
		expectedRow        []string
		expectedErr        error
		expectedErrorState bool
	}{
		{
			name: "happy path",
			args: args{
				filePath: TEST_FOLDER + "paths.csv",
			},
			expectedRow:        []string{"1000", "37.966660", "23.728308", "1405594957"},
			expectedErr:        nil,
			expectedErrorState: false,
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(t *testing.T) {
			reader := reader.New()

			rowChannel := make(chan []string)
			go func() {
				defer close(rowChannel)
				rowChannel <- testCase.expectedRow
			}()

			errorChannel := make(chan error)
			go func() {
				defer close(errorChannel)
				errorChannel <- testCase.expectedErr
			}()

			gotRowChannel, gotErrorChannel, err := reader.Read(testCase.args.filePath)
			if (err != nil) != testCase.expectedErrorState {
				t.Errorf("Read() err: %v, want err: %v", err, testCase.expectedErrorState)
				return

			}
			if !reflect.DeepEqual(<-gotRowChannel, <-rowChannel) {
				t.Errorf("Read() got row: %v, want row: %v", <-gotRowChannel, <-rowChannel)
			}

			if !reflect.DeepEqual(<-gotErrorChannel, <-errorChannel) {
				t.Errorf("Read() got err: %v, want err: %v", <-gotErrorChannel, <-errorChannel)
			}
		})
	}
}
