package models

type Course struct {
	Id         int `gorm:"primaryKey,not null,Unique"`
	Name       string
	Student_id int `gorm:"foreignKey referencing STUDENT Id"`
}
