package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize the router with default settings
	router := gin.Default()

	//Initialize routes
	initializeRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
