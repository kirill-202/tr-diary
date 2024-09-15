package main 


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	//"strings"
	"net"
	//"os"
	//"html/template"
)

const port = "8081"
const mySupportMail = "supportemail@gmail.com"

//This function assumes that all templates are in the same folder
// func getTemplatePaths(path string) ([]string, error){
// 	var filenames []string

// 	entries, err := os.ReadDir(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't read directory")
// 	}
// 	for _, entry := range entries {
// 		if entry.IsDir() {
// 			continue
// 		}
// 		filenames = append(filenames, path + "/"+ entry.Name())
// 	}
// 	return filenames, nil
// }


func LoggerMiddleWare(context *gin.Context) {
	log.Printf("Full request path %s", context.FullPath())
	log.Printf("The IP address in the header")
}

func RateLimiterMiddleWare(context *gin.Context) {

	ipAddress := context.Request.RemoteAddr
	// IpPortDelim := strings.LastIndex(ipAddress, ":")
	// IP, Port := ipAddress[:IpPortDelim], ipAddress[IpPortDelim+1:]


	// println(IP, Port)
	ip, port, err := net.SplitHostPort(ipAddress)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	println(ip, port)


}

func main() {

	name := "Kirill"
	welcomeString := fmt.Sprintf("Be sure to know what you wish for %s", name)

	if port == "" {
		port := "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := gin.Default()
	router.Use(LoggerMiddleWare)
	router.Use(RateLimiterMiddleWare)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(context *gin.Context){
		context.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Homepage", 
			"Message": welcomeString,
			"SidebarContent": "My sidebar mock menu",
		})
	})

	router.GET("/about", func(context *gin.Context) {
        data := gin.H{
            "Title": "About",
            "Description": "Learn more about us on this page.",
        }
        context.HTML(http.StatusOK, "about.html", data)
    })


	fmt.Println("The test Go server is running on port: ", port)
	router.Run(":"+ port) 
  }