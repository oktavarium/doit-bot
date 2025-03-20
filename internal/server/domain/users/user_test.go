package users

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreUserFromDB(t *testing.T) {
	type args struct {
		id       string
		tgId     int64
		chatTgId int64
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr error
	}{
		{
			name: "ok",
			args: args{
				id:       "01959671-5efc-72a3-b157-838fb211840e",
				tgId:     10,
				chatTgId: 10,
				username: "user",
			},
			want: &User{
				id:       "01959671-5efc-72a3-b157-838fb211840e",
				tgId:     10,
				chatTgId: 10,
				username: "user",
				_valid:   true,
			},
			wantErr: nil,
		},
		{
			name: "bad tgid",
			args: args{
				id:       "01959671-5efc-72a3-b157-838fb211840e",
				tgId:     -10,
				chatTgId: 10,
				username: "user",
			},
			want:    nil,
			wantErr: ErrBadTgId,
		},
		{
			name: "bad chat tgid",
			args: args{
				id:       "01959671-5efc-72a3-b157-838fb211840e",
				tgId:     10,
				chatTgId: -10,
				username: "user",
			},
			want:    nil,
			wantErr: ErrBadTgId,
		},
		{
			name: "bad username",
			args: args{
				id:       "01959671-5efc-72a3-b157-838fb211840e",
				tgId:     10,
				chatTgId: 10,
				username: "",
			},
			want:    nil,
			wantErr: ErrEmptyUsername,
		},
		{
			name: "bad id",
			args: args{
				id:       "019591840e",
				tgId:     10,
				chatTgId: 10,
				username: "user",
			},
			want:    nil,
			wantErr: ErrBadId,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := RestoreUserFromDB(tt.args.id, tt.args.tgId, tt.args.chatTgId, tt.args.username)
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreUserFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
