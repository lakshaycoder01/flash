package service

import (
	"github.com/lakshaycoder01/server/app/db/product_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func SearchProduct(request *query.SearchProduct) (*query.Response, error) {

	response := new(query.Response)

	products, e := product_repo.FindSearchProducts(request)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.FindCustomerUsingID:: Unable to find product")
		response.Status = "failure"
		response.Message = "no product with given data"
		return response, nil
	}

	result := make([]*map[string]interface{}, 0)

	for _, product := range products {

		prod := map[string]interface{}{
			"product_id": product.ID,
			"brand":      product.Brand,
			"price":      product.Price,
			"quantity":   product.Quantity,
			"name":       product.Name,
		}

		result = append(result, &prod)
	}

	response.Status = "success"
	response.Message = "products with search results"
	response.Data = result

	return response, nil
}
