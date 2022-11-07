package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/toluhikay/authentication/controllers"
	"github.com/toluhikay/authentication/middleware"
)


func UserRouter(incomingRoute *gin.Engine){
	incomingRoute.Use(middleware.Authenticate())
	incomingRoute.GET("/users", controllers.GetUsers())
	incomingRoute.GET("/users/:user_id", controllers.GetUser())
}
