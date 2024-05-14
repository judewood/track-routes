package fileStore

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
)

type FileStore struct {
	inputFile  string
	outputFile string
}

func NewCsv(inputFile, outputFile string) domain.FileStore {
	return &FileStore{
		inputFile:  inputFile,
		outputFile: outputFile,
	}
}

func (f *FileStore) ReadFile() (*[]models.RouteSection, error) {
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

func (f *FileStore) WriteFile(records *[]models.RouteDistance) (int, error) {
	file, err := os.Create(f.outputFile)
	if err != nil {
		// handle error
		log.Fatal("Could not create output file")
		return 0, err
	}

	// Ensure file is closed before exiting function
	defer file.Close()
	err = gocsv.MarshalFile(records, file)
	if err != nil {
		log.Fatal("Error while writing output file", err)
		return 0, err
	}
	return len(*records), nil
}

func WriteDebug(filename string, items *[]string, ) (int, error) {
	file, err := os.Create(filename)
	if err != nil {
		// handle error
		msg := fmt.Sprintf("Could not create file %s", filename)
		log.Fatal(msg)
		return 0, err
	}

	// Ensure file is closed before exiting function
	defer file.Close()
	for _, line := range *items {
        _, err := file.WriteString(line + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
	return len(*items), nil
}
