package handlers

import "errors"

const maxSummarySize = 32

func validateCreateTaskRequest(request createTaskRequest) error {
	if len(request.Summary) > maxSummarySize {
		return errors.New("too long summary")
	}
	return nil
}

func validateUpdateTaskRequest(request updateTaskRequest) error {
	if request.Summary != nil && len(*request.Summary) > maxSummarySize {
		return errors.New("too long summary")
	}
	return nil
}
