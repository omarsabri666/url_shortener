package token

type TokenResponse struct {
	Success bool              `json:"success"`
	Data    map[string]string `json:"data ,omitempty"`
	Message string            `json:"message"`
}
