package task

import "gorm.io/gorm"

type Repository interface {
	Create(task *Task) error
	FindAll() ([]Task, error)
	FindOne(id int) (Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) Repository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(task *Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) FindAll() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindOne(id int) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Update(changes *Task) error {
	// get data record
	var task Task
	if err := r.db.First(&task, changes.ID).Error; err != nil {
		return err
	}

	// assign data record to struct task
	db := r.db.Model(task)

	// update by request API
	return db.Updates(*changes).Error
}

func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(&Task{}, id).Error
}
