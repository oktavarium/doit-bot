package users

import (
	"github.com/google/uuid"
	"github.com/oktavarium/doit-bot/internal/doiterr"
)

func generateId() (string, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return "", doiterr.WrapError(ErrInternalError, err)
	}

	return newId.String(), nil
}
