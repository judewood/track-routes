package main

import (
	"fmt"
	"log"

	"github.com/judewood/routeDistances/fileStore"
	"github.com/judewood/routeDistances/inputHandling"
	"github.com/judewood/routeDistances/outputHandling"
)

func main() {
	outputFile := "./output/sample-output.csv"
	fileHandler := fileStore.NewCsv("./input/Tracks.csv", outputFile)
	inputHandler := inputHandling.New(fileHandler)
	outputHandler := outputHandling.New(fileHandler)
	fmt.Println("Getting and formatting input route sections")
	inputData, err := inputHandler.GetInputData()
		if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Printf("\n\nCalculating and outputting shortest routes to %s", outputFile)
	numRecords, err := outputHandler.OutputRoutes(inputData, outputHandling.GetSampleRoutes())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("\n", numRecords, "routes written to", outputFile)
}