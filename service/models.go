package service

type movies struct {
	Name    string   `json:"name,omitempty" validate:"required"`
	Ratings int      `json:"ratings,omitempty" validate:"gte=1, lte=5"`
	Actors  []string `json:"actors,omitempty" validate:"required"`
	Watched bool     `json:"watched,omitempty" validate:"required"`
}
