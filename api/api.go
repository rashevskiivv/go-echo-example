package api

import (
	"github.com/labstack/echo"
	"github.com/rashevsiivv/echo-example/handlers"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
