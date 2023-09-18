package routes

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) *echo.Echo {	
	// students route
	studentGroup := e.Group("/students")
	StudentsRoute(studentGroup)
	
	return e
}