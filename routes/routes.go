package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)


func InitRoutes() {
	
	gin.SetMode(os.Getenv("APP_MODE"))
	
	server := gin.Default()
	Api(server)
	
	PORT := os.Getenv("SERVER_PORT")
	serverUrl := fmt.Sprintf("http://localhost:%s", PORT)

	fmt.Println("Server listening on: ", serverUrl)
	server.Run(":"+PORT)
}