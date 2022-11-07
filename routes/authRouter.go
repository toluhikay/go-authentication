package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/toluhikay/authentication/controllers"
)


func AuthRouter(incomingRoute *gin.Engine){
	incomingRoute.POST("users/signup", controllers.SignUp())
	incomingRoute.POST("users/signin", controllers.SignIn())
}
