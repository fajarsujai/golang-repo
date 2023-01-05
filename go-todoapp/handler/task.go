package handler

import (
	"net/http"
	"todoapp/helper"
	"todoapp/task"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	service task.Service
}

func NewTaskHandler(service task.Service) *taskHandler {
	return &taskHandler{service}
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	// c.Header("Content-Type", "application/json")
	// c.Header("Content-Type"), "application/json")
	c.Request.Header.Add("Content-Type", "application/json")
	tasks, err := h.service.GetTasks()

	if err != nil {
		response := helper.APIResponse("Error to get tasks 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseGet("Success", task.FormatTasks(tasks))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) GetTask(c *gin.Context) {
	var input task.GetTaskDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	taskDetail, err := h.service.GetTaskByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponseGet("Success", task.FormatTaskDetail(taskDetail))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) CreateTask(c *gin.Context) {

	var input task.CreateTaskInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponseError("Failed", "Error Selain yang tercantum di sini", "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	err = h.service.CreateTask(input)

	if err != nil {
		response := helper.APIResponse("Failed to create task 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponseCreate("Success")
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdatedTask(c *gin.Context) {
	var inputID task.GetTaskDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData task.UpdateTaskInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		panic(err)
		// response := helper.APIResponseError("Failed", "Error Selain yang tercantum di sini", "Ketentuan Path Param / Query Param  untuk Pemanggilan API tidak sesuai", nil)
		// c.JSON(http.StatusOK, response)
		// return
	}

	err = h.service.UpdateTask(inputID, inputData)

	if err != nil {
		response := helper.APIResponse("Failed update task", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponseCreate("Success")
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) DeletedTask(c *gin.Context) {
	var input task.GetTaskDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteTaskByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to delete task 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseCreate("Success")
	c.JSON(http.StatusOK, response)
}
