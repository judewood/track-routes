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

func (s *OutputStruct) OutputRoutes(inputData *[]models.RouteSection, routes *[]models.StartEnd) (int, error) {
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
			From: "BERKHMD",
			To:   "TRING",
		},
		{
			From: "HYWRDSH",
			To:   "KEYMERJ",
		},
		{
			From: "BERKHMD",
			To:   "HEMLHMP",
		},
		{
			From: "BHAMNWS",
			To:   "BHAMINT",
		},
		{
			From: "BERKHMD",
			To:   "WATFDJ",
		},
		{
			From: "EUSTON",
			To:   "BERKHMD",
		},
		{
			From: "MNCRPIC",
			To:   "CRDFCEN",
		},
		{
			From: "KNGX",
			To:   "EDINBUR",
		},
		{
			From: "THURSO",
			To:   "PENZNCE",
		},
		{
			From: "PHBR",
			To:   "RYDP",
		},
	}
	return &outputSample
}

func getResult(route models.StartEnd, inputData *[]models.RouteSection, doneChan chan models.RouteDistance) {
	inputGraph := routes.CreateInputGraph(inputData, route.From, route.To)
	routesGraph := routes.CreateGraph(inputGraph)
	routeIsConnected := routes.TIPLOCsAreConnected(route.From, route.To, routesGraph, len(*inputData))
	if !routeIsConnected {
		fmt.Printf("TIPLOCs %s and %s are not connected. Setting distance to -1", route.From, route.To)
		unconnected := models.RouteDistance{
			From:      route.From,
			To:        route.To,
			Distance:  -1,
			NumTracks: -1,
		}
		doneChan <- unconnected
		return
	}
	numTracks, distance := routes.GetShortestRoute(route.From, route.To, routesGraph)

	r := models.RouteDistance{
		From:      route.From,
		To:        route.To,
		Distance:  distance,
		NumTracks: numTracks,
	}
	fmt.Printf("Route for %s to %s is %v with num tracks %v", route.From, route.To, r.Distance, r.NumTracks)
	doneChan <- r
}
