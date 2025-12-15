package url

type URL struct {
	ShortUrl  string  `json:"short_url"`
	LongUrl   string  `json:"long_url"`
	CreatedAt string  `json:"created_at"`
	UserId    *string `json:"user_id"`
}
