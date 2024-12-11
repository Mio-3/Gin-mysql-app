package repository

import (
	"testing"
	"todo-app/db"
	"todo-app/models"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	database, err := db.Init()

	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	return database
}

func TestTodoRepository_Create(t *testing.T) {
	database := setupTestDB(t)

	repo := NewTodoRepository(database)

	todo := &models.Todo{
		Title:       "title",
		Description: "description",
		Completed:   false,
	}

	createdTodo, err := repo.Create(todo)
	if err != nil {
		t.Errorf("failed to create todo: %v", err)
	}

	if createdTodo.ID == 0 {
		t.Errorf("failed to create todo: invalid ID")
	}

	fetchedTodo, err := repo.FindByID(createdTodo.ID)
	if err != nil {
		t.Errorf("failed to fetch todo: %v", err)
	}

	if fetchedTodo.Title != todo.Title {
		t.Errorf("expected title %s, got %s", todo.Title, fetchedTodo.Title)
	}
}

func TestTodoRepository_FindAll(t *testing.T) {
	database := setupTestDB(t)
	repo := NewTodoRepository(database)

	todos := []*models.Todo{
		{Title: "タスク1", Description: "説明1", Completed: false},
		{Title: "タスク2", Description: "説明2", Completed: false},
	}

	for _, todo := range todos {
		_, err := repo.Create(todo)
		if err != nil {
			t.Fatalf("failed to create todo: %v", err)
		}
	}

	fetchedTodos, err := repo.FindAll()

	if err != nil {
		t.Errorf("failed to fetch todos: %v", err)
	}

	if len(fetchedTodos) < 2 {
		t.Errorf("expected at least 2 todos, got %d", len(fetchedTodos))
	}
}
