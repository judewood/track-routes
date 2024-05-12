package routes

import "sync"

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*Node)
	for _, v := range data.Graph {
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
		g.AddEdge(nodes[v.To], nodes[v.From], v.DistanceFrom, v.DistanceTo)
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
func (g *ItemGraph) AddEdge(n1, n2 *Node, distanceFrom, distanceTo int) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:         n2,
		DistanceFrom: distanceFrom,
		DistanceTo:  distanceTo,
	}

	ed2 := Edge{
		Node:         n1,
		DistanceFrom: distanceFrom,
		DistanceTo:  distanceTo,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
}
