package service

import "github.com/lakshaycoder01/server/app/db/product_repo"

func GetProductCatalog(tCount int) ([]*map[string]interface{}, error) {

	products, e := product_repo.FindProducts(tCount)
	if e != nil {
		return nil, e
	}

	resultProducts := make([]*map[string]interface{}, 0)
	for _, product := range products {

		productData := map[string]interface{}{
			"name":     product.Name,
			"brand":    product.Brand,
			"quantity": product.Quantity,
			"price":    product.Price,
		}

		resultProducts = append(resultProducts, &productData)
	}

	return resultProducts, nil
}
