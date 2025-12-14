package errs
func BadRequest(details string) *AppError {

	return New(400, "Bad Request", details, nil)
}
func NotFound(details string) *AppError {

	return New(404, "Not Found", details, nil)
}
func Conflict(details string) *AppError {
    return  New(409, "Conflict", details, nil)
}
func InternalServerError(details string) *AppError {
	return  New(500, "Internal Server Error", details, nil)
}
func RateLimitExceeded(details string) *AppError {
	return  New(429, "Rate Limit Exceeded", details, nil)
}
func Unauthorized(details string) *AppError {
	return  New(401, "Unauthorized", details, nil)
}