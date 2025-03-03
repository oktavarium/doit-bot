package handlers

import "errors"

const maxSummarySize = 32

func validateCreateTaskRequest(request createTaskRequest) error {
	if len(request.Summary) > maxSummarySize {
		return errors.New("too long summary")
	}
	return nil
}

func validateUpdateTaskByIdRequest(request updateTaskByIdRequest) error {
	if request.Id == "" {
		return errors.New("empty task id")
	}
	if request.Summary != nil && len(*request.Summary) > maxSummarySize {
		return errors.New("too long summary")
	}
	return nil
}

func validateSetTaskDoneByIdRequest(request setTaskDoneByIdRequest) error {
	if request.Id == "" {
		return errors.New("empty task id")
	}
	return nil
}

func validateGetTaskByIdRequest(request getTaskByIdRequest) error {
	if request.Id == "" {
		return errors.New("empty task id")
	}
	return nil
}

func validateDeleteTaskByIdRequest(request deleteTaskByIdRequest) error {
	if request.Id == "" {
		return errors.New("empty task id")
	}
	return nil
}
