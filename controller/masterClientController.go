package controller

import (
	"echosample/model"
	"echosample/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type masterClientController struct {
	masterClientUseCase usecase.MasterClientUseCaseInterface
}

func MasterClientController(masterClientUseCase usecase.MasterClientUseCaseInterface) *masterClientController {
	return &masterClientController{masterClientUseCase}
}

// Insert Master client godoc
// @Summary Insert the master client data.
// @Description Insert the master client data.
// @Tags master-client
// @Accept json
// @Produce json
// @Param request body model.MasterClient true "Param Body [RAW]"
// @Success 200 {object} model.MasterClient
// @Router /v1/master-client [post]
func (controller *masterClientController) CreateMasterClient(context echo.Context) error {

	var request *model.MasterClient

	error := context.Bind(&request)
	description := []string{}
	var status *model.StandardResponse

	error = controller.masterClientUseCase.CreateMasterClient()

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
