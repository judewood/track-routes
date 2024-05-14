package routes

import (
	"sync"

	"github.com/judewood/routeDistances/models"
)

type ItemGraph struct {
	TIPLOCs       []*string
	RouteSections map[string][]*models.RouteSection
	Lock          sync.RWMutex
}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	TIPLOCs := make(map[string]string) //dictionary key: TIPLOC and value = TIPLOC
	for _, v := range data.RouteSections {
		if _, found := TIPLOCs[v.To]; !found {
			TIPLOCs[v.To] = v.To
			g.AddTIPLOC(v.To)
		}
		if _, found := TIPLOCs[v.From]; !found {
			TIPLOCs[v.From] = v.From
			g.AddTIPLOC(v.From)
		}
		g.AddRouteSection(TIPLOCs[v.To], TIPLOCs[v.From], v.Distance, v.LineCode)
	}
	return &g
}

// AddTIPLOC adds a TIPLOC to the graph
func (g *ItemGraph) AddTIPLOC(n string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.TIPLOCs = append(g.TIPLOCs, &n)
}

// AddRouteSection adds a route section to the graph
func (g *ItemGraph) AddRouteSection(from, to string, distanceFrom int, lineCode string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.RouteSections == nil {
		g.RouteSections = make(map[string][]*models.RouteSection)
	}
	routeSection := models.RouteSection{
		From:     to,
		To:       from,
		Distance: distanceFrom,
		LineCode: lineCode,
	}
	g.RouteSections[from] = append(g.RouteSections[from], &routeSection)
}
