package planner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_SetName(t *testing.T) {
	type list struct {
		ownerId     string
		name        string
		description string
	}
	type args struct {
		actorId string
		name    string
	}
	tests := []struct {
		name    string
		list    list
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "new list name",
			},
			wantErr: false,
		},
		{
			name: "empty list name",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "",
			},
			wantErr: true,
		},
		{
			name: "too long list name",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "tooooooooooo loooooooooooong taaaaaaaaaaaaaaask naaaaaaaaaaameeeeeeeee",
			},
			wantErr: true,
		},
		{
			name: "list name not changed",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "list name",
			},
			wantErr: true,
		},
		{
			name: "forbidden",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-738fb211840e",
				name:    "new list name",
			},
			wantErr: true,
		},
		{
			name: "bad actor id",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId: "019572a3-b157-738fb211840e",
				name:    "new list name",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := newList(
				tt.list.ownerId,
				tt.list.name,
				tt.list.description,
			)
			assert.NoError(t, err)

			if err := tr.SetName(tt.args.actorId, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("List.SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_SetDescription(t *testing.T) {
	type list struct {
		ownerId     string
		name        string
		description string
	}
	type args struct {
		actorId     string
		description string
	}
	tests := []struct {
		name    string
		list    list
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-838fb211840e",
				description: "new list description",
			},
			wantErr: false,
		},
		{
			name: "very big description",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-838fb211840e",
				description: "new list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list description",
			},
			wantErr: true,
		},
		{
			name: "forbidden",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-738fb211840e",
				description: "new list description",
			},
			wantErr: true,
		},
		{
			name: "bad actor id",
			list: list{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				name:        "list name",
				description: "list description",
			},
			args: args{
				actorId:     "01959671-5efc-72b211840e",
				description: "new list description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := newList(
				tt.list.ownerId,
				tt.list.name,
				tt.list.description,
			)
			assert.NoError(t, err)

			if err := tr.SetDescription(tt.args.actorId, tt.args.description); (err != nil) != tt.wantErr {
				t.Errorf("List.SetDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newList(t *testing.T) {
	type args struct {
		ownerId     string
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				name:        "name",
				description: "description",
			},
			wantErr: false,
		},
		{
			name: "bad owner id",
			args: args{
				ownerId:     "01959671-5efc157-738fb211840e",
				name:        "name",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "empty name",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				name:        "",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "too long name",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				name:        "tooooooooooo loooooooooooong taaaaaaaaaaaaaaask naaaaaaaaaaameeeeeeeee",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "too long description",
			args: args{
				ownerId: "01959671-5efc-72a3-b157-738fb211840e",
				name:        "list name",
				description: "new list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list descriptionnew list description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newList(tt.args.ownerId, tt.args.name, tt.args.description)
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
