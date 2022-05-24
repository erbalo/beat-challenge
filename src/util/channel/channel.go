package channel

func MergeErrorChannels(channels ...<-chan error) <-chan error {
	errorChannel := make(chan error)

	for _, channel := range channels {
		go func(ch <-chan error) {
			for err := range ch {
				errorChannel <- err
			}
		}(channel)
	}

	return errorChannel
}
