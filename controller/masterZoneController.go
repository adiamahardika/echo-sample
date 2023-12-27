package controller

import (
	"echosample/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type masterZoneController struct {
}

func MasterZoneController() *masterZoneController {
	return &masterZoneController{}
}

// Insert Master zone godoc
// @Summary Insert the master zone data.
// @Description Insert the master zone data.
// @Tags master-zone
// @Accept */*
// @Produce json
// @Success 200 {object} model.MasterZone
// @Router /v1/master-zone [post]
func (controller *masterZoneController) CreateMasterZone(context echo.Context) error {

	var request *model.MasterZone

	error := context.Bind(&request)
	description := []string{}
	var status *model.StandardResponse

	if error == nil {

		description = append(description, "Success")
		status = &model.StandardResponse{
			HttpStatus:  http.StatusOK,
			Description: description,
		}

		return context.JSON(http.StatusOK, status)

	} else {

		description = append(description, error.Error())
		status = &model.StandardResponse{
			HttpStatus:  http.StatusInternalServerError,
			Description: description,
		}
		return context.JSON(http.StatusInternalServerError, status)
	}
}
