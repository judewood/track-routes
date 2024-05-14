package routes

import (
	"sync"
)

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*Node)
	for _, v := range data.InputData {
		if _, found := nodes[v.To]; !found {
			nA := Node{v.To}
			nodes[v.To] = &nA
			g.AddNode(&nA)
		}
		if _, found := nodes[v.From]; !found {
			nA := Node{v.From}
			nodes[v.From] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.To], nodes[v.From], v.DistanceFrom, v.LineCode)
	}
	return &g
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.Nodes = append(g.Nodes, n)
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(fromNode, toNode *Node, distanceFrom int, lineCode string) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		FromNode:     toNode,
		ToNode:       fromNode,
		DistanceFrom: distanceFrom,
		LineCode:     lineCode,
	}
	g.Edges[*fromNode] = append(g.Edges[*fromNode], &ed1)
}
