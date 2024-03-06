package services

import (
	"gitlab.com/AndresRamirez9912/vozy-tech-evaluation/app/models"
)

type TaskManager struct {
	Semaphore      chan bool
	UserRepository models.Repository
}

type ServiceInterface interface {
	CreateTask(errCh chan error, task models.Task)
	GetTask(errCh chan error, taskId string, taskCh chan models.Task)
	UpdateTask(errCh chan error, task models.Task, taskId string)
	DeleteTask(errCh chan error, taskId string)

	SignUpUser(errCh chan error, user *models.User)
	LogInUser(errCh chan error, userCh chan *models.User, userName string)
}

func NewManager(maxWorkers int, repo models.Repository) *TaskManager {
	return &TaskManager{
		Semaphore:      make(chan bool, maxWorkers),
		UserRepository: repo,
	}
}
func (manager *TaskManager) CreateTask(errCh chan error, task models.Task) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		errCh <- manager.UserRepository.CreateTask(task)
	}()
}

func (manager *TaskManager) GetTask(errCh chan error, taskId string, taskCh chan models.Task) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		task, err := manager.UserRepository.GetTask(taskId)
		errCh <- err
		taskCh <- *task
	}()
}

func (manager *TaskManager) UpdateTask(errCh chan error, task models.Task, taskId string) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		errCh <- manager.UserRepository.UpdateTask(task, taskId)
	}()
}

func (manager *TaskManager) DeleteTask(errCh chan error, taskId string) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		errCh <- manager.UserRepository.DeleteTask(taskId)
	}()
}

func (manager *TaskManager) SignUpUser(errCh chan error, user *models.User) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		errCh <- manager.UserRepository.SignUpUser(user)
	}()
}

func (manager *TaskManager) LogInUser(errCh chan error, userCh chan *models.User, user string) {
	manager.Semaphore <- true
	go func() {
		defer func() {
			<-manager.Semaphore
			close(errCh)
		}()
		user, err := manager.UserRepository.LogInUser(user)
		errCh <- err
		userCh <- user
	}()
}
