package routes

import (
	"fmt"
	"math"

	"github.com/judewood/routeDistances/fileStore"
	"github.com/judewood/routeDistances/models"
)

func TIPLOCsAreConnected(startTIPLOC string, endTIPLOC string, g *ItemGraph, routeSectionCount int) bool {
	visited := make(map[string]bool)
	q := Queue{}
	pq := q.NewQ()
	start := models.RouteSection{
		From:               startTIPLOC,
		CumulativeDistance: 0,
		LineCode:           "",
	}
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue() //take next item form the queue
		if visited[v.From] {
			continue
		}
		visited[v.From] = true          //and mark it as visited
		near := g.RouteSections[v.From] //get its neighbours
		for _, val := range near {
			if val.From == endTIPLOC {
				return true
			}
			if !visited[val.From] {
				data := models.RouteSection{
					From:               val.From,
					CumulativeDistance: 0,
					LineCode:           val.LineCode,
				}
				pq.Enqueue(data) //add not visited TIPLOC to the queue
			}
		}
	}
	return false
}

func GetShortestRoute(startTIPLOC string, endTIPLOC string, g *ItemGraph) (int, int) {
	var sections []string

	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	q := Queue{}
	pq := q.NewQ()
	startRouteSection := models.RouteSection{
		From:               startTIPLOC,
		To:                 "", // there is no previous TIPLOC
		CumulativeDistance: 0,
		LineCode:           "",
	}
	for _, nval := range g.TIPLOCs {
		dist[*nval] = math.MaxInt64 //set initial distances to each TIPLOC to max value
	}
	dist[startTIPLOC] = startRouteSection.CumulativeDistance // then set distance to start TIPLOC to zero
	pq.Enqueue(startRouteSection)
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.From] {
			continue
		}
		visited[v.From] = true
		near := g.RouteSections[v.From] //get sections connected to current TIPLOC
		for _, val := range near {
			if !visited[val.From] {
				if dist[v.From]+val.Distance < dist[val.From] {
					if v.From != startTIPLOC {
						debugOutput := fmt.Sprintf("queueing %s from %s. From: %v, line code %s. Queue size: %v", v.From, v.To, v.Distance, v.LineCode, pq.Size())
						//fmt.Println(debugOutput)
						sections = append(sections, debugOutput)
					}

					store := models.RouteSection{
						From:               val.From,
						To:                 val.To,
						CumulativeDistance: dist[v.From] + val.Distance,
						Distance:           val.Distance,
						LineCode:           val.LineCode,
					}
					dist[val.From] = dist[v.From] + val.Distance
					prev[val.From] = v.From
					pq.Enqueue(store)
				}
			}
		}
	}
	//fmt.Println("prevs", prev)
	pathVal := prev[endTIPLOC] //start at the end
	var finalArr []string
	finalArr = append(finalArr, endTIPLOC)
	for pathVal != startTIPLOC {
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
	numTrackSections := len(finalArr) - 1 //one less than the number of TIPLOCs in the route
	return numTrackSections, dist[endTIPLOC]
}
