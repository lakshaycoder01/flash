package service

import (
	"github.com/lakshaycoder01/server/app/db/customer_repo"
	"github.com/lakshaycoder01/server/app/db/product_customer_repo"
	"github.com/lakshaycoder01/server/app/db/product_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func CancelProduct(request *query.CancelProductRequest) (*query.Response, error) {

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
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find customer")
		response.Status = "failure"
		response.Message = "no product with given data"
		return response, nil
	}

	customerProduct, e := product_customer_repo.FindProductwithCustomer(request.ProductID, request.CustomerID)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
		response.Status = "failure"
		response.Message = "no product with given data"
		return response, nil
	}

	if customerProduct.Quantity == request.Quantity {
		updates := map[string]interface{}{
			"status": "CANCELLED",
			"price":  0.0,
		}

		if e := product_customer_repo.Update(updates, request.ProductID, customer.ID); e != nil {
			log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
			response.Status = "failure"
			response.Message = "unable to cancel product"
			return response, nil
		}

		if e := product_repo.UpdateProductQuantity(request.ProductID, request.Quantity); e != nil {
			log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
			response.Status = "failure"
			response.Message = "unable to cancel product"
			return response, nil
		}
	} else {

		updates := map[string]interface{}{
			"quantity": customerProduct.Quantity - request.Quantity,
		}

		if e := product_customer_repo.Update(updates, request.ProductID, customer.ID); e != nil {
			log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
			response.Status = "failure"
			response.Message = "unable to cancel product"
			return response, nil
		}

		if e := product_repo.UpdateProductQuantity(request.ProductID, product.Quantity+request.Quantity); e != nil {
			log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
			response.Status = "failure"
			response.Message = "unable to cancel product"
			return response, nil
		}
	}

	response.Status = "succcess"
	response.Message = "Your order is cancelled, We hope to meet you in future"

	return response, nil

}
