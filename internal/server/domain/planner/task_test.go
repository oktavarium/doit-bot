package planner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask_SetName(t *testing.T) {
	type task struct {
		ownerId     string
		listId      *string
		name        string
		description string
	}
	type args struct {
		actorId string
		name    string
	}
	tests := []struct {
		name    string
		task    task
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "new task name",
			},
			wantErr: false,
		},
		{
			name: "empty task name",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "",
			},
			wantErr: true,
		},
		{
			name: "too long task name",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "tooooooooooo loooooooooooong taaaaaaaaaaaaaaask naaaaaaaaaaameeeeeeeee",
			},
			wantErr: true,
		},
		{
			name: "task name not changed",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				name:    "task name",
			},
			wantErr: true,
		},
		{
			name: "forbidden",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-738fb211840e",
				name:    "new task name",
			},
			wantErr: true,
		},
		{
			name: "bad actor id",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId: "019572a3-b157-738fb211840e",
				name:    "new task name",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := newTask(
				tt.task.ownerId,
				tt.task.listId,
				tt.task.name,
				tt.task.description,
			)
			assert.NoError(t, err)

			if err := tr.SetName(tt.args.actorId, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Task.SetName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTask_SetDescription(t *testing.T) {
	type task struct {
		ownerId     string
		listId      *string
		name        string
		description string
	}
	type args struct {
		actorId     string
		description string
	}
	tests := []struct {
		name    string
		task    task
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-838fb211840e",
				description: "new task description",
			},
			wantErr: false,
		},
		{
			name: "very big description",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-838fb211840e",
				description: "new task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task description",
			},
			wantErr: true,
		},
		{
			name: "forbidden",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId:     "01959671-5efc-72a3-b157-738fb211840e",
				description: "new task description",
			},
			wantErr: true,
		},
		{
			name: "bad actor id",
			task: task{
				ownerId:     "01959671-5efc-72a3-b157-838fb211840e",
				listId:      nil,
				name:        "task name",
				description: "task description",
			},
			args: args{
				actorId:     "01959671-5efc-72b211840e",
				description: "new task description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := newTask(
				tt.task.ownerId,
				tt.task.listId,
				tt.task.name,
				tt.task.description,
			)
			assert.NoError(t, err)

			if err := tr.SetDescription(tt.args.actorId, tt.args.description); (err != nil) != tt.wantErr {
				t.Errorf("Task.SetDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTask_SetListId(t *testing.T) {
	type task struct {
		ownerId string
		listId  *string
		name    string
	}
	type args struct {
		actorId string
		listId  string
	}
	tests := []struct {
		name    string
		task    task
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			task: task{
				ownerId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  nil,
				name:    "task name",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  "11959671-5efc-72a3-b157-838fb211840e",
			},
			wantErr: false,
		},
		{
			name: "bad list id",
			task: task{
				ownerId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  nil,
				name:    "task name",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  "19596asdfadf71-5efc-72a3-b157-838fb211840e",
			},
			wantErr: true,
		},
		{
			name: "list id not chaned",
			task: task{
				ownerId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  pointerToString("11959671-5efc-72a3-b157-838fb211840e"),
				name:    "task name",
			},
			args: args{
				actorId: "01959671-5efc-72a3-b157-838fb211840e",
				listId:  "11959671-5efc-72a3-b157-838fb211840e",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr, err := newTask(
				tt.task.ownerId,
				tt.task.listId,
				tt.task.name,
				"",
			)
			assert.NoError(t, err)

			if err := tr.SetListId(tt.args.actorId, tt.args.listId); (err != nil) != tt.wantErr {
				t.Errorf("Task.SetListId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newTask(t *testing.T) {
	type args struct {
		ownerId     string
		listId      *string
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
				listId:      nil,
				name:        "name",
				description: "description",
			},
			wantErr: false,
		},
		{
			name: "bad owner id",
			args: args{
				ownerId:     "01959671-5efc157-738fb211840e",
				listId:      nil,
				name:        "name",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "empty name",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				listId:      nil,
				name:        "",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "too long name",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				listId:      nil,
				name:        "tooooooooooo loooooooooooong taaaaaaaaaaaaaaask naaaaaaaaaaameeeeeeeee",
				description: "description",
			},
			wantErr: true,
		},
		{
			name: "too long description",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				listId:      nil,
				name:        "task name",
				description: "new task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task descriptionnew task description",
			},
			wantErr: true,
		},
		{
			name: "bad list id",
			args: args{
				ownerId:     "01959671-5efc-72a3-b157-738fb211840e",
				listId:      pointerToString("01959671-5efc-72a3738fb211840e"),
				name:        "name",
				description: "description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newTask(tt.args.ownerId, tt.args.listId, tt.args.name, tt.args.description)
			if !tt.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func pointerToString(s string) *string {
	return &s
}
