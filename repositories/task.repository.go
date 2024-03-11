package repositories

import (
	"github.com/mesxx/Fiber_Task_Management_API/models"

	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		Create(task *models.Task) (*models.Task, error)
		GetAll() ([]models.Task, error)
		GetAllByUser(userID uint) ([]models.Task, error)
		GetByID(id uint, userID uint) (*models.Task, error)
		UpdateByID(task *models.Task, values map[string]interface{}) (*models.Task, error)
		DeleteByID(task *models.Task) (*models.Task, error)
		DeleteAll() error
	}

	taskRepository struct {
		DB *gorm.DB
	}
)

func NewTaskRepositoy(db *gorm.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (repository taskRepository) Create(task *models.Task) (*models.Task, error) {
	if err := repository.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := repository.DB.Preload("User").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repository taskRepository) GetAllByUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := repository.DB.Preload("User").Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repository taskRepository) GetByID(id uint, userID uint) (*models.Task, error) {
	var task models.Task
	if err := repository.DB.Preload("User").Where("ID = ?", id).Where("user_id = ?", userID).Find(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (repository taskRepository) UpdateByID(task *models.Task, values map[string]interface{}) (*models.Task, error) {
	if err := repository.DB.Table("tasks").Preload("User").Where("ID = ?", task.ID).Updates(values).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository taskRepository) DeleteByID(task *models.Task) (*models.Task, error) {
	if err := repository.DB.Preload("User").Where("ID = ?", task.ID).Delete(task, task.ID).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository taskRepository) DeleteAll() error {
	if err := repository.DB.Where("1 = 1").Delete(&models.Task{}).Error; err != nil {
		return err
	}
	return nil
}
