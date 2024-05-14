package outputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/mocks"
	"github.com/judewood/routeDistances/models"
	"github.com/stretchr/testify/assert"
)

func TestOutputRoutes(t *testing.T) {

	inputData := []models.RouteSection{
		{
			To:       "A",
			From:     "B",
			Distance: 4,
		},
		{
			To:       "A",
			From:     "C",
			Distance: 2,
		},
		{
			To:       "B",
			From:     "C",
			Distance: 1,
		},
	}

	sampleRoutes := []models.StartEnd{
		{
			From: "A",
			To:   "B",
		},
		{
			From: "C",
			To:   "B",
		}, {
			From: "A",
			To:   "C",
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

func TestUnconnectedStartAndEnd(t *testing.T) {
	inputData := []models.RouteSection{
		{
			To:       "A",
			From:     "B",
			Distance: 4,
		},
		{
			To:       "X",
			From:     "Y",
			Distance: 2,
		},
	}

	sampleRoutes := []models.StartEnd{
		{
			From: "A",
			To:   "X",
		},
	}
	mockFileStore := new(mocks.FileStore)
	mockFileStore.On("WriteFile").Return(0, nil)

	outputHandler := New(mockFileStore)
	numRecords, err := outputHandler.OutputRoutes(&inputData, &sampleRoutes)
	assert.NoError(t, err)
	assert.Equal(t, 0, numRecords)
	mockFileStore.AssertExpectations(t)
	mockFileStore.AssertNumberOfCalls(t, "WriteFile", 1)
}
