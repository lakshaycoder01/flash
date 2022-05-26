package service

import (
	"github.com/lakshaycoder01/server/app/db/product_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func AddProduct(requestData *query.AddProductRequest) (*query.CreationResponse, error) {

	response := new(query.CreationResponse)

	createData := map[string]interface{}{
		"brand":    requestData.Brand,
		"name":     requestData.Name,
		"price":    requestData.Price,
		"quantity": requestData.Quantity,
	}

	product, e := product_repo.CreateProduct(createData)
	if e != nil {
		log.Error().Err(e).Msgf("product_repo.CreateProduct:: Unable to add product")
		response.Status = "failure"
		response.Message = "unable to add product"
		return response, nil
	}

	productResponse := map[string]interface{}{
		"id":    product.ID,
		"brand": product.Brand,
		"name":  product.Name,
	}

	response.Status = "success"
	response.Message = "product has been added to our system"
	response.Data = productResponse
	return response, nil
}
