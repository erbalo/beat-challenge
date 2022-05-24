package writer

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type Writer interface {
	Write(in <-chan []string) (<-chan int, <-chan error, error)
}

type writerImpl struct {
	outputFile string
}

func New(outputFile string) Writer {
	return &writerImpl{
		outputFile: outputFile,
	}
}

func (writer *writerImpl) Write(fareChannel <-chan []string) (<-chan int, <-chan error, error) {
	path, err := filepath.Abs(writer.outputFile)
	if err != nil {
		return nil, nil, fmt.Errorf("not valid path %v with error: %v", writer.outputFile, err)
	}

	directory := filepath.Dir(path)
	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't create the path %v, with error: %v", writer.outputFile, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create the file with error: %v", err)
	}

	csvWriter := csv.NewWriter(file)
	successChannel := make(chan int)
	errorChannel := make(chan error)

	go func() {
		for row := range fareChannel {
			err := csvWriter.Write(row)
			if err != nil {
				errorChannel <- fmt.Errorf("cannot write with error: %v", err)
			}
		}

		csvWriter.Flush()
		file.Close()
		successChannel <- 0

		close(successChannel)
		close(errorChannel)
	}()

	return successChannel, errorChannel, nil
}
