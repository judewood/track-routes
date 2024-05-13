package routes

import (
	"fmt"
	"math"
)

func NodesAreConnected(startNode *Node, endNode *Node, g *ItemGraph, routeSectionCount int) bool {
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
		if v.Node.Value == "BERKHMD" || v.Node.Value == "BONENDJ" || v.Node.Value == "BORN421" || v.Node.Value == "HEMLHMP" {
			fmt.Println("dequeueing", *v)
			fmt.Println("queue size", pq.Size())
		}

		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]
		for _, val := range near {
			if !visited[val.FromNode.Value] {
				// if val.FromNode.Value == "BERKHMD" || val.FromNode.Value == "BONENDJ" || val.FromNode.Value == "BORN421" || val.FromNode.Value == "HEMLHMP" {
				// 	fmt.Println("visiting", val.FromNode.Value, ", ", val.Node2.Value, "DistanceFrom", val.DistanceFrom, "DistanceTo", val.DistanceTo)

				// }
				if dist[v.Node.Value]+val.DistanceFrom < dist[val.FromNode.Value] {
					if val.FromNode.Value == "BERKHMD" || val.FromNode.Value == "BONENDJ" || val.FromNode.Value == "BORN421" || val.FromNode.Value == "HEMLHMP" {
						fmt.Println("queueing", val.FromNode.Value, ", ", val.Node2.Value,  "DistanceFrom", val.DistanceFrom, "DistanceTo", val.DistanceTo)
					}

					store := Vertex{
						Node:     val.FromNode,
						Distance: dist[v.Node.Value] + val.DistanceFrom,
					}
					dist[val.FromNode.Value] = dist[v.Node.Value] + val.DistanceFrom
					prev[val.FromNode.Value] = v.Node.Value
					if val.FromNode.Value == "BERKHMD" || val.FromNode.Value == "BONENDJ" || val.FromNode.Value == "BORN421" || val.FromNode.Value == "HEMLHMP" {
						fmt.Println("queueing", val.FromNode.Value, ", ", val.Node2.Value, "DistanceFrom", val.DistanceFrom, "DistanceTo", val.DistanceTo)
						fmt.Println("store", store.Node.Value, store.Distance)
					}
					pq.Enqueue(store)
				}
			}
		}
	}
	pathVal := prev[endNode.Value]
	//fmt.Println("pathVal", prev[endNode.Value])
	var finalArr []string
	fmt.Println("finalArr 1", finalArr)
	finalArr = append(finalArr, endNode.Value)
	for pathVal != startNode.Value {
		finalArr = append(finalArr, pathVal)
		pathVal = prev[pathVal]
		fmt.Println("pathVal", pathVal)
	}
	finalArr = append(finalArr, pathVal)
	//fmt.Println("finalArr", finalArr)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	//print the route for debugging
	for _, f := range finalArr {
		fmt.Println(f)
	}

	numTracks := len(finalArr) - 1
	return numTracks, dist[endNode.Value]

}

type Node struct {
	Value string
}

type Edge struct {
	FromNode     *Node
	Node2        *Node
	DistanceFrom int
	DistanceTo   int
}

type Vertex struct {
	Node     *Node
	Distance int
}

type PriorityQueue []*Vertex

type InputData struct {
	To           string
	From         string
	DistanceFrom int
	DistanceTo   int
}
