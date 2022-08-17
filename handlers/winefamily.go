package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gprisco/decanto-winefamily-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWinefamilies(c *gin.Context) {
	if c.Query("page") == "" {
		c.JSON(http.StatusBadRequest, "'page' is required and should be > 0")
		return
	}

	if c.Query("limit") == "" {
		c.JSON(http.StatusBadRequest, "'limit' is required and should be > 0")
		return
	}

	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if page <= 0 {
		c.JSON(http.StatusBadRequest, "'page' is required and should be > 0")
		return
	}

	if limit <= 0 {
		c.JSON(http.StatusBadRequest, "'limit' is required and should be > 0")
		return
	}

	winefamilies := services.GetWinefamilies(page, limit)
	c.JSON(http.StatusOK, winefamilies)
}

func GetWinefamily(c *gin.Context) {
	winefamilyId, err := primitive.ObjectIDFromHex(c.Param("winefamilyId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	winefamily := services.GetWinefamily(winefamilyId)
	c.JSON(http.StatusOK, winefamily)
}
