package routes

import (
	"example/goUdemyRest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	uId := c.GetInt64("userId")
	eId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	e, err := models.GetEventById(eId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event."})
	}

	err = e.Register(uId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register event."})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event registered successfully"})

}

func cancelRegistration(c *gin.Context) {
	uId := c.GetInt64("userId")
	eId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	var event models.Event
	event.ID = eId

	err = event.CancelRegistration(uId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not cancel registration", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration cancelled"})

}
