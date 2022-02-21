package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilesh-gupta/LibraryManagementSystem/handler"
)

func BookRouter(router *gin.Engine, api handler.Handler) {
	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	router.GET("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
}
