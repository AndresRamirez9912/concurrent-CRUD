package services

import (
	"testing"

	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

type mockRepository struct{}

func (r *mockRepository) CreateTask(task models.Task) error {
	return nil
}

func (r *mockRepository) GetTask(taskId string) (*models.Task, error) {
	return &models.Task{}, nil
}

func (r *mockRepository) UpdateTask(task models.Task, taskId string) error {
	return nil
}

func (r *mockRepository) DeleteTask(taskId string) error {
	return nil
}

func (r *mockRepository) SignUpUser(user *models.User) error {
	return nil
}

func (r *mockRepository) LogInUser(userName string) (*models.User, error) {
	return &models.User{}, nil
}

func TestCreateTask(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)

	manager.CreateTask(errCh, models.Task{})

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
}

func TestGetTask(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)
	taskCh := make(chan models.Task)

	manager.GetTask(errCh, "1", taskCh)

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
	if myTask := <-taskCh; myTask.Id == " " {
		t.Errorf("Expected a Task, got %v", myTask)
	}
}

func TestUpdateTask(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)

	manager.UpdateTask(errCh, models.Task{}, "1")

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
}

func TestDeleteTask(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)

	manager.DeleteTask(errCh, "1")

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
}

func TestSignUpUser(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)

	manager.SignUpUser(errCh, &models.User{})

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
}

func TestLogInUser(t *testing.T) {
	manager := NewManager(10, &mockRepository{})
	errCh := make(chan error)
	userCh := make(chan *models.User)

	manager.LogInUser(errCh, userCh, "")

	if err := <-errCh; err != nil {
		t.Errorf("Expected nul error, got %t", err)
	}
}
