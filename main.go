package main

import (
	"fmt"

	"github.com/judewood/routeDistances/file"
)




func main() {
	fileHandler := file.NewCsvFileHandler("./input/Tracks.csv")

	records := fileHandler.ReadFile()

	// Loop to iterate through
	// and print each of the string slice
	for _, record := range records {
		fmt.Println(record)
	}

}