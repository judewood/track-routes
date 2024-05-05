package mocks

import (
	"fmt"
	"errors"

	"github.com/judewood/routeDistances/models"
	mocks "github.com/stretchr/testify/mock"
)

// FileStore mock implementation of FileStore
type FileStore struct {
	mocks.Mock
}

// ReadFile returns error if provided in constructor, else returns provided route sections
func (f *FileStore) ReadFile() (*[]models.RouteSection, error) {
	fmt.Println("Mocked ReadFile function")
	args := f.Called()
	return args.Get(0).(*[]models.RouteSection), args.Error(1)
}

// WriteFile returns error if provided in constructor, else returns number of lines written to file
func (f *FileStore) WriteFile(records *[]models.RouteDistance) (int, error) {
	fmt.Println("Mocked WriteFile function")
	args := f.Called()
	return args.Get(0).(int), args.Error(1)
}
