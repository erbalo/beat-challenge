package configuration

import (
	"beat-challenge/src/services/fare"
	"beat-challenge/src/services/location"
	"beat-challenge/src/services/ride"
	"beat-challenge/src/util/channel"
	"beat-challenge/src/util/reader"
	"beat-challenge/src/util/writer"
	"fmt"
)

const (
	OUTPUT_DIRECTORY = "./output/"
)

type AppConfiguration interface {
	OrchestrateDependencies(file, outputName string)
}

type configurationImpl struct {
}

func New() AppConfiguration {
	return &configurationImpl{}
}

func (configuration *configurationImpl) OrchestrateDependencies(file, outputName string) {
	reader := reader.New()
	writer := writer.New(OUTPUT_DIRECTORY + outputName)

	rideService := ride.NewService()
	locationService := location.New()
	fareService := fare.New(locationService)

	if err := orchestrateExecution(file, reader, writer, rideService, fareService); err != nil {
		panic(err)
	}
}

func orchestrateExecution(file string, reader reader.Reader, writer writer.Writer, rideService ride.Service, fareService fare.Service) error {
	rowChannel, rowChannelError, err := reader.Read(file)
	if err != nil {
		return fmt.Errorf("failed to read file, with error: %v", err)
	}

	rideChannel, rideChannelError := rideService.Create(rowChannel)
	fareChannel := fareService.Estimate(rideChannel)

	success, writeChannelError, err := writer.Write(fareChannel)
	if err != nil {
		return fmt.Errorf("failed to initialize writing channel, with error: %v", err)
	}

	select {
	case <-success:
	case err := <-channel.MergeErrorChannels(rowChannelError, rideChannelError, writeChannelError):
		if err != nil {
			return fmt.Errorf("failed to calculate estimation, exists error from channerls, with error %v", err)
		}
	}

	return nil
}
