package routes

import (
	"echosample/controller"
	"echosample/usecase"
	"net/http"

	"github.com/labstack/echo/v4"

	_ "echosample/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {

	masterClientUseCase := usecase.MasterClientUseCase()
	masterClientController := controller.MasterClientController(masterClientUseCase)

	masterZoneController := controller.MasterZoneController()

	v1 := e.Group("/v1")
	{
		v1.GET("/", HealthCheck)
		v1.GET("/swagger/*", echoSwagger.WrapHandler)

		masterClient := v1.Group("/master-client")
		{
			masterClient.POST("/", masterClientController.CreateMasterClient)
		}

		masterZone := v1.Group("/master-zone")
		{
			masterZone.POST("/", masterZoneController.CreateMasterZone)
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
