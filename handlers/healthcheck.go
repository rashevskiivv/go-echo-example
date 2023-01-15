package handlers

import (
	"github.com/labstack/echo"
	"github.com/rashevsiivv/echo-example/vo"
	"net/http"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	resp := vo.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
