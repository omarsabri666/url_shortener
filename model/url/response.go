package url

type CreateUrlResponse struct {
	ShortUrl string `json:"short_url"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
}
type GetUrlResponse struct {
	LongUrl string `json:"long_url"`
}
