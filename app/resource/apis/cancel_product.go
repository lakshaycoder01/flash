package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/lakshaycoder01/server/app/service"
	"github.com/lakshaycoder01/server/app/utils"
	"github.com/phuslu/log"
)

func CancelProduct(c echo.Context) error {

	request := new(query.CancelProductRequest)

	if e := c.Bind(request); e != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, e)
	}

	response, e := service.CancelProduct(request)
	if e != nil {
		log.Error().Err(e).Msgf("service.AddProduct:: Unable to add product to our system")
		return utils.ErrorResponse(c, http.StatusBadGateway, e)
	}

	if response.Status == "failure" {
		c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusAccepted, response)
}
