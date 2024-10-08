package api

type ID struct {
	Value int64 `json:"-" uri:"id,min=0" binding:"required"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
