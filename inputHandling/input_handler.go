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
		item := routes.InputData{From: v.From, To: v.To, Distance: v.DistanceFrom}
		inputData = append(inputData, item)
	}
	return &inputData
}

func ApplyFilter(input *[]models.RouteSection) *[]models.RouteSection {
	var distinct []models.RouteSection

	for _, v := range *input {
		v.From = clean(v.From)
		v.To = clean(v.To)
		skip := false
		for _, u := range distinct {
			if (v.From == u.From && v.To == u.To)  {
				//add in the distance for the revers direction - may not be the same as the forward one 
				v.DistanceTo = u.DistanceFrom
				skip = true  
				break
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
