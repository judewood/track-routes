package routes

import (
	"fmt"
	"math"

	"github.com/judewood/routeDistances/fileStore"
)

func TIPLOCsAreConnected(startNode string, endNode string, g *ItemGraph, routeSectionCount int) bool {
	visited := make(map[string]bool)
	q := Queue{}
	pq := q.NewQ()
	start := Edge{
		From:     startNode,
		Distance: 0,
		LineCode: "",
	}
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue() //take next item form the queue
		if visited[v.From] {
			continue
		}
		visited[v.From] = true  //and mark it as visited
		near := g.Edges[v.From] //get its neighbours
		for _, val := range near {
			if val.From == endNode {
				return true
			}
			if !visited[val.From] {
				data := Edge{
					From:     val.From,
					Distance: 0,
					LineCode: val.LineCode,
				}
				pq.Enqueue(data) //add not visited TIPLOC to the queue
			}
		}
	}
	return false
}

func GetShortestPath(startNode string, endNode string, g *ItemGraph) (int, int) {
	var sections []string

	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	q := Queue{}
	pq := q.NewQ()
	startEdge := Edge{
		From:     startNode,
		To:       "", // there is no previous TIPLOC
		Distance: 0,
		LineCode: "",
	}
	for _, nval := range g.TIPLOCs {
		dist[*nval] = math.MaxInt64 //set initial distances to each TIPLOC to max value
	}
	dist[startNode] = startEdge.Distance // then set distance to start TIPLOC to zero
	pq.Enqueue(startEdge)
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.From] {
			continue
		}
		visited[v.From] = true
		near := g.Edges[v.From] //get edges for current TIPLOC
		for _, val := range near {
			if !visited[val.From] {
				if dist[v.From]+val.DistanceFrom < dist[val.From] {
					if v.From != startNode {
						debugOutput := fmt.Sprintf("queueing %s from %s. From: %v, line code %s. Queue size: %v", v.From, v.To, v.DistanceFrom, v.LineCode, pq.Size())
						//fmt.Println(debugOutput)
						sections = append(sections, debugOutput)
					}

					store := Edge{
						From:         val.From,
						To:           val.To,
						Distance:     dist[v.From] + val.DistanceFrom,
						DistanceFrom: val.DistanceFrom,
						LineCode:     val.LineCode,
					}
					dist[val.From] = dist[v.From] + val.DistanceFrom
					prev[val.From] = v.From
					pq.Enqueue(store)
				}
			}
		}
	}
	//fmt.Println("prevs", prev)
	pathVal := prev[endNode] //start at the end
	var finalArr []string
	finalArr = append(finalArr, endNode)
	for pathVal != startNode {
		//step back though the previous KV pairs of track sections in the prev array
		//and append them
		finalArr = append(finalArr, pathVal)
		pathVal = prev[pathVal]
	}
	finalArr = append(finalArr, pathVal) //this will be the start TIPLOC
	//reverse the array so it is ordered start to end
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	fileStore.WriteDebug("finalArray.txt", &finalArr)
	fileStore.WriteDebug("debugOutput.txt", &sections)
	numTrackSections := len(finalArr) - 1 //one less than the number of nodes
	return numTrackSections, dist[endNode]
}

type Edge struct {
	From         string
	To           string
	Distance     int
	DistanceFrom int
	LineCode     string
}
