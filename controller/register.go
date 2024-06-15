package controller

// import (
// 	"demo/models"
// 	"demo/pkg/config"
// 	"demo/pkg/helpers"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-acme/lego/log"
// 	"github.com/gorilla/mux"
// )

// func Register(ctx *gin.Context) {
// 	var err error
// 	var userData helpers.RegisterRequestBody
// 	err = ctx.ShouldBind(&userData)
// 	if err != nil {
// 		fmt.Println("error in binding", err.Error())
// 		ctx.JSON(http.StatusBadRequest, "user not registered")
// 	}

// 	db := config.DB

// 	user := models.USER{
// 		Name:       userData.Name,
// 		Adress:     userData.Adress,
// 		Phone:      userData.Phone,
// 		Created_at: time.Now(),
// 	}

// 	db.Create(&user)
// 	ctx.JSON(http.StatusOK, "user registered")
// 	// return
// }

// func UpdateUser(ctx *gin.Context) {
// 	var err error
// 	db := config.DB
// 	var userData helpers.ModifyUserRequest
// 	var data models.USER
// 	err = ctx.ShouldBind(&userData)
// 	if err != nil {
// 		fmt.Println("error in binding", err.Error())
// 		ctx.JSON(http.StatusBadRequest, "user not updated")
// 	}
// 	vars := mux.Vars(ctx.Request)
// 	id := vars["id"]

// 	result := db.First(&data, id)
// 	if result.Error != nil {
// 		log.Println("User not found", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Println("userData", data)
// 	userData.Updated_at = time.Now()

// 	db.Model(&data).Updates(userData)

// 	ctx.JSON(http.StatusOK, "user updated ")

// }

// func DeleteUser(ctx *gin.Context) {

// 	var err error
// 	var userData models.USER
// 	db := config.DB
// 	var data int

// 	err = ctx.ShouldBind(&data)
// 	if err != nil {
// 		fmt.Println("error in binding", err.Error())
// 		ctx.JSON(http.StatusBadRequest, "user not updated")
// 	}
// 	// vars := mux.Vars(ctx.Request)
// 	id := ctx.Param("id")
// 	fmt.Println(id)
// 	result := db.First(&userData, id)
// 	if result.Error != nil || result.RowsAffected <= 0 {
// 		log.Println("User not found", http.StatusNotFound)
// 		ctx.JSON(http.StatusNotFound, "user not found ")
// 		return
// 	}

// 	if len(id) <= 0 {
// 		ctx.JSON(http.StatusBadRequest, "user not found ")
// 		return
// 	}

// 	deleteResult := db.Delete(&userData)
// 	if deleteResult.Error != nil {
// 		log.Println("Failed to delete user", http.StatusInternalServerError)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, "user Deleted ")
// }
