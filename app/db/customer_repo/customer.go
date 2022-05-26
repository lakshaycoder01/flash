package customer_repo

import (
	"github.com/lakshaycoder01/server/app/config"
	"github.com/lakshaycoder01/server/app/models"
)

func Create(createData map[string]interface{}) (*models.Customer, error) {
	customer := new(models.Customer)
	e := config.WriteDB().
		Model(&models.Customer{}).
		Create(createData).
		Where("email = ?", createData["email"].(string)).
		First(customer).
		Error

	return customer, e

}

func FindCustomerUsingID(customerID int64) (*models.Customer, error) {

	customer := new(models.Customer)

	e := config.ReadDB().
		Model(&models.Customer{}).
		Where("id = ?", customerID).
		First(&customer).Error
	return customer, e
}
