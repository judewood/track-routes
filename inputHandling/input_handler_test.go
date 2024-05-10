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
			From:         "abc  ",
			To:           "XYz",
			Distance:     99,
			PassengerUse: "Y",
		},
		{ //duplicate of above
			From:         "ABC",
			To:           "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but with start and end swapped
			From:         "XYZ",
			To:           "ABC",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ // duplicate but not a passenger route
			From:         "XYZ",
			To:           "ABC",
			Distance:     123,
			PassengerUse: "",
		},
		{
			From:         "ABC",
			To:           "XYZ",
			Distance:     98,
			PassengerUse: "Y",
		},
		{
			From:         "DEF",
			To:           "XYZ",
			Distance:     123,
			PassengerUse: "Y",
		},
		{ //not a passenger route
			From:         "ABC",
			To:           "xYz",
			Distance:     66,
			PassengerUse: "not y",
		},
	}

	var expectedOutputData = []models.RouteSection{
		{
			From:         "ABC",
			To:           "XYZ",
			Distance:     99,
			PassengerUse: "Y",
		},
		{
			From:         "DEF",
			To:           "XYZ",
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
			From:         "A",
			To:           "B",
			Distance:     4,
			PassengerUse: "Y",
		},
		{
			From:         "A",
			To:           "C",
			Distance:     2,
			PassengerUse: "Y",
		},
		{
			From:         "B",
			To:           "C",
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
