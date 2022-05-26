package service

import (
	"github.com/lakshaycoder01/server/app/db/customer_repo"
	"github.com/lakshaycoder01/server/app/db/product_customer_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func GetCustomerOrders(customerID int64) (*query.Response, error) {

	response := new(query.Response)

	customer, e := customer_repo.FindCustomerUsingID(customerID)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find customer")
		response.Status = "failure"
		response.Message = "no customer with given data"
		return response, nil
	}

	customerProduct, e := product_customer_repo.FindCustomerProduct(customer.ID)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find customer")
		response.Status = "failure"
		response.Message = "No product for this user"
		return response, nil
	}

	productResult := make([]*map[string]interface{}, 0)

	for _, product := range customerProduct {

		products := map[string]interface{}{
			"customer_id":      customer.ID,
			"customer_name":    customer.Name,
			"product_id":       product.ProductID,
			"product_name":     product.ProductName,
			"product_price":    product.Price,
			"product_quantity": product.Quantity,
			"product_status":   product.Status,
		}

		productResult = append(productResult, &products)
	}

	response.Status = "success"
	response.Message = "Customer product"
	response.Data = productResult

	return response, nil
}
