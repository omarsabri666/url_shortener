package repository

import (
	"context"

	"github.com/omarsabri666/url_shorter/model/url"
)

type URLRepository interface {
	CreateURL(req url.URL, c context.Context) error
	GetURL(shortURL string) (*url.URL, error)
	IncrementCounter() (int64, error)
}
