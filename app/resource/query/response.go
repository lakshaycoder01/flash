package query

type Response struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    []*map[string]interface{} `json:"data"`
}

type CreationResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
