package planner

import (
	"errors"

	"github.com/google/uuid"
)

func generateId() (string, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return "", errors.Join(ErrInternalError, err)
	}

	return newId.String(), nil
}
