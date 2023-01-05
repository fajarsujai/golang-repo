package main

import (
	"log"
	"todoapp/handler"
	"todoapp/task"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "fajarsujai:Secret$123@tcp(127.0.0.1:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	taskRepository := task.NewRepository(db)

	taskService := task.NewService(taskRepository)

	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	apitask := router.Group("/task")
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	apitask.GET("/get", taskHandler.GetTasks)
	apitask.GET("/get/:id", taskHandler.GetTask)
	apitask.POST("/add", taskHandler.CreateTask)
	apitask.PUT("/update/:id", taskHandler.UpdatedTask)
	apitask.DELETE("/delete/:id", taskHandler.DeletedTask)

	router.Run()
}
