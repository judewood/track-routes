package inputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/mocks"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
	"github.com/stretchr/testify/assert"
)

func TestApplyFilter(t *testing.T) {
	var inputData = []models.RouteSection{
		{
			Start:        "abc  ",
			End:          "XYz",
			Distance:     99,
			PassengerUse: "Y",
		},
		{ //duplicate of above
			Start:        "ABC",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but with start and end swapped
			Start:        "XYZ",
			End:          "ABC",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but not a passenger route
			Start:        "XYZ",
			End:          "ABC",
			Distance:     123,
			PassengerUse: "",
		},
		{
			Start:        "ABC",
			End:          "XYZ",
			Distance:     98,
			PassengerUse: "Y",
		},
		{
			Start:        "DEF",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ //not a passenger route
			Start:        "ABC",
			End:          "xYz",
			Distance:     66,
			PassengerUse: "not y",
		},
	}

	var expectedOutputData = []models.RouteSection{
		{
			Start:        "ABC",
			End:          "XYZ",
			Distance:     99,
			PassengerUse: "Y",
		},
		{
			Start:        "DEF",
			End:          "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
	}

	res := *(ApplyFilter(&inputData))
	for i, v := range expectedOutputData {
		if v != res[i] {
			fmt.Printf("\n ApplyFilter test failed index: %v, %v %v", i, v, res[i])
			t.Fail()
			return
		}
	}
}

func TestGetInputData(t *testing.T) {
	var fileData = []models.RouteSection{
		{
			Start:        "A",
			End:          "B",
			Distance:     4,
			PassengerUse: "Y",
		},
		{
			Start:        "A",
			End:          "C",
			Distance:     2,
			PassengerUse: "Y",
		},
		{
			Start:        "B",
			End:          "C",
			Distance:     1,
			PassengerUse: "Y",
		},
	}

	expectedInputData := []routes.InputData{
		{
			Source:      "A",
			Destination: "B",
			Weight:      4,
		},
		{
			Source:      "A",
			Destination: "C",
			Weight:      2,
		},
		{
			Source:      "B",
			Destination: "C",
			Weight:      1,
		},
	}

	mockFileStore := new(mocks.FileStore)
	mockFileStore.On("ReadFile").Return(&fileData, nil)

	inputHandler := New(mockFileStore)
	inputData, err := inputHandler.GetInputData()

	assert.NoError(t, err)
	fmt.Println(inputData)
	assert.Equal(t, &expectedInputData, inputData)
	mockFileStore.AssertExpectations(t)
	mockFileStore.AssertNumberOfCalls(t, "ReadFile", 1)
}
