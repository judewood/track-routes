package inputHandling

import (
	"fmt"
	"testing"

	"github.com/judewood/routeDistances/mocks"
	"github.com/judewood/routeDistances/models"
	"github.com/judewood/routeDistances/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetInputData(t *testing.T) {
	var fileData = []models.RouteSection{
		{
			From:         "A",
			To:           "B",
			DistanceFrom: 4,
		},
		{
			From:         "C",
			To:           "B",
			DistanceFrom: 1,
		},
		{
			From:         "B",
			To:           "C",
			DistanceFrom: 2,
		},
	}

	expectedInputData := []routes.Edge{
		{
			From:         "A",
			To:           "B",
			DistanceFrom: 4,
		},
		{
			From:         "C",
			To:           "B",
			DistanceFrom: 1,
		},
		{
			From:         "B",
			To:           "C",
			DistanceFrom: 2,
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
