package controller

import (
	"demo/models"
	"demo/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(ctx *gin.Context) {
	var err error
	db := config.DB
	var student models.StudentRequestBody

	err = ctx.ShouldBind(&student)
	if err != nil {
		fmt.Print("error in binding", err.Error())
		return
	}

	StidentData := models.Student{
		Name:  student.Name,
		Phone: student.Phone,
		Email: student.Email,
	}
	result := db.Create(&StidentData)

	if result.Error != nil || result.RowsAffected == 0 {

		fmt.Println("failed to create student")
		ctx.JSON(http.StatusBadRequest, "failed to  create student")
		return
	}
	for _, course := range student.Courses {
		crs := models.Course{
			Name:       course.Name,
			Student_id: StidentData.Id,
		}

		result = db.Create(&crs)

		if result.Error != nil || result.RowsAffected == 0 {

			fmt.Println("failed to create course", err)
			ctx.JSON(http.StatusBadRequest, "failed to  create course")
			return
		}
	}

	ctx.JSON(http.StatusOK, "student created")
}

func UpdateStudent(ctx *gin.Context) {
	var err error
	db := config.DB
	var studentRequest models.StudentRequestBody

	// Bind the request body to the studentRequest struct
	err = ctx.ShouldBind(&studentRequest)
	if err != nil {
		fmt.Print("error in binding", err.Error())
		ctx.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	// Get the student ID from the URL parameter
	studentID := ctx.Param("id")

	var existingStudent models.Student
	// Check if the student exists
	result := db.First(&existingStudent, studentID)
	if result.Error != nil {
		fmt.Println("student not found")
		ctx.JSON(http.StatusNotFound, "student not found")
		return
	}

	// Update the student's basic details
	existingStudent.Name = studentRequest.Name
	existingStudent.Phone = studentRequest.Phone
	existingStudent.Email = studentRequest.Email

	// Save the updated student details
	result = db.Save(&existingStudent)
	if result.Error != nil {
		fmt.Println("failed to update student")
		ctx.JSON(http.StatusInternalServerError, "failed to update student")
		return
	}

	// Delete existing courses associated with the student
	db.Where("student_id = ?", studentID).Delete(&models.Course{})

	// Add the updated courses
	for _, course := range studentRequest.Courses {
		crs := models.Course{
			Name:       course.Name,
			Student_id: existingStudent.Id,
		}

		result = db.Create(&crs)
		if result.Error != nil {
			fmt.Println("failed to create course", err)
			ctx.JSON(http.StatusInternalServerError, "failed to update courses")
			return
		}
	}

	ctx.JSON(http.StatusOK, "student updated")
}

// delete api
func DeleteStudent(ctx *gin.Context) {
	db := config.DB
	studentID := ctx.Param("id")

	var existingStudent models.Student
	result := db.First(&existingStudent, studentID)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, "student not found")
		} else {
			ctx.JSON(http.StatusInternalServerError, "failed to find student")
		}
		return
	}

	result = db.Where("student_id = ?", studentID).Delete(&models.Course{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to delete courses")
		return
	}

	result = db.Delete(&existingStudent)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to delete student")
		return
	}

	ctx.JSON(http.StatusOK, "student and courses deleted")
}

func GetStudentByID(ctx *gin.Context) {
	db := config.DB
	studentID := ctx.Param("id")

	var student models.Student
	result := db.First(&student, studentID)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, "student not found")
		} else {
			ctx.JSON(http.StatusInternalServerError, "failed to find student")
		}
		return
	}

	var courses []models.Course
	db.Where("student_id = ?", studentID).Find(&courses)

	studentDetails := struct {
		Student models.Student
		Courses []models.Course
	}{
		Student: student,
		Courses: courses,
	}

	ctx.JSON(http.StatusOK, studentDetails)
}

func GetAllStudents(ctx *gin.Context) {
	db := config.DB

	var students []models.Student
	result := db.Find(&students)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to retrieve students")
		return
	}

	var studentsDetails []struct {
		Student models.Student
		Courses []models.Course
	}

	for _, student := range students {
		var courses []models.Course
		db.Where("student_id = ?", student.Id).Find(&courses)

		studentDetail := struct {
			Student models.Student
			Courses []models.Course
		}{
			Student: student,
			Courses: courses,
		}

		studentsDetails = append(studentsDetails, studentDetail)
	}

	ctx.JSON(http.StatusOK, studentsDetails)
}
