package file

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
)

type FileStruct struct {
	inputFile string
	outputFile string
}

func NewCsv(inputFile, outputFile string) domain.FileHandler {
	return &FileStruct{
		inputFile: inputFile,
		outputFile: outputFile,
	}
}

func (f *FileStruct) ReadFile() *[]models.RouteSection {
	// os.Open() opens file in read-only mode and returns a pointer of type os.File
	file, err := os.Open(f.inputFile)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file just before exiting this function for any reason
	defer file.Close()

	routeSections := []models.RouteSection{}

	if err := gocsv.UnmarshalFile(file, &routeSections); err != nil {
		panic(err)
	}
	return &routeSections
}

func (f *FileStruct) WriteFile(records *[]models.RouteDistance) {
	 fmt.Println("hello")
// to download file inside downloads folder
file, err := os.Create(f.outputFile)
if err != nil {
 // handle error
 fmt.Println("oops")
 log.Fatal("Could not create output file ")
}
defer file.Close()
fmt.Printf("records jude %v", records)
gocsv.MarshalFile(records, file)
}

