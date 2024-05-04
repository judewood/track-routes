package inputHandling

import (
	"fmt"
	"strings"

	"github.com/judewood/routeDistances/domain"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
)

type InputStruct struct {
	fileHandler domain.FileHandler
}

func New(fileHandler domain.FileHandler) *InputStruct {
	return &InputStruct{
		fileHandler : fileHandler,
	}
}

func (d *InputStruct) GetInputData() *[]routes.InputData {
	routeSections := d.fileHandler.ReadFile()
	deduplicatedData := d.ApplyFilter(routeSections)
	inputData := func() *[]routes.InputData {
		var inputData []routes.InputData
		for _, v := range *(*[]models.RouteSection)(deduplicatedData) {
			item := routes.InputData{Source: v.Start, Destination: v.End, Weight: v.Distance}
			inputData = append(inputData, item)
		}
		return &inputData
	}()
	return inputData
}


func (d *InputStruct) ApplyFilter(input *[]models.RouteSection) *[]models.RouteSection {
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
						fmt.Printf("different distances found for same route section. \n Picking shorter distance \n %v,  \n %v", u, v)
						v.Distance = min(v.Distance, u.Distance)
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
