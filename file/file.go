package file

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/judewood/routeDistances/domain"
)

type FileReader struct{
    fileName string 
}
func NewCsvFileHandler(fileName string) domain.FileHandler {
	return &FileReader {
		fileName: fileName,
	}
}

func (f *FileReader)  ReadFile() [][]string {
	// os.Open() opens file in read-only mode and returns a pointer of type os.File
	file, err := os.Open(f.fileName)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file just before exiting this function for any reason
	defer file.Close()

    // creates a new csv.Reader that reads from the file
	reader := csv.NewReader(file)

	//skip first line
	reader.Read()

	// Read all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}
	return records
}