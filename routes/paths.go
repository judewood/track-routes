package routes

import (
	"fmt"
	"math"

	"github.com/judewood/routeDistances/fileStore"
)

func NodesAreConnected(startNode *Node, endNode *Node, g *ItemGraph, routeSectionCount int) bool {
	visited := make(map[string]bool)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
		LineCode: "",
	}
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue() //take next node form the queue
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true //and mark it as visited
		near := g.Edges[*v.Node]     //get its neighbours
		for _, val := range near {
			if val.FromNode.Value == endNode.Value {
				return true
			}
			if !visited[val.FromNode.Value] {
				vertex := Vertex{
					Node:     val.FromNode,
					Distance: 0,
					LineCode: val.LineCode,
				}
				pq.Enqueue(vertex) //add not visited node to the queue
			}
		}
	}
	return false
}

func GetShortestPath(startNode *Node, endNode *Node, g *ItemGraph) (int, int) {
	var sections []string

	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
		LineCode: "",
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
			if !visited[val.FromNode.Value] {
				if dist[v.Node.Value]+val.DistanceFrom < dist[val.FromNode.Value] {
					if v.Node.Value != startNode.Value {
						debugOutput := fmt.Sprintf("queueing %s from %s. From: %v, line code %s. Queue size: %v", v.Node.Value, v.Node2.Value, v.DistanceFrom, v.LineCode, pq.Size())
						//fmt.Println(debugOutput)
						sections = append(sections, debugOutput)
					}

					store := Vertex{
						Node:         val.FromNode,
						Node2:        val.ToNode,
						Distance:     dist[v.Node.Value] + val.DistanceFrom,
						DistanceFrom: val.DistanceFrom,
						LineCode:     val.LineCode,
					}
					dist[val.FromNode.Value] = dist[v.Node.Value] + val.DistanceFrom
					prev[val.FromNode.Value] = v.Node.Value
					pq.Enqueue(store)
				} else {

				}
			}
		}
	}
	//fmt.Println("prevs", prev)
	pathVal := prev[endNode.Value] //start at the end
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathVal != startNode.Value {
		//step back though the previous KV pairs of track sections in the prev array
		//and append them
		finalArr = append(finalArr, pathVal)
		pathVal = prev[pathVal]
	}
	finalArr = append(finalArr, pathVal) //this will be the start node
	//reverse the array so it is ordered start to end
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	fileStore.WriteDebug("finalArray.txt", &finalArr)
	fileStore.WriteDebug("debugOutput.txt", &sections)
	numTrackSections := len(finalArr) - 1 //one less than the number of nodes
	return numTrackSections, dist[endNode.Value]
}

type Node struct {
	Value string
}

type Edge struct {
	FromNode     *Node
	ToNode       *Node
	DistanceFrom int
	LineCode     string
}

type Vertex struct {
	Node         *Node
	Node2        *Node
	Distance     int
	DistanceFrom int
	LineCode     string
}

type PriorityQueue []*Vertex

type InputData struct {
	To           string
	From         string
	DistanceFrom int
	LineCode     string
}
