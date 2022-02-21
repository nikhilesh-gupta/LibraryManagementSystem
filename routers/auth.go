package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilesh-gupta/LibraryManagementSystem/handler"
)

func AuthRouter(router *gin.Engine, api handler.Handler) {
	router.POST("/signup", api.SignUp)
	router.POST("/signin", api.SignIn)
}
