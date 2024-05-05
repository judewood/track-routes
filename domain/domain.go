package domain

import "github.com/judewood/routeDistances/models"

// FileStore provides Read and Write methods for files
type FileStore interface {
	ReadFile() (*[]models.RouteSection, error)
	WriteFile(records *[]models.RouteDistance) (int, error)
}
