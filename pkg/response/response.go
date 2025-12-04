package response

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func Success[T any](message string, data T, status int) APIResponse[T] {
	return APIResponse[T]{
		Status:  status,
		Success: true,
		Message: message,
		Data:    &data,
	}
}

func Error(message string, status int) APIResponse[any] {
	return APIResponse[any]{
		Status:  status,
		Success: false,
		Message: message,
		Data:    nil,
	}
}
