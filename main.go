package main

import (
	"github.com/gin-gonic/gin/routing"
	"github.com/gin-gonic/gin/config"
)

func main() {
	config.LoadENV()
	db := config.DBConnect()
	defer db.Close()
	r := routing.RegisterRoutes(db)
	r.Run(":8080")
}