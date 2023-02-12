package typed_errors

type ErrorType uint

const (
	InternalError = ErrorType(iota)
	Forbidden
	Conflict
	NotFound
)

type customError struct {
	errorType ErrorType
}

func (error customError) Error() string {
	return ""
}

func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}
	return InternalError
}
