package customer_repo

import (
	"github.com/lakshaycoder01/server/app/config"
	"github.com/lakshaycoder01/server/app/models"
)

func Create(createData map[string]interface{}) error {

	e := config.WriteDB().
		Model(&models.Customer{}).
		Create(createData).Error

	return e

}

func FindCustomerUsingID(customerID int64) (*models.Customer, error) {

	customer := new(models.Customer)

	e := config.ReadDB().
		Model(&models.Customer{}).
		Where("id = ?", customerID).
		First(&customer).Error
	return customer, e
}
