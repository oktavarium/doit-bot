package handlers

func newStatusResponse(code int64, message string) Status {
	return Status{
		Code:    code,
		Message: message,
	}
}
