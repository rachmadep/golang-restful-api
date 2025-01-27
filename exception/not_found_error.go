package exception

type NotFoundError struct {
	Error string `json:"error"`
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{
		Error: error,
	}
}