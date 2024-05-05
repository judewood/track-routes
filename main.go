package main

import (
	"fmt"
	"log"

	"github.com/judewood/routeDistances/file"
	"github.com/judewood/routeDistances/inputHandling"
	"github.com/judewood/routeDistances/outputHandling"
)

func main() {
	fileHandler := file.NewCsv("./input/Tracks.csv", "./output/sample-output.csv")
	inputHandler := inputHandling.New(fileHandler)
	outputHandler := outputHandling.New(fileHandler)
	fmt.Println("Getting and formatting input route sections")
	inputData, err := inputHandler.GetInputData()
		if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Printf("\n\nCalculating and outputting shortest routes to %s", "./output/sample-output.csv")
	outputHandler.OutputRoutes(inputData)
}