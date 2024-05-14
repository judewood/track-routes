package routes

import "github.com/judewood/routeDistances/models"

type InputGraph struct {
	RouteSections []models.RouteSection
	From          string
	To            string
}

// Create an array of adjacent TIPLOC pairs and the distance between them in the direction of travel
func CreateInputGraph(routeSections *[]models.RouteSection, from, to string) InputGraph {
	inputGraph := InputGraph{
		RouteSections: *routeSections,
		From:          from,
		To:            to,
	}
	return inputGraph
}
