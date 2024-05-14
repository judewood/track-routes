package routes

type InputGraph struct {
	Edges []Edge
	From  string
	To    string
}

// Create an array of adjacent TIPLOC pairs and the distance between them in the direction of travel
func CreateInputGraph(inputData *[]Edge, from, to string) InputGraph {
	inputGraph := InputGraph{
		Edges: *inputData,
		From:  from,
		To:    to,
	}
	return inputGraph
}
