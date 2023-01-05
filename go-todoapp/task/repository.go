package task

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Task, error)
	FindByID(ID int) (Task, error)
	FindObjectiveByTaskID(TaskID int) ([]Objective, error)
	Save(task Task) (Task, error)
	SaveObjective(objective Objective) error
	UpdateObjective(objective Objective) error
	Update(task Task) error
	DeleteByID(ID int) error
}

type repository struct {
	db *gorm.DB
}

//buat instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Task, error) {
	var task []Task

	err := r.db.Preload("Objectives").Find(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *repository) FindByID(ID int) (Task, error) {
	var task Task
	err := r.db.Where("id = ?", ID).Preload("Objectives").Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) FindObjectiveByTaskID(TaskID int) ([]Objective, error) {
	var objective []Objective
	err := r.db.Where("task_id = ?", TaskID).Find(&objective).Error

	if err != nil {
		return objective, err
	}

	return objective, nil
}

func (r *repository) Save(task Task) (Task, error) {
	err := r.db.Create(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Update(task Task) error {
	err := r.db.Save(&task).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateObjective(objective Objective) error {
	err := r.db.Save(&objective).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteByID(ID int) error {
	var task Task
	err := r.db.Where("id = ?", ID).Delete(&task).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SaveObjective(objective Objective) error {
	err := r.db.Create(&objective).Error

	if err != nil {
		return err
	}

	return nil
}
