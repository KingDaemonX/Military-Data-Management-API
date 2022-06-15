package responses

type Response struct {
	Status  uint                   `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
