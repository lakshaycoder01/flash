package product_customer_repo

import (
	"github.com/lakshaycoder01/server/app/config"
	"github.com/lakshaycoder01/server/app/models"
)

func Create(createData map[string]interface{}) error {
	e := config.WriteDB().
		Model(&models.CustomerProduct{}).
		Create(createData).Error

	return e

}

func Update(updates map[string]interface{}, productID int64, customerID int64) error {
	e := config.WriteDB().
		Model(&models.CustomerProduct{}).
		Where("product_id = ?", productID).
		Where("customer_id = ?", customerID).
		Updates(updates).Error

	return e
}

func FindProductwithCustomer(productID int64, customerID int64) (*models.CustomerProduct, error) {
	product := new(models.CustomerProduct)

	e := config.ReadDB().
		Model(&models.CustomerProduct{}).
		Where("product_id = ?", productID).
		Where("customer_id = ?", customerID).
		Where("status !=", "CANCELLED").
		First(&product).Error

	return product, e
}

func FindCustomerProduct(customerID int64) ([]*models.CustomerProduct, error) {

	customerProduct := make([]*models.CustomerProduct, 0)

	e := config.ReadDB().
		Model(&models.CustomerProduct{}).
		Where("customer_id = ?", customerID).
		Where("status !=", "CANCELLED").
		First(&customerProduct).Error

	return customerProduct, e

}
