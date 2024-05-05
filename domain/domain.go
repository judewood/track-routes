package domain

import "github.com/judewood/routeDistances/models"

type FileHandler interface {
	ReadFile() (*[]models.RouteSection, error)
	WriteFile(records *[]models.RouteDistance) error
}


type ApplyFilter interface {
	Apply(*[]models.RouteSection) *[]models.RouteSection
}

