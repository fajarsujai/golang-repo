package task

// type TaskFormatter struct {
// 	ID         int    `json:"id"`
// 	Title      string `json:"title"`
// 	IsFinished int    `json:"is_finished"`
// }

// func FormatTask(task Task) TaskFormatter {
// 	taskFormatter := TaskFormatter{}
// 	taskFormatter.ID = task.ID
// 	taskFormatter.Title = task.Title
// 	taskFormatter.IsFinished = task.IsFinished

// 	return taskFormatter
// }

// func FormatTasks(tasks []Task) []TaskFormatter {
// 	tasksFormatter := []TaskFormatter{}

// 	for _, task := range tasks {
// 		taskFormatter := FormatTask(task)
// 		tasksFormatter = append(tasksFormatter, taskFormatter)
// 	}

// 	return tasksFormatter
// }

type TasksFormatter struct {
	TaskDetailFormatter []TaskDetailFormatter `json:"List_Data"`
}

func FormatTasks(tasks []Task) TasksFormatter {
	tasksFormatter := TasksFormatter{}

	taskes := []TaskDetailFormatter{}
	for _, task := range tasks {
		taskFormatter := FormatTaskDetail(task)

		taskes = append(taskes, taskFormatter)
	}

	tasksFormatter.TaskDetailFormatter = taskes

	return tasksFormatter
}

type TaskDetailFormatter struct {
	ID          int                  `json:"Task_Id"`
	Title       string               `json:"Title"`
	ActionTime  int64                `json:"Action_Time"`
	CreatedTime int64                `json:"Created_Time"`
	UpdatedTime int64                `json:"Updated_Time"`
	IsFinishend bool                 `json:"Is_Finished"`
	Objectives  []ObjectiveFormatter `json:"Objective_Lists"`
}

type ObjectiveFormatter struct {
	ObjectiveName string `json:"Objective_Name"`
	IsFinishend   bool   `json:"Is_Finished"`
}

func FormatTaskDetail(task Task) TaskDetailFormatter {
	taskDetailFormatter := TaskDetailFormatter{}
	taskDetailFormatter.ID = task.ID
	taskDetailFormatter.Title = task.Title
	taskDetailFormatter.ActionTime = task.ActionTime
	taskDetailFormatter.CreatedTime = task.CreatedTime
	taskDetailFormatter.UpdatedTime = task.UpdatedTime
	taskDetailFormatter.IsFinishend = task.IsFinished

	objectives := []ObjectiveFormatter{}
	for _, objective := range task.Objectives {
		objectiveFormatter := ObjectiveFormatter{}
		objectiveFormatter.ObjectiveName = objective.ObjectiveName
		objectiveFormatter.IsFinishend = objective.IsFinished

		objectives = append(objectives, objectiveFormatter)
	}

	taskDetailFormatter.Objectives = objectives

	return taskDetailFormatter

}
