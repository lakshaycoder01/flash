package query

type AddProductRequest struct {
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddCustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BuyProduct struct {
	ProductID  int64 `json:"product_id"`
	CustomerID int64 `json:"customer_id"`
	Quantity   *int  `json:"quantity"` // *int means if quantity is optional key from api request, if there we will update that number of quantity, default 1
}

type CancelProductRequest struct {
	ProductID  int64 `json:"product_id"`
	CustomerID int64 `json:"customer_id"`
	Quantity   int   `json:"quantity"`
}

type SearchProduct struct {
	BrandName   *string  `json:"brand_name"`
	Price       *float64 `json:"price"`
	ProductName *string  `json:"product_name"`
}
