package usecases

import (
	"database/sql"
	"errors"

	"github.com/mesxx/Fiber_Task_Management_API/models"
	"github.com/mesxx/Fiber_Task_Management_API/repositories"
)

type (
	TaskUsecase interface {
		Create(requestCreateTask *models.RequestCreateTask) (*models.Task, error)
		GetAll() ([]models.Task, error)
		GetAllByUser(userID uint) ([]models.Task, error)
		GetByID(id uint, userID uint) (*models.Task, error)
		UpdateByID(task *models.Task, values map[string]interface{}) (*models.Task, error)
		DeleteByID(task *models.Task) (*models.Task, error)
		DeleteAll() error
	}

	taskUsecase struct {
		TaskRepository repositories.TaskRepository
	}
)

func NewTaskUsecase(repository repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{
		TaskRepository: repository,
	}
}

func (usecase taskUsecase) Create(requestCreateTask *models.RequestCreateTask) (*models.Task, error) {
	task := models.Task{
		UserID:      requestCreateTask.UserID,
		Status:      "active",
		Title:       requestCreateTask.Title,
		Description: sql.NullString{String: requestCreateTask.Description, Valid: requestCreateTask.Description != ""},
		Image:       sql.NullString{String: requestCreateTask.Image, Valid: requestCreateTask.Image != ""},
	}
	return usecase.TaskRepository.Create(&task)
}

func (usecase taskUsecase) GetAll() ([]models.Task, error) {
	return usecase.TaskRepository.GetAll()
}

func (usecase taskUsecase) GetAllByUser(userID uint) ([]models.Task, error) {
	return usecase.TaskRepository.GetAllByUser(userID)
}

func (usecase taskUsecase) GetByID(id uint, userID uint) (*models.Task, error) {
	getByID, err := usecase.TaskRepository.GetByID(id, userID)
	if err != nil {
		return nil, err
	} else if getByID.ID == 0 {
		return nil, errors.New("task is invalid, please try again")
	}
	return getByID, nil
}

func (usecase taskUsecase) UpdateByID(task *models.Task, values map[string]interface{}) (*models.Task, error) {
	return usecase.TaskRepository.UpdateByID(task, values)
}

func (usecase taskUsecase) DeleteByID(task *models.Task) (*models.Task, error) {
	return usecase.TaskRepository.DeleteByID(task)
}

func (usecase taskUsecase) DeleteAll() error {
	return usecase.TaskRepository.DeleteAll()
}
