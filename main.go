package main

import (
	"project/config"
	"project/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// echo init
	e := echo.New()
	
	// set database
	config.InitDatabase()
	
	// route list
	routes.Route(e)
	
	// set endpoint
	e.Logger.Fatal(e.Start(":8000"))
}