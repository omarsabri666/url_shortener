package url

// type CreateURLRequest struct {
// 	LongUrl string  `json:"long_url" binding:"required,valid_long_url"`
// 	Alias   *string `json:"alias,omitempty" binding:"omitempty,alias,min=3,max=7"`
// }

// type GetUrlRequest struct {
// 	ShortUrl string `json:"short_url" binding:"required"`
// }

type CreateURLRequest struct {
	LongUrl string  `json:"long_url" validate:"required,valid_long_url"`
	Alias   *string `json:"alias,omitempty" validate:"omitempty,alias,min=3,max=7"`
}

// type GetUrlRequest struct {
// 	ShortUrl string `json:"short_url" validate:"required"`
// }
