package service

import (
	"github.com/lakshaycoder01/server/app/db/customer_repo"
	"github.com/lakshaycoder01/server/app/db/product_customer_repo"
	"github.com/lakshaycoder01/server/app/db/product_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func BuyProduct(request *query.BuyProduct) (*query.Response, error) {

	response := new(query.Response)

	customer, e := customer_repo.FindCustomerUsingID(request.CustomerID)

	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find customer")
		response.Status = "failure"
		response.Message = "no customer with given data"
		return response, nil
	}

	product, e := product_repo.FindProduct(request.ProductID)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
		response.Status = "failure"
		response.Message = "no product with given data"
		return response, nil
	}

	if product.Quantity <= 0 {
		response.Status = "failure"
		response.Status = "Sorry we are out of quantity for this product"
	}

	go func() {

		quantity := 1

		if request.Quantity != nil {
			quantity = *request.Quantity
		}

		if e := product_repo.UpdateProductQuantity(request.ProductID, product.Quantity-quantity); e != nil {
			log.Error().Err(e).Msgf("product_repo.UpdateProductQuantity:: Unable to update product quantity")
		}

		createData := map[string]interface{}{
			"customer_id":  customer.ID,
			"product_id":   product.ID,
			"product_name": product.Name,
			"brand":        product.Brand,
			"price":        float64(quantity) * product.Price,
			"quantity":     quantity,
			"status":       "CONFIRMED",
		}

		if e := product_customer_repo.Create(createData); e != nil {
			log.Error().Err(e).Msgf("product_customer_repo.Create:: Unable to add product bought by customer")
		}

	}()

	response.Status = "success"
	response.Message = "Your order is confirmed"

	return response, nil
}
