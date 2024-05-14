package inputHandling

import (
	"fmt"
	"strings"

	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/fileStore"
	"github.com/judewood/routeDistances/models"
)

type InputStruct struct {
	fileHandler domain.FileStore
}

func New(fileHandler domain.FileStore) *InputStruct {
	return &InputStruct{
		fileHandler: fileHandler,
	}
}

func (d *InputStruct) GetInputData() (*[]models.RouteSection, error) {
	input, err := d.fileHandler.ReadFile()
	if err != nil {
		return &[]models.RouteSection{}, err
	}
	noDuplicates := RemoveDuplicates(input)
	inputData := createInputData(noDuplicates)
	return inputData, nil
}

func createInputData(routeSections *[]models.RouteSection) *[]models.RouteSection {
	var inputData []models.RouteSection
	for _, v := range *(*[]models.RouteSection)(routeSections) {
		item := models.RouteSection{To: v.To, From: v.From, Distance: v.Distance}
		inputData = append(inputData, item)
	}
	return &inputData
}

func RemoveDuplicates(input *[]models.RouteSection) *[]models.RouteSection {
	var duplicates []string
	var distinct []models.RouteSection

	for _, v := range *input {
		v.To = clean(v.To)
		v.From = clean(v.From)
		skip := false
		for _, u := range distinct {
			if v.To == u.To && v.From == u.From {
				if v.Distance != u.Distance {
					duplicate := fmt.Sprintf("%s to %s duplicate found. 1st Distance: %v, Line Code %s . 2nd Distance: %v, Line Code %s", v.To, v.From, v.Distance, v.LineCode, u.Distance, u.LineCode)
					duplicates = append(duplicates, duplicate)
					//use the shortest
					v.Distance = min(v.Distance, u.Distance)
				}
				skip = true
				break
			}
		}
		if !skip {
			distinct = append(distinct, v)
		}
	}
	if len(duplicates) > 0 {
		fileStore.WriteDebug("duplicates.txt", &duplicates)
	}
	return &distinct
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clean(input string) string {
	trimmed := strings.Trim(input, " ")
	return strings.ToUpper(trimmed)
}
