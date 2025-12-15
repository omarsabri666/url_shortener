package global

import (
	"time"
)

// ACCESS_TOKEN_EXP := time

const (
	AccessTokenExp  = 1 * time.Hour
	RefreshTokenExp = 30 * 24 * time.Hour
)
