package res

import (
	"github.com/aroxu/CEasy-Backend/models"
)

//Empty returns empty structure
type Empty struct{}

//Search returns CBS search result
type Search struct {
	Count int64              `json:"count"`
	Data  []models.CeasyData `json:"data"`
}

//Location returns location data based on cbs content
type Location struct {
	Data []string `json:"data"`
}
