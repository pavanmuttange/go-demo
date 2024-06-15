package main

import (
	"demo/controller"
	"demo/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()
	router := gin.Default()
	router.GET("/health", controller.Ping)
	router.GET("/create", controller.CreateTable)
	router.POST("/students", controller.CreateStudent)
	router.PUT("/students/:id", controller.UpdateStudent)
	router.DELETE("/students/:id", controller.DeleteStudent)
	router.GET("/students/:id", controller.GetStudentByID)
	router.GET("/students", controller.GetAllStudents)
	router.Run(":8001")
}
