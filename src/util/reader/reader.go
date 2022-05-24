package reader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	BUFFER_SIZE = 50
)

type Reader interface {
	Read(filePath string) (<-chan []string, <-chan error, error)
}

type readerImpl struct{}

func New() Reader {
	return &readerImpl{}
}

func (reader *readerImpl) Read(filePath string) (<-chan []string, <-chan error, error) {
	path, err := filepath.Abs(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("not valid path %v, with error: %v", filePath, err)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file %v with error: %v", path, err)
	}

	csvReader := csv.NewReader(bufio.NewReader(file))

	rowChannel := make(chan []string, BUFFER_SIZE)
	errorChannel := make(chan error)

	go func() {
		for {
			row, err := csvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				errorChannel <- fmt.Errorf("failed to read from file %v with error: %v", path, err)
			}

			rowChannel <- row
		}

		file.Close()
		close(rowChannel)
		close(errorChannel)
	}()

	return rowChannel, errorChannel, nil
}
