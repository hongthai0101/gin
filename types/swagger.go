package types

type httpError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type Swagger struct {
	HttpError httpError
}
