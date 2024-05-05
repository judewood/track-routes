package inputHandling

import (
	"fmt"
	"strings"

	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
)

type InputStruct struct {
	fileHandler domain.FileStore
}

func New(fileHandler domain.FileStore) *InputStruct {
	return &InputStruct{
		fileHandler: fileHandler,
	}
}

func (d *InputStruct) GetInputData() (*[]routes.InputData, error) {
	input, err := d.fileHandler.ReadFile()
	if err != nil {
		return &[]routes.InputData{}, err
	}
	routeSections := ApplyFilter(input)
	inputData := createInputData(routeSections)
	return inputData, nil
}

func createInputData(routeSections *[]models.RouteSection) *[]routes.InputData {
	var inputData []routes.InputData
	for _, v := range *(*[]models.RouteSection)(routeSections) {
		item := routes.InputData{Source: v.Start, Destination: v.End, Weight: v.Distance}
		inputData = append(inputData, item)
	}
	return &inputData
}

func ApplyFilter(input *[]models.RouteSection) *[]models.RouteSection {
	var distinct []models.RouteSection

	for _, v := range *input {
		v.Start = clean(v.Start)
		v.End = clean(v.End)
		v.PassengerUse = clean(v.PassengerUse)
		skip := false
		if v.PassengerUse != "Y" {
			skip = true
		} else {
			for _, u := range distinct {
				if (v.Start == u.Start && v.End == u.End) || (v.Start == u.End && v.End == u.Start) {
					if v.Distance != u.Distance {
						fmt.Printf("\nMultiple distances for %s to %s. (%v, %v) ", v.Start, v.End, v.Distance, u.Distance)
						v.Distance = minPositiveValue(v.Distance, u.Distance)
						fmt.Printf("Using %v", v.Distance)
					}
					skip = true
					break
				}
			}
		}
		if !skip {
			distinct = append(distinct, v)
		}
	}
	return &distinct
}

func minPositiveValue(a, b int) int {
	if a <= 0 {
		return b
	}
	if a < b || b == 0 {
		return a
	}
	return b
}

func clean(input string) string {
	trimmed := strings.Trim(input, " ")
	return strings.ToUpper(trimmed)
}
