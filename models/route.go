package models

type RouteSection struct {
	From         string `csv:"FROM_TIPLOC"` // .csv column headers
	To           string `csv:"TO_TIPLOC"`
	DistanceFrom int    `csv:"DISTANCE"`
	LineCode     string `csv:"LINE_CODE"`
}

type StartEnd struct {
	From string
	To   string
}

type RouteDistance struct {
	From      string `csv:"FROM_TIPLOC"` // .csv column headers
	To        string `csv:"TO_TIPLOC"`
	Distance  int    `csv:"DISTANCE"`
	NumTracks int    `csv:"Num Tracks"`
}
