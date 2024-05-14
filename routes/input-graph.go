package routes

type InputGraph struct {
	InputData []InputData `json:"graph"`
	From      string      `json:"from"`
	To        string      `json:"to"`
}

// Create an array of adjacent node pairs and the distance or weight between them
func CreateInputGraph(inputData *[]InputData, from, to string) InputGraph {
	inputGraph := InputGraph{
		InputData: *inputData,
		From:      from,
		To:        to,
	}
	return inputGraph
}
