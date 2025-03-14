package planner

func newStatusResponse(code int, message string) Status {
	return Status{
		Code:    code,
		Message: message,
	}
}
