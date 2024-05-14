package routes

import (
	"sync"
)

type ItemGraph struct {
	TIPLOCs []*string
	Edges   map[string][]*Edge
	Lock    sync.RWMutex
}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]string)
	for _, v := range data.Edges {
		if _, found := nodes[v.To]; !found {
			nodes[v.To] = v.To
			g.AddTIPLOC(&v.To)
		}
		if _, found := nodes[v.From]; !found {
			nodes[v.From] = v.From
			g.AddTIPLOC(&v.From)
		}
		g.AddRouteSection(nodes[v.To], nodes[v.From], v.DistanceFrom, v.LineCode)
	}
	return &g
}

// AddTIPLOC adds a TIPLOC to the graph
func (g *ItemGraph) AddTIPLOC(n *string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.TIPLOCs = append(g.TIPLOCs, n)
}

// AddRouteSection adds an edge to the graph
func (g *ItemGraph) AddRouteSection(from, to string, distanceFrom int, lineCode string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[string][]*Edge)
	}
	edge := Edge{
		From:         to,
		To:           from,
		DistanceFrom: distanceFrom,
		LineCode:     lineCode,
	}
	g.Edges[from] = append(g.Edges[from], &edge)
}
