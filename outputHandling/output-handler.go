package outputHandling

import (
	"fmt"

	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
)

type OutputStruct struct {
	fileHandler domain.FileStore
}

func New(fileHandler domain.FileStore) *OutputStruct {
	return &OutputStruct{
		fileHandler: fileHandler,
	}
}

func (s *OutputStruct) OutputRoutes(inputData *[]routes.InputData, routes *[]models.StartEnd) (int, error) {
	doneChannels := make([]chan models.RouteDistance, len(*routes))
	for i := range doneChannels {
		doneChannels[i] = make(chan models.RouteDistance)
	}

	var results []models.RouteDistance
	for i, val := range *routes {
		fmt.Println()
		go getResult(val, inputData, doneChannels[i])
		results = append(results, <-doneChannels[i])
	}
	numRecords, err := s.fileHandler.WriteFile(&results)
	if err != nil {
		return 0, err
	}
	return numRecords, nil
}

func GetSampleRoutes() *[]models.StartEnd {
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
		{
			Start: "PHBR",
			End:   "RYDP",
		},
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
	nodesAreConnected := routes.NodesAreConnected(&node1, &node2, routesGraph, len(*inputData))
	if !nodesAreConnected {
		fmt.Printf("TIPLOCs %s and %s are not connected. Setting distance to -1", node1.Value, node2.Value)
		unconnected := models.RouteDistance{
			Start:     route.Start,
			End:       route.End,
			Distance:  -1,
			NumTracks: -1,
		}
		doneChan <- unconnected
		return
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
