package product_repo

import (
	"github.com/lakshaycoder01/server/app/config"
	"github.com/lakshaycoder01/server/app/models"
	"github.com/lakshaycoder01/server/app/resource/query"
)

func FindProducts(count int) ([]*models.Products, error) {

	products := make([]*models.Products, 0)

	query := config.ReadDB().
		Model(&models.Products{}).
		Limit(count)

	if count > 0 {
		query.Limit(count)
	}

	e := query.Find(&products).Error

	return products, e
}

func FindProduct(productID int64) (*models.Products, error) {

	product := new(models.Products)

	e := config.ReadDB().
		Model(&models.Products{}).
		Where("id = ?", productID).
		First(&product).Error
	return product, e

}

func UpdateProductQuantity(productID int64, quantities int) error {
	e := config.WriteDB().
		Model(&models.Products{}).
		Where("id = ?", productID).
		Update("quantity", quantities).Error
	return e
}

func CreateProduct(createData map[string]interface{}) (*models.Products, error) {
	product := new(models.Products)
	e := config.ReadDB().
		Model(&models.Products{}).
		Create(createData).
		Where("brand = ?", createData["brand"]).
		Where("name = ?", createData["name"]).
		First(product).
		Error

	return product, e
}

func FindSearchProducts(request *query.SearchProduct) ([]*models.Products, error) {

	products := make([]*models.Products, 0)

	query := config.ReadDB().
		Model(&models.Products{})

	if request.BrandName != nil {
		query = query.Where("brand = ?", *request.BrandName)
	}

	if request.Price != nil {
		query = query.Where("price = ?", *request.Price)
	}

	if request.ProductName != nil {
		query = query.Where("price = ?", *request.Price)
	}

	e := query.Find(&products).Error

	return products, e
}
