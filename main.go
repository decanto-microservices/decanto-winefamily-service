package main

import (
	"net/http"

	"github.com/Gprisco/decanto-winefamily-service/consul"
	"github.com/Gprisco/decanto-winefamily-service/env"
	"github.com/Gprisco/decanto-winefamily-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	consul.Register()

	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/check", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl, handlers.GetWinefamilies)
	r.GET(baseUrl+"/:winefamilyId", handlers.GetWinefamily)

	r.Run(env.GetInstance().Port)
}
