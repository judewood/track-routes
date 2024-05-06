package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/judewood/routeDistances/fileStore"
	"github.com/judewood/routeDistances/inputHandling"
	"github.com/judewood/routeDistances/outputHandling"
)

func main() {
	outputFile := "./output/sample-output.csv"
	inputFile :=  "./input/Tracks.csv"
	fmt.Printf("\n\nCalculating and outputting shortest routes to %s", outputFile)
	fmt.Print("\nPress 'Enter' to continue...")
  	bufio.NewReader(os.Stdin).ReadBytes('\n') 
	fileHandler := fileStore.NewCsv(inputFile, outputFile)
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
	fmt.Printf("\n %v Routes written to %s", numRecords, outputFile)
}