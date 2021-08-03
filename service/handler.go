package service

import (
	"fmt"
	"go/format"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MovieService) AddMovie(c *gin.Context){

	var m movies
	err := c.ShouldBind(&m)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid response",
			"err": err.Error(),
		})
		return
	}

	s.Validator.Struct(m)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "A required field is missing or outside of the accepted parameters",
			"err": err.Error(),
		})
		return
	}

	fmt.Println(m)
	datastore[m.Name] = m
	c.JSON(http.StatusCreated, gin.H{
		"message": "Movie added.",
	})
}
