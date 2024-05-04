package models

type RouteStruct struct  {
    Start string  `csv:"FROM_TIPLOC"` // .csv column headers
    End string    `csv:"TO_TIPLOC"`
	Distance int  `csv:"DISTANCE"`

}