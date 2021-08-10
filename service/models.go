package service

import "github.com/go-playground/validator/v10"

type movies struct {
	Name    string   `json:"name,omitempty" validate:"required"`
	Ratings int      `json:"ratings,omitempty" validate:"gte=1,lte=5"`
	Actors  []string `json:"actors,omitempty" validate:"required"`
	Watched bool     `json:"watched,omitempty"`
}

type MovieService struct {
	Validator *validator.Validate
}

var datastore = defaultData

var defaultData = map[string]movies{
	"Space Jam": {
		Name:    "Space Jam",
		Ratings: 5,
		Actors:  []string{"Michael Jordan", "Bugs Bunny"},
		Watched: true,
	},
	"Scarface": {
		Name:    "Scarface",
		Ratings: 5,
		Actors:  []string{"Al Pacino"},
		Watched: true,
	},
}
