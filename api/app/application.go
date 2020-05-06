package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	// https://github.com/gin-gonic/gin#quick-start
	router = gin.Default()
)

// StartApp - Start our application
func StartApp() {
	mapUrls()

	log.Println("Start App...")
	router.Run(":8080")
}
