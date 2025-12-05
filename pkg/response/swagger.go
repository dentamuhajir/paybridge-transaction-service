package response

// SwaggerSuccessResponse is a universal success wrapper for Swagger documentation.
type SwaggerSuccessResponse struct {
	Status  int    `json:"status" example:"200"`
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"success"`
	Data    any    `json:"data"`
}

// SwaggerErrorResponse is a universal error wrapper for Swagger.
type SwaggerErrorResponse struct {
	Status  int    `json:"status" example:"400"`
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"bad request"`
	Data    any    `json:"data"` // always null on errors
}
