package service

import (
	"github.com/lakshaycoder01/server/app/db/customer_repo"
	"github.com/lakshaycoder01/server/app/resource/query"
	"github.com/phuslu/log"
)

func AddCustomer(requestData *query.AddCustomerRequest) (*query.CreationResponse, error) {

	response := new(query.CreationResponse)

	createData := map[string]interface{}{
		"name":  requestData.Name,
		"email": requestData.Email,
	}

	customer, e := customer_repo.Create(createData)
	if e != nil {
		log.Error().Err(e).Msgf("customer_repo.Create:: Unable to add customer")
		response.Status = "failure"
		response.Message = "unable to add customer"
		return response, nil
	}

	customerResponse := map[string]interface{}{
		"id":    customer.ID,
		"name":  customer.Name,
		"email": customer.Email,
	}

	response.Status = "success"
	response.Message = "product has been added to our system"
	response.Data = customerResponse

	return response, nil
}
