package outputhandling

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
	sampleRoutes := s.GetSampleRoutes()	
	doneChannels := make([]chan models.RouteDistance, len(*sampleRoutes))
	for i := range doneChannels {
		doneChannels[i] = make(chan models.RouteDistance)
	}

	var results []models.RouteDistance
	for i, val := range *sampleRoutes {
		go s.GetResult(val, inputData, doneChannels[i])
		results = append(results, <-doneChannels[i])
	}
	fmt.Printf("The results %v ", results)
	s.fileHandler.WriteFile(&results)
}

func (s *OutputStruct) GetSampleRoutes() *[]models.StartEnd {
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

func (s *OutputStruct) GetResult(route models.StartEnd, inputData *[]routes.InputData, doneChan chan models.RouteDistance) {
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

