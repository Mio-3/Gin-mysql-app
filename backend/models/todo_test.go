package models

import (
	"fmt"
	"testing"
)

func TestTodoValidation(t *testing.T) {
	tests := []struct {
		name  string
		todo Todo
		wantErr bool
	} {
		{
			name: "valid todo",
			todo : Todo{
				Title: "title",
			  Description: "description",
				Completed: false,
			},
			wantErr: false,
		},{
			name: "empty title",
			todo: Todo {
				Title: "",
				Description: "description",
				Completed: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.todo.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Todo.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}