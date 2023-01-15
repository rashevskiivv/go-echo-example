package api

import (
	"github.com/labstack/echo"
	handlers2 "github.com/rashevsiivv/echo-example/handlers"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers2.HealthCheck)

	e.GET("/cats/:data", handlers2.GetCats)
	e.POST("/cats", handlers2.AddCat)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers2.MainAdmin)
}
