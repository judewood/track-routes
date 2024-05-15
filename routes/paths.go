package routes

import (
	"math"

	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/utils"
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
		v := pq.Dequeue() //take next item from the queue
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

func GetShortestRoute(startTIPLOC string, endTIPLOC string, g *ItemGraph) (int, int, *[]models.RouteSection) {
	var routeDetail []models.RouteSection
	var totalDistance int = 0
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]models.RouteSection)
	q := Queue{}
	pq := q.NewQ()
	startRouteSection := models.RouteSection{
		To:                 startTIPLOC,
		From:               "", // there is no previous TIPLOC
		Distance: 0,
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

	pathVal := prev[endTIPLOC] //start at the end
	totalDistance += pathVal.Distance
	//debugSection := fmt.Sprintf("%s,%s,%v,%s,%v", pathVal.From, pathVal.To, pathVal.Distance, pathVal.LineCode, totalDistance)
	routeDetail = append(routeDetail, pathVal)
	var finalArr []string
	finalArr = append(finalArr, endTIPLOC)
	for pathVal.From != startTIPLOC {
		//step back though the previous pairs of track sections in the route
		//and append them
		finalArr = append(finalArr, pathVal.From)
		pathVal = prev[pathVal.From]
		routeDetail = append(routeDetail, pathVal)
	}
	finalArr = append(finalArr, pathVal.From) //this will be the start TIPLOC
	reversedFinalArr := utils.ReverseCollection(&finalArr);
	reversedRouteDetail := utils.ReverseCollection(&routeDetail)
	numTrackSections := len(*reversedFinalArr) - 1 //one less than the number of TIPLOCs in the route
	return numTrackSections, dist[endTIPLOC],reversedRouteDetail
}
