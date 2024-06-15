package controller

import (
	"demo/models"
	"demo/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ping") //for API testing
}

func CreateTable(ctx *gin.Context) {

	student := new(models.Student)
	course := new(models.Course)

	config.DB.Migrator().CreateTable(&student)
	config.DB.Migrator().CreateTable(&course)

	ctx.JSON(http.StatusOK, "table created successfully")
}
