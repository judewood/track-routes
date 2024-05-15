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
		From:               "", //No from value as this is our start point
		To:                 startTIPLOC,
		CumulativeDistance: 0,
		LineCode:           "",
	}
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		v := pq.Dequeue() //take next item form the queue
		if visited[v.To] {
			continue
		}
		visited[v.To] = true          //and mark it as visited
		near := g.RouteSections[v.To] //get its neighbours
		for _, val := range near {
			if val.To == endTIPLOC {
				return true
			}
			if !visited[val.To] {
				data := models.RouteSection{
					To:                 val.To,
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
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]models.RouteSection)
	q := Queue{}
	pq := q.NewQ()
	startRouteSection := models.RouteSection{
		To:                 startTIPLOC,
		From:               "", // there is no previous TIPLOC
		CumulativeDistance: 0,
		LineCode:           "",
	}
	for _, tiploc := range g.TIPLOCs {
		dist[*tiploc] = math.MaxInt64 //set initial distances to each TIPLOC to max value
	}
	dist[startTIPLOC] = 0 // then set distance to start TIPLOC to zero
	pq.Enqueue(startRouteSection)

	for !pq.IsEmpty() {
		currTIPLOC := pq.Dequeue()
		if visited[currTIPLOC.To] {
			continue
		}
		visited[currTIPLOC.To] = true
		connectedTIPLOCs := g.RouteSections[currTIPLOC.To]
		for _, connectedTIPLOC := range connectedTIPLOCs {
			if visited[connectedTIPLOC.To] {
				continue
			}

			newDistance := dist[currTIPLOC.To] + connectedTIPLOC.Distance
			if newDistance < dist[connectedTIPLOC.To] {
				routeSection := models.RouteSection{
					To:                 connectedTIPLOC.To,
					From:               connectedTIPLOC.From,
					CumulativeDistance: newDistance,
					Distance:           connectedTIPLOC.Distance,
					LineCode:           connectedTIPLOC.LineCode,
				}
				dist[connectedTIPLOC.To] = newDistance  // update distance to the connected TIPLOC
				prev[connectedTIPLOC.To] = routeSection // Add connected TIPLOC to the shortest route
				pq.Enqueue(routeSection)
			}

		}
	}
	var debugOutput []string
	pathVal := prev[endTIPLOC] //start at the end
	debugSection := fmt.Sprintf("Route section. From %s to %s. Section distance %v. Line Code %s", pathVal.From, pathVal.To, pathVal.Distance, pathVal.LineCode)
	debugOutput = append(debugOutput, debugSection)
	var finalArr []string
	finalArr = append(finalArr, endTIPLOC)
	for pathVal.From != startTIPLOC {
		//step back though the previous pairs of track sections in the route
		//and append them
		finalArr = append(finalArr, pathVal.From)
		pathVal = prev[pathVal.From]
		debugSection := fmt.Sprintf("Route section. From %s to %s. Section distance %v. Line Code %s", pathVal.From, pathVal.To, pathVal.Distance, pathVal.LineCode)
		debugOutput = append(debugOutput, debugSection)
	}
	finalArr = append(finalArr, pathVal.From) //this will be the start TIPLOC
	//reverse the array so it is ordered start to end
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	fileStore.WriteDebug("debugOutput.txt", &debugOutput)
	numTrackSections := len(finalArr) - 1 //one less than the number of TIPLOCs in the route
	return numTrackSections, dist[endTIPLOC]
}
