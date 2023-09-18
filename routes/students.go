package routes

import (
	"project/controllers"

	"github.com/labstack/echo/v4"
)

func StudentsRoute(g *echo.Group) {
	g.GET("", controllers.StudentGetData)
	g.POST("", controllers.StudentAddData)
	g.GET("/:id", controllers.StudentGetDetailData)
	g.DELETE("/:id", controllers.StudentDeleteData)
	g.PUT("/:id", controllers.StudentUpdateData)
}