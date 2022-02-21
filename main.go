package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nikhilesh-gupta/LibraryManagementSystem/database"
	"github.com/nikhilesh-gupta/LibraryManagementSystem/handler"
	"github.com/nikhilesh-gupta/LibraryManagementSystem/routers"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func AuthMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		token := c.GetHeader("token")
		fmt.Println("got token:	" + token)
		isValid, err := handler.ValidateToken(token)
		if err != nil && !isValid {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}
func main() {
	database.Setup()
	engine := gin.Default() //get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	} //get the customized engine
	engine.Use(AuthMiddleware())
	routers.BookRouter(engine, api)
	routers.AuthRouter(engine, api)

	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}
