package main

import (
	"todo/config"
	"todo/internal/routes"
	databse "todo/pkg/database"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
	databse.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	// defer databsse.DB.close()
	r.Run()
}
