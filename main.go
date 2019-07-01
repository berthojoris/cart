package main

import (
	"gin-framework/routing"
	"gin-framework/config"
)

func main() {
	config.LoadENV()
	db := config.DBConnect()
	defer db.Close()
	r := routing.RegisterRoutes(db)
	r.Run(":8080")
}