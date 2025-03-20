package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTgId(t *testing.T) {
	tests := []struct {
		name    string
		tgId    int64
		wantErr error
	}{
		{
			name:    "ok",
			tgId:    256,
			wantErr: nil,
		},
		{
			name:    "bad tgid",
			tgId:    0,
			wantErr: ErrBadTgId,
		},
		{
			name:    "negative tgid",
			tgId:    -5,
			wantErr: ErrBadTgId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTgId(tt.tgId)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestValidateId(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{
			name:    "ok",
			id:      "01959671-5efc-72a3-b157-838fb211840e",
			wantErr: nil,
		},
		{
			name:    "empty id",
			id:      "",
			wantErr: ErrBadId,
		},
		{
			name:    "wrong id",
			id:      "34534245254",
			wantErr: ErrBadId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateId(tt.id)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  error
	}{
		{
			name:     "ok",
			username: "name",
			wantErr:  nil,
		},
		{
			name:     "empty username",
			username: "",
			wantErr:  ErrEmptyUsername,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validatUsername(tt.username)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
