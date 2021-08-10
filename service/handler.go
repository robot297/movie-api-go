package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create / Post
func (s *MovieService) AddMovie(c *gin.Context) {

	var m movies
	err := c.ShouldBind(&m)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid response",
			"err":     err.Error(),
		})
		return
	}

	err = s.Validator.Struct(m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "A required field is missing or outside of the accepted parameters",
			"err":     err.Error(),
		})
		return
	}

	fmt.Println(m)
	datastore[m.Name] = m
	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie added.",
	})
}

// Read / Get all movies
func (s *MovieService) GetMovies(c *gin.Context) {
	var movies []movies
	for _, v := range datastore {
		movies = append(movies, v)
	}
	c.JSON(http.StatusOK, movies)
}

// Read / Get an individual movie
func (s *MovieService) GetMovie(c *gin.Context) {
	c.JSON(http.StatusOK, datastore[c.Param("name")])
}

// Update / Put
func (s *MovieService) UpdateMovie(c *gin.Context) {
	if val, ok := datastore[c.Param("name")]; ok {
		err := c.ShouldBind(&val)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid response",
				"err":     err.Error(),
			})
			return
		}
		datastore[c.Param("name")] = val
		c.JSON(http.StatusOK, datastore[c.Param("name")])
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Unable to find requested movie.",
	})
}

// Delete / Delete
func (s *MovieService) DeleteMovie(c *gin.Context) {
	if _, ok := datastore[c.Param("name")]; ok {
		delete(datastore, c.Param("name"))
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Movie deleted.",
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Unable to find requested movie.",
	})
}

// Reset datastore
func (s *MovieService) ResetData(c *gin.Context) {
	datastore = map[string]movies{
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
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Data reset.",
	})
}
