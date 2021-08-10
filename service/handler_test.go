package service

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func newMovieModel() Movies {
	return movies{
		Name:         "Test Movie",
		Ratings:      5,
		Actors:       ["Brad Pitt"],
		Watched:      true,
	}
}

func TestMovieService_AddMovie(t *testing.T) {
	type fields struct {
		Validator *validator.Validate
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MovieService{
				Validator: tt.fields.Validator,
			}
			s.AddMovie(tt.args.c)
		})
	}
}
