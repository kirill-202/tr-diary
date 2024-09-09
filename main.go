package main 


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)


const port = "8081"


func main() {


	if port == "" {
		port := "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
	  context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	fmt.Println("The test Go server is running on port: ", port)
	router.Run(":"+ port) 
  }