package models

type Student struct {
	Id    int `gorm:"primaryKey,not null,Unique"`
	Name  string
	Phone string
	Email string
}

func (e *Student) TableName() string {
	return "STUDENT"
}

type StudentRequestBody struct {
	Name    string
	Phone   string
	Email   string
	Courses []*Course
}
