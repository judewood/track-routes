package routes

import (

	"math"
	"sync"
)

func NodesAreConnected(startNode *Node, endNode *Node, g *ItemGraph, routeSectionCount int) (bool) {
	visited := make(map[string]bool)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue() //take next node form the queue
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true  //and mark it as visited
		near := g.Edges[*v.Node]  //get its neighbours
		for _, val := range near {
			if(val.Node.Value == endNode.Value) {
				return true
			}
			if !visited[val.Node.Value] {
				vertex := Vertex{
					Node:     val.Node,
					Distance: 0,
				}
				pq.Enqueue(vertex) //add not visited node to the queue
			}
		}
	}
	return false
}


func GetShortestPath(startNode *Node, endNode *Node, g *ItemGraph) (int, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]

		for _, val := range near {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight
					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
			}
		}
	}
	pathVal := prev[endNode.Value]
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathVal != startNode.Value {
		finalArr = append(finalArr, pathVal)
		pathVal = prev[pathVal]
	}
	finalArr = append(finalArr, pathVal)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	// print the route for debugging
	// for _, f := range finalArr {
	// 	fmt.Println(f)
	// }
	
	numTracks := len(finalArr)-1 //one less than the number of nodes
	return numTracks, dist[endNode.Value]

}

type Node struct {
	Value string
}

type Edge struct {
	Node   *Node
	Weight int
}

type Vertex struct {
	Node     *Node
	Distance int
}

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

type InputGraph struct {
	Graph []InputData `json:"graph"`
	From  string      `json:"from"`
	To    string      `json:"to"`
}

type InputData struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
}