package service

import (
	"github.com/lakshaycoder01/server/app/db/product_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func AddProduct(requestData *query.AddProductRequest) (*query.Response, error) {

	response := new(query.Response)

	createData := map[string]interface{}{
		"brand":    requestData.Brand,
		"name":     requestData.Name,
		"price":    requestData.Price,
		"quantity": requestData.Quantity,
	}

	if e := product_repo.CreateProduct(createData); e != nil {
		log.Error().Err(e).Msgf("product_repo.CreateProduct:: Unable to add product")
		response.Status = "failure"
		response.Message = "unable to add product"
		return response, nil
	}

	response.Status = "success"
	response.Message = "product has been added to our system"
	return response, nil
}
