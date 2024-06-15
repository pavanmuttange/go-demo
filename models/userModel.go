package models

import "time"

type USER struct {
	Id         int `gorm:"primaryKey,Serial,not null,Unique,id"`
	Name       string
	Phone      string
	Adress     string
	Created_at time.Time
	Updated_at time.Time
}

func (e *USER) TableName() string {
	return "USER"
}
