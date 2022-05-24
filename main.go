package main

import (
	"beat-challenge/src/configuration"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("f", "", "-f paths.csv")
	output := flag.String("o", "", "-o result.csv")
	flag.Parse()

	if *file == "" {
		fmt.Println("parameter 'f' is required to read file")
		os.Exit(1)
	}

	if *output == "" {
		fmt.Println("parameter 'o' is required to save the result")
		os.Exit(1)
	}

	appConfiguration := configuration.New()
	appConfiguration.OrchestrateDependencies(*file, *output)
}
