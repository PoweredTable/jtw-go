package main

import (
	"github.com/gin-gonic/gin"
	"jtw-go/initializers"
	"jtw-go/routes"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDB()

}

func main() {
	defer initializers.DB.Close()

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Fatal(r.Run(":8080"))
}
