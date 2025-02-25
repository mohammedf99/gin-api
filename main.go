package main

import (
	"example/goUdemyRest/db"
	"example/goUdemyRest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
