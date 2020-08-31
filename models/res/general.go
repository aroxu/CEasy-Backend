package res

import (
	"github.com/B1ackAnge1/CEasy-Backend/models"
)

type Empty struct{}

type Search struct {
	Count int64              `json:"count"`
	Data  []models.CeasyData `json:"data"`
}

type Location struct {
	Data []string `json:data`
}
