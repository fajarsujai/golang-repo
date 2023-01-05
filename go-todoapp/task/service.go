package task

import (
	"time"
)

type Service interface {
	GetTasks() ([]Task, error)
	GetTaskByID(input GetTaskDetailInput) (Task, error)
	CreateTask(input CreateTaskInput) error
	UpdateTask(inputID GetTaskDetailInput, inputData UpdateTaskInput) error
	DeleteTaskByID(input GetTaskDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTasks() ([]Task, error) {

	tasks, err := s.repository.FindAll()

	if err != nil {
		return tasks, err
	}

	return tasks, nil

}

func (s *service) GetTaskByID(input GetTaskDetailInput) (Task, error) {
	task, err := s.repository.FindByID(input.ID)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) CreateTask(input CreateTaskInput) error {
	t := time.Now()

	// Calling Unix method
	unix := t.Unix()

	// Prints output
	// unixTime := fmt.Printf("%v\n", unix)

	task := Task{}
	task.Title = input.Title
	task.ActionTime = unix
	task.CreatedTime = unix
	task.UpdatedTime = unix

	newTask, err := s.repository.Save(task)

	if err != nil {
		return err
	}

	objective := Objective{}
	for _, objectname := range input.ObjectiveList {
		objective.ObjectiveName = objectname
		objective.TaskID = newTask.ID
		err := s.repository.SaveObjective(objective)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) UpdateTask(inputID GetTaskDetailInput, inputData UpdateTaskInput) error {
	t := time.Now()

	// Calling Unix method
	unix := t.Unix()

	task, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return err
	}

	task.Title = inputData.Title
	task.UpdatedTime = unix

	err = s.repository.Update(task)

	if err != nil {
		return err
	}

	// fmt.Println(inputData.ObjectiveList[0].ObjectiveName)

	objectives, err := s.repository.FindObjectiveByTaskID(inputID.ID)

	if err != nil {
		return err
	}

	objective := Objective{}
	for key, obj := range objectives {
		objective.ObjectiveName = inputData.ObjectiveList[key].ObjectiveName
		objective.IsFinished = inputData.ObjectiveList[key].IsFinished
		objective.ID = obj.ID
		objective.TaskID = inputID.ID
		err := s.repository.UpdateObjective(objective)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteTaskByID(input GetTaskDetailInput) error {
	err := s.repository.DeleteByID(input.ID)

	if err != nil {
		return err
	}

	return nil
}
