package apis

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lakshaycoder01/server/app/service"
	"github.com/lakshaycoder01/server/app/utils"
)

func ViewProduct(c echo.Context) error {

	tCount := utils.ParseInt(c.QueryParam("count"), 3)

	result, e := service.GetProductCatalog(tCount)
	if e != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, result)
}
