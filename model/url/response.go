package url

type CreateUrlResponse struct {
	Data    map[string]string `json:"data ,omitempty"`
	Error   interface{}       `json:"error ,omitempty"`
	Success bool              `json:"success"`
	Message string            `json:"message"`
}
type GetUrlResponse struct {
	Error   interface{} `json:"error ,omitempty"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
}
