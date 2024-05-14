package models

type RouteSection struct {
	From               string `csv:"FROM_TIPLOC"` // .csv column headers
	To                 string `csv:"TO_TIPLOC"`
	CumulativeDistance int
	Distance           int    `csv:"DISTANCE"`
	LineCode           string `csv:"LINE_CODE"`
}

type StartEnd struct {
	From string
	To   string
}

// RouteDistance is the structure for the output CSV file
type RouteDistance struct {
	From      string `csv:"FROM_TIPLOC"` // .csv column headers
	To        string `csv:"TO_TIPLOC"`
	Distance  int    `csv:"DISTANCE"`
	NumTracks int    `csv:"Num Tracks"`
}
