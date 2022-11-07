package main

import (
	"os"

	"github.com/toluhikay/authentication/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = ":8000"
	}

	router := gin.New()
	router.Use(gin.Logger()) 

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("api-1", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message from creating api-1":"Success Creating api-1",
		})
	})

	router.GET("api-2", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message creating api-2":"Success creating api-2",
		})
	})
	
	router.Run(":" + port )
}
