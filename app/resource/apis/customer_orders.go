package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lakshaycoder01/server/app/service"
	"github.com/lakshaycoder01/server/app/utils"
	"github.com/phuslu/log"
)

func CustomerOrders(c echo.Context) error {

	customerID := int64(utils.ParseInt(c.Param("customer_id"), 0))

	response, e := service.GetCustomerOrders(customerID)
	if e != nil {
		log.Error().Err(e).Msgf("service.AddProduct:: Unable to add product to our system")
		return utils.ErrorResponse(c, http.StatusBadGateway, e)
	}

	if response.Status == "failure" {
		c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusAccepted, response)
}
