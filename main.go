package main 


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"os"
)

const port = "8081"

//This function assumes that all templates are in the same folder
func getTemplatePaths(path string) ([]string, error){
	var filenames []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("can't read directory")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filenames = append(filenames, path + "/"+ entry.Name())
	}
	return filenames, nil
}




func main() {

	name := "Kirill"
	welcomeString := fmt.Sprintf("Be sure to know what you wish for %s", name)
	templateNames, err := getTemplatePaths("./templates")

	if err != nil {
		log.Fatal("can't read templates")
	}

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

	router.LoadHTMLFiles(templateNames...)

	router.GET("/", func(context *gin.Context){
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": welcomeString, 
		})
	})


	fmt.Println("The test Go server is running on port: ", port)
	router.Run(":"+ port) 
  }