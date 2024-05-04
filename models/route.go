package models

type RouteSection struct {
	Start        string `csv:"FROM_TIPLOC"` // .csv column headers
	End          string `csv:"TO_TIPLOC"`
	Distance     int    `csv:"DISTANCE"`
	PassengerUse string `csv:"PASSENGER_USE"`
}

type StartEnd struct {
	Start string
	End   string
}

type RouteDistance struct {
	Start     string `csv:"FROM_TIPLOC"` // .csv column headers
	End       string `csv:"TO_TIPLOC"`
	Distance  int    `csv:"DISTANCE"`
	NumTracks int    `csv:"Num Tracks"`
}
