package outputHandling

import (
	"fmt"

	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
)

type OutputStruct struct {
	fileHandler domain.FileHandler
}

func New(fileHandler domain.FileHandler) *OutputStruct {
	return &OutputStruct{
		fileHandler: fileHandler,
	}
}

func (s *OutputStruct) OutputRoutes(inputData *[]routes.InputData) {
	sampleRoutes := getSampleRoutes()
	doneChannels := make([]chan models.RouteDistance, len(*sampleRoutes))
	for i := range doneChannels {
		doneChannels[i] = make(chan models.RouteDistance)
	}

	var results []models.RouteDistance
	for i, val := range *sampleRoutes {
		fmt.Println()
		go getResult(val, inputData, doneChannels[i])
		results = append(results, <-doneChannels[i])
	}
	s.fileHandler.WriteFile(&results)
}

func getSampleRoutes() *[]models.StartEnd {
	var outputSample = []models.StartEnd{
		{
			Start: "BERKHMD",
			End:   "TRING",
		},
		{
			Start: "HYWRDSH",
			End:   "KEYMERJ",
		},
		{
			Start: "BERKHMD",
			End:   "HEMLHMP",
		},
		{
			Start: "BHAMNWS",
			End:   "BHAMINT",
		},
		{
			Start: "BERKHMD",
			End:   "WATFDJ",
		},
		{
			Start: "EUSTON",
			End:   "BERKHMD",
		},
		{
			Start: "MNCRPIC",
			End:   "CRDFCEN",
		},
		{
			Start: "KNGX",
			End:   "EDINBUR",
		},
		{
			Start: "THURSO",
			End:   "PENZNCE",
		},
		// {
		// 	Start: "PHBR",
		// 	End:   "RYDP",
		// },
	}
	return &outputSample
}

func getResult(route models.StartEnd, inputData *[]routes.InputData, doneChan chan models.RouteDistance) {
	inputGraph := routes.CreateInputGraph(inputData, route.Start, route.End)
	routesGraph := routes.CreateGraph(inputGraph)

	node1 := routes.Node{
		Value: route.Start,
	}
	node2 := routes.Node{
		Value: route.End,
	}
	numTracks, distance := routes.GetShortestPath(&node1, &node2, routesGraph)

	r := models.RouteDistance{
		Start:     route.Start,
		End:       route.End,
		Distance:  distance,
		NumTracks: numTracks,
	}
	fmt.Printf("Route for %s to %s is %v with num tracks %v", route.Start, route.End, r.Distance, r.NumTracks)
	doneChan <- r
}
