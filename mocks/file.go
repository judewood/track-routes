package mocks

import (
	"fmt"

	"github.com/judewood/routeDistances/models"
	mocks "github.com/stretchr/testify/mock"
)

// FileStore mock implementation of FileStore
type FileStore struct {
	mocks.Mock
}

// ReadFile returns error if provided in constructor, else returns provided route sections
func (f *FileStore) ReadFile() (*[]models.RouteSection, error) {
	fmt.Println("\nMocked ReadFile function")
	args := f.Called()
	return args.Get(0).(*[]models.RouteSection), args.Error(1)
}

// WriteOutputFile returns error if provided in constructor, else returns number of lines written to file
func (f *FileStore) WriteOutputFile(records *[]models.RouteDistance) (int, error) {
	fmt.Println("\nMocked WriteFile function")
	args := f.Called()
	return args.Get(0).(int), args.Error(1)
}

// WriteOutputFile returns error if provided in constructor, else returns number of lines written to file
func (f *FileStore) WriteDetailFile(filename string, records *[]models.RouteSection) (int, error) {
	fmt.Println("\nMocked WriteFile function")
	args := f.Called()
	return args.Get(0).(int), args.Error(1)
}
