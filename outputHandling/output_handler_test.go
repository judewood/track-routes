package outputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/mocks"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
	"github.com/stretchr/testify/assert"
)

func TestOutputRoutes(t *testing.T) {

	inputData := []routes.InputData{
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

	sampleRoutes := []models.StartEnd{
		{
			Start: "A",
			End:   "B",
		},
				{
			Start: "C",
			End:   "B",
		},			{
			Start: "A",
			End:   "C",
		},
	}

	mockFileStore := new(mocks.FileStore)
	mockFileStore.On("WriteFile").Return(3, nil)

	outputHandler := New(mockFileStore)
	numRecords, err := outputHandler.OutputRoutes(&inputData, &sampleRoutes)

	assert.NoError(t, err)
	fmt.Println(inputData)
	assert.Equal(t, 3, numRecords)
	mockFileStore.AssertExpectations(t)
	mockFileStore.AssertNumberOfCalls(t, "WriteFile", 1)

}