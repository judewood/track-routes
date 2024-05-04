package domain

import "github.com/judewood/routeDistances/models"

type FileHandler interface {
	ReadFile() *[]models.RouteSection
	WriteFile(records *[]models.RouteDistance) 
}


type ApplyFilter interface {
	Apply(*[]models.RouteSection) *[]models.RouteSection
}

