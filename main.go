package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/robot297/movie-api-go/service"
)

func main() {
	svc := service.MovieService{
		Validator: validator.New(),
	}
	router := gin.Default()
	router.POST("/movie", svc.AddMovie)
	router.GET("/movie", svc.GetMovies)
	router.GET("/movie/:name", svc.GetMovie)
	router.PUT("/movie/:name", svc.UpdateMovie)
	router.DELETE("/movie/:name", svc.DeleteMovie)
	router.Run()
}
