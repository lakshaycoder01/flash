package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lakshaycoder01/server/app/resource/apis"
)

func status(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func addRoutes(e *echo.Echo) {

	e.GET("/status", status)
	v2 := e.Group("/v2/flash_monk")

	addFlaskMonkRoutes(v2)

}

func addFlaskMonkRoutes(router *echo.Group) {
	router.POST("/add_customer", apis.AddCustomer)
	router.POST("/add_product", apis.AddProduct)
	router.GET("/view_products", apis.ViewProduct)
	router.POST("/buy_product", apis.BuyProduct)
	router.POST("/cancel_product", apis.CancelProduct)
	router.GET("/fetch_customer_orders/:customerID", apis.CustomerOrders)
	router.POST("/search_products", apis.SearchProduct)
}
