package task

type GetTaskDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTaskInput struct {
	Title         string   `json:"Title" binding:"required"`
	ObjectiveList []string `json:"Objective_List" binding:"required"`
}

type UpdateTaskInput struct {
	Title         string               `json:"Title" binding:"required"`
	ObjectiveList []ObjectiveListInput `json:"Objective_List" binding:"required"`
}

type ObjectiveListInput struct {
	ObjectiveName string `json:"Objective_Name" `
	IsFinished    bool   `json:"Is_Finished"`
}
