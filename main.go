package main

import (
	"fmt"

	"github.com/judewood/routeDistances/file"
	"github.com/judewood/routeDistances/inputHandling"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/outputHandling"
	"github.com/judewood/routeDistances/routes"
)

func main() {
	fileHandler := file.NewCsv("./input/Tracks.csv", "./output/sample-output.csv")
	inputHandler := inputHandling.New(fileHandler)
	outputHandler := outputhandling.New(fileHandler)

	inputData := inputHandler.GetInputData()

	outputHandler.OutputRoutes(inputData)
	// sampleRoutes := outputHandler.GetSampleRoutes()	
	// doneChannels := make([]chan models.RouteDistance, len(*sampleRoutes))
	// for i := range doneChannels {
	// 	doneChannels[i] = make(chan models.RouteDistance)
	// }

	// var results []models.RouteDistance
	// for i, val := range *sampleRoutes {
	// 	go outputHandler.GetResult(val, inputData, doneChannels[i])
	// 	results = append(results, <-doneChannels[i])
	// }
	// fmt.Printf("The results %v ", results)
	// fileHandler.WriteFile(&results)
}

func GetResult(route models.StartEnd, inputData *[]routes.InputData, doneChan chan models.RouteDistance) {
	fmt.Printf("Getting shortest route for %s to %s", route.Start, route.End)
	fmt.Println()
	//time.Sleep(1 * time.Second)

	inputGraph := routes.CreateInputGraphJude(inputData, route.Start, route.End)
	k := routes.CreateGraph(inputGraph)

	node1 := routes.Node{
		Value: route.Start,
	}
	node2 := routes.Node{
		Value: route.End,
	}

	numTracks, distance := routes.GetShortestPath(&node1, &node2, k)

	r := models.RouteDistance{
		Start:     route.Start,
		End:       route.End,
		Distance:  distance,
		NumTracks: numTracks,
	}
	doneChan <- r
}
