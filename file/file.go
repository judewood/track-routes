package file

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
)

type FileStruct struct {
	inputFile  string
	outputFile string
}

func NewCsv(inputFile, outputFile string) domain.FileHandler {
	return &FileStruct{
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}

func (f *FileStruct) ReadFile() (*[]models.RouteSection, error) {
	// open file in read-only mode
	file, err := os.Open(f.inputFile)
	if err != nil {
		log.Fatal("Error while reading the file", err)
		return &[]models.RouteSection{}, err
	}

	// Ensure file is closed before exiting function
	defer file.Close()

	routeSections := []models.RouteSection{}
	err = gocsv.UnmarshalFile(file, &routeSections)
	if err != nil {
		log.Fatal("Error while reading the file", err)
		return &[]models.RouteSection{}, err
	}
	return &routeSections, nil
}

func (f *FileStruct) WriteFile(records *[]models.RouteDistance) error {
	file, err := os.Create(f.outputFile)
	if err != nil {
		// handle error
		log.Fatal("Could not create output file")
		return err
	}

	// Ensure file is closed before exiting function
	defer file.Close()
	err = gocsv.MarshalFile(records, file)
	if err != nil {
		log.Fatal("Error while writing output file", err)
		return err
	}
	return nil
}
