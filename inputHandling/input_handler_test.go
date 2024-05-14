package inputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/mocks"
	"github.com/judewood/routeDistances/models"
	"github.com/stretchr/testify/assert"
)

func TestGetInputData(t *testing.T) {
	var fileData = []models.RouteSection{
		{
			From:     "A",
			To:       "B",
			Distance: 4,
		},
		{
			From:     "C",
			To:       "B",
			Distance: 1,
		},
		{
			From:     "B",
			To:       "C",
			Distance: 2,
		},
	}

	expectedInputData := []models.RouteSection{
		{
			From:     "A",
			To:       "B",
			Distance: 4,
		},
		{
			From:     "C",
			To:       "B",
			Distance: 1,
		},
		{
			From:     "B",
			To:       "C",
			Distance: 2,
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
