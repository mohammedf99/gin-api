package routes

import (
	"example/goUdemyRest/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	e, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	r, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. it might not be in the database"})
		return
	}

	c.JSON(http.StatusOK, r)

}

func createEvent(c *gin.Context) {

	// CREATED MIDDLEWARE INSTEAD
	// token := c.Request.Header.Get("Authorization")

	// if token == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "user not authorized"})
	// 	return
	// }

	// uId, err := utils.VerifyJWTToken(token)

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "user not authorized"})
	// 	return
	// }

	var e models.Event
	err := c.ShouldBindJSON(&e)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not bind request data.", "error": err.Error()})
		return
	}

	// Log the incoming event data
	fmt.Printf("Incoming Event Data: %+v\n", e)

	// uId, _ := c.Get("userId")
	// e.UserID = uId.(int64)

	// OR:

	uId := c.GetInt64("userId")
	e.UserID = uId

	err = e.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": e})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	uId := c.GetInt64("userId")
	e, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find/fetch event"})
		return
	}

	if e.UserID != uId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid PUT request made"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully", "event": updatedEvent})

}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	// Course implementation
	uId := c.GetInt64("userId")
	e, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find/fetch data."})
		return
	}

	if e.UserID != uId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}

	err = e.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})

	// My implementation
	// err = models.DeleteEventById(id)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully.", "eventId": id})
}
